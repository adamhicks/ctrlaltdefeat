// Code generated by protoc-gen-go. DO NOT EDIT.
// source: player.proto

package playerpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import reflexpb "github.com/luno/reflex/reflexpb"

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

type GetRoundReq struct {
	RoundId              int64    `protobuf:"varint,1,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
	Player               string   `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoundReq) Reset()         { *m = GetRoundReq{} }
func (m *GetRoundReq) String() string { return proto.CompactTextString(m) }
func (*GetRoundReq) ProtoMessage()    {}
func (*GetRoundReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_player_b38b3b0f861e50a7, []int{0}
}
func (m *GetRoundReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoundReq.Unmarshal(m, b)
}
func (m *GetRoundReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoundReq.Marshal(b, m, deterministic)
}
func (dst *GetRoundReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoundReq.Merge(dst, src)
}
func (m *GetRoundReq) XXX_Size() int {
	return xxx_messageInfo_GetRoundReq.Size(m)
}
func (m *GetRoundReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoundReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoundReq proto.InternalMessageInfo

func (m *GetRoundReq) GetRoundId() int64 {
	if m != nil {
		return m.RoundId
	}
	return 0
}

func (m *GetRoundReq) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

type RoundInfo struct {
	Rank                 int64       `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Parts                []*PartInfo `protobuf:"bytes,2,rep,name=parts,proto3" json:"parts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RoundInfo) Reset()         { *m = RoundInfo{} }
func (m *RoundInfo) String() string { return proto.CompactTextString(m) }
func (*RoundInfo) ProtoMessage()    {}
func (*RoundInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_player_b38b3b0f861e50a7, []int{1}
}
func (m *RoundInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundInfo.Unmarshal(m, b)
}
func (m *RoundInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundInfo.Marshal(b, m, deterministic)
}
func (dst *RoundInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundInfo.Merge(dst, src)
}
func (m *RoundInfo) XXX_Size() int {
	return xxx_messageInfo_RoundInfo.Size(m)
}
func (m *RoundInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoundInfo proto.InternalMessageInfo

func (m *RoundInfo) GetRank() int64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *RoundInfo) GetParts() []*PartInfo {
	if m != nil {
		return m.Parts
	}
	return nil
}

type PartInfo struct {
	Player               string   `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Part                 int64    `protobuf:"varint,2,opt,name=part,proto3" json:"part,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartInfo) Reset()         { *m = PartInfo{} }
func (m *PartInfo) String() string { return proto.CompactTextString(m) }
func (*PartInfo) ProtoMessage()    {}
func (*PartInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_player_b38b3b0f861e50a7, []int{2}
}
func (m *PartInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartInfo.Unmarshal(m, b)
}
func (m *PartInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartInfo.Marshal(b, m, deterministic)
}
func (dst *PartInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartInfo.Merge(dst, src)
}
func (m *PartInfo) XXX_Size() int {
	return xxx_messageInfo_PartInfo.Size(m)
}
func (m *PartInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PartInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PartInfo proto.InternalMessageInfo

func (m *PartInfo) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

func (m *PartInfo) GetPart() int64 {
	if m != nil {
		return m.Part
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRoundReq)(nil), "playerpb.GetRoundReq")
	proto.RegisterType((*RoundInfo)(nil), "playerpb.RoundInfo")
	proto.RegisterType((*PartInfo)(nil), "playerpb.PartInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlayerClient is the client API for Player service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlayerClient interface {
	Stream(ctx context.Context, in *reflexpb.StreamRequest, opts ...grpc.CallOption) (Player_StreamClient, error)
	GetRoundParts(ctx context.Context, in *GetRoundReq, opts ...grpc.CallOption) (*RoundInfo, error)
}

type playerClient struct {
	cc *grpc.ClientConn
}

func NewPlayerClient(cc *grpc.ClientConn) PlayerClient {
	return &playerClient{cc}
}

func (c *playerClient) Stream(ctx context.Context, in *reflexpb.StreamRequest, opts ...grpc.CallOption) (Player_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Player_serviceDesc.Streams[0], "/playerpb.Player/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &playerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Player_StreamClient interface {
	Recv() (*reflexpb.Event, error)
	grpc.ClientStream
}

type playerStreamClient struct {
	grpc.ClientStream
}

func (x *playerStreamClient) Recv() (*reflexpb.Event, error) {
	m := new(reflexpb.Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *playerClient) GetRoundParts(ctx context.Context, in *GetRoundReq, opts ...grpc.CallOption) (*RoundInfo, error) {
	out := new(RoundInfo)
	err := c.cc.Invoke(ctx, "/playerpb.Player/GetRoundParts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServer is the server API for Player service.
type PlayerServer interface {
	Stream(*reflexpb.StreamRequest, Player_StreamServer) error
	GetRoundParts(context.Context, *GetRoundReq) (*RoundInfo, error)
}

func RegisterPlayerServer(s *grpc.Server, srv PlayerServer) {
	s.RegisterService(&_Player_serviceDesc, srv)
}

func _Player_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(reflexpb.StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlayerServer).Stream(m, &playerStreamServer{stream})
}

type Player_StreamServer interface {
	Send(*reflexpb.Event) error
	grpc.ServerStream
}

type playerStreamServer struct {
	grpc.ServerStream
}

func (x *playerStreamServer) Send(m *reflexpb.Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Player_GetRoundParts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoundReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).GetRoundParts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playerpb.Player/GetRoundParts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).GetRoundParts(ctx, req.(*GetRoundReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Player_serviceDesc = grpc.ServiceDesc{
	ServiceName: "playerpb.Player",
	HandlerType: (*PlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoundParts",
			Handler:    _Player_GetRoundParts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Player_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "player.proto",
}

func init() { proto.RegisterFile("player.proto", fileDescriptor_player_b38b3b0f861e50a7) }

var fileDescriptor_player_b38b3b0f861e50a7 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x6c, 0x5a, 0x08, 0xe9, 0x2b, 0x08, 0xe9, 0x21, 0x20, 0x64, 0x8a, 0x32, 0x65, 0x40, 0x2e,
	0x2a, 0x52, 0x37, 0x24, 0x16, 0x84, 0xb2, 0x55, 0xe6, 0x03, 0x90, 0x43, 0x5c, 0xa8, 0x48, 0x6d,
	0xd7, 0x71, 0x10, 0x4c, 0xfc, 0x3a, 0xb2, 0x9d, 0x90, 0x4c, 0xbe, 0x77, 0x7e, 0xa7, 0xbb, 0x77,
	0x70, 0xaa, 0x6a, 0xf6, 0xc3, 0x35, 0x51, 0x5a, 0x1a, 0x89, 0x91, 0x9f, 0x54, 0x99, 0xdc, 0xbe,
	0xef, 0xcc, 0x47, 0x5b, 0x92, 0x37, 0xb9, 0x5f, 0xd6, 0xad, 0x90, 0x4b, 0xcd, 0xb7, 0x35, 0xff,
	0xee, 0x1e, 0x55, 0x76, 0xc0, 0xeb, 0xb2, 0x47, 0x58, 0x3c, 0x73, 0x43, 0x65, 0x2b, 0x2a, 0xca,
	0x0f, 0x78, 0x03, 0x91, 0xb6, 0xf8, 0x75, 0x57, 0xc5, 0x41, 0x1a, 0xe4, 0x33, 0x7a, 0xe2, 0xe6,
	0xa2, 0xc2, 0x2b, 0x08, 0xbd, 0x47, 0x3c, 0x4d, 0x83, 0x7c, 0x4e, 0xbb, 0x29, 0x2b, 0x60, 0xee,
	0xe4, 0x85, 0xd8, 0x4a, 0x44, 0x38, 0xd2, 0x4c, 0x7c, 0x76, 0x5a, 0x87, 0x31, 0x87, 0x63, 0xc5,
	0xb4, 0x69, 0xe2, 0x69, 0x3a, 0xcb, 0x17, 0x2b, 0x24, 0x7d, 0x54, 0xb2, 0x61, 0xda, 0x58, 0x19,
	0xf5, 0x0b, 0xd9, 0x1a, 0xa2, 0x9e, 0x1a, 0xd9, 0x05, 0x63, 0x3b, 0xeb, 0x60, 0x97, 0x5d, 0x88,
	0x19, 0x75, 0x78, 0xf5, 0x0b, 0xe1, 0xc6, 0xff, 0xae, 0x21, 0x7c, 0x31, 0x9a, 0xb3, 0x3d, 0x5e,
	0x93, 0xfe, 0x60, 0xe2, 0x19, 0xca, 0x0f, 0x2d, 0x6f, 0x4c, 0x72, 0x3e, 0x7c, 0x3c, 0x7d, 0x71,
	0x61, 0xb2, 0xc9, 0x5d, 0x80, 0x0f, 0x70, 0xd6, 0xd7, 0x60, 0x13, 0x34, 0x78, 0x39, 0xa4, 0x1c,
	0xf5, 0x93, 0x5c, 0x0c, 0xf4, 0xff, 0xd1, 0xd9, 0xa4, 0x0c, 0x5d, 0x99, 0xf7, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x54, 0x7c, 0x78, 0xba, 0x94, 0x01, 0x00, 0x00,
}
