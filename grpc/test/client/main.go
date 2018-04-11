package main

import (
	"context"
	"fmt"

	"github.com/gonethopper/libs/grpc"
	"github.com/wuqifei/server_lib/libgrpc/test/pb"
	"google.golang.org/grpc/grpclog"
)

func main() {

	options := &grpc.ClientOptions{}
	options.Address = "127.0.0.1:9999"
	client := grpc.NewClient(options)
	// 初始化客户端

	in := &pb.HelloRequest{}
	in.Name = "game"

	l := pb.NewHelloClient(client.ClientConn)
	res, err := l.SayHello(context.Background(), in)

	if err != nil {
		grpclog.Fatalln(err)
	}
	fmt.Println(res.Message)

	state := client.ClientConn.GetState()
	log.Debug("rpc status:%s", state.String())
}
