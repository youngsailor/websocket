/**
 * @Author: 591245853@qq.com
 * @Description:
 * @File: serverMsg
 * @Date: 2022/11/18 21:26
 */

package util

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/olahol/melody"
	"github.com/youngsailor/websocket/vars"
)

// BroadcastJson Broadcast broadcasts a text message to all sessions.
func BroadcastJson(bizType string, msg string, data interface{}) error {
	serverMsg := ServerMsg{
		BizType: bizType,
		Code:    0,
		Msg:     msg,
		Data:    data,
	}
	marshalData, err := json.Marshal(serverMsg)
	if err != nil {
		return gerror.Wrap(err, "broadcast json marshal error")
	}
	return vars.WsHub.Broadcast(marshalData)
}

// BroadcastOthersJson BroadcastOthers broadcasts a text message to all sessions except session s.
func BroadcastOthersJson(s *melody.Session, bizType string, msg string, data interface{}) error {
	serverMsg := ServerMsg{
		BizType: bizType,
		Code:    0,
		Msg:     msg,
		Data:    data,
	}
	marshalData, err := json.Marshal(serverMsg)
	if err != nil {
		return gerror.Wrap(err, "broadcast others json marshal error")
	}

	return vars.WsHub.BroadcastOthers(marshalData, s)
}

// BroadcastMultipleJson broadcasts a text message to multiple sessions given in the sessions slice.
func BroadcastMultipleJson(sessions []*melody.Session, bizType string, msg string, data interface{}) error {
	serverMsg := ServerMsg{
		BizType: bizType,
		Code:    0,
		Msg:     msg,
		Data:    data,
	}

	marshalData, err := json.Marshal(serverMsg)
	if err != nil {
		return gerror.Wrap(err, "broadcast others json marshal error")
	}
	return vars.WsHub.BroadcastMultiple(marshalData, sessions)
}

// BroadcastFilterJson broadcasts a text message to all sessions that fn returns true for.
func BroadcastFilterJson(msg interface{}, fn func(*melody.Session) bool) error {
	marshalData, err := json.Marshal(msg)
	if err != nil {
		return gerror.Wrap(err, "broadcast filter json marshal error")
	}
	return vars.WsHub.BroadcastFilter(marshalData, fn)
}
