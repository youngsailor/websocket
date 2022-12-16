package iface

import "context"

// IServer 定义服务器接口
type IServer interface {
	// AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(ctx context.Context, bizType string, router IRouter)
}
