// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authorization.proto

package authorization

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

type AuthCookie struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Secret               string   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthCookie) Reset()         { *m = AuthCookie{} }
func (m *AuthCookie) String() string { return proto.CompactTextString(m) }
func (*AuthCookie) ProtoMessage()    {}
func (*AuthCookie) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dbbe58d1e51a797, []int{0}
}

func (m *AuthCookie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthCookie.Unmarshal(m, b)
}
func (m *AuthCookie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthCookie.Marshal(b, m, deterministic)
}
func (m *AuthCookie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthCookie.Merge(m, src)
}
func (m *AuthCookie) XXX_Size() int {
	return xxx_messageInfo_AuthCookie.Size(m)
}
func (m *AuthCookie) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthCookie.DiscardUnknown(m)
}

var xxx_messageInfo_AuthCookie proto.InternalMessageInfo

func (m *AuthCookie) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *AuthCookie) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type Id struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dbbe58d1e51a797, []int{1}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Username struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Username) Reset()         { *m = Username{} }
func (m *Username) String() string { return proto.CompactTextString(m) }
func (*Username) ProtoMessage()    {}
func (*Username) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dbbe58d1e51a797, []int{2}
}

func (m *Username) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Username.Unmarshal(m, b)
}
func (m *Username) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Username.Marshal(b, m, deterministic)
}
func (m *Username) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Username.Merge(m, src)
}
func (m *Username) XXX_Size() int {
	return xxx_messageInfo_Username.Size(m)
}
func (m *Username) XXX_DiscardUnknown() {
	xxx_messageInfo_Username.DiscardUnknown(m)
}

var xxx_messageInfo_Username proto.InternalMessageInfo

func (m *Username) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthCookie)(nil), "authorization.AuthCookie")
	proto.RegisterType((*Id)(nil), "authorization.Id")
	proto.RegisterType((*Username)(nil), "authorization.Username")
}

func init() { proto.RegisterFile("authorization.proto", fileDescriptor_1dbbe58d1e51a797) }

var fileDescriptor_1dbbe58d1e51a797 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2c, 0x2d, 0xc9,
	0xc8, 0x2f, 0xca, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x54, 0xb2, 0xe0, 0xe2, 0x72, 0x2c, 0x2d, 0xc9, 0x70, 0xce, 0xcf, 0xcf, 0xce,
	0x4c, 0x15, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x53, 0x93, 0x8b, 0x52, 0x4b, 0x24, 0x98, 0xc0, 0xa2,
	0x50, 0x9e, 0x92, 0x2c, 0x17, 0x93, 0x67, 0x8a, 0x90, 0x38, 0x17, 0x7b, 0x69, 0x71, 0x6a, 0x51,
	0x7c, 0x66, 0x0a, 0x58, 0x13, 0x73, 0x10, 0x1b, 0x88, 0xeb, 0x99, 0xa2, 0x24, 0xc7, 0xc5, 0x11,
	0x5a, 0x9c, 0x5a, 0x94, 0x97, 0x98, 0x8b, 0xd5, 0x58, 0xa3, 0xf9, 0x8c, 0x5c, 0xdc, 0x60, 0x9b,
	0x33, 0x52, 0x93, 0xb3, 0x53, 0x8b, 0x84, 0x1c, 0xb9, 0xf8, 0xdd, 0x53, 0x4b, 0x3c, 0x53, 0xdc,
	0x8a, 0xf2, 0x73, 0xa1, 0xae, 0x91, 0xd4, 0x43, 0xf5, 0x00, 0xc2, 0xa1, 0x52, 0x82, 0x68, 0x52,
	0x9e, 0x29, 0x4a, 0x0c, 0x42, 0xbe, 0x5c, 0xa2, 0xee, 0xa9, 0x25, 0x30, 0x5b, 0x89, 0x33, 0x48,
	0x1c, 0x4d, 0x0a, 0xa6, 0x5b, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0x60, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8d, 0xef, 0xe0, 0xe6, 0x47, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthCheckerClient is the client API for AuthChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthCheckerClient interface {
	GetIdFromCookie(ctx context.Context, in *AuthCookie, opts ...grpc.CallOption) (*Id, error)
	GetUsernameFromCookie(ctx context.Context, in *AuthCookie, opts ...grpc.CallOption) (*Username, error)
}

type authCheckerClient struct {
	cc *grpc.ClientConn
}

func NewAuthCheckerClient(cc *grpc.ClientConn) AuthCheckerClient {
	return &authCheckerClient{cc}
}

func (c *authCheckerClient) GetIdFromCookie(ctx context.Context, in *AuthCookie, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/authorization.AuthChecker/GetIdFromCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetUsernameFromCookie(ctx context.Context, in *AuthCookie, opts ...grpc.CallOption) (*Username, error) {
	out := new(Username)
	err := c.cc.Invoke(ctx, "/authorization.AuthChecker/GetUsernameFromCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthCheckerServer is the server API for AuthChecker service.
type AuthCheckerServer interface {
	GetIdFromCookie(context.Context, *AuthCookie) (*Id, error)
	GetUsernameFromCookie(context.Context, *AuthCookie) (*Username, error)
}

func RegisterAuthCheckerServer(s *grpc.Server, srv AuthCheckerServer) {
	s.RegisterService(&_AuthChecker_serviceDesc, srv)
}

func _AuthChecker_GetIdFromCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetIdFromCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthChecker/GetIdFromCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetIdFromCookie(ctx, req.(*AuthCookie))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetUsernameFromCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetUsernameFromCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthChecker/GetUsernameFromCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetUsernameFromCookie(ctx, req.(*AuthCookie))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.AuthChecker",
	HandlerType: (*AuthCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIdFromCookie",
			Handler:    _AuthChecker_GetIdFromCookie_Handler,
		},
		{
			MethodName: "GetUsernameFromCookie",
			Handler:    _AuthChecker_GetUsernameFromCookie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization.proto",
}
