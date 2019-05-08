// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mgr.proto

package mgr

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

type NewScore struct {
	Score                string   `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewScore) Reset()         { *m = NewScore{} }
func (m *NewScore) String() string { return proto.CompactTextString(m) }
func (*NewScore) ProtoMessage()    {}
func (*NewScore) Descriptor() ([]byte, []int) {
	return fileDescriptor_b589f75408ae08b4, []int{0}
}

func (m *NewScore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewScore.Unmarshal(m, b)
}
func (m *NewScore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewScore.Marshal(b, m, deterministic)
}
func (m *NewScore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewScore.Merge(m, src)
}
func (m *NewScore) XXX_Size() int {
	return xxx_messageInfo_NewScore.Size(m)
}
func (m *NewScore) XXX_DiscardUnknown() {
	xxx_messageInfo_NewScore.DiscardUnknown(m)
}

var xxx_messageInfo_NewScore proto.InternalMessageInfo

func (m *NewScore) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

type UserScore struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	AddScore             int32    `protobuf:"varint,2,opt,name=add_score,json=addScore,proto3" json:"add_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserScore) Reset()         { *m = UserScore{} }
func (m *UserScore) String() string { return proto.CompactTextString(m) }
func (*UserScore) ProtoMessage()    {}
func (*UserScore) Descriptor() ([]byte, []int) {
	return fileDescriptor_b589f75408ae08b4, []int{1}
}

func (m *UserScore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserScore.Unmarshal(m, b)
}
func (m *UserScore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserScore.Marshal(b, m, deterministic)
}
func (m *UserScore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserScore.Merge(m, src)
}
func (m *UserScore) XXX_Size() int {
	return xxx_messageInfo_UserScore.Size(m)
}
func (m *UserScore) XXX_DiscardUnknown() {
	xxx_messageInfo_UserScore.DiscardUnknown(m)
}

var xxx_messageInfo_UserScore proto.InternalMessageInfo

func (m *UserScore) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserScore) GetAddScore() int32 {
	if m != nil {
		return m.AddScore
	}
	return 0
}

func init() {
	proto.RegisterType((*NewScore)(nil), "mgr.NewScore")
	proto.RegisterType((*UserScore)(nil), "mgr.UserScore")
}

func init() { proto.RegisterFile("mgr.proto", fileDescriptor_b589f75408ae08b4) }

var fileDescriptor_b589f75408ae08b4 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x4d, 0x2f, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4d, 0x2f, 0x52, 0x52, 0xe0, 0xe2, 0xf0, 0x4b,
	0x2d, 0x0f, 0x4e, 0xce, 0x2f, 0x4a, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x06, 0x31, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x17, 0x2e, 0xce, 0xd0, 0xe2, 0xd4, 0x22, 0x88, 0x12,
	0x29, 0x2e, 0x8e, 0xd2, 0xe2, 0xd4, 0xa2, 0xbc, 0xc4, 0x5c, 0x98, 0x2a, 0x38, 0x5f, 0x48, 0x9a,
	0x8b, 0x33, 0x31, 0x25, 0x25, 0x1e, 0x62, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x47, 0x62,
	0x4a, 0x0a, 0x58, 0xa3, 0x91, 0x1b, 0x97, 0x00, 0xdc, 0x94, 0xd0, 0x82, 0x94, 0xc4, 0x92, 0xd4,
	0x22, 0x21, 0x23, 0x2e, 0x7e, 0x08, 0x13, 0x61, 0x3e, 0x9f, 0x1e, 0xc8, 0x7d, 0x70, 0xbe, 0x14,
	0x2f, 0x98, 0x0f, 0x73, 0xa1, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xed, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0xf0, 0xec, 0xf1, 0xc8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserScoreUpdaterClient is the client API for UserScoreUpdater service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserScoreUpdaterClient interface {
	UpdateUserScore(ctx context.Context, in *UserScore, opts ...grpc.CallOption) (*NewScore, error)
}

type userScoreUpdaterClient struct {
	cc *grpc.ClientConn
}

func NewUserScoreUpdaterClient(cc *grpc.ClientConn) UserScoreUpdaterClient {
	return &userScoreUpdaterClient{cc}
}

func (c *userScoreUpdaterClient) UpdateUserScore(ctx context.Context, in *UserScore, opts ...grpc.CallOption) (*NewScore, error) {
	out := new(NewScore)
	err := c.cc.Invoke(ctx, "/mgr.UserScoreUpdater/UpdateUserScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserScoreUpdaterServer is the server API for UserScoreUpdater service.
type UserScoreUpdaterServer interface {
	UpdateUserScore(context.Context, *UserScore) (*NewScore, error)
}

func RegisterUserScoreUpdaterServer(s *grpc.Server, srv UserScoreUpdaterServer) {
	s.RegisterService(&_UserScoreUpdater_serviceDesc, srv)
}

func _UserScoreUpdater_UpdateUserScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserScore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserScoreUpdaterServer).UpdateUserScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mgr.UserScoreUpdater/UpdateUserScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserScoreUpdaterServer).UpdateUserScore(ctx, req.(*UserScore))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserScoreUpdater_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mgr.UserScoreUpdater",
	HandlerType: (*UserScoreUpdaterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUserScore",
			Handler:    _UserScoreUpdater_UpdateUserScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mgr.proto",
}
