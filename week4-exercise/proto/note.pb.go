// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/note.proto

/*
Package note is a generated protocol buffer package.

It is generated from these files:
	proto/note.proto

It has these top-level messages:
	Note
	NoteReq
	NoteFindReq
*/
package note

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Note struct {
	Id        int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title     string                     `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Completed string                     `protobuf:"bytes,3,opt,name=completed" json:"completed,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Note) Reset()                    { *m = Note{} }
func (m *Note) String() string            { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()               {}
func (*Note) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Note) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Note) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Note) GetCompleted() string {
	if m != nil {
		return m.Completed
	}
	return ""
}

func (m *Note) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Note) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type NoteReq struct {
	Title     string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Completed bool   `protobuf:"varint,2,opt,name=completed" json:"completed,omitempty"`
}

func (m *NoteReq) Reset()                    { *m = NoteReq{} }
func (m *NoteReq) String() string            { return proto.CompactTextString(m) }
func (*NoteReq) ProtoMessage()               {}
func (*NoteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NoteReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NoteReq) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type NoteFindReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *NoteFindReq) Reset()                    { *m = NoteFindReq{} }
func (m *NoteFindReq) String() string            { return proto.CompactTextString(m) }
func (*NoteFindReq) ProtoMessage()               {}
func (*NoteFindReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NoteFindReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Note)(nil), "note.Note")
	proto.RegisterType((*NoteReq)(nil), "note.NoteReq")
	proto.RegisterType((*NoteFindReq)(nil), "note.NoteFindReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NoteService service

type NoteServiceClient interface {
	Create(ctx context.Context, in *NoteReq, opts ...grpc.CallOption) (*Note, error)
	Find(ctx context.Context, in *NoteFindReq, opts ...grpc.CallOption) (*Note, error)
}

type noteServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoteServiceClient(cc *grpc.ClientConn) NoteServiceClient {
	return &noteServiceClient{cc}
}

func (c *noteServiceClient) Create(ctx context.Context, in *NoteReq, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := grpc.Invoke(ctx, "/note.NoteService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) Find(ctx context.Context, in *NoteFindReq, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := grpc.Invoke(ctx, "/note.NoteService/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NoteService service

type NoteServiceServer interface {
	Create(context.Context, *NoteReq) (*Note, error)
	Find(context.Context, *NoteFindReq) (*Note, error)
}

func RegisterNoteServiceServer(s *grpc.Server, srv NoteServiceServer) {
	s.RegisterService(&_NoteService_serviceDesc, srv)
}

func _NoteService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Create(ctx, req.(*NoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteFindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Find(ctx, req.(*NoteFindReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "note.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NoteService_Create_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _NoteService_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/note.proto",
}

func init() { proto.RegisterFile("proto/note.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xc7, 0x7f, 0x9b, 0x5f, 0x5a, 0xed, 0x94, 0xfa, 0x67, 0x51, 0x08, 0xa1, 0x62, 0x09, 0x08,
	0x3d, 0x65, 0xb1, 0x9e, 0x04, 0x3d, 0x14, 0xc1, 0xa3, 0x87, 0xe8, 0x5d, 0xb6, 0xd9, 0xb1, 0x2c,
	0x24, 0xd9, 0x98, 0x4e, 0x7b, 0x11, 0x2f, 0xbe, 0x82, 0x2f, 0xe4, 0x3b, 0xf8, 0x0a, 0x3e, 0x88,
	0xec, 0x26, 0x21, 0x2a, 0x88, 0xb7, 0x9d, 0x99, 0xef, 0x77, 0xe6, 0x33, 0x3b, 0xb0, 0x57, 0x56,
	0x86, 0x8c, 0x28, 0x0c, 0x61, 0xec, 0x9e, 0xdc, 0xb7, 0xef, 0xf0, 0x78, 0x69, 0xcc, 0x32, 0x43,
	0xe1, 0x72, 0x8b, 0xf5, 0x83, 0x20, 0x9d, 0xe3, 0x8a, 0x64, 0x5e, 0xd6, 0xb2, 0x70, 0xdc, 0x08,
	0x64, 0xa9, 0x85, 0x2c, 0x0a, 0x43, 0x92, 0xb4, 0x29, 0x56, 0x75, 0x35, 0x7a, 0x63, 0xe0, 0xdf,
	0x18, 0x42, 0xbe, 0x03, 0x9e, 0x56, 0x01, 0x9b, 0xb0, 0x69, 0x2f, 0xf1, 0xb4, 0xe2, 0x07, 0xd0,
	0x23, 0x4d, 0x19, 0x06, 0xde, 0x84, 0x4d, 0x07, 0x49, 0x1d, 0xf0, 0x31, 0x0c, 0x52, 0x93, 0x97,
	0x19, 0x12, 0xaa, 0xe0, 0xbf, 0xab, 0x74, 0x09, 0x7e, 0x0e, 0x90, 0x56, 0x28, 0x09, 0xd5, 0xbd,
	0xa4, 0xc0, 0x9f, 0xb0, 0xe9, 0x70, 0x16, 0xc6, 0xf5, 0xfc, 0xb8, 0x05, 0x8c, 0xef, 0x5a, 0xc0,
	0x64, 0xd0, 0xa8, 0xe7, 0x64, 0xad, 0xeb, 0x52, 0xb5, 0xd6, 0xde, 0xdf, 0xd6, 0x46, 0x3d, 0xa7,
	0xe8, 0x12, 0xb6, 0xec, 0x06, 0x09, 0x3e, 0x76, 0xd0, 0xec, 0x57, 0x68, 0xbb, 0xce, 0xf6, 0x17,
	0xe8, 0xe8, 0x08, 0x86, 0xd6, 0x7e, 0xad, 0x0b, 0x65, 0x5b, 0xfc, 0xf8, 0x87, 0x59, 0x55, 0x97,
	0x6f, 0xb1, 0xda, 0xe8, 0x14, 0xf9, 0x09, 0xf4, 0xaf, 0x1c, 0x34, 0x1f, 0xc5, 0xee, 0x16, 0xcd,
	0xe8, 0x10, 0xba, 0x30, 0xfa, 0xc7, 0x2f, 0xc0, 0xb7, 0x0d, 0xf9, 0x7e, 0x97, 0x6d, 0x06, 0x7c,
	0x13, 0x1e, 0xbe, 0xbc, 0x7f, 0xbc, 0x7a, 0xbb, 0x7c, 0x24, 0x36, 0xa7, 0xee, 0xb4, 0xe2, 0x49,
	0xab, 0xe7, 0x45, 0xdf, 0x2d, 0x7c, 0xf6, 0x19, 0x00, 0x00, 0xff, 0xff, 0x40, 0xfd, 0x87, 0x98,
	0xf4, 0x01, 0x00, 0x00,
}
