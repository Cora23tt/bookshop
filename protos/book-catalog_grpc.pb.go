// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protos/book-catalog.proto

package protos

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

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	GetBookById(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Book, error)
	CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookID, error)
	DeleteBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*BookID, error)
	EditBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error)
	CreateGenre(ctx context.Context, in *Genre, opts ...grpc.CallOption) (*Genre, error)
	GetGenres(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GenreResponse, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBookById(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/protos.BookService/GetBookById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookID, error) {
	out := new(BookID)
	err := c.cc.Invoke(ctx, "/protos.BookService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*BookID, error) {
	out := new(BookID)
	err := c.cc.Invoke(ctx, "/protos.BookService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) EditBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/protos.BookService/EditBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) CreateGenre(ctx context.Context, in *Genre, opts ...grpc.CallOption) (*Genre, error) {
	out := new(Genre)
	err := c.cc.Invoke(ctx, "/protos.BookService/CreateGenre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetGenres(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GenreResponse, error) {
	out := new(GenreResponse)
	err := c.cc.Invoke(ctx, "/protos.BookService/GetGenres", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	GetBookById(context.Context, *BookID) (*Book, error)
	CreateBook(context.Context, *Book) (*BookID, error)
	DeleteBook(context.Context, *BookID) (*BookID, error)
	EditBook(context.Context, *Book) (*Book, error)
	CreateGenre(context.Context, *Genre) (*Genre, error)
	GetGenres(context.Context, *Empty) (*GenreResponse, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) GetBookById(context.Context, *BookID) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookById not implemented")
}
func (UnimplementedBookServiceServer) CreateBook(context.Context, *Book) (*BookID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServiceServer) DeleteBook(context.Context, *BookID) (*BookID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookServiceServer) EditBook(context.Context, *Book) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditBook not implemented")
}
func (UnimplementedBookServiceServer) CreateGenre(context.Context, *Genre) (*Genre, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGenre not implemented")
}
func (UnimplementedBookServiceServer) GetGenres(context.Context, *Empty) (*GenreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenres not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_GetBookById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BookService/GetBookById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookById(ctx, req.(*BookID))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BookService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BookService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*BookID))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_EditBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).EditBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BookService/EditBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).EditBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_CreateGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Genre)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BookService/CreateGenre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateGenre(ctx, req.(*Genre))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BookService/GetGenres",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetGenres(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookById",
			Handler:    _BookService_GetBookById_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
		{
			MethodName: "EditBook",
			Handler:    _BookService_EditBook_Handler,
		},
		{
			MethodName: "CreateGenre",
			Handler:    _BookService_CreateGenre_Handler,
		},
		{
			MethodName: "GetGenres",
			Handler:    _BookService_GetGenres_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/book-catalog.proto",
}
