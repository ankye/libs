package db

import "github.com/gonethopper/libs/grpc/pb"

//Consumer 定义接口
type Consumer interface {
	GetDBConfig() *DBGroup
	ConsumerDBMessage(message *pb.SSMessage)
}
