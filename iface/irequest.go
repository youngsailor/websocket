package iface

import (
	"github.com/olahol/melody"
)

/*
IRequest 接口：
实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
*/
type IRequest interface {
	GetSession() *melody.Session                    //获取请求连接信息
	GetData() []byte                                //获取请求消息的数据
	GetBizType() string                             //获取请求的消息ID
	GetMelody() *melody.Melody                      //获取当前melody
	Get(key string) (value interface{}, exist bool) //获取当前玩家信息
}
