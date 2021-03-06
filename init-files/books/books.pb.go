// Code generated by protoc-gen-go. DO NOT EDIT.
// source: books/books.proto

package books

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_c4ef05cb7f62cdaf, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Book struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_c4ef05cb7f62cdaf, []int{1}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (dst *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(dst, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type BookList struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookList) Reset()         { *m = BookList{} }
func (m *BookList) String() string { return proto.CompactTextString(m) }
func (*BookList) ProtoMessage()    {}
func (*BookList) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_c4ef05cb7f62cdaf, []int{2}
}
func (m *BookList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookList.Unmarshal(m, b)
}
func (m *BookList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookList.Marshal(b, m, deterministic)
}
func (dst *BookList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookList.Merge(dst, src)
}
func (m *BookList) XXX_Size() int {
	return xxx_messageInfo_BookList.Size(m)
}
func (m *BookList) XXX_DiscardUnknown() {
	xxx_messageInfo_BookList.DiscardUnknown(m)
}

var xxx_messageInfo_BookList proto.InternalMessageInfo

func (m *BookList) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type BookIdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookIdRequest) Reset()         { *m = BookIdRequest{} }
func (m *BookIdRequest) String() string { return proto.CompactTextString(m) }
func (*BookIdRequest) ProtoMessage()    {}
func (*BookIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_c4ef05cb7f62cdaf, []int{3}
}
func (m *BookIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookIdRequest.Unmarshal(m, b)
}
func (m *BookIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookIdRequest.Marshal(b, m, deterministic)
}
func (dst *BookIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookIdRequest.Merge(dst, src)
}
func (m *BookIdRequest) XXX_Size() int {
	return xxx_messageInfo_BookIdRequest.Size(m)
}
func (m *BookIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BookIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BookIdRequest proto.InternalMessageInfo

func (m *BookIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "books.Empty")
	proto.RegisterType((*Book)(nil), "books.Book")
	proto.RegisterType((*BookList)(nil), "books.BookList")
	proto.RegisterType((*BookIdRequest)(nil), "books.BookIdRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error)
	Insert(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Book, error)
	Delete(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Empty, error)
	Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BookService_WatchClient, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/books.BookService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Insert(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/books.BookService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Get(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/books.BookService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Delete(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/books.BookService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BookService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BookService_serviceDesc.Streams[0], "/books.BookService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_WatchClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookServiceWatchClient struct {
	grpc.ClientStream
}

func (x *bookServiceWatchClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	List(context.Context, *Empty) (*BookList, error)
	Insert(context.Context, *Book) (*Empty, error)
	Get(context.Context, *BookIdRequest) (*Book, error)
	Delete(context.Context, *BookIdRequest) (*Empty, error)
	Watch(*Empty, BookService_WatchServer) error
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Insert(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Get(ctx, req.(*BookIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Delete(ctx, req.(*BookIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).Watch(m, &bookServiceWatchServer{stream})
}

type BookService_WatchServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookServiceWatchServer struct {
	grpc.ServerStream
}

func (x *bookServiceWatchServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "books.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BookService_List_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _BookService_Insert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BookService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _BookService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "books/books.proto",
}

func init() { proto.RegisterFile("books/books.proto", fileDescriptor_books_c4ef05cb7f62cdaf) }

var fileDescriptor_books_c4ef05cb7f62cdaf = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xb3, 0x49, 0x37, 0xea, 0xc4, 0x3f, 0x38, 0x14, 0x09, 0xb9, 0x18, 0x17, 0xd4, 0x20,
	0x18, 0xa5, 0x7e, 0x03, 0xa9, 0x48, 0xc1, 0x53, 0x3c, 0x78, 0x6e, 0x9b, 0x81, 0x2e, 0xad, 0x6e,
	0x4d, 0xa6, 0x82, 0x1f, 0xd7, 0x6f, 0x22, 0xbb, 0x09, 0x92, 0x04, 0xbc, 0x2c, 0xbc, 0x9d, 0x37,
	0xbf, 0xf7, 0x60, 0xe0, 0x74, 0x61, 0xcc, 0xba, 0xbe, 0x73, 0x6f, 0xbe, 0xad, 0x0c, 0x1b, 0x94,
	0x4e, 0xa8, 0x3d, 0x90, 0x4f, 0xef, 0x5b, 0xfe, 0x56, 0x53, 0x18, 0x3d, 0x1a, 0xb3, 0xc6, 0x63,
	0xf0, 0x75, 0x19, 0x8b, 0x54, 0x64, 0xb2, 0xf0, 0x75, 0x89, 0x63, 0x90, 0xac, 0x79, 0x43, 0xb1,
	0x9f, 0x8a, 0xec, 0xa0, 0x68, 0x04, 0x9e, 0x41, 0x38, 0xdf, 0xf1, 0xca, 0x54, 0x71, 0xe0, 0xbe,
	0x5b, 0xa5, 0x6e, 0x61, 0xdf, 0x52, 0x5e, 0x74, 0xcd, 0x78, 0x01, 0x4d, 0x46, 0x2c, 0xd2, 0x20,
	0x8b, 0x26, 0x51, 0xde, 0xc4, 0xdb, 0x79, 0xd1, 0xa6, 0x9f, 0xc3, 0x91, 0x95, 0xb3, 0xb2, 0xa0,
	0xcf, 0x1d, 0xd5, 0x3c, 0x4c, 0x9f, 0xfc, 0x08, 0x88, 0xac, 0xe3, 0x95, 0xaa, 0x2f, 0xbd, 0x24,
	0xbc, 0x86, 0x91, 0x63, 0x1f, 0xb6, 0x30, 0xd7, 0x3d, 0x39, 0xe9, 0xa0, 0xed, 0x58, 0x79, 0x78,
	0x09, 0xe1, 0xec, 0xa3, 0xa6, 0x8a, 0xb1, 0x9b, 0x9b, 0xf4, 0xf6, 0x94, 0x87, 0x37, 0x10, 0x3c,
	0x13, 0xe3, 0xb8, 0xe3, 0xf9, 0x2b, 0x93, 0x74, 0x37, 0x95, 0x87, 0x39, 0x84, 0x53, 0xda, 0x10,
	0xd3, 0x3f, 0xf6, 0x21, 0xfb, 0x0a, 0xe4, 0xdb, 0x9c, 0x97, 0xab, 0x41, 0xd9, 0x3e, 0xf5, 0x5e,
	0x2c, 0x42, 0x77, 0x90, 0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x43, 0x65, 0x5e, 0xa5,
	0x01, 0x00, 0x00,
}
