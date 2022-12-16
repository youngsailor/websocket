/**
 * @Author: 591245853@qq.com
 * @Description:
 * @File: iconnhandler
 * @Date: 2022/12/16 16:09
 */

package iface

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/olahol/melody"
	"net/http"
)

// IConnHandler Caller implements this interface to handle websocket connection.
type IConnHandler interface {
	HandleClose(*melody.Session, int, string) error
	HandleConnect(*melody.Session)
	HandleDisconnect(*melody.Session)
	HandleError(*melody.Session, error)
	HandleMessage(*melody.Session, []byte)
	HandleMessageBinary(*melody.Session, []byte)
	HandlePong(*melody.Session)
	HandleSentMessage(*melody.Session, []byte)
	HandleSentMessageBinary(*melody.Session, []byte)
	HandleRequest(w http.ResponseWriter, r *http.Request) error
	HandleRequestWithKeys(w http.ResponseWriter, r *http.Request, keys map[string]interface{}) error
	HandleHttpRequest(r *ghttp.Request)
}
