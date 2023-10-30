// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: questionanswering/v1/questionanswering.proto

package questionansweringv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	QuestionAnsweringService_Answer_FullMethodName = "/questionanswering.v1.QuestionAnsweringService/Answer"
)

// QuestionAnsweringServiceClient is the client API for QuestionAnsweringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionAnsweringServiceClient interface {
	Answer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error)
}

type questionAnsweringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionAnsweringServiceClient(cc grpc.ClientConnInterface) QuestionAnsweringServiceClient {
	return &questionAnsweringServiceClient{cc}
}

func (c *questionAnsweringServiceClient) Answer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error) {
	out := new(AnswerResponse)
	err := c.cc.Invoke(ctx, QuestionAnsweringService_Answer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionAnsweringServiceServer is the server API for QuestionAnsweringService service.
// All implementations must embed UnimplementedQuestionAnsweringServiceServer
// for forward compatibility
type QuestionAnsweringServiceServer interface {
	Answer(context.Context, *AnswerRequest) (*AnswerResponse, error)
	mustEmbedUnimplementedQuestionAnsweringServiceServer()
}

// UnimplementedQuestionAnsweringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionAnsweringServiceServer struct {
}

func (UnimplementedQuestionAnsweringServiceServer) Answer(context.Context, *AnswerRequest) (*AnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Answer not implemented")
}
func (UnimplementedQuestionAnsweringServiceServer) mustEmbedUnimplementedQuestionAnsweringServiceServer() {
}

// UnsafeQuestionAnsweringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionAnsweringServiceServer will
// result in compilation errors.
type UnsafeQuestionAnsweringServiceServer interface {
	mustEmbedUnimplementedQuestionAnsweringServiceServer()
}

func RegisterQuestionAnsweringServiceServer(s grpc.ServiceRegistrar, srv QuestionAnsweringServiceServer) {
	s.RegisterService(&QuestionAnsweringService_ServiceDesc, srv)
}

func _QuestionAnsweringService_Answer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionAnsweringServiceServer).Answer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionAnsweringService_Answer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionAnsweringServiceServer).Answer(ctx, req.(*AnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuestionAnsweringService_ServiceDesc is the grpc.ServiceDesc for QuestionAnsweringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuestionAnsweringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "questionanswering.v1.QuestionAnsweringService",
	HandlerType: (*QuestionAnsweringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Answer",
			Handler:    _QuestionAnsweringService_Answer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "questionanswering/v1/questionanswering.proto",
}
