// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/whopper.proto

package api

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

// HubClient is the client API for Hub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubClient interface {
	RunHub(ctx context.Context, in *HubRequest, opts ...grpc.CallOption) (*HubResponse, error)
}

type hubClient struct {
	cc grpc.ClientConnInterface
}

func NewHubClient(cc grpc.ClientConnInterface) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) RunHub(ctx context.Context, in *HubRequest, opts ...grpc.CallOption) (*HubResponse, error) {
	out := new(HubResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.Hub/RunHub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HubServer is the server API for Hub service.
// All implementations must embed UnimplementedHubServer
// for forward compatibility
type HubServer interface {
	RunHub(context.Context, *HubRequest) (*HubResponse, error)
	mustEmbedUnimplementedHubServer()
}

// UnimplementedHubServer must be embedded to have forward compatible implementations.
type UnimplementedHubServer struct {
}

func (UnimplementedHubServer) RunHub(context.Context, *HubRequest) (*HubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunHub not implemented")
}
func (UnimplementedHubServer) mustEmbedUnimplementedHubServer() {}

// UnsafeHubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubServer will
// result in compilation errors.
type UnsafeHubServer interface {
	mustEmbedUnimplementedHubServer()
}

func RegisterHubServer(s grpc.ServiceRegistrar, srv HubServer) {
	s.RegisterService(&Hub_ServiceDesc, srv)
}

func _Hub_RunHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).RunHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.Hub/RunHub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).RunHub(ctx, req.(*HubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hub_ServiceDesc is the grpc.ServiceDesc for Hub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "whopper.api.Hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunHub",
			Handler:    _Hub_RunHub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/whopper.proto",
}

// DiscovererClient is the client API for Discoverer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscovererClient interface {
	Discover(ctx context.Context, in *DiscovererRequest, opts ...grpc.CallOption) (*DiscovererResponse, error)
}

type discovererClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscovererClient(cc grpc.ClientConnInterface) DiscovererClient {
	return &discovererClient{cc}
}

func (c *discovererClient) Discover(ctx context.Context, in *DiscovererRequest, opts ...grpc.CallOption) (*DiscovererResponse, error) {
	out := new(DiscovererResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.Discoverer/Discover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscovererServer is the server API for Discoverer service.
// All implementations must embed UnimplementedDiscovererServer
// for forward compatibility
type DiscovererServer interface {
	Discover(context.Context, *DiscovererRequest) (*DiscovererResponse, error)
	mustEmbedUnimplementedDiscovererServer()
}

// UnimplementedDiscovererServer must be embedded to have forward compatible implementations.
type UnimplementedDiscovererServer struct {
}

func (UnimplementedDiscovererServer) Discover(context.Context, *DiscovererRequest) (*DiscovererResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Discover not implemented")
}
func (UnimplementedDiscovererServer) mustEmbedUnimplementedDiscovererServer() {}

// UnsafeDiscovererServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscovererServer will
// result in compilation errors.
type UnsafeDiscovererServer interface {
	mustEmbedUnimplementedDiscovererServer()
}

func RegisterDiscovererServer(s grpc.ServiceRegistrar, srv DiscovererServer) {
	s.RegisterService(&Discoverer_ServiceDesc, srv)
}

func _Discoverer_Discover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscovererRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscovererServer).Discover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.Discoverer/Discover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscovererServer).Discover(ctx, req.(*DiscovererRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Discoverer_ServiceDesc is the grpc.ServiceDesc for Discoverer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Discoverer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "whopper.api.Discoverer",
	HandlerType: (*DiscovererServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Discover",
			Handler:    _Discoverer_Discover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/whopper.proto",
}

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
	err := c.cc.Invoke(ctx, "/whopper.api.Downloader/Download", in, out, opts...)
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
		FullMethod: "/whopper.api.Downloader/Download",
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
	ServiceName: "whopper.api.Downloader",
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

// ParserClient is the client API for Parser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParserClient interface {
	Parse(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error)
}

type parserClient struct {
	cc grpc.ClientConnInterface
}

func NewParserClient(cc grpc.ClientConnInterface) ParserClient {
	return &parserClient{cc}
}

func (c *parserClient) Parse(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error) {
	out := new(ParserResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.Parser/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParserServer is the server API for Parser service.
// All implementations must embed UnimplementedParserServer
// for forward compatibility
type ParserServer interface {
	Parse(context.Context, *ParserRequest) (*ParserResponse, error)
	mustEmbedUnimplementedParserServer()
}

// UnimplementedParserServer must be embedded to have forward compatible implementations.
type UnimplementedParserServer struct {
}

func (UnimplementedParserServer) Parse(context.Context, *ParserRequest) (*ParserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (UnimplementedParserServer) mustEmbedUnimplementedParserServer() {}

// UnsafeParserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParserServer will
// result in compilation errors.
type UnsafeParserServer interface {
	mustEmbedUnimplementedParserServer()
}

func RegisterParserServer(s grpc.ServiceRegistrar, srv ParserServer) {
	s.RegisterService(&Parser_ServiceDesc, srv)
}

func _Parser_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.Parser/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserServer).Parse(ctx, req.(*ParserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Parser_ServiceDesc is the grpc.ServiceDesc for Parser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Parser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "whopper.api.Parser",
	HandlerType: (*ParserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _Parser_Parse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/whopper.proto",
}

// TranslatorClient is the client API for Translator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslatorClient interface {
	Translate(ctx context.Context, in *TranslatorRequest, opts ...grpc.CallOption) (*TranslatorResponse, error)
}

type translatorClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslatorClient(cc grpc.ClientConnInterface) TranslatorClient {
	return &translatorClient{cc}
}

func (c *translatorClient) Translate(ctx context.Context, in *TranslatorRequest, opts ...grpc.CallOption) (*TranslatorResponse, error) {
	out := new(TranslatorResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.Translator/Translate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslatorServer is the server API for Translator service.
// All implementations must embed UnimplementedTranslatorServer
// for forward compatibility
type TranslatorServer interface {
	Translate(context.Context, *TranslatorRequest) (*TranslatorResponse, error)
	mustEmbedUnimplementedTranslatorServer()
}

// UnimplementedTranslatorServer must be embedded to have forward compatible implementations.
type UnimplementedTranslatorServer struct {
}

func (UnimplementedTranslatorServer) Translate(context.Context, *TranslatorRequest) (*TranslatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (UnimplementedTranslatorServer) mustEmbedUnimplementedTranslatorServer() {}

// UnsafeTranslatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslatorServer will
// result in compilation errors.
type UnsafeTranslatorServer interface {
	mustEmbedUnimplementedTranslatorServer()
}

func RegisterTranslatorServer(s grpc.ServiceRegistrar, srv TranslatorServer) {
	s.RegisterService(&Translator_ServiceDesc, srv)
}

func _Translator_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.Translator/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).Translate(ctx, req.(*TranslatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Translator_ServiceDesc is the grpc.ServiceDesc for Translator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Translator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "whopper.api.Translator",
	HandlerType: (*TranslatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _Translator_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/whopper.proto",
}

// AnalyzerClient is the client API for Analyzer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyzerClient interface {
	Analyze(ctx context.Context, in *AnalyzerRequest, opts ...grpc.CallOption) (*AnalyzerResponse, error)
}

type analyzerClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyzerClient(cc grpc.ClientConnInterface) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) Analyze(ctx context.Context, in *AnalyzerRequest, opts ...grpc.CallOption) (*AnalyzerResponse, error) {
	out := new(AnalyzerResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.Analyzer/Analyze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyzerServer is the server API for Analyzer service.
// All implementations must embed UnimplementedAnalyzerServer
// for forward compatibility
type AnalyzerServer interface {
	Analyze(context.Context, *AnalyzerRequest) (*AnalyzerResponse, error)
	mustEmbedUnimplementedAnalyzerServer()
}

// UnimplementedAnalyzerServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyzerServer struct {
}

func (UnimplementedAnalyzerServer) Analyze(context.Context, *AnalyzerRequest) (*AnalyzerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedAnalyzerServer) mustEmbedUnimplementedAnalyzerServer() {}

// UnsafeAnalyzerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyzerServer will
// result in compilation errors.
type UnsafeAnalyzerServer interface {
	mustEmbedUnimplementedAnalyzerServer()
}

func RegisterAnalyzerServer(s grpc.ServiceRegistrar, srv AnalyzerServer) {
	s.RegisterService(&Analyzer_ServiceDesc, srv)
}

func _Analyzer_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.Analyzer/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Analyze(ctx, req.(*AnalyzerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Analyzer_ServiceDesc is the grpc.ServiceDesc for Analyzer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Analyzer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "whopper.api.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _Analyzer_Analyze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/whopper.proto",
}
