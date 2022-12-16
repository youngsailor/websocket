package util

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/olahol/melody"
)

const (
	SuccessCode int = 0
	ErrorCode   int = -1
)

type ServerMsg struct {
	BizType string `json:"biz_type"` //业务消息类型
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"message"`
}

var serverMsg = new(ServerMsg)

func WriteJsonExit(s *melody.Session, bizType string, code int, msg string, data ...interface{}) error {
	return serverMsg.jsonExit(s, bizType, code, msg, data...)
}

func WriteJson(s *melody.Session, bizType string, code int, msg string, data ...interface{}) error {
	return serverMsg.wJson(s, bizType, code, msg, data...)
}

func WriteSusJson(s *melody.Session, bizType string, msg string, data ...interface{}) error {
	return serverMsg.susJson(false, s, bizType, msg, data...)
}

func WriteSusJsonExit(s *melody.Session, bizType string, msg string, data ...interface{}) error {
	return serverMsg.susJson(true, s, bizType, msg, data...)
}

func WriteFailJson(s *melody.Session, bizType string, msg string, data ...interface{}) error {
	return serverMsg.failJson(false, s, bizType, msg, data...)
}

func WriteFailJsonExit(s *melody.Session, bizType string, msg string, data ...interface{}) error {
	return serverMsg.failJson(true, s, bizType, msg, data...)
}

// WriteJsonExit 返回JSON数据并退出当前HTTP执行函数。
func (m *ServerMsg) jsonExit(s *melody.Session, bizType string, code int, msg string, data ...interface{}) error {
	err := m.wJson(s, bizType, code, msg, data...)
	if err != nil {
		return err
	}

	return s.Close()
}

// wJson 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// code:  状态码(200:成功,302跳转，和http请求状态码一至);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func (m *ServerMsg) wJson(s *melody.Session, bizType string, code int, msg string, data ...interface{}) error {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	serverMsg = &ServerMsg{
		BizType: bizType,
		Code:    code,
		Msg:     msg,
		Data:    responseData,
	}
	marshalMsg, err := json.Marshal(serverMsg)
	if err != nil {
		return gerror.Wrap(err, "session json marshal wrong")
	}

	err = s.Write(marshalMsg)
	if err != nil {
		return gerror.Wrap(err, "session write wrong")
	}
	return nil
}

// WriteSusJson 成功返回JSON
func (m *ServerMsg) susJson(isExit bool, s *melody.Session, bizType string, msg string, data ...interface{}) error {
	if isExit {
		return m.jsonExit(s, bizType, SuccessCode, msg, data...)
	}
	return m.wJson(s, bizType, SuccessCode, msg, data...)
}

// failJson 失败返回JSON
func (m *ServerMsg) failJson(isExit bool, s *melody.Session, bizType string, msg string, data ...interface{}) error {
	if isExit {
		return m.jsonExit(s, bizType, ErrorCode, msg, data...)
	}
	return m.wJson(s, bizType, ErrorCode, msg, data...)
}
