// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts.proto

package accounts

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetAccountRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountRequest) Reset()         { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()    {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}

func (m *GetAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountRequest.Unmarshal(m, b)
}
func (m *GetAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountRequest.Merge(m, src)
}
func (m *GetAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountRequest.Size(m)
}
func (m *GetAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountRequest proto.InternalMessageInfo

func (m *GetAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Account struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAccountRequest)(nil), "accounts.GetAccountRequest")
	proto.RegisterType((*Account)(nil), "accounts.Account")
}

func init() { proto.RegisterFile("accounts.proto", fileDescriptor_e1e7723af4c007b7) }

var fileDescriptor_e1e7723af4c007b7 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x64,
	0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b,
	0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xa0, 0xea, 0x94, 0xd4, 0xb9, 0x04, 0xdd, 0x53, 0x4b, 0x1c, 0x21,
	0x8a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x59, 0x2e, 0x76, 0xa8, 0x2a, 0x6c,
	0xd2, 0x46, 0xe9, 0x5c, 0x1c, 0x50, 0xe9, 0x62, 0xa1, 0x68, 0x2e, 0x2e, 0x84, 0x99, 0x42, 0xd2,
	0x7a, 0x70, 0xa7, 0x61, 0xd8, 0x24, 0x25, 0x88, 0x90, 0x84, 0xca, 0x28, 0xc9, 0x36, 0x5d, 0x7e,
	0x32, 0x99, 0x49, 0x5c, 0x48, 0x54, 0xbf, 0xcc, 0x40, 0xbf, 0x1a, 0x64, 0xb8, 0x2d, 0x4c, 0x8d,
	0xbe, 0x56, 0x6d, 0x12, 0x1b, 0xd8, 0xdd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x21,
	0x70, 0x63, 0xf1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsClient interface {
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error)
}

type accountsClient struct {
	cc *grpc.ClientConn
}

func NewAccountsClient(cc *grpc.ClientConn) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/accounts.Accounts/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
type AccountsServer interface {
	GetAccount(context.Context, *GetAccountRequest) (*Account, error)
}

// UnimplementedAccountsServer can be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (*UnimplementedAccountsServer) GetAccount(ctx context.Context, req *GetAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounts/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _Accounts_GetAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}