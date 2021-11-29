// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package climate_whopper

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

// DownloaderClient is the client API for Downloader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownloaderClient interface {
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
}

type downloaderClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloaderClient(cc grpc.ClientConnInterface) DownloaderClient {
	return &downloaderClient{cc}
}

func (c *downloaderClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/Downloader/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloaderServer is the server API for Downloader service.
// All implementations must embed UnimplementedDownloaderServer
// for forward compatibility
type DownloaderServer interface {
	Download(context.Context, *DownloadRequest) (*DownloadResponse, error)
	mustEmbedUnimplementedDownloaderServer()
}

// UnimplementedDownloaderServer must be embedded to have forward compatible implementations.
type UnimplementedDownloaderServer struct {
}

func (UnimplementedDownloaderServer) Download(context.Context, *DownloadRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedDownloaderServer) mustEmbedUnimplementedDownloaderServer() {}

// UnsafeDownloaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DownloaderServer will
// result in compilation errors.
type UnsafeDownloaderServer interface {
	mustEmbedUnimplementedDownloaderServer()
}

func RegisterDownloaderServer(s grpc.ServiceRegistrar, srv DownloaderServer) {
	s.RegisterService(&Downloader_ServiceDesc, srv)
}

func _Downloader_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloaderServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Downloader/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloaderServer).Download(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Downloader_ServiceDesc is the grpc.ServiceDesc for Downloader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Downloader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Downloader",
	HandlerType: (*DownloaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Download",
			Handler:    _Downloader_Download_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/whopper.proto",
}
