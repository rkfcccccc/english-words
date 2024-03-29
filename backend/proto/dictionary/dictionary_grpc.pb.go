// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: dictionary.proto

package dictionary

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

// DictionaryServiceClient is the client API for DictionaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictionaryServiceClient interface {
	Create(ctx context.Context, in *Word, opts ...grpc.CallOption) (*WordId, error)
	GetById(ctx context.Context, in *WordId, opts ...grpc.CallOption) (*WordEntry, error)
	GetByName(ctx context.Context, in *Word, opts ...grpc.CallOption) (*WordEntry, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type dictionaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDictionaryServiceClient(cc grpc.ClientConnInterface) DictionaryServiceClient {
	return &dictionaryServiceClient{cc}
}

func (c *dictionaryServiceClient) Create(ctx context.Context, in *Word, opts ...grpc.CallOption) (*WordId, error) {
	out := new(WordId)
	err := c.cc.Invoke(ctx, "/dictionary.DictionaryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryServiceClient) GetById(ctx context.Context, in *WordId, opts ...grpc.CallOption) (*WordEntry, error) {
	out := new(WordEntry)
	err := c.cc.Invoke(ctx, "/dictionary.DictionaryService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryServiceClient) GetByName(ctx context.Context, in *Word, opts ...grpc.CallOption) (*WordEntry, error) {
	out := new(WordEntry)
	err := c.cc.Invoke(ctx, "/dictionary.DictionaryService/GetByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/dictionary.DictionaryService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictionaryServiceServer is the server API for DictionaryService service.
// All implementations must embed UnimplementedDictionaryServiceServer
// for forward compatibility
type DictionaryServiceServer interface {
	Create(context.Context, *Word) (*WordId, error)
	GetById(context.Context, *WordId) (*WordEntry, error)
	GetByName(context.Context, *Word) (*WordEntry, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	mustEmbedUnimplementedDictionaryServiceServer()
}

// UnimplementedDictionaryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDictionaryServiceServer struct {
}

func (UnimplementedDictionaryServiceServer) Create(context.Context, *Word) (*WordId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDictionaryServiceServer) GetById(context.Context, *WordId) (*WordEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedDictionaryServiceServer) GetByName(context.Context, *Word) (*WordEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedDictionaryServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedDictionaryServiceServer) mustEmbedUnimplementedDictionaryServiceServer() {}

// UnsafeDictionaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictionaryServiceServer will
// result in compilation errors.
type UnsafeDictionaryServiceServer interface {
	mustEmbedUnimplementedDictionaryServiceServer()
}

func RegisterDictionaryServiceServer(s grpc.ServiceRegistrar, srv DictionaryServiceServer) {
	s.RegisterService(&DictionaryService_ServiceDesc, srv)
}

func _DictionaryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictionary.DictionaryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServiceServer).Create(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictionaryService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WordId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictionary.DictionaryService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServiceServer).GetById(ctx, req.(*WordId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictionaryService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictionary.DictionaryService/GetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServiceServer).GetByName(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictionaryService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictionary.DictionaryService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DictionaryService_ServiceDesc is the grpc.ServiceDesc for DictionaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DictionaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dictionary.DictionaryService",
	HandlerType: (*DictionaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DictionaryService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _DictionaryService_GetById_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _DictionaryService_GetByName_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _DictionaryService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dictionary.proto",
}
