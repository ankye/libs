package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/gonethopper/libs/grpc"
	"github.com/gonethopper/libs/grpc/pb"
	"google.golang.org/grpc/grpclog"
)

func main() {

	options := &grpc.ClientOptions{}
	options.Address = "127.0.0.1:9999"
	client := grpc.NewClient(options)
	// 初始化客户端
	c := pb.NewMessageClient(client.ClientConn)

	stream, err := c.Request(context.Background())
	if err != nil {
		grpclog.Fatalln(err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %d ", in.MessageID)
		}
	}()
	for i := 1; i < 100; i++ {
		msg := &pb.SSMessage{MessageID: int32(i)}
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	stream.CloseSend()
	<-waitc

}
