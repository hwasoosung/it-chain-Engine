// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Envelope
	Empty
	Message
	Block
	Transaction
	PeerTable
*/
package message

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

type Envelope struct {
	// marshalled Message
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// signed Message
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// sender's public key
	Pubkey []byte `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Envelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Envelope) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Message struct {
	Channel []byte `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// Types that are valid to be assigned to Content:
	//	*Message_Block
	//	*Message_Transaction
	//	*Message_PeerTable
	Content isMessage_Content `protobuf_oneof:"content"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isMessage_Content interface {
	isMessage_Content()
}

type Message_Block struct {
	Block *Block `protobuf:"bytes,2,opt,name=block,oneof"`
}
type Message_Transaction struct {
	Transaction *Transaction `protobuf:"bytes,3,opt,name=transaction,oneof"`
}
type Message_PeerTable struct {
	PeerTable *PeerTable `protobuf:"bytes,4,opt,name=peerTable,oneof"`
}

func (*Message_Block) isMessage_Content()       {}
func (*Message_Transaction) isMessage_Content() {}
func (*Message_PeerTable) isMessage_Content()   {}

func (m *Message) GetContent() isMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Message) GetChannel() []byte {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *Message) GetBlock() *Block {
	if x, ok := m.GetContent().(*Message_Block); ok {
		return x.Block
	}
	return nil
}

func (m *Message) GetTransaction() *Transaction {
	if x, ok := m.GetContent().(*Message_Transaction); ok {
		return x.Transaction
	}
	return nil
}

func (m *Message) GetPeerTable() *PeerTable {
	if x, ok := m.GetContent().(*Message_PeerTable); ok {
		return x.PeerTable
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Block)(nil),
		(*Message_Transaction)(nil),
		(*Message_PeerTable)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// content
	switch x := m.Content.(type) {
	case *Message_Block:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case *Message_Transaction:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transaction); err != nil {
			return err
		}
	case *Message_PeerTable:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PeerTable); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Content has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 2: // content.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Block)
		err := b.DecodeMessage(msg)
		m.Content = &Message_Block{msg}
		return true, err
	case 3: // content.transaction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transaction)
		err := b.DecodeMessage(msg)
		m.Content = &Message_Transaction{msg}
		return true, err
	case 4: // content.peerTable
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PeerTable)
		err := b.DecodeMessage(msg)
		m.Content = &Message_PeerTable{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// content
	switch x := m.Content.(type) {
	case *Message_Block:
		s := proto.Size(x.Block)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Transaction:
		s := proto.Size(x.Transaction)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_PeerTable:
		s := proto.Size(x.PeerTable)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Block struct {
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Transaction struct {
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PeerTable struct {
}

func (m *PeerTable) Reset()                    { *m = PeerTable{} }
func (m *PeerTable) String() string            { return proto.CompactTextString(m) }
func (*PeerTable) ProtoMessage()               {}
func (*PeerTable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*Envelope)(nil), "message.Envelope")
	proto.RegisterType((*Empty)(nil), "message.Empty")
	proto.RegisterType((*Message)(nil), "message.Message")
	proto.RegisterType((*Block)(nil), "message.Block")
	proto.RegisterType((*Transaction)(nil), "message.Transaction")
	proto.RegisterType((*PeerTable)(nil), "message.PeerTable")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessageService service

type MessageServiceClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (MessageService_StreamClient, error)
	// Ping is used to probe a remote peer's aliveness
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (MessageService_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MessageService_serviceDesc.Streams[0], c.cc, "/message.MessageService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceStreamClient{stream}
	return x, nil
}

type MessageService_StreamClient interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type messageServiceStreamClient struct {
	grpc.ClientStream
}

func (x *messageServiceStreamClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceStreamClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/message.MessageService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessageService service

type MessageServiceServer interface {
	Stream(MessageService_StreamServer) error
	// Ping is used to probe a remote peer's aliveness
	Ping(context.Context, *Empty) (*Empty, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).Stream(&messageServiceStreamServer{stream})
}

type MessageService_StreamServer interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type messageServiceStreamServer struct {
	grpc.ServerStream
}

func (x *messageServiceStreamServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceStreamServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _MessageService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _MessageService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x13, 0x6d, 0x1b, 0x33, 0xb1, 0x05, 0x07, 0x91, 0x50, 0x3c, 0xc8, 0x1e, 0xa4, 0xa7,
	0x22, 0xd1, 0x83, 0xe7, 0x42, 0xa1, 0x17, 0xa1, 0xa4, 0x3d, 0x79, 0xdb, 0xc4, 0x21, 0x86, 0x26,
	0xbb, 0xcb, 0x66, 0x5b, 0xc8, 0xf7, 0xf3, 0x83, 0x49, 0xb6, 0xf9, 0x53, 0xf4, 0xf8, 0xe6, 0xcd,
	0xe3, 0xb7, 0xf3, 0x16, 0xa6, 0x25, 0x55, 0x15, 0xcf, 0x68, 0xa9, 0xb4, 0x34, 0x12, 0xbd, 0x56,
	0xb2, 0x4f, 0xb8, 0x59, 0x8b, 0x13, 0x15, 0x52, 0x11, 0x86, 0xe0, 0x29, 0x5e, 0x17, 0x92, 0x7f,
	0x85, 0xee, 0x93, 0xbb, 0xb8, 0x8d, 0x3b, 0x89, 0x8f, 0xe0, 0x57, 0x79, 0x26, 0xb8, 0x39, 0x6a,
	0x0a, 0xaf, 0xac, 0x37, 0x0c, 0xf0, 0x01, 0x26, 0xea, 0x98, 0x1c, 0xa8, 0x0e, 0xaf, 0xad, 0xd5,
	0x2a, 0xe6, 0xc1, 0x78, 0x5d, 0x2a, 0x53, 0xb3, 0x1f, 0x17, 0xbc, 0x8f, 0x33, 0xb0, 0x81, 0xa4,
	0xdf, 0x5c, 0x08, 0x2a, 0x3a, 0x48, 0x2b, 0xf1, 0x19, 0xc6, 0x49, 0x21, 0xd3, 0x83, 0x05, 0x04,
	0xd1, 0x6c, 0xd9, 0x3d, 0x79, 0xd5, 0x4c, 0x37, 0x4e, 0x7c, 0xb6, 0xf1, 0x1d, 0x02, 0xa3, 0xb9,
	0xa8, 0x78, 0x6a, 0x72, 0x29, 0x2c, 0x33, 0x88, 0xee, 0xfb, 0xed, 0xfd, 0xe0, 0x6d, 0x9c, 0xf8,
	0x72, 0x15, 0x23, 0xf0, 0x15, 0x91, 0xde, 0xf3, 0xa4, 0xa0, 0x70, 0x64, 0x73, 0xd8, 0xe7, 0xb6,
	0x9d, 0xb3, 0x71, 0xe2, 0x61, 0x6d, 0xe5, 0x83, 0x97, 0x4a, 0x61, 0x48, 0x98, 0xe6, 0x1e, 0xfb,
	0x14, 0x36, 0x85, 0xe0, 0x82, 0xc2, 0x02, 0xf0, 0xfb, 0x70, 0xa4, 0x60, 0xd6, 0x9e, 0xba, 0x23,
	0x7d, 0xca, 0x53, 0xc2, 0x37, 0x98, 0xec, 0x8c, 0x26, 0x5e, 0xe2, 0x5d, 0x0f, 0xeb, 0x3a, 0x9f,
	0xff, 0x1f, 0x31, 0x67, 0xe1, 0xbe, 0xb8, 0xb8, 0x80, 0xd1, 0x36, 0x17, 0x19, 0x0e, 0x35, 0xd8,
	0x2e, 0xe7, 0x7f, 0x34, 0x73, 0x92, 0x89, 0xfd, 0xd2, 0xd7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x97, 0xb9, 0x1d, 0x78, 0xe3, 0x01, 0x00, 0x00,
}
