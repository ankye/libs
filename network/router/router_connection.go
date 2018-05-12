package router

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/gonethopper/libs/ds/queue"
	libgrpc "github.com/gonethopper/libs/grpc"
	"github.com/gonethopper/libs/grpc/pb"
	log "github.com/gonethopper/libs/logs"
	"google.golang.org/grpc/metadata"
)

//NewRouterConnection 创建router连接
func NewRouterConnection(serverID int, address string, queueSize uint32) *RouterConnection {
	r := new(RouterConnection)
	r.Address = address
	r.ServerID = serverID
	r.Queue = queue.NewQueue(1, queueSize, false, r)
	r.QuitChan = make(chan struct{})

	options := &libgrpc.ClientOptions{}
	options.Address = r.Address
	client, err := libgrpc.NewClient(options)
	if err != nil {
		log.Error("server connect error :", err)
		return nil

	}
	r.Client = client

	return r
}

//RouterConnection router connection
type RouterConnection struct {
	Client   *libgrpc.Client
	Stream   pb.Router_ProxyClient
	Address  string
	ServerID int
	Queue    *queue.Queue
	QuitChan chan struct{}
}

//Connect 连接到服务器
func (r *RouterConnection) Connect() error {

	// 初始化客户端
	c := pb.NewRouterClient(r.Client.ClientConn)

	header := r.GetHeader()
	// 开启流
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	stream, err := c.Proxy(ctx)
	if err != nil {
		return err
	}
	r.Stream = stream
	go r.ReceiveServerMessage()
	return nil

}

//DisConnect 断开连接
func (r *RouterConnection) DisConnect() error {
	if r.Stream != nil {
		if err := r.Stream.CloseSend(); err != nil {
			r.Stream = nil
			return err
		}
		r.Stream = nil
	}

	return nil
}

//IsConnected 是否已经连接上
func (r *RouterConnection) IsConnected() bool {
	if r.Stream == nil {
		return false
	}

	return true
}

//ReceiveServerMessage receive server message and send to logic queue
func (r *RouterConnection) ReceiveServerMessage() {
	for {

		if r.Stream == nil {
			return
		}
		fmt.Println("hello world")
		in, err := r.Stream.Recv()
		if err == io.EOF {
			// read done.
			log.Error("RPC Stream close address:[%s]", r.Address)
			r.Stream = nil
			return
		}
		if err != nil {
			log.Error("Router connection closed by server,Failed to receive a note : %v", err)
			r.Stream = nil
			return
		}
		if in != nil {
			log.Info("Got message %d ", in.TransType)
			//send to logic queue

			r.PostMessage(in)
		}

		// deliver the data to the input queue of agent()

	}
}

//Consume 队列消费者,需要发送给router
func (r *RouterConnection) Consume(queueID, lowerSequence, upperSequence int64) {

	if r.Queue == nil {
		return
	}
	for sequence := lowerSequence; sequence <= upperSequence; sequence++ {
		message := r.Queue.Pop(sequence) // see performance note on producer sample above
		//fmt.Println("get message:", message)
		if message != nil {
			if r.Stream == nil {
				fmt.Println("stream closed")
				continue
			}
			if err := r.Stream.Send(message.(*pb.SSRouter)); err != nil {
				log.Error("Failed to send a note: %v", err)
			}
		}
		//	time.Sleep(1 * time.Second)
		// handle the incoming message with your application code
	}

}

//GetHeader get gprc header
func (r *RouterConnection) GetHeader() metadata.MD {
	m := metadata.New(
		map[string]string{
			"sid": strconv.Itoa(GetRouterManagerInstance().GetLocalServerID()),
		})
	return m
}

//Produce 数据入队列
func (r *RouterConnection) Produce(m *pb.SSRouter) {
	if r.Queue == nil {
		return
	}
	if err := r.Queue.Push(m); err != nil {
		log.Error("queue error %s", err)
	}
}

//PostMessage 投递消息到主逻辑，把router message转换为ssmessage
func (r *RouterConnection) PostMessage(m *pb.SSRouter) {
	buf := m.Body
	message := &pb.SSMessage{}
	if err := proto.Unmarshal(buf, message); err != nil {
		log.Error("unmarshaling error: ", err)
	} // 测试结果
	GetRouterManagerInstance().Consume(message)

}
