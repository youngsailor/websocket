package impl

import (
	"github.com/olahol/melody"
	"github.com/youngsailor/websocket/iface"
)

type Request struct {
	msg     iface.IMessage  // 当前客户端请求的数据
	session *melody.Session // 当前session
	m       *melody.Melody  // 当前melody
}

func NewRequest(msg iface.IMessage, s *melody.Session, m *melody.Melody) iface.IRequest {
	return &Request{
		msg:     msg,
		session: s,
		m:       m,
	}
}

// 获取请求连接信息
func (r *Request) GetSession() *melody.Session {
	return r.session
}

// 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// 获取请求的消息的ID
func (r *Request) GetBizType() string {
	return r.msg.GetBizType()
}

// GetMelody returns melody object
func (r *Request) GetMelody() *melody.Melody {
	return r.m
}

// Get returns the value for the given key, ie: (value, true). If the value does not exists it returns (nil, false)
func (r *Request) Get(key string) (value interface{}, exist bool) {
	return r.GetSession().Get(key)
}
