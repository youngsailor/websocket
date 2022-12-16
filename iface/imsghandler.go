package iface

import "context"

/*
	IMsgHandler 消息管理抽象层
*/
type IMsgHandler interface {
	DoMsgHandler(ctx context.Context, request IRequest)            //马上以非阻塞方式处理消息
	AddRouter(ctx context.Context, bizType string, router IRouter) //为消息添加具体的处理逻辑
	StartWorkerPool(ctx context.Context)                           //启动worker工作池
	SendMsgToTaskQueue(request IRequest) (err error)               //将消息交给TaskQueue,由worker进行处理
	HandleMsg(ctx context.Context, request IRequest)               //综合处理消息
}
