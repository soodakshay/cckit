// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chaincode.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	peer "github.com/hyperledger/fabric/protos/peer"
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

type ChaincodeInput struct {
	// Chaincode name
	Chaincode string `protobuf:"bytes,1,opt,name=chaincode,proto3" json:"chaincode,omitempty"`
	// Channel name
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	// Input contains the arguments for invocation.
	Args [][]byte `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// TransientMap contains data (e.g. cryptographic material) that might be used
	// to implement some form of application-level confidentiality. The contents
	// of this field are supposed to always be omitted from the transaction and
	// excluded from the ledger.
	Transient            map[string][]byte `protobuf:"bytes,4,rep,name=transient,proto3" json:"transient,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ChaincodeInput) Reset()         { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()    {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{0}
}

func (m *ChaincodeInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInput.Unmarshal(m, b)
}
func (m *ChaincodeInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInput.Marshal(b, m, deterministic)
}
func (m *ChaincodeInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInput.Merge(m, src)
}
func (m *ChaincodeInput) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInput.Size(m)
}
func (m *ChaincodeInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInput.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInput proto.InternalMessageInfo

func (m *ChaincodeInput) GetChaincode() string {
	if m != nil {
		return m.Chaincode
	}
	return ""
}

func (m *ChaincodeInput) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ChaincodeInput) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ChaincodeInput) GetTransient() map[string][]byte {
	if m != nil {
		return m.Transient
	}
	return nil
}

type ChaincodeLocator struct {
	// Chaincode name
	Chaincode string `protobuf:"bytes,1,opt,name=chaincode,proto3" json:"chaincode,omitempty"`
	// Channel name
	Channel              string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeLocator) Reset()         { *m = ChaincodeLocator{} }
func (m *ChaincodeLocator) String() string { return proto.CompactTextString(m) }
func (*ChaincodeLocator) ProtoMessage()    {}
func (*ChaincodeLocator) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{1}
}

func (m *ChaincodeLocator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeLocator.Unmarshal(m, b)
}
func (m *ChaincodeLocator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeLocator.Marshal(b, m, deterministic)
}
func (m *ChaincodeLocator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeLocator.Merge(m, src)
}
func (m *ChaincodeLocator) XXX_Size() int {
	return xxx_messageInfo_ChaincodeLocator.Size(m)
}
func (m *ChaincodeLocator) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeLocator.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeLocator proto.InternalMessageInfo

func (m *ChaincodeLocator) GetChaincode() string {
	if m != nil {
		return m.Chaincode
	}
	return ""
}

func (m *ChaincodeLocator) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func init() {
	proto.RegisterType((*ChaincodeInput)(nil), "service.ChaincodeInput")
	proto.RegisterMapType((map[string][]byte)(nil), "service.ChaincodeInput.TransientEntry")
	proto.RegisterType((*ChaincodeLocator)(nil), "service.ChaincodeLocator")
}

func init() { proto.RegisterFile("chaincode.proto", fileDescriptor_97136ef4b384cc22) }

var fileDescriptor_97136ef4b384cc22 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x25, 0x4d, 0xff, 0x90, 0xf9, 0x95, 0xfe, 0xca, 0x22, 0x1a, 0x83, 0x87, 0xd2, 0x83, 0xf4,
	0x94, 0x48, 0xbd, 0x48, 0x55, 0x44, 0xb4, 0x87, 0x8a, 0x07, 0x0d, 0xde, 0xcb, 0x76, 0x3b, 0x36,
	0xa1, 0x71, 0x37, 0xec, 0x6e, 0x02, 0xf9, 0x7c, 0x7e, 0x11, 0x3f, 0x8a, 0x34, 0xdd, 0xa4, 0x14,
	0x11, 0xa4, 0xb7, 0x99, 0x7d, 0xf3, 0xde, 0xcc, 0xdb, 0x07, 0xff, 0x59, 0x44, 0x63, 0xce, 0xc4,
	0x12, 0xfd, 0x54, 0x0a, 0x2d, 0x48, 0x47, 0xa1, 0xcc, 0x63, 0x86, 0xde, 0xfd, 0x2a, 0xd6, 0x51,
	0xb6, 0xf0, 0x99, 0xf8, 0x08, 0xa2, 0x22, 0x45, 0x99, 0xe0, 0x72, 0x85, 0x32, 0x78, 0xa7, 0x0b,
	0x19, 0xb3, 0xa0, 0x9c, 0x56, 0x41, 0x8a, 0x28, 0x37, 0x75, 0x2a, 0x14, 0x4d, 0xe6, 0x12, 0x55,
	0x2a, 0xb8, 0x32, 0x5a, 0xde, 0xdd, 0xdf, 0x25, 0xea, 0x33, 0xe6, 0x98, 0x23, 0xd7, 0x5b, 0x81,
	0xe1, 0x97, 0x05, 0xbd, 0x87, 0x0a, 0x99, 0xf1, 0x34, 0xd3, 0xe4, 0x0c, 0x9c, 0x7a, 0xd6, 0xb5,
	0x06, 0xd6, 0xc8, 0x09, 0x77, 0x0f, 0xc4, 0x85, 0x0e, 0x8b, 0x28, 0xe7, 0x98, 0xb8, 0x8d, 0x12,
	0xab, 0x5a, 0x42, 0xa0, 0x49, 0xe5, 0x4a, 0xb9, 0xf6, 0xc0, 0x1e, 0x75, 0xc3, 0xb2, 0x26, 0x8f,
	0xe0, 0x68, 0x49, 0xb9, 0x8a, 0x91, 0x6b, 0xb7, 0x39, 0xb0, 0x47, 0xff, 0xc6, 0xe7, 0xbe, 0xf1,
	0xef, 0xef, 0xef, 0xf5, 0xdf, 0xaa, 0xc1, 0x29, 0xd7, 0xb2, 0x08, 0x77, 0x44, 0xef, 0x06, 0x7a,
	0xfb, 0x20, 0xe9, 0x83, 0xbd, 0xc6, 0xc2, 0x5c, 0xb7, 0x29, 0xc9, 0x11, 0xb4, 0x72, 0x9a, 0x64,
	0x58, 0x5e, 0xd5, 0x0d, 0xb7, 0xcd, 0xa4, 0x71, 0x65, 0x0d, 0x9f, 0xa0, 0x5f, 0x6f, 0x7a, 0x16,
	0x8c, 0x6a, 0x21, 0x0f, 0xf5, 0x38, 0xfe, 0xb4, 0xc0, 0xa9, 0xc5, 0xc8, 0x04, 0x5a, 0xaf, 0x19,
	0xca, 0x82, 0x9c, 0xfc, 0xe2, 0xc9, 0x73, 0xb7, 0xdf, 0xac, 0xfc, 0x17, 0x13, 0x60, 0x68, 0xf2,
	0x23, 0xd7, 0xd0, 0x9e, 0xf1, 0x5c, 0xac, 0xf1, 0x10, 0xf2, 0x2d, 0xb4, 0xa7, 0x9b, 0x10, 0x15,
	0x39, 0xfd, 0x49, 0x36, 0x1e, 0xbd, 0xe3, 0x8a, 0x5e, 0x23, 0x25, 0xe7, 0xc2, 0x5a, 0xb4, 0x4b,
	0xe0, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x63, 0x4d, 0xd9, 0x47, 0x9b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChaincodeClient is the client API for Chaincode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChaincodeClient interface {
	// Query chaincode on home peer. Do NOT send to orderer.
	Query(ctx context.Context, in *ChaincodeInput, opts ...grpc.CallOption) (*peer.ProposalResponse, error)
	// Invoke chaincode on peers, according to endorsement policy and the SEND to orderer
	Invoke(ctx context.Context, in *ChaincodeInput, opts ...grpc.CallOption) (*peer.ProposalResponse, error)
	// Chaincode events stream
	Events(ctx context.Context, in *ChaincodeLocator, opts ...grpc.CallOption) (Chaincode_EventsClient, error)
}

type chaincodeClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeClient(cc *grpc.ClientConn) ChaincodeClient {
	return &chaincodeClient{cc}
}

func (c *chaincodeClient) Query(ctx context.Context, in *ChaincodeInput, opts ...grpc.CallOption) (*peer.ProposalResponse, error) {
	out := new(peer.ProposalResponse)
	err := c.cc.Invoke(ctx, "/service.Chaincode/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeClient) Invoke(ctx context.Context, in *ChaincodeInput, opts ...grpc.CallOption) (*peer.ProposalResponse, error) {
	out := new(peer.ProposalResponse)
	err := c.cc.Invoke(ctx, "/service.Chaincode/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeClient) Events(ctx context.Context, in *ChaincodeLocator, opts ...grpc.CallOption) (Chaincode_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chaincode_serviceDesc.Streams[0], "/service.Chaincode/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chaincode_EventsClient interface {
	Recv() (*peer.ChaincodeEvent, error)
	grpc.ClientStream
}

type chaincodeEventsClient struct {
	grpc.ClientStream
}

func (x *chaincodeEventsClient) Recv() (*peer.ChaincodeEvent, error) {
	m := new(peer.ChaincodeEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChaincodeServer is the server API for Chaincode service.
type ChaincodeServer interface {
	// Query chaincode on home peer. Do NOT send to orderer.
	Query(context.Context, *ChaincodeInput) (*peer.ProposalResponse, error)
	// Invoke chaincode on peers, according to endorsement policy and the SEND to orderer
	Invoke(context.Context, *ChaincodeInput) (*peer.ProposalResponse, error)
	// Chaincode events stream
	Events(*ChaincodeLocator, Chaincode_EventsServer) error
}

func RegisterChaincodeServer(s *grpc.Server, srv ChaincodeServer) {
	s.RegisterService(&_Chaincode_serviceDesc, srv)
}

func _Chaincode_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChaincodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Chaincode/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).Query(ctx, req.(*ChaincodeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chaincode_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChaincodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Chaincode/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).Invoke(ctx, req.(*ChaincodeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chaincode_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChaincodeLocator)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChaincodeServer).Events(m, &chaincodeEventsServer{stream})
}

type Chaincode_EventsServer interface {
	Send(*peer.ChaincodeEvent) error
	grpc.ServerStream
}

type chaincodeEventsServer struct {
	grpc.ServerStream
}

func (x *chaincodeEventsServer) Send(m *peer.ChaincodeEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _Chaincode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Chaincode",
	HandlerType: (*ChaincodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Chaincode_Query_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _Chaincode_Invoke_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _Chaincode_Events_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chaincode.proto",
}
