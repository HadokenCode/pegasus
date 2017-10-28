// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pegasus_serve.proto

/*
Package serve is a generated protocol buffer package.

It is generated from these files:
	pegasus_serve.proto

It has these top-level messages:
	HandlerRequest
	HandlerReply
*/
package serve

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HandlerRequest struct {
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Options []byte `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	Path    string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
}

func (m *HandlerRequest) Reset()                    { *m = HandlerRequest{} }
func (m *HandlerRequest) String() string            { return proto.CompactTextString(m) }
func (*HandlerRequest) ProtoMessage()               {}
func (*HandlerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HandlerRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *HandlerRequest) GetOptions() []byte {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *HandlerRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type HandlerReply struct {
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Options []byte `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (m *HandlerReply) Reset()                    { *m = HandlerReply{} }
func (m *HandlerReply) String() string            { return proto.CompactTextString(m) }
func (*HandlerReply) ProtoMessage()               {}
func (*HandlerReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HandlerReply) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *HandlerReply) GetOptions() []byte {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*HandlerRequest)(nil), "serve.HandlerRequest")
	proto.RegisterType((*HandlerReply)(nil), "serve.HandlerReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Serve service

type ServeClient interface {
	HandlerSync(ctx context.Context, in *HandlerRequest, opts ...grpc.CallOption) (*HandlerReply, error)
	Handler(ctx context.Context, opts ...grpc.CallOption) (Serve_HandlerClient, error)
}

type serveClient struct {
	cc *grpc.ClientConn
}

func NewServeClient(cc *grpc.ClientConn) ServeClient {
	return &serveClient{cc}
}

func (c *serveClient) HandlerSync(ctx context.Context, in *HandlerRequest, opts ...grpc.CallOption) (*HandlerReply, error) {
	out := new(HandlerReply)
	err := grpc.Invoke(ctx, "/serve.Serve/HandlerSync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serveClient) Handler(ctx context.Context, opts ...grpc.CallOption) (Serve_HandlerClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Serve_serviceDesc.Streams[0], c.cc, "/serve.Serve/Handler", opts...)
	if err != nil {
		return nil, err
	}
	x := &serveHandlerClient{stream}
	return x, nil
}

type Serve_HandlerClient interface {
	Send(*HandlerRequest) error
	Recv() (*HandlerReply, error)
	grpc.ClientStream
}

type serveHandlerClient struct {
	grpc.ClientStream
}

func (x *serveHandlerClient) Send(m *HandlerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serveHandlerClient) Recv() (*HandlerReply, error) {
	m := new(HandlerReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Serve service

type ServeServer interface {
	HandlerSync(context.Context, *HandlerRequest) (*HandlerReply, error)
	Handler(Serve_HandlerServer) error
}

func RegisterServeServer(s *grpc.Server, srv ServeServer) {
	s.RegisterService(&_Serve_serviceDesc, srv)
}

func _Serve_HandlerSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServeServer).HandlerSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serve.Serve/HandlerSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServeServer).HandlerSync(ctx, req.(*HandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Serve_Handler_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServeServer).Handler(&serveHandlerServer{stream})
}

type Serve_HandlerServer interface {
	Send(*HandlerReply) error
	Recv() (*HandlerRequest, error)
	grpc.ServerStream
}

type serveHandlerServer struct {
	grpc.ServerStream
}

func (x *serveHandlerServer) Send(m *HandlerReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serveHandlerServer) Recv() (*HandlerRequest, error) {
	m := new(HandlerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Serve_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serve.Serve",
	HandlerType: (*ServeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandlerSync",
			Handler:    _Serve_HandlerSync_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Handler",
			Handler:       _Serve_Handler_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pegasus_serve.proto",
}

func init() { proto.RegisterFile("pegasus_serve.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0x4d, 0x4f,
	0x2c, 0x2e, 0x2d, 0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x73, 0x94, 0x22, 0xb8, 0xf8, 0x3c, 0x12, 0xf3, 0x52, 0x72, 0x52, 0x8b, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0xf3, 0xf3, 0x4a, 0x52, 0xf3, 0x4a,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0x90, 0x4c, 0x7e, 0x41, 0x49, 0x66, 0x7e,
	0x5e, 0xb1, 0x04, 0x13, 0x44, 0x06, 0xca, 0x15, 0x12, 0xe2, 0x62, 0x29, 0x48, 0x2c, 0xc9, 0x90,
	0x60, 0x56, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x9c, 0xb8, 0x78, 0xe0, 0x26, 0x17, 0xe4,
	0x54, 0x92, 0x63, 0xae, 0x51, 0x23, 0x23, 0x17, 0x6b, 0x30, 0xc8, 0x9d, 0x42, 0xd6, 0x5c, 0xdc,
	0x50, 0xd3, 0x82, 0x2b, 0xf3, 0x92, 0x85, 0x44, 0xf5, 0x20, 0x7e, 0x41, 0x75, 0xbb, 0x94, 0x30,
	0xba, 0x70, 0x41, 0x4e, 0xa5, 0x12, 0x83, 0x90, 0x35, 0x17, 0x3b, 0x54, 0x84, 0x34, 0x8d, 0x1a,
	0x8c, 0x06, 0x8c, 0x4e, 0xb2, 0x5c, 0xdc, 0xc9, 0xf9, 0xb9, 0x7a, 0xd0, 0x30, 0x74, 0xe2, 0x01,
	0xbb, 0x27, 0x00, 0xc2, 0x0b, 0x60, 0x4c, 0x62, 0x03, 0x07, 0xa7, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x0f, 0xc2, 0x6e, 0xce, 0x65, 0x01, 0x00, 0x00,
}