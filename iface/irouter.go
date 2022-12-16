package iface

import "context"

/*
	路由接口， 这里面路由是 使用框架者给该链接自定的 处理业务方法
	路由里的IRequest 则包含用该链接的链接信息和该链接的请求数据信息
*/

// BaseRouter 实现router时，先嵌入这个基类，然后根据需要对这个基类的方法进行重写
type BaseRouter struct{}

// PreHandle 这里之所以BaseRouter的方法都为空，
// 是因为有的Router不希望有PreHandle或PostHandle
// 所以Router全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化
func (br *BaseRouter) PreHandle(ctx context.Context, req IRequest)  {}
func (br *BaseRouter) Handle(ctx context.Context, req IRequest)     {}
func (br *BaseRouter) PostHandle(ctx context.Context, req IRequest) {}

type IRouter interface {
	PreHandle(ctx context.Context, request IRequest)  //在处理conn业务之前的钩子方法
	Handle(ctx context.Context, request IRequest)     //处理conn业务的方法
	PostHandle(ctx context.Context, request IRequest) //处理conn业务之后的钩子方法
}
