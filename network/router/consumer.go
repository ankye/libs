package router

import "github.com/gonethopper/libs/grpc/pb"

//Consumer 定义接口
type Consumer interface {
	GetServerID() int
	ConsumerRouterMessage(message *pb.SSMessage)
}
