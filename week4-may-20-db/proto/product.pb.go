// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/product.proto

/*
Package product is a generated protocol buffer package.

It is generated from these files:
	proto/product.proto

It has these top-level messages:
	AddReq
	AddRes
	UpdateReq
	UpdateRes
	DeleteReq
	DeleteRes
	GetReq
	GetRes
	Product
	Error
*/
package product

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

type AddReq struct {
	Sku   string  `protobuf:"bytes,2,opt,name=sku" json:"sku,omitempty"`
	Name  string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,4,opt,name=price" json:"price,omitempty"`
	Qty   int32   `protobuf:"varint,5,opt,name=qty" json:"qty,omitempty"`
}

func (m *AddReq) Reset()                    { *m = AddReq{} }
func (m *AddReq) String() string            { return proto.CompactTextString(m) }
func (*AddReq) ProtoMessage()               {}
func (*AddReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddReq) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *AddReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddReq) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *AddReq) GetQty() int32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

type AddRes struct {
	Product *Product `protobuf:"bytes,1,opt,name=product" json:"product,omitempty"`
	Error   *Error   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *AddRes) Reset()                    { *m = AddRes{} }
func (m *AddRes) String() string            { return proto.CompactTextString(m) }
func (*AddRes) ProtoMessage()               {}
func (*AddRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddRes) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *AddRes) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type UpdateReq struct {
	Product *Product `protobuf:"bytes,1,opt,name=product" json:"product,omitempty"`
}

func (m *UpdateReq) Reset()                    { *m = UpdateReq{} }
func (m *UpdateReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()               {}
func (*UpdateReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateReq) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type UpdateRes struct {
	Product *Product `protobuf:"bytes,1,opt,name=product" json:"product,omitempty"`
	Error   *Error   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *UpdateRes) Reset()                    { *m = UpdateRes{} }
func (m *UpdateRes) String() string            { return proto.CompactTextString(m) }
func (*UpdateRes) ProtoMessage()               {}
func (*UpdateRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateRes) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *UpdateRes) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type DeleteReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteReq) Reset()                    { *m = DeleteReq{} }
func (m *DeleteReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()               {}
func (*DeleteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteRes struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   *Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *DeleteRes) Reset()                    { *m = DeleteRes{} }
func (m *DeleteRes) String() string            { return proto.CompactTextString(m) }
func (*DeleteRes) ProtoMessage()               {}
func (*DeleteRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DeleteRes) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetRes struct {
	Product *Product `protobuf:"bytes,1,opt,name=product" json:"product,omitempty"`
	Error   *Error   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *GetRes) Reset()                    { *m = GetRes{} }
func (m *GetRes) String() string            { return proto.CompactTextString(m) }
func (*GetRes) ProtoMessage()               {}
func (*GetRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetRes) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *GetRes) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Product struct {
	Id    int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Sku   string  `protobuf:"bytes,2,opt,name=sku" json:"sku,omitempty"`
	Name  string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,4,opt,name=price" json:"price,omitempty"`
	Qty   int32   `protobuf:"varint,5,opt,name=qty" json:"qty,omitempty"`
}

func (m *Product) Reset()                    { *m = Product{} }
func (m *Product) String() string            { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()               {}
func (*Product) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetQty() int32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

type Error struct {
	Code     int32             `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message  string            `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*AddReq)(nil), "AddReq")
	proto.RegisterType((*AddRes)(nil), "AddRes")
	proto.RegisterType((*UpdateReq)(nil), "UpdateReq")
	proto.RegisterType((*UpdateRes)(nil), "UpdateRes")
	proto.RegisterType((*DeleteReq)(nil), "DeleteReq")
	proto.RegisterType((*DeleteRes)(nil), "DeleteRes")
	proto.RegisterType((*GetReq)(nil), "GetReq")
	proto.RegisterType((*GetRes)(nil), "GetRes")
	proto.RegisterType((*Product)(nil), "Product")
	proto.RegisterType((*Error)(nil), "Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProductService service

type ProductServiceClient interface {
	Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRes, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRes, error) {
	out := new(AddRes)
	err := grpc.Invoke(ctx, "/ProductService/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error) {
	out := new(UpdateRes)
	err := grpc.Invoke(ctx, "/ProductService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error) {
	out := new(DeleteRes)
	err := grpc.Invoke(ctx, "/ProductService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error) {
	out := new(GetRes)
	err := grpc.Invoke(ctx, "/ProductService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceServer interface {
	Add(context.Context, *AddReq) (*AddRes, error)
	Update(context.Context, *UpdateReq) (*UpdateRes, error)
	Delete(context.Context, *DeleteReq) (*DeleteRes, error)
	Get(context.Context, *GetReq) (*GetRes, error)
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Add(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ProductService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProductService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/product.proto",
}

func init() { proto.RegisterFile("proto/product.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcb, 0x8a, 0xd4, 0x40,
	0x14, 0x25, 0x9d, 0x49, 0xd2, 0x7d, 0x1b, 0x07, 0x29, 0x07, 0x0d, 0x51, 0x30, 0x64, 0x21, 0xd9,
	0x58, 0x2d, 0xed, 0x46, 0xc6, 0xd5, 0xa8, 0xc3, 0xc0, 0xc0, 0x80, 0x94, 0xe8, 0x56, 0xd2, 0xa9,
	0x6b, 0x0c, 0x3d, 0x99, 0xca, 0x54, 0x55, 0x06, 0xb2, 0x72, 0xe7, 0x97, 0xf8, 0x5d, 0x7e, 0x84,
	0x5f, 0x20, 0xf5, 0xe8, 0x0e, 0x82, 0xa0, 0x8b, 0xde, 0x9d, 0x73, 0x5f, 0xe7, 0xde, 0xe2, 0x14,
	0x3c, 0xe8, 0xa5, 0xd0, 0x62, 0xd5, 0x4b, 0xc1, 0x87, 0x5a, 0x53, 0xcb, 0xb2, 0xe7, 0x4d, 0xab,
	0xbf, 0x0e, 0x1b, 0x5a, 0x8b, 0x6e, 0xd5, 0x88, 0xc6, 0xa6, 0xb5, 0xd8, 0x0c, 0x5f, 0x2c, 0x73,
	0x1d, 0x06, 0xb9, 0xf2, 0xe2, 0x13, 0xc4, 0x67, 0x9c, 0x33, 0xbc, 0x25, 0xf7, 0x21, 0x54, 0xdb,
	0x21, 0x9d, 0xe5, 0x41, 0xb9, 0x60, 0x06, 0x12, 0x02, 0x47, 0x37, 0x55, 0x87, 0x69, 0x68, 0x43,
	0x16, 0x93, 0x13, 0x88, 0x7a, 0xd9, 0xd6, 0x98, 0x1e, 0xe5, 0x41, 0x39, 0x63, 0x8e, 0x98, 0xde,
	0x5b, 0x3d, 0xa6, 0x51, 0x1e, 0x94, 0x11, 0x33, 0xb0, 0xb8, 0xf4, 0x73, 0x15, 0x29, 0x20, 0xf1,
	0x1b, 0xa6, 0x41, 0x1e, 0x94, 0xcb, 0xf5, 0x9c, 0xbe, 0x77, 0x9c, 0xed, 0x12, 0xe4, 0x09, 0x44,
	0x28, 0xa5, 0x90, 0x56, 0x7d, 0xb9, 0x8e, 0xe9, 0xb9, 0x61, 0xcc, 0x05, 0x8b, 0x15, 0x2c, 0x3e,
	0xf6, 0xbc, 0xd2, 0x68, 0xd6, 0xfc, 0x8f, 0x71, 0xc5, 0xd5, 0xd4, 0x70, 0x08, 0xfd, 0xc7, 0xb0,
	0x78, 0x87, 0xd7, 0xe8, 0xf4, 0x8f, 0x61, 0xd6, 0x72, 0x3b, 0x29, 0x62, 0xb3, 0x96, 0x17, 0x6f,
	0xa7, 0xa4, 0x22, 0x29, 0x24, 0x6a, 0xa8, 0x6b, 0x54, 0xca, 0x56, 0xcc, 0xd9, 0x8e, 0xfe, 0x43,
	0x21, 0x85, 0xf8, 0x02, 0xf5, 0xdf, 0xc6, 0x5f, 0xfa, 0xcc, 0x21, 0xee, 0xf8, 0x06, 0x89, 0xef,
	0x20, 0xcf, 0x26, 0x99, 0x37, 0x0f, 0x7f, 0xfd, 0x7c, 0x4a, 0x1a, 0x21, 0xbb, 0xd3, 0xa2, 0x97,
	0x6d, 0x57, 0xc9, 0xf1, 0xf3, 0x16, 0xc7, 0xc2, 0xc8, 0x1f, 0xd8, 0x14, 0x3f, 0x02, 0x88, 0xec,
	0x46, 0x66, 0x4a, 0x2d, 0x38, 0xfa, 0x43, 0x2d, 0x36, 0x8f, 0xd7, 0xa1, 0x52, 0x55, 0x83, 0x5e,
	0x6f, 0x47, 0xc9, 0x0b, 0x98, 0x77, 0xa8, 0x2b, 0x5e, 0xe9, 0x2a, 0x0d, 0xf3, 0xb0, 0x5c, 0xae,
	0x4f, 0xdc, 0x65, 0xf4, 0xca, 0x87, 0xcf, 0x6f, 0xb4, 0x1c, 0xd9, 0xbe, 0x2a, 0x7b, 0x0d, 0xf7,
	0xfe, 0x48, 0x99, 0x65, 0xb6, 0x38, 0x5a, 0xbd, 0x05, 0x33, 0xd0, 0x2c, 0x7d, 0x57, 0x5d, 0x0f,
	0x3b, 0x31, 0x47, 0x4e, 0x67, 0xaf, 0x82, 0xf5, 0xf7, 0x00, 0x8e, 0xfd, 0x43, 0x7d, 0x40, 0x79,
	0x67, 0x6e, 0x79, 0x04, 0xe1, 0x19, 0xe7, 0x24, 0xa1, 0xee, 0xb3, 0x64, 0x1e, 0x28, 0x92, 0x43,
	0xec, 0xac, 0x46, 0x80, 0xee, 0x4d, 0x9a, 0x4d, 0xd8, 0x56, 0x38, 0x83, 0x10, 0xa0, 0x7b, 0x1b,
	0x65, 0x13, 0x56, 0x66, 0xf8, 0x05, 0x6a, 0x92, 0x50, 0xe7, 0x81, 0xcc, 0x03, 0xb5, 0x89, 0xed,
	0x1f, 0x7d, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x7a, 0x5b, 0xe2, 0xe9, 0x03, 0x00, 0x00,
}