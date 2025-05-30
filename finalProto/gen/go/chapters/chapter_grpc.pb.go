// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: chapters/chapter.proto

package chaptersv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChaptersService_CreateChapter_FullMethodName  = "/chapters.ChaptersService/CreateChapter"
	ChaptersService_GetChapterByID_FullMethodName = "/chapters.ChaptersService/GetChapterByID"
	ChaptersService_ListChapters_FullMethodName   = "/chapters.ChaptersService/ListChapters"
	ChaptersService_UpdateChapter_FullMethodName  = "/chapters.ChaptersService/UpdateChapter"
	ChaptersService_DeleteChapter_FullMethodName  = "/chapters.ChaptersService/DeleteChapter"
	ChaptersService_AddPage_FullMethodName        = "/chapters.ChaptersService/AddPage"
)

// ChaptersServiceClient is the client API for ChaptersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChaptersServiceClient interface {
	CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error)
	GetChapterByID(ctx context.Context, in *GetChapterByIDRequest, opts ...grpc.CallOption) (*Chapter, error)
	ListChapters(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChapterList, error)
	UpdateChapter(ctx context.Context, in *UpdateChapterRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteChapter(ctx context.Context, in *DeleteChapterRequest, opts ...grpc.CallOption) (*Empty, error)
	AddPage(ctx context.Context, in *AddPageRequest, opts ...grpc.CallOption) (*AddPageResponse, error)
}

type chaptersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChaptersServiceClient(cc grpc.ClientConnInterface) ChaptersServiceClient {
	return &chaptersServiceClient{cc}
}

func (c *chaptersServiceClient) CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChapterResponse)
	err := c.cc.Invoke(ctx, ChaptersService_CreateChapter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaptersServiceClient) GetChapterByID(ctx context.Context, in *GetChapterByIDRequest, opts ...grpc.CallOption) (*Chapter, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Chapter)
	err := c.cc.Invoke(ctx, ChaptersService_GetChapterByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaptersServiceClient) ListChapters(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChapterList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChapterList)
	err := c.cc.Invoke(ctx, ChaptersService_ListChapters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaptersServiceClient) UpdateChapter(ctx context.Context, in *UpdateChapterRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChaptersService_UpdateChapter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaptersServiceClient) DeleteChapter(ctx context.Context, in *DeleteChapterRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChaptersService_DeleteChapter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaptersServiceClient) AddPage(ctx context.Context, in *AddPageRequest, opts ...grpc.CallOption) (*AddPageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPageResponse)
	err := c.cc.Invoke(ctx, ChaptersService_AddPage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChaptersServiceServer is the server API for ChaptersService service.
// All implementations must embed UnimplementedChaptersServiceServer
// for forward compatibility.
type ChaptersServiceServer interface {
	CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error)
	GetChapterByID(context.Context, *GetChapterByIDRequest) (*Chapter, error)
	ListChapters(context.Context, *Empty) (*ChapterList, error)
	UpdateChapter(context.Context, *UpdateChapterRequest) (*Empty, error)
	DeleteChapter(context.Context, *DeleteChapterRequest) (*Empty, error)
	AddPage(context.Context, *AddPageRequest) (*AddPageResponse, error)
	mustEmbedUnimplementedChaptersServiceServer()
}

// UnimplementedChaptersServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChaptersServiceServer struct{}

func (UnimplementedChaptersServiceServer) CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChapter not implemented")
}
func (UnimplementedChaptersServiceServer) GetChapterByID(context.Context, *GetChapterByIDRequest) (*Chapter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChapterByID not implemented")
}
func (UnimplementedChaptersServiceServer) ListChapters(context.Context, *Empty) (*ChapterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChapters not implemented")
}
func (UnimplementedChaptersServiceServer) UpdateChapter(context.Context, *UpdateChapterRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChapter not implemented")
}
func (UnimplementedChaptersServiceServer) DeleteChapter(context.Context, *DeleteChapterRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChapter not implemented")
}
func (UnimplementedChaptersServiceServer) AddPage(context.Context, *AddPageRequest) (*AddPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPage not implemented")
}
func (UnimplementedChaptersServiceServer) mustEmbedUnimplementedChaptersServiceServer() {}
func (UnimplementedChaptersServiceServer) testEmbeddedByValue()                         {}

// UnsafeChaptersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChaptersServiceServer will
// result in compilation errors.
type UnsafeChaptersServiceServer interface {
	mustEmbedUnimplementedChaptersServiceServer()
}

func RegisterChaptersServiceServer(s grpc.ServiceRegistrar, srv ChaptersServiceServer) {
	// If the following call pancis, it indicates UnimplementedChaptersServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChaptersService_ServiceDesc, srv)
}

func _ChaptersService_CreateChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaptersServiceServer).CreateChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaptersService_CreateChapter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaptersServiceServer).CreateChapter(ctx, req.(*CreateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaptersService_GetChapterByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChapterByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaptersServiceServer).GetChapterByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaptersService_GetChapterByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaptersServiceServer).GetChapterByID(ctx, req.(*GetChapterByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaptersService_ListChapters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaptersServiceServer).ListChapters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaptersService_ListChapters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaptersServiceServer).ListChapters(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaptersService_UpdateChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaptersServiceServer).UpdateChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaptersService_UpdateChapter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaptersServiceServer).UpdateChapter(ctx, req.(*UpdateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaptersService_DeleteChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaptersServiceServer).DeleteChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaptersService_DeleteChapter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaptersServiceServer).DeleteChapter(ctx, req.(*DeleteChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaptersService_AddPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaptersServiceServer).AddPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaptersService_AddPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaptersServiceServer).AddPage(ctx, req.(*AddPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChaptersService_ServiceDesc is the grpc.ServiceDesc for ChaptersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChaptersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chapters.ChaptersService",
	HandlerType: (*ChaptersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChapter",
			Handler:    _ChaptersService_CreateChapter_Handler,
		},
		{
			MethodName: "GetChapterByID",
			Handler:    _ChaptersService_GetChapterByID_Handler,
		},
		{
			MethodName: "ListChapters",
			Handler:    _ChaptersService_ListChapters_Handler,
		},
		{
			MethodName: "UpdateChapter",
			Handler:    _ChaptersService_UpdateChapter_Handler,
		},
		{
			MethodName: "DeleteChapter",
			Handler:    _ChaptersService_DeleteChapter_Handler,
		},
		{
			MethodName: "AddPage",
			Handler:    _ChaptersService_AddPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chapters/chapter.proto",
}
