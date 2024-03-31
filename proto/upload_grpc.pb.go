// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/upload.proto

package proto

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

// StreamUploadClient is the client API for StreamUpload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamUploadClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (StreamUpload_UploadClient, error)
}

type streamUploadClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamUploadClient(cc grpc.ClientConnInterface) StreamUploadClient {
	return &streamUploadClient{cc}
}

func (c *streamUploadClient) Upload(ctx context.Context, opts ...grpc.CallOption) (StreamUpload_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamUpload_ServiceDesc.Streams[0], "/StreamUpload/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamUploadUploadClient{stream}
	return x, nil
}

type StreamUpload_UploadClient interface {
	Send(*UploadReq) error
	CloseAndRecv() (*UploadResp, error)
	grpc.ClientStream
}

type streamUploadUploadClient struct {
	grpc.ClientStream
}

func (x *streamUploadUploadClient) Send(m *UploadReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamUploadUploadClient) CloseAndRecv() (*UploadResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamUploadServer is the server API for StreamUpload service.
// All implementations must embed UnimplementedStreamUploadServer
// for forward compatibility
type StreamUploadServer interface {
	Upload(StreamUpload_UploadServer) error
	mustEmbedUnimplementedStreamUploadServer()
}

// UnimplementedStreamUploadServer must be embedded to have forward compatible implementations.
type UnimplementedStreamUploadServer struct {
}

func (UnimplementedStreamUploadServer) Upload(StreamUpload_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedStreamUploadServer) mustEmbedUnimplementedStreamUploadServer() {}

// UnsafeStreamUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamUploadServer will
// result in compilation errors.
type UnsafeStreamUploadServer interface {
	mustEmbedUnimplementedStreamUploadServer()
}

func RegisterStreamUploadServer(s grpc.ServiceRegistrar, srv StreamUploadServer) {
	s.RegisterService(&StreamUpload_ServiceDesc, srv)
}

func _StreamUpload_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamUploadServer).Upload(&streamUploadUploadServer{stream})
}

type StreamUpload_UploadServer interface {
	SendAndClose(*UploadResp) error
	Recv() (*UploadReq, error)
	grpc.ServerStream
}

type streamUploadUploadServer struct {
	grpc.ServerStream
}

func (x *streamUploadUploadServer) SendAndClose(m *UploadResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamUploadUploadServer) Recv() (*UploadReq, error) {
	m := new(UploadReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamUpload_ServiceDesc is the grpc.ServiceDesc for StreamUpload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamUpload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StreamUpload",
	HandlerType: (*StreamUploadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _StreamUpload_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/upload.proto",
}
