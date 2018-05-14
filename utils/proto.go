package utils

import "github.com/gogo/protobuf/proto"

//NewSSMessageBody 通过messageID创建消息体
func NewSSMessageBody(messageID int32) (proto.Message, error) {
	return nil, nil
}

//EncodeMessage proto message 转化为byte流
func EncodeMessage(message proto.Message) *[]byte {
	bytes, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}

	return &bytes
}

//DecodeMessage decode byte to proto message
func DecodeMessage(b []byte, v proto.Message) error {
	return proto.Unmarshal(b, v)
}
