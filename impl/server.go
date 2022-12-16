package impl

import (
	"context"
	"github.com/youngsailor/websocket/config"
	"github.com/youngsailor/websocket/iface"
)

type Server struct {
	//服务器IP地址
	IP string
	//端口号
	Port int
	//Server的消息管理模块
	MsgHandler iface.IMsgHandler
}

func (s *Server) AddRouter(ctx context.Context, bizType string, router iface.IRouter) {
	s.MsgHandler.AddRouter(ctx, bizType, router)
}

func NewServer(ctx context.Context) iface.IServer {
	return &Server{
		IP:         config.WsConf.Ip,
		Port:       config.WsConf.Port,
		MsgHandler: NewMsgHandle(ctx),
	}
}
