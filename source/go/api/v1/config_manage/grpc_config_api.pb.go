// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_config_api.proto

package config_manage // import "github.com/polarismesh/specification/source/go/api/v1/config_manage"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PolarisConfigGRPCClient is the client API for PolarisConfigGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolarisConfigGRPCClient interface {
	// 拉取配置
	GetConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 创建配置
	CreateConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 更新配置
	UpdateConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 发布配置
	PublishConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 订阅配置变更
	WatchConfigFiles(ctx context.Context, in *ClientWatchConfigFileRequest, opts ...grpc.CallOption) (*ConfigClientResponse, error)
}

type polarisConfigGRPCClient struct {
	cc *grpc.ClientConn
}

func NewPolarisConfigGRPCClient(cc *grpc.ClientConn) PolarisConfigGRPCClient {
	return &polarisConfigGRPCClient{cc}
}

func (c *polarisConfigGRPCClient) GetConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/GetConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) CreateConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/CreateConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) UpdateConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/UpdateConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) PublishConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/PublishConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) WatchConfigFiles(ctx context.Context, in *ClientWatchConfigFileRequest, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/WatchConfigFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolarisConfigGRPCServer is the server API for PolarisConfigGRPC service.
type PolarisConfigGRPCServer interface {
	// 拉取配置
	GetConfigFile(context.Context, *ClientConfigFileInfo) (*ConfigClientResponse, error)
	// 创建配置
	CreateConfigFile(context.Context, *ClientConfigFileInfo) (*ConfigClientResponse, error)
	// 更新配置
	UpdateConfigFile(context.Context, *ClientConfigFileInfo) (*ConfigClientResponse, error)
	// 发布配置
	PublishConfigFile(context.Context, *ClientConfigFileInfo) (*ConfigClientResponse, error)
	// 订阅配置变更
	WatchConfigFiles(context.Context, *ClientWatchConfigFileRequest) (*ConfigClientResponse, error)
}

func RegisterPolarisConfigGRPCServer(s *grpc.Server, srv PolarisConfigGRPCServer) {
	s.RegisterService(&_PolarisConfigGRPC_serviceDesc, srv)
}

func _PolarisConfigGRPC_GetConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConfigFileInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).GetConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/GetConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).GetConfigFile(ctx, req.(*ClientConfigFileInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_CreateConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConfigFileInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).CreateConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/CreateConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).CreateConfigFile(ctx, req.(*ClientConfigFileInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_UpdateConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConfigFileInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).UpdateConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/UpdateConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).UpdateConfigFile(ctx, req.(*ClientConfigFileInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_PublishConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConfigFileInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).PublishConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/PublishConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).PublishConfigFile(ctx, req.(*ClientConfigFileInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_WatchConfigFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientWatchConfigFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).WatchConfigFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/WatchConfigFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).WatchConfigFiles(ctx, req.(*ClientWatchConfigFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolarisConfigGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PolarisConfigGRPC",
	HandlerType: (*PolarisConfigGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfigFile",
			Handler:    _PolarisConfigGRPC_GetConfigFile_Handler,
		},
		{
			MethodName: "CreateConfigFile",
			Handler:    _PolarisConfigGRPC_CreateConfigFile_Handler,
		},
		{
			MethodName: "UpdateConfigFile",
			Handler:    _PolarisConfigGRPC_UpdateConfigFile_Handler,
		},
		{
			MethodName: "PublishConfigFile",
			Handler:    _PolarisConfigGRPC_PublishConfigFile_Handler,
		},
		{
			MethodName: "WatchConfigFiles",
			Handler:    _PolarisConfigGRPC_WatchConfigFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_config_api.proto",
}

func init() {
	proto.RegisterFile("grpc_config_api.proto", fileDescriptor_grpc_config_api_1a375b8f68f49d73)
}

var fileDescriptor_grpc_config_api_1a375b8f68f49d73 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4a, 0xc4, 0x30,
	0x14, 0x87, 0xfd, 0x03, 0x2e, 0x02, 0xc2, 0xb4, 0x20, 0x0c, 0x5d, 0x89, 0x07, 0x48, 0xa8, 0x82,
	0x07, 0x98, 0xa2, 0xe3, 0xec, 0x4a, 0x45, 0x04, 0x37, 0x25, 0x8d, 0xaf, 0xed, 0x83, 0x36, 0x2f,
	0x26, 0x69, 0x2f, 0xe1, 0x01, 0xbc, 0xae, 0xd8, 0x14, 0xac, 0x0a, 0x6e, 0x66, 0xb6, 0xdf, 0x2f,
	0x7c, 0x59, 0x7c, 0x8f, 0x5d, 0x34, 0xd6, 0xa8, 0x52, 0x91, 0xae, 0xb1, 0x29, 0xa5, 0x41, 0x6e,
	0x2c, 0x79, 0x8a, 0x4f, 0xc6, 0x34, 0x89, 0x66, 0x5a, 0x63, 0x07, 0x01, 0x27, 0xc9, 0x02, 0x95,
	0x16, 0x9c, 0x21, 0xed, 0xe6, 0xed, 0xfa, 0xfd, 0x94, 0x45, 0x39, 0x75, 0xd2, 0xa2, 0xcb, 0xa6,
	0x57, 0xdb, 0x22, 0xcf, 0xe2, 0x3b, 0x76, 0xbe, 0x05, 0x1f, 0xc0, 0x3d, 0x76, 0x10, 0xaf, 0xf9,
	0x98, 0xf2, 0xac, 0x43, 0xd0, 0x0b, 0xba, 0xd3, 0x35, 0x25, 0x61, 0x99, 0x58, 0xd8, 0x8b, 0xf9,
	0x83, 0xab, 0xa3, 0xf8, 0x81, 0xad, 0x32, 0x0b, 0xd2, 0xc3, 0x21, 0x4c, 0x4f, 0xe6, 0xf5, 0x10,
	0xa6, 0x1d, 0x8b, 0xf2, 0xa1, 0xea, 0xd0, 0xb5, 0x7b, 0xab, 0x72, 0xb6, 0x7a, 0x96, 0x5e, 0x2d,
	0x44, 0x2e, 0xbe, 0xfc, 0x36, 0xfd, 0xda, 0x0a, 0x78, 0x1b, 0xc0, 0xf9, 0xff, 0x8c, 0x9b, 0x8f,
	0x63, 0x76, 0xab, 0xa8, 0xe7, 0x1e, 0xb4, 0x02, 0xed, 0xb9, 0x09, 0x65, 0xb8, 0x33, 0xa0, 0xb0,
	0x46, 0x25, 0x3d, 0x92, 0xe6, 0x5f, 0xc1, 0xc7, 0x94, 0x87, 0xaa, 0xbc, 0x97, 0x5a, 0x36, 0xb0,
	0x59, 0xff, 0xa9, 0xf8, 0x08, 0x76, 0x44, 0x05, 0x2f, 0x59, 0x83, 0xbe, 0x1d, 0x2a, 0xae, 0xa8,
	0x17, 0xb3, 0xb0, 0x07, 0xd7, 0x8a, 0x1f, 0x52, 0xe1, 0x68, 0xb0, 0x0a, 0x44, 0x43, 0x42, 0x1a,
	0x14, 0x63, 0x2a, 0xe6, 0xa3, 0x09, 0xfa, 0xea, 0x6c, 0x3a, 0x97, 0x9b, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd0, 0xb6, 0x70, 0x3c, 0x7a, 0x02, 0x00, 0x00,
}
