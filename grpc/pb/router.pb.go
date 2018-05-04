// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SSRouter_TransferType int32

const (
	SSRouter_TRANSTYPHEARTBEAT    SSRouter_TransferType = 0
	SSRouter_TRANSTYPEBYP2P       SSRouter_TransferType = 1
	SSRouter_TRANSTYPEBYP2G       SSRouter_TransferType = 2
	SSRouter_TRANSTYPEBYBROADCAST SSRouter_TransferType = 3
	SSRouter_TRANSTYPEBYKEY       SSRouter_TransferType = 4
)

var SSRouter_TransferType_name = map[int32]string{
	0: "TRANSTYPHEARTBEAT",
	1: "TRANSTYPEBYP2P",
	2: "TRANSTYPEBYP2G",
	3: "TRANSTYPEBYBROADCAST",
	4: "TRANSTYPEBYKEY",
}
var SSRouter_TransferType_value = map[string]int32{
	"TRANSTYPHEARTBEAT":    0,
	"TRANSTYPEBYP2P":       1,
	"TRANSTYPEBYP2G":       2,
	"TRANSTYPEBYBROADCAST": 3,
	"TRANSTYPEBYKEY":       4,
}

func (x SSRouter_TransferType) String() string {
	return proto.EnumName(SSRouter_TransferType_name, int32(x))
}
func (SSRouter_TransferType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// SSRouter Server Message结构
type SSRouter struct {
	SrcSID    int32                 `protobuf:"varint,1,opt,name=srcSID" json:"srcSID,omitempty"`
	SrcType   int32                 `protobuf:"varint,2,opt,name=srcType" json:"srcType,omitempty"`
	DestSID   int32                 `protobuf:"varint,3,opt,name=destSID" json:"destSID,omitempty"`
	DestType  int32                 `protobuf:"varint,4,opt,name=destType" json:"destType,omitempty"`
	TransType SSRouter_TransferType `protobuf:"varint,5,opt,name=transType,enum=pb.SSRouter_TransferType" json:"transType,omitempty"`
	Uid       int64                 `protobuf:"varint,6,opt,name=uid" json:"uid,omitempty"`
	BodySize  int32                 `protobuf:"varint,7,opt,name=bodySize" json:"bodySize,omitempty"`
	Body      []byte                `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *SSRouter) Reset()                    { *m = SSRouter{} }
func (m *SSRouter) String() string            { return proto.CompactTextString(m) }
func (*SSRouter) ProtoMessage()               {}
func (*SSRouter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SSRouter) GetSrcSID() int32 {
	if m != nil {
		return m.SrcSID
	}
	return 0
}

func (m *SSRouter) GetSrcType() int32 {
	if m != nil {
		return m.SrcType
	}
	return 0
}

func (m *SSRouter) GetDestSID() int32 {
	if m != nil {
		return m.DestSID
	}
	return 0
}

func (m *SSRouter) GetDestType() int32 {
	if m != nil {
		return m.DestType
	}
	return 0
}

func (m *SSRouter) GetTransType() SSRouter_TransferType {
	if m != nil {
		return m.TransType
	}
	return SSRouter_TRANSTYPHEARTBEAT
}

func (m *SSRouter) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SSRouter) GetBodySize() int32 {
	if m != nil {
		return m.BodySize
	}
	return 0
}

func (m *SSRouter) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*SSRouter)(nil), "pb.SSRouter")
	proto.RegisterEnum("pb.SSRouter_TransferType", SSRouter_TransferType_name, SSRouter_TransferType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Router service

type RouterClient interface {
	// 定义Proxy方法
	Proxy(ctx context.Context, opts ...grpc.CallOption) (Router_ProxyClient, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Proxy(ctx context.Context, opts ...grpc.CallOption) (Router_ProxyClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Router_serviceDesc.Streams[0], c.cc, "/pb.Router/Proxy", opts...)
	if err != nil {
		return nil, err
	}
	x := &routerProxyClient{stream}
	return x, nil
}

type Router_ProxyClient interface {
	Send(*SSRouter) error
	Recv() (*SSRouter, error)
	grpc.ClientStream
}

type routerProxyClient struct {
	grpc.ClientStream
}

func (x *routerProxyClient) Send(m *SSRouter) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routerProxyClient) Recv() (*SSRouter, error) {
	m := new(SSRouter)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Router service

type RouterServer interface {
	// 定义Proxy方法
	Proxy(Router_ProxyServer) error
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Proxy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouterServer).Proxy(&routerProxyServer{stream})
}

type Router_ProxyServer interface {
	Send(*SSRouter) error
	Recv() (*SSRouter, error)
	grpc.ServerStream
}

type routerProxyServer struct {
	grpc.ServerStream
}

func (x *routerProxyServer) Send(m *SSRouter) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routerProxyServer) Recv() (*SSRouter, error) {
	m := new(SSRouter)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Router",
	HandlerType: (*RouterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Proxy",
			Handler:       _Router_Proxy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "router.proto",
}

func init() { proto.RegisterFile("router.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4f, 0x83, 0x30,
	0x14, 0xc6, 0x57, 0x60, 0x6c, 0xbe, 0x90, 0xa5, 0xbe, 0xa8, 0xa9, 0x3b, 0x91, 0x9d, 0xf0, 0x42,
	0xcc, 0x76, 0xf0, 0x0c, 0x8e, 0xa8, 0x31, 0x51, 0xd2, 0xf6, 0x32, 0x6f, 0x32, 0x30, 0xd9, 0x45,
	0x48, 0x61, 0x89, 0xe8, 0x7f, 0xec, 0x5f, 0x61, 0xda, 0x81, 0xa2, 0xde, 0xbe, 0xef, 0xfb, 0x7d,
	0x8f, 0x47, 0x5b, 0xf0, 0x54, 0xb9, 0x6f, 0x0a, 0x15, 0x56, 0xaa, 0x6c, 0x4a, 0xb4, 0xaa, 0x6c,
	0xf1, 0x69, 0xc1, 0x54, 0x08, 0x6e, 0x62, 0x3c, 0x03, 0xb7, 0x56, 0x5b, 0x71, 0xb7, 0x66, 0xc4,
	0x27, 0xc1, 0x98, 0x77, 0x0e, 0x19, 0x4c, 0x6a, 0xb5, 0x95, 0x6d, 0x55, 0x30, 0xcb, 0x80, 0xde,
	0x6a, 0x92, 0x17, 0x75, 0xa3, 0x47, 0xec, 0x03, 0xe9, 0x2c, 0xce, 0x61, 0xaa, 0xa5, 0x19, 0x72,
	0x0c, 0xfa, 0xf6, 0x78, 0x05, 0x47, 0x8d, 0x7a, 0x7e, 0xad, 0x0d, 0x1c, 0xfb, 0x24, 0x98, 0x2d,
	0xcf, 0xc3, 0x2a, 0x0b, 0xfb, 0x1f, 0x09, 0xa5, 0xa6, 0x2f, 0x85, 0xd2, 0x05, 0xfe, 0xd3, 0x45,
	0x0a, 0xf6, 0x7e, 0x97, 0x33, 0xd7, 0x27, 0x81, 0xcd, 0xb5, 0xd4, 0x6b, 0xb2, 0x32, 0x6f, 0xc5,
	0xee, 0xbd, 0x60, 0x93, 0xc3, 0x9a, 0xde, 0x23, 0x82, 0xa3, 0x35, 0x9b, 0xfa, 0x24, 0xf0, 0xb8,
	0xd1, 0x8b, 0x0f, 0xf0, 0x86, 0x1f, 0xc7, 0x53, 0x38, 0x96, 0x3c, 0x7a, 0x10, 0x72, 0x93, 0xde,
	0x26, 0x11, 0x97, 0x71, 0x12, 0x49, 0x3a, 0x42, 0x84, 0x59, 0x1f, 0x27, 0xf1, 0x26, 0x5d, 0xa6,
	0x94, 0xfc, 0xcb, 0x6e, 0xa8, 0x85, 0x0c, 0x4e, 0x06, 0x59, 0xcc, 0x1f, 0xa3, 0xf5, 0x75, 0x24,
	0x24, 0xb5, 0xff, 0xb4, 0xef, 0x93, 0x0d, 0x75, 0x96, 0x2b, 0x70, 0xbb, 0x9b, 0xbe, 0x80, 0x71,
	0xaa, 0xca, 0xb7, 0x16, 0xbd, 0xe1, 0xb9, 0xe7, 0xbf, 0xdc, 0x62, 0x14, 0x90, 0x4b, 0x12, 0x3b,
	0x4f, 0x56, 0x95, 0x65, 0xae, 0x79, 0xb2, 0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xa7,
	0x55, 0xdf, 0xc2, 0x01, 0x00, 0x00,
}