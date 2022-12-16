/**
 * @Author: 591245853@qq.com
 * @Description:
 * @File: message
 * @Date: 2022/11/16 12:00
 */

package types

type Message struct {
	BizType string `json:"biz_type"` //业务消息类型
	Data    string `json:"data"`     //消息的内容
}
