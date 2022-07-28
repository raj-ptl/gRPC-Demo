// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: proto/average.proto

package pb

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

// GetAverageServiceClient is the client API for GetAverageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetAverageServiceClient interface {
	GetAverage(ctx context.Context, opts ...grpc.CallOption) (GetAverageService_GetAverageClient, error)
}

type getAverageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGetAverageServiceClient(cc grpc.ClientConnInterface) GetAverageServiceClient {
	return &getAverageServiceClient{cc}
}

func (c *getAverageServiceClient) GetAverage(ctx context.Context, opts ...grpc.CallOption) (GetAverageService_GetAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &GetAverageService_ServiceDesc.Streams[0], "/GetAverageService/GetAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &getAverageServiceGetAverageClient{stream}
	return x, nil
}

type GetAverageService_GetAverageClient interface {
	Send(*GetAverageRequest) error
	CloseAndRecv() (*GetAverageResponse, error)
	grpc.ClientStream
}

type getAverageServiceGetAverageClient struct {
	grpc.ClientStream
}

func (x *getAverageServiceGetAverageClient) Send(m *GetAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *getAverageServiceGetAverageClient) CloseAndRecv() (*GetAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GetAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GetAverageServiceServer is the server API for GetAverageService service.
// All implementations must embed UnimplementedGetAverageServiceServer
// for forward compatibility
type GetAverageServiceServer interface {
	GetAverage(GetAverageService_GetAverageServer) error
	mustEmbedUnimplementedGetAverageServiceServer()
}

// UnimplementedGetAverageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGetAverageServiceServer struct {
}

func (UnimplementedGetAverageServiceServer) GetAverage(GetAverageService_GetAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAverage not implemented")
}
func (UnimplementedGetAverageServiceServer) mustEmbedUnimplementedGetAverageServiceServer() {}

// UnsafeGetAverageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetAverageServiceServer will
// result in compilation errors.
type UnsafeGetAverageServiceServer interface {
	mustEmbedUnimplementedGetAverageServiceServer()
}

func RegisterGetAverageServiceServer(s grpc.ServiceRegistrar, srv GetAverageServiceServer) {
	s.RegisterService(&GetAverageService_ServiceDesc, srv)
}

func _GetAverageService_GetAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GetAverageServiceServer).GetAverage(&getAverageServiceGetAverageServer{stream})
}

type GetAverageService_GetAverageServer interface {
	SendAndClose(*GetAverageResponse) error
	Recv() (*GetAverageRequest, error)
	grpc.ServerStream
}

type getAverageServiceGetAverageServer struct {
	grpc.ServerStream
}

func (x *getAverageServiceGetAverageServer) SendAndClose(m *GetAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *getAverageServiceGetAverageServer) Recv() (*GetAverageRequest, error) {
	m := new(GetAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GetAverageService_ServiceDesc is the grpc.ServiceDesc for GetAverageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetAverageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GetAverageService",
	HandlerType: (*GetAverageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAverage",
			Handler:       _GetAverageService_GetAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/average.proto",
}
