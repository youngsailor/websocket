package impl

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/youngsailor/websocket/config"
	"github.com/youngsailor/websocket/iface"
	"time"
)

var workerID uint64

type MsgHandler struct {
	Apis           map[string]iface.IRouter //存放每个bizType 所对应的处理方法的map属性
	WorkerPoolSize uint64                   //业务工作Worker池的数量
	TaskQueue      []chan iface.IRequest    //Worker负责取任务的消息队列
}

func (mh *MsgHandler) HandleMsg(ctx context.Context, request iface.IRequest) {
	if config.WsConf.WorkerPoolSize > 0 {
		//已经启动工作池机制，将消息交给Worker处理
		err := mh.SendMsgToTaskQueue(request)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
	} else {
		//从绑定好的消息和对应的处理方法中执行对应的Handle方法
		go mh.DoMsgHandler(ctx, request)
	}
}

func NewMsgHandle(ctx context.Context) *MsgHandler {
	m := &MsgHandler{
		Apis:           make(map[string]iface.IRouter),
		WorkerPoolSize: config.WsConf.WorkerPoolSize,
		//一个worker对应一个queue
		TaskQueue: make([]chan iface.IRequest, config.WsConf.WorkerPoolSize),
	}
	m.StartWorkerPool(ctx)
	return m
}

// SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
func (mh *MsgHandler) SendMsgToTaskQueue(request iface.IRequest) (err error) {
	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则

	//分发方法一：得到需要处理此条连接的workerID
	//sessionIdValue, exists := request.GetSession().Get("session_id")
	//if !exists {
	//	return fmt.Errorf("session_id not exists")
	//}
	//sessionId, ok := sessionIdValue.(snowflake.ID)
	//if !ok {
	//	return fmt.Errorf("session_id type error")
	//}
	//workerID := uint64(sessionId.Int64()) % mh.WorkerPoolSize

	///分发方法二：根据时间戳来分配
	//workerID := uint64(time.Now().UnixMilli()) % mh.WorkerPoolSize

	//分发方法三：轮询
	workerID++
	if workerID == mh.WorkerPoolSize {
		workerID = 0
	}

	//fmt.Println("Add ConnID=", request.GetSession().GetConnID()," request bizType=", request.GetBizType(), "to workerID=", workerID)
	//将请求消息发送给任务队列
	select {
	case mh.TaskQueue[workerID] <- request:
		//fmt.Println("workerID=", workerID, "queue len=", len(mh.TaskQueue[workerID]))
		//default:
		//	// todo channle满了发送通知
		//	return fmt.Errorf("workerID = %d task queue full", workerID)
	}
	return nil
}

// DoMsgHandler 马上以非阻塞方式处理消息
func (mh *MsgHandler) DoMsgHandler(ctx context.Context, request iface.IRequest) {
	time.AfterFunc(time.Duration(config.WsConf.Timeout)*time.Second, func() {
		g.Log().Error(ctx, "request bizType=", request.GetBizType(), "timeout")
		return
	})
	handler, ok := mh.Apis[request.GetBizType()]
	if !ok {
		g.Log().Error(ctx, "api biz_type = ", request.GetBizType(), " is not FOUND!")
		return
	}
	//执行对应处理方法
	handler.PreHandle(ctx, request)
	handler.Handle(ctx, request)
	handler.PostHandle(ctx, request)
}

// AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandler) AddRouter(ctx context.Context, bizType string, router iface.IRouter) {
	//1 判断当前msg绑定的API处理方法是否已经存在
	if _, ok := mh.Apis[bizType]; ok {
		panic("repeated api , bizType = " + bizType)
	}
	//2 添加msg与api的绑定关系
	mh.Apis[bizType] = router
	g.Log().Info(ctx, "Add api bizType = ", bizType)
}

// StartOneWorker 启动一个Worker工作流程
func (mh *MsgHandler) StartOneWorker(ctx context.Context, workerID int, taskQueue chan iface.IRequest) {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Error(ctx, "workerID = ", workerID, " panic = ", r)
		}
	}()
	g.Log().Info(ctx, "Worker ID = ", workerID, " is started.")
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			mh.DoMsgHandler(ctx, request)
		}
	}
}

// StartWorkerPool 启动worker工作池
func (mh *MsgHandler) StartWorkerPool(ctx context.Context) {
	//遍历需要启动worker的数量，依此启动
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//给当前worker对应的任务队列开辟空间
		mh.TaskQueue[i] = make(chan iface.IRequest, config.WsConf.MaxWorkTaskLen)
		//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go mh.StartOneWorker(ctx, i, mh.TaskQueue[i])
	}
}
