package main

import (
	"fmt"
	"io"

	"github.com/gonethopper/libs/grpc"
	"github.com/gonethopper/libs/grpc/pb"
	"github.com/gonethopper/libs/signal"
)

type messageServer struct {
}

func (s *messageServer) Request(stream pb.Message_RequestServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(in.MessageID)

		if err := stream.Send(in); err != nil {
			return err
		}

	}
}

func main() {
	options := &grpc.ServerOptions{}
	options.Address = "127.0.0.1:9999"
	server := grpc.NewServer(options)
	pb.RegisterMessageServer(server.Server, &messageServer{})
	go server.RPCServe()

	signal.InitSignal()
}
