// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cluster/cluster.proto

package cluster

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	"github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GrpcRequest struct {
	Envelope             *types.Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GrpcRequest) Reset()         { *m = GrpcRequest{} }
func (m *GrpcRequest) String() string { return proto.CompactTextString(m) }
func (*GrpcRequest) ProtoMessage()    {}
func (*GrpcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca74c088e8b0dfed, []int{0}
}
func (m *GrpcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcRequest.Unmarshal(m, b)
}
func (m *GrpcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcRequest.Marshal(b, m, deterministic)
}
func (m *GrpcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcRequest.Merge(m, src)
}
func (m *GrpcRequest) XXX_Size() int {
	return xxx_messageInfo_GrpcRequest.Size(m)
}
func (m *GrpcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcRequest proto.InternalMessageInfo

func (m *GrpcRequest) GetEnvelope() *types.Envelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

type GrpcResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code                 int64    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcResponse) Reset()         { *m = GrpcResponse{} }
func (m *GrpcResponse) String() string { return proto.CompactTextString(m) }
func (*GrpcResponse) ProtoMessage()    {}
func (*GrpcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca74c088e8b0dfed, []int{1}
}
func (m *GrpcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcResponse.Unmarshal(m, b)
}
func (m *GrpcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcResponse.Marshal(b, m, deterministic)
}
func (m *GrpcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcResponse.Merge(m, src)
}
func (m *GrpcResponse) XXX_Size() int {
	return xxx_messageInfo_GrpcResponse.Size(m)
}
func (m *GrpcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcResponse proto.InternalMessageInfo

func (m *GrpcResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GrpcResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GrpcResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GrpcRequest)(nil), "GrpcRequest")
	proto.RegisterType((*GrpcResponse)(nil), "GrpcResponse")
}

func init() { proto.RegisterFile("cluster/cluster.proto", fileDescriptor_ca74c088e8b0dfed) }

var fileDescriptor_ca74c088e8b0dfed = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x4f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x75, 0xad, 0x68, 0x9d, 0x44, 0x90, 0xc5, 0x42, 0xe8, 0x29, 0xe4, 0x62, 0x0e, 0x65, 0x0a,
	0xf5, 0x22, 0x78, 0x8b, 0x88, 0xd7, 0xb0, 0x7f, 0x10, 0x37, 0x43, 0x2f, 0x49, 0x66, 0xdd, 0xd9,
	0x0a, 0xfe, 0xbd, 0x24, 0x9b, 0x48, 0x4e, 0xf3, 0xe6, 0xf1, 0xde, 0xcc, 0x7b, 0xb0, 0xb3, 0xdd,
	0x45, 0x02, 0xf9, 0xe3, 0x3c, 0xd1, 0x79, 0x0e, 0xbc, 0x7f, 0x0a, 0xbf, 0x8e, 0xe4, 0x48, 0xc3,
	0x0f, 0x75, 0xec, 0x28, 0xb2, 0xc5, 0x1b, 0x24, 0x9f, 0xde, 0x59, 0x43, 0xdf, 0x17, 0x92, 0xa0,
	0x0f, 0xb0, 0x5d, 0x04, 0x99, 0xca, 0x55, 0x99, 0x9c, 0x1e, 0xd1, 0x72, 0xdf, 0xf3, 0x80, 0x1f,
	0x33, 0x6f, 0xfe, 0x15, 0x45, 0x0d, 0x69, 0x34, 0x8b, 0xe3, 0x41, 0x48, 0x67, 0x70, 0xd7, 0x93,
	0x48, 0x73, 0x8e, 0xe6, 0x7b, 0xb3, 0xac, 0x5a, 0xc3, 0x8d, 0xe5, 0x96, 0xb2, 0xeb, 0x5c, 0x95,
	0x1b, 0x33, 0xe1, 0x91, 0x6b, 0x9b, 0xd0, 0x64, 0x9b, 0x5c, 0x95, 0xa9, 0x99, 0xf0, 0xe9, 0x15,
	0xb6, 0x55, 0x23, 0x34, 0x5e, 0xd5, 0x07, 0x48, 0x84, 0x86, 0x76, 0x89, 0x96, 0xe2, 0x2a, 0xe8,
	0xfe, 0x01, 0xd7, 0x9f, 0x8b, 0xab, 0xea, 0x19, 0x76, 0x96, 0x7b, 0xb4, 0xd4, 0x75, 0x78, 0xf6,
	0xce, 0xe2, 0xdc, 0xbe, 0x4a, 0xdf, 0x23, 0xa8, 0xc7, 0xbe, 0xb5, 0xfa, 0xba, 0x9d, 0x8a, 0xbf,
	0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0x49, 0xd7, 0x8f, 0x00, 0x27, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BaseGrpcClient is the client API for BaseGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BaseGrpcClient interface {
	SendRequest(ctx context.Context, in *GrpcRequest, opts ...grpc.CallOption) (*GrpcResponse, error)
}

type baseGrpcClient struct {
	cc *grpc.ClientConn
}

func NewBaseGrpcClient(cc *grpc.ClientConn) BaseGrpcClient {
	return &baseGrpcClient{cc}
}

func (c *baseGrpcClient) SendRequest(ctx context.Context, in *GrpcRequest, opts ...grpc.CallOption) (*GrpcResponse, error) {
	out := new(GrpcResponse)
	err := c.cc.Invoke(ctx, "/BaseGrpc/sendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseGrpcServer is the server API for BaseGrpc service.
type BaseGrpcServer interface {
	SendRequest(context.Context, *GrpcRequest) (*GrpcResponse, error)
}

// UnimplementedBaseGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedBaseGrpcServer struct {
}

func (*UnimplementedBaseGrpcServer) SendRequest(ctx context.Context, req *GrpcRequest) (*GrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}

func RegisterBaseGrpcServer(s *grpc.Server, srv BaseGrpcServer) {
	s.RegisterService(&_BaseGrpc_serviceDesc, srv)
}

func _BaseGrpc_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseGrpcServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BaseGrpc/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseGrpcServer).SendRequest(ctx, req.(*GrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BaseGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BaseGrpc",
	HandlerType: (*BaseGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendRequest",
			Handler:    _BaseGrpc_SendRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster/cluster.proto",
}
