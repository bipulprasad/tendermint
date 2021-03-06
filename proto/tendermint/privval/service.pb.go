// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/privval/service.proto

package privval

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

func init() { proto.RegisterFile("tendermint/privval/service.proto", fileDescriptor_7afe74f9f46d3dc9) }

var fileDescriptor_7afe74f9f46d3dc9 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x28, 0xca, 0x2c, 0x2b, 0x4b, 0xcc, 0xd1, 0x2f,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0xa8,
	0xd0, 0x83, 0xaa, 0x90, 0x92, 0xc3, 0xa2, 0xab, 0xa4, 0xb2, 0x20, 0xb5, 0x18, 0xa2, 0xc7, 0x68,
	0x09, 0x13, 0x97, 0x40, 0x40, 0x51, 0x66, 0x59, 0x58, 0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0x7e,
	0x91, 0x63, 0x80, 0xa7, 0x50, 0x10, 0x17, 0xa7, 0x7b, 0x6a, 0x49, 0x40, 0x69, 0x92, 0x77, 0x6a,
	0xa5, 0x90, 0xa2, 0x1e, 0xa6, 0xb1, 0x7a, 0x10, 0xb9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x29, 0x25, 0x7c, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc2, 0xb9, 0x38, 0x82, 0x33,
	0xd3, 0xf3, 0xc2, 0xf2, 0x4b, 0x52, 0x85, 0x94, 0xb1, 0xa9, 0x87, 0xc9, 0xc2, 0x0c, 0x55, 0xc3,
	0xa5, 0x28, 0x35, 0x05, 0xa2, 0x0c, 0x6a, 0x70, 0x32, 0x17, 0x0f, 0x48, 0x34, 0xa0, 0x28, 0xbf,
	0x20, 0xbf, 0x38, 0x31, 0x47, 0x48, 0x1d, 0x97, 0x3e, 0x98, 0x0a, 0x98, 0x05, 0x5a, 0xb8, 0x2d,
	0x40, 0x28, 0x85, 0x58, 0xe2, 0x14, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0x96, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x48, 0x61, 0x8d,
	0x12, 0xec, 0xf9, 0x25, 0xf9, 0xfa, 0x98, 0xf1, 0x90, 0xc4, 0x06, 0x96, 0x31, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x60, 0x24, 0x48, 0xda, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrivValidatorAPIClient is the client API for PrivValidatorAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrivValidatorAPIClient interface {
	GetPubKey(ctx context.Context, in *PubKeyRequest, opts ...grpc.CallOption) (*PubKeyResponse, error)
	SignVote(ctx context.Context, in *SignVoteRequest, opts ...grpc.CallOption) (*SignedVoteResponse, error)
	SignProposal(ctx context.Context, in *SignProposalRequest, opts ...grpc.CallOption) (*SignedProposalResponse, error)
}

type privValidatorAPIClient struct {
	cc *grpc.ClientConn
}

func NewPrivValidatorAPIClient(cc *grpc.ClientConn) PrivValidatorAPIClient {
	return &privValidatorAPIClient{cc}
}

func (c *privValidatorAPIClient) GetPubKey(ctx context.Context, in *PubKeyRequest, opts ...grpc.CallOption) (*PubKeyResponse, error) {
	out := new(PubKeyResponse)
	err := c.cc.Invoke(ctx, "/tendermint.privval.PrivValidatorAPI/GetPubKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privValidatorAPIClient) SignVote(ctx context.Context, in *SignVoteRequest, opts ...grpc.CallOption) (*SignedVoteResponse, error) {
	out := new(SignedVoteResponse)
	err := c.cc.Invoke(ctx, "/tendermint.privval.PrivValidatorAPI/SignVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privValidatorAPIClient) SignProposal(ctx context.Context, in *SignProposalRequest, opts ...grpc.CallOption) (*SignedProposalResponse, error) {
	out := new(SignedProposalResponse)
	err := c.cc.Invoke(ctx, "/tendermint.privval.PrivValidatorAPI/SignProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivValidatorAPIServer is the server API for PrivValidatorAPI service.
type PrivValidatorAPIServer interface {
	GetPubKey(context.Context, *PubKeyRequest) (*PubKeyResponse, error)
	SignVote(context.Context, *SignVoteRequest) (*SignedVoteResponse, error)
	SignProposal(context.Context, *SignProposalRequest) (*SignedProposalResponse, error)
}

// UnimplementedPrivValidatorAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPrivValidatorAPIServer struct {
}

func (*UnimplementedPrivValidatorAPIServer) GetPubKey(ctx context.Context, req *PubKeyRequest) (*PubKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPubKey not implemented")
}
func (*UnimplementedPrivValidatorAPIServer) SignVote(ctx context.Context, req *SignVoteRequest) (*SignedVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignVote not implemented")
}
func (*UnimplementedPrivValidatorAPIServer) SignProposal(ctx context.Context, req *SignProposalRequest) (*SignedProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignProposal not implemented")
}

func RegisterPrivValidatorAPIServer(s *grpc.Server, srv PrivValidatorAPIServer) {
	s.RegisterService(&_PrivValidatorAPI_serviceDesc, srv)
}

func _PrivValidatorAPI_GetPubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivValidatorAPIServer).GetPubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.privval.PrivValidatorAPI/GetPubKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivValidatorAPIServer).GetPubKey(ctx, req.(*PubKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivValidatorAPI_SignVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivValidatorAPIServer).SignVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.privval.PrivValidatorAPI/SignVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivValidatorAPIServer).SignVote(ctx, req.(*SignVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivValidatorAPI_SignProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivValidatorAPIServer).SignProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.privval.PrivValidatorAPI/SignProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivValidatorAPIServer).SignProposal(ctx, req.(*SignProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivValidatorAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.privval.PrivValidatorAPI",
	HandlerType: (*PrivValidatorAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPubKey",
			Handler:    _PrivValidatorAPI_GetPubKey_Handler,
		},
		{
			MethodName: "SignVote",
			Handler:    _PrivValidatorAPI_SignVote_Handler,
		},
		{
			MethodName: "SignProposal",
			Handler:    _PrivValidatorAPI_SignProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tendermint/privval/service.proto",
}
