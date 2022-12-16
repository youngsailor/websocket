# websocket
a small go websocket library with business router
一个封装的包含业务路由的go websocket库。

例子：https://github.com/youngsailor/websocket-example

业务路由参考：
```go
func WebsocketRouterInit(ctx context.Context, wsServer iface.IServer) {
	wsServer.AddRouter(ctx, "print", &controller.PrintRouter.Print)
}
```
