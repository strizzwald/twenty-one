// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateGameRequest struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	GameId               string   `protobuf:"bytes,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGameRequest) Reset()         { *m = CreateGameRequest{} }
func (m *CreateGameRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGameRequest) ProtoMessage()    {}
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *CreateGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGameRequest.Unmarshal(m, b)
}
func (m *CreateGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGameRequest.Marshal(b, m, deterministic)
}
func (m *CreateGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGameRequest.Merge(m, src)
}
func (m *CreateGameRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGameRequest.Size(m)
}
func (m *CreateGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGameRequest proto.InternalMessageInfo

func (m *CreateGameRequest) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *CreateGameRequest) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type CreateGameResponse struct {
	GameId               string   `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGameResponse) Reset()         { *m = CreateGameResponse{} }
func (m *CreateGameResponse) String() string { return proto.CompactTextString(m) }
func (*CreateGameResponse) ProtoMessage()    {}
func (*CreateGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *CreateGameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGameResponse.Unmarshal(m, b)
}
func (m *CreateGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGameResponse.Marshal(b, m, deterministic)
}
func (m *CreateGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGameResponse.Merge(m, src)
}
func (m *CreateGameResponse) XXX_Size() int {
	return xxx_messageInfo_CreateGameResponse.Size(m)
}
func (m *CreateGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGameResponse proto.InternalMessageInfo

func (m *CreateGameResponse) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type StartGameRequest struct {
	GameId               string   `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartGameRequest) Reset()         { *m = StartGameRequest{} }
func (m *StartGameRequest) String() string { return proto.CompactTextString(m) }
func (*StartGameRequest) ProtoMessage()    {}
func (*StartGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}

func (m *StartGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartGameRequest.Unmarshal(m, b)
}
func (m *StartGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartGameRequest.Marshal(b, m, deterministic)
}
func (m *StartGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartGameRequest.Merge(m, src)
}
func (m *StartGameRequest) XXX_Size() int {
	return xxx_messageInfo_StartGameRequest.Size(m)
}
func (m *StartGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartGameRequest proto.InternalMessageInfo

func (m *StartGameRequest) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type StartGameResponse struct {
	GameStarted          bool     `protobuf:"varint,1,opt,name=game_started,json=gameStarted,proto3" json:"game_started,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartGameResponse) Reset()         { *m = StartGameResponse{} }
func (m *StartGameResponse) String() string { return proto.CompactTextString(m) }
func (*StartGameResponse) ProtoMessage()    {}
func (*StartGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3}
}

func (m *StartGameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartGameResponse.Unmarshal(m, b)
}
func (m *StartGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartGameResponse.Marshal(b, m, deterministic)
}
func (m *StartGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartGameResponse.Merge(m, src)
}
func (m *StartGameResponse) XXX_Size() int {
	return xxx_messageInfo_StartGameResponse.Size(m)
}
func (m *StartGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartGameResponse proto.InternalMessageInfo

func (m *StartGameResponse) GetGameStarted() bool {
	if m != nil {
		return m.GameStarted
	}
	return false
}

func init() {
	proto.RegisterType((*CreateGameRequest)(nil), "rpc.CreateGameRequest")
	proto.RegisterType((*CreateGameResponse)(nil), "rpc.CreateGameResponse")
	proto.RegisterType((*StartGameRequest)(nil), "rpc.StartGameRequest")
	proto.RegisterType((*StartGameResponse)(nil), "rpc.StartGameResponse")
}

func init() {
	proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769)
}

var fileDescriptor_38fc58335341d769 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0xf2, 0xe4, 0x12, 0x74,
	0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x75, 0x4f, 0xcc, 0x4d, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x92, 0xe6, 0xe2, 0x2c, 0xc8, 0x49, 0xac, 0x4c, 0x2d, 0x8a, 0xcf, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x80, 0x08, 0x78, 0xa6, 0x08, 0x89, 0x73, 0xb1, 0x83, 0x0c, 0x01,
	0x49, 0x31, 0x81, 0xa5, 0xd8, 0x40, 0x5c, 0xcf, 0x14, 0x25, 0x5d, 0x2e, 0x21, 0x64, 0xa3, 0x8a,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x91, 0x95, 0x33, 0xa2, 0x28, 0xd7, 0xe6, 0x12, 0x08, 0x2e, 0x49,
	0x2c, 0x2a, 0x41, 0xb6, 0x18, 0xa7, 0x62, 0x33, 0x2e, 0x41, 0x24, 0xc5, 0x50, 0xa3, 0x15, 0xb9,
	0x78, 0xc0, 0xaa, 0x8b, 0x41, 0x32, 0xa9, 0x10, 0x2d, 0x1c, 0x41, 0xdc, 0x20, 0xb1, 0x60, 0x88,
	0x90, 0x51, 0x2b, 0x23, 0x17, 0x37, 0x48, 0x4f, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90,
	0x15, 0x17, 0xbb, 0x5f, 0x6a, 0x39, 0x48, 0x44, 0x48, 0x4c, 0xaf, 0xa8, 0x20, 0x59, 0x0f, 0xc3,
	0xf3, 0x52, 0xe2, 0x18, 0xe2, 0x50, 0xeb, 0xac, 0xb8, 0x38, 0xe1, 0x6e, 0x10, 0x12, 0x05, 0xab,
	0x42, 0xf7, 0x80, 0x94, 0x18, 0xba, 0x30, 0x44, 0x6f, 0x12, 0x1b, 0x38, 0xc8, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x90, 0xb2, 0x45, 0x9d, 0x80, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceClient interface {
	NewGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error)
	StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) NewGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error) {
	out := new(CreateGameResponse)
	err := c.cc.Invoke(ctx, "/rpc.GameService/NewGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error) {
	out := new(StartGameResponse)
	err := c.cc.Invoke(ctx, "/rpc.GameService/StartGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
type GameServiceServer interface {
	NewGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error)
	StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error)
}

// UnimplementedGameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (*UnimplementedGameServiceServer) NewGame(ctx context.Context, req *CreateGameRequest) (*CreateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewGame not implemented")
}
func (*UnimplementedGameServiceServer) StartGame(ctx context.Context, req *StartGameRequest) (*StartGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_NewGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).NewGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.GameService/NewGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).NewGame(ctx, req.(*CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.GameService/StartGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).StartGame(ctx, req.(*StartGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGame",
			Handler:    _GameService_NewGame_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _GameService_StartGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}
