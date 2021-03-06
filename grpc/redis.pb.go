// Code generated by protoc-gen-go. DO NOT EDIT.
// source: redis.proto

package neptune

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PingReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{1}
}

func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (m *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(m, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DoRegistRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Auth                 string   `protobuf:"bytes,2,opt,name=Auth,proto3" json:"Auth,omitempty"`
	InitialCap           uint32   `protobuf:"varint,3,opt,name=InitialCap,proto3" json:"InitialCap,omitempty"`
	MaxCap               uint32   `protobuf:"varint,4,opt,name=MaxCap,proto3" json:"MaxCap,omitempty"`
	IdleTimeout          int64    `protobuf:"zigzag64,5,opt,name=IdleTimeout,proto3" json:"IdleTimeout,omitempty"`
	DBNum                uint32   `protobuf:"varint,6,opt,name=DBNum,proto3" json:"DBNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoRegistRequest) Reset()         { *m = DoRegistRequest{} }
func (m *DoRegistRequest) String() string { return proto.CompactTextString(m) }
func (*DoRegistRequest) ProtoMessage()    {}
func (*DoRegistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{2}
}

func (m *DoRegistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoRegistRequest.Unmarshal(m, b)
}
func (m *DoRegistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoRegistRequest.Marshal(b, m, deterministic)
}
func (m *DoRegistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoRegistRequest.Merge(m, src)
}
func (m *DoRegistRequest) XXX_Size() int {
	return xxx_messageInfo_DoRegistRequest.Size(m)
}
func (m *DoRegistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoRegistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoRegistRequest proto.InternalMessageInfo

func (m *DoRegistRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DoRegistRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *DoRegistRequest) GetInitialCap() uint32 {
	if m != nil {
		return m.InitialCap
	}
	return 0
}

func (m *DoRegistRequest) GetMaxCap() uint32 {
	if m != nil {
		return m.MaxCap
	}
	return 0
}

func (m *DoRegistRequest) GetIdleTimeout() int64 {
	if m != nil {
		return m.IdleTimeout
	}
	return 0
}

func (m *DoRegistRequest) GetDBNum() uint32 {
	if m != nil {
		return m.DBNum
	}
	return 0
}

type DoRequest struct {
	RegistId             string   `protobuf:"bytes,1,opt,name=RegistId,proto3" json:"RegistId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Params               []byte   `protobuf:"bytes,3,opt,name=Params,proto3" json:"Params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoRequest) Reset()         { *m = DoRequest{} }
func (m *DoRequest) String() string { return proto.CompactTextString(m) }
func (*DoRequest) ProtoMessage()    {}
func (*DoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{3}
}

func (m *DoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoRequest.Unmarshal(m, b)
}
func (m *DoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoRequest.Marshal(b, m, deterministic)
}
func (m *DoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoRequest.Merge(m, src)
}
func (m *DoRequest) XXX_Size() int {
	return xxx_messageInfo_DoRequest.Size(m)
}
func (m *DoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoRequest proto.InternalMessageInfo

func (m *DoRequest) GetRegistId() string {
	if m != nil {
		return m.RegistId
	}
	return ""
}

func (m *DoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DoRequest) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

type DoReleaseRequest struct {
	RegistId             string   `protobuf:"bytes,1,opt,name=RegistId,proto3" json:"RegistId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoReleaseRequest) Reset()         { *m = DoReleaseRequest{} }
func (m *DoReleaseRequest) String() string { return proto.CompactTextString(m) }
func (*DoReleaseRequest) ProtoMessage()    {}
func (*DoReleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{4}
}

func (m *DoReleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReleaseRequest.Unmarshal(m, b)
}
func (m *DoReleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReleaseRequest.Marshal(b, m, deterministic)
}
func (m *DoReleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReleaseRequest.Merge(m, src)
}
func (m *DoReleaseRequest) XXX_Size() int {
	return xxx_messageInfo_DoReleaseRequest.Size(m)
}
func (m *DoReleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoReleaseRequest proto.InternalMessageInfo

func (m *DoReleaseRequest) GetRegistId() string {
	if m != nil {
		return m.RegistId
	}
	return ""
}

type DoResponse struct {
	Response             []byte   `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoResponse) Reset()         { *m = DoResponse{} }
func (m *DoResponse) String() string { return proto.CompactTextString(m) }
func (*DoResponse) ProtoMessage()    {}
func (*DoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{5}
}

func (m *DoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoResponse.Unmarshal(m, b)
}
func (m *DoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoResponse.Marshal(b, m, deterministic)
}
func (m *DoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoResponse.Merge(m, src)
}
func (m *DoResponse) XXX_Size() int {
	return xxx_messageInfo_DoResponse.Size(m)
}
func (m *DoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoResponse proto.InternalMessageInfo

func (m *DoResponse) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "neptune.PingRequest")
	proto.RegisterType((*PingReply)(nil), "neptune.PingReply")
	proto.RegisterType((*DoRegistRequest)(nil), "neptune.DoRegistRequest")
	proto.RegisterType((*DoRequest)(nil), "neptune.DoRequest")
	proto.RegisterType((*DoReleaseRequest)(nil), "neptune.DoReleaseRequest")
	proto.RegisterType((*DoResponse)(nil), "neptune.DoResponse")
}

func init() { proto.RegisterFile("redis.proto", fileDescriptor_d954120a2319ff8f) }

var fileDescriptor_d954120a2319ff8f = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x6e, 0xb4, 0xbf, 0x93, 0x8a, 0x32, 0x16, 0x59, 0x7b, 0x90, 0xb8, 0x20, 0xe4, 0x94, 0x43,
	0x3d, 0x8a, 0x87, 0x6a, 0x11, 0x7a, 0xb0, 0x94, 0xe8, 0x0b, 0xac, 0x64, 0xa8, 0x81, 0x24, 0x1b,
	0xb3, 0x09, 0xd8, 0xa7, 0xf2, 0xd9, 0x7c, 0x03, 0xd9, 0xdd, 0xa4, 0xa4, 0x05, 0xc5, 0xdb, 0x7c,
	0x33, 0xdf, 0xcc, 0x7c, 0x3b, 0xdf, 0x82, 0x5b, 0x50, 0x14, 0xab, 0x20, 0x2f, 0x64, 0x29, 0x71,
	0x90, 0x51, 0x5e, 0x56, 0x19, 0xf1, 0x6b, 0x70, 0xd7, 0x71, 0xb6, 0x09, 0xe9, 0xa3, 0x22, 0x55,
	0x22, 0x42, 0x37, 0x13, 0x29, 0x31, 0xc7, 0x73, 0xfc, 0x51, 0x68, 0x62, 0x7e, 0x03, 0x23, 0x4b,
	0xc9, 0x93, 0x2d, 0x32, 0x18, 0xa4, 0xa4, 0x94, 0xd8, 0x34, 0x9c, 0x06, 0xf2, 0x2f, 0x07, 0x4e,
	0x17, 0x32, 0xa4, 0x4d, 0xac, 0xca, 0x66, 0x1c, 0x83, 0xc1, 0x3c, 0x8a, 0x0a, 0x52, 0xaa, 0x61,
	0xd7, 0x50, 0x2f, 0x9a, 0x57, 0xe5, 0x3b, 0x3b, 0xb2, 0x8b, 0x74, 0x8c, 0x57, 0x00, 0xcb, 0x2c,
	0x2e, 0x63, 0x91, 0x3c, 0x8a, 0x9c, 0x1d, 0x7b, 0x8e, 0x7f, 0x12, 0xb6, 0x32, 0x78, 0x01, 0xfd,
	0x67, 0xf1, 0xa9, 0x6b, 0x5d, 0x53, 0xab, 0x11, 0x7a, 0xe0, 0x2e, 0xa3, 0x84, 0x5e, 0xe3, 0x94,
	0x64, 0x55, 0xb2, 0x9e, 0xe7, 0xf8, 0x18, 0xb6, 0x53, 0x38, 0x81, 0xde, 0xe2, 0x61, 0x55, 0xa5,
	0xac, 0x6f, 0x1a, 0x2d, 0xe0, 0x2f, 0x30, 0xd2, 0x82, 0xad, 0xd4, 0x29, 0x0c, 0xad, 0xf6, 0x65,
	0x54, 0x6b, 0xdd, 0x61, 0x2d, 0x76, 0xa5, 0xaf, 0x52, 0x8b, 0xd5, 0xb1, 0x16, 0xb3, 0x16, 0x85,
	0x48, 0x95, 0x11, 0x3a, 0x0e, 0x6b, 0xc4, 0x03, 0x38, 0xd3, 0x43, 0x13, 0x12, 0x8a, 0xfe, 0x31,
	0x9b, 0xfb, 0x00, 0x9a, 0xaf, 0x72, 0x99, 0x29, 0xb2, 0x4c, 0x1b, 0x1b, 0xe6, 0x38, 0xdc, 0xe1,
	0xd9, 0xb7, 0x03, 0xee, 0xca, 0xda, 0xb6, 0x96, 0x32, 0xc1, 0x19, 0x74, 0xb5, 0x2f, 0x38, 0x09,
	0x6a, 0x33, 0x83, 0x96, 0x93, 0x53, 0x3c, 0xc8, 0xe6, 0xc9, 0x96, 0x77, 0xf0, 0x0e, 0x86, 0x8d,
	0x47, 0xc8, 0x76, 0x8c, 0x03, 0xdb, 0xa6, 0xe7, 0x7b, 0x15, 0xbb, 0x9e, 0x77, 0xf0, 0xde, 0xde,
	0xcb, 0x3c, 0x0d, 0x2f, 0xf7, 0x38, 0xed, 0xe7, 0xfe, 0xd6, 0x3e, 0x83, 0xde, 0x42, 0x3e, 0x55,
	0x19, 0xe2, 0x5e, 0xfd, 0xaf, 0x9e, 0xb7, 0xbe, 0xf9, 0xae, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0x7e, 0xc3, 0xf5, 0xbd, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NeptunePoolClient is the client API for NeptunePool service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NeptunePoolClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	DoRegist(ctx context.Context, in *DoRegistRequest, opts ...grpc.CallOption) (*DoResponse, error)
	DoRelease(ctx context.Context, in *DoReleaseRequest, opts ...grpc.CallOption) (*DoResponse, error)
	DoFun(ctx context.Context, in *DoRequest, opts ...grpc.CallOption) (*DoResponse, error)
}

type neptunePoolClient struct {
	cc *grpc.ClientConn
}

func NewNeptunePoolClient(cc *grpc.ClientConn) NeptunePoolClient {
	return &neptunePoolClient{cc}
}

func (c *neptunePoolClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/neptune.NeptunePool/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neptunePoolClient) DoRegist(ctx context.Context, in *DoRegistRequest, opts ...grpc.CallOption) (*DoResponse, error) {
	out := new(DoResponse)
	err := c.cc.Invoke(ctx, "/neptune.NeptunePool/DoRegist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neptunePoolClient) DoRelease(ctx context.Context, in *DoReleaseRequest, opts ...grpc.CallOption) (*DoResponse, error) {
	out := new(DoResponse)
	err := c.cc.Invoke(ctx, "/neptune.NeptunePool/DoRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neptunePoolClient) DoFun(ctx context.Context, in *DoRequest, opts ...grpc.CallOption) (*DoResponse, error) {
	out := new(DoResponse)
	err := c.cc.Invoke(ctx, "/neptune.NeptunePool/DoFun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NeptunePoolServer is the server API for NeptunePool service.
type NeptunePoolServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	DoRegist(context.Context, *DoRegistRequest) (*DoResponse, error)
	DoRelease(context.Context, *DoReleaseRequest) (*DoResponse, error)
	DoFun(context.Context, *DoRequest) (*DoResponse, error)
}

func RegisterNeptunePoolServer(s *grpc.Server, srv NeptunePoolServer) {
	s.RegisterService(&_NeptunePool_serviceDesc, srv)
}

func _NeptunePool_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeptunePoolServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptune.NeptunePool/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeptunePoolServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NeptunePool_DoRegist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoRegistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeptunePoolServer).DoRegist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptune.NeptunePool/DoRegist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeptunePoolServer).DoRegist(ctx, req.(*DoRegistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NeptunePool_DoRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeptunePoolServer).DoRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptune.NeptunePool/DoRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeptunePoolServer).DoRelease(ctx, req.(*DoReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NeptunePool_DoFun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeptunePoolServer).DoFun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptune.NeptunePool/DoFun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeptunePoolServer).DoFun(ctx, req.(*DoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NeptunePool_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neptune.NeptunePool",
	HandlerType: (*NeptunePoolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _NeptunePool_Ping_Handler,
		},
		{
			MethodName: "DoRegist",
			Handler:    _NeptunePool_DoRegist_Handler,
		},
		{
			MethodName: "DoRelease",
			Handler:    _NeptunePool_DoRelease_Handler,
		},
		{
			MethodName: "DoFun",
			Handler:    _NeptunePool_DoFun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redis.proto",
}
