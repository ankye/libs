package utils

import (
	"encoding/json"

	"github.com/gogo/protobuf/proto"
	"github.com/gonethopper/libs/common"
	"github.com/gonethopper/libs/grpc/body"
)

//NewRequsetBody 通过messageID创建request消息体
func NewRequsetBody(messageID int32) (proto.Message, error) {
	switch messageID {
	case common.SSMessageIDUserLogin:
		return &body.UserLoginReq{}, nil
	case common.SSMessageIDUserRegister:
		return &body.UserRegisterReq{}, nil
	}
	return nil, nil
}

//NewResponseBody 通过messageID创建response消息体
func NewResponseBody(messageID int32) (proto.Message, error) {
	switch messageID {
	case common.SSMessageIDUserLogin:
		return &body.UserLoginResp{}, nil
	case common.SSMessageIDUserRegister:
		return &body.UserRegisterResp{}, nil
	}
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

//JSON2PB json convert protobuf
func JSON2PB(jstr string, pb proto.Message) error {
	// json字符串转pb
	return json.Unmarshal([]byte(jstr), &pb)
}

//PB2JSON protobuf convert to json
func PB2JSON(pb proto.Message) (string, error) {
	// pb转json字符串
	jsonStr, err := json.Marshal(pb)
	if err == nil {
		jstr := string(jsonStr)
		return jstr, nil
	}

	return "", err
}
