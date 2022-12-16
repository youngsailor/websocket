package impl

type Message struct {
	BizType string //业务消息ID
	Data    []byte //消息的内容
	DataLen uint32 //消息的长度
}

// NewMsg 创建一个Message消息包
func NewMsg(bizType string, data []byte) *Message {
	return &Message{
		BizType: bizType,
		Data:    data,
		DataLen: uint32(len(data)),
	}
}

// GetBizType 获取消息类型
func (msg *Message) GetBizType() string {
	return msg.BizType
}

// GetDataLen 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// GetData 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}
