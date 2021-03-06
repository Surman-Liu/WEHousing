// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetSession

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Request struct {
	Sessionid            string   `protobuf:"bytes,1,opt,name=sessionid,proto3" json:"sessionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetSession.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetSession.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xb1, 0xae, 0x82, 0x40,
	0x10, 0x45, 0x1f, 0x4f, 0x05, 0x99, 0x72, 0x62, 0x94, 0xa8, 0x05, 0xd2, 0x68, 0xb5, 0x26, 0xfa,
	0x0d, 0xc6, 0xc6, 0x0a, 0x7a, 0x13, 0x94, 0x09, 0x21, 0x11, 0x16, 0x67, 0x56, 0xe3, 0xe7, 0x9b,
	0xec, 0x62, 0xa8, 0xb4, 0x9a, 0xb9, 0x27, 0xb7, 0x38, 0x17, 0x16, 0x2d, 0x6b, 0xa3, 0xb7, 0xf4,
	0xca, 0xeb, 0xf6, 0x46, 0x9f, 0xab, 0x2c, 0xc5, 0x59, 0xa9, 0x55, 0x5d, 0x5d, 0x59, 0x2b, 0xe1,
	0xa7, 0x3a, 0x92, 0xc9, 0x48, 0xa4, 0xd2, 0x4d, 0xb2, 0x86, 0x20, 0xa5, 0xfb, 0x83, 0xc4, 0xe0,
	0x12, 0x42, 0x71, 0xb4, 0x2a, 0x22, 0x2f, 0xf6, 0x36, 0x61, 0xda, 0x83, 0xe4, 0x04, 0xe3, 0x94,
	0xa4, 0xd5, 0x8d, 0x10, 0x4e, 0x60, 0x44, 0xcc, 0x8d, 0xee, 0x5a, 0x2e, 0xe0, 0x14, 0x7c, 0x62,
	0xae, 0xa5, 0x8c, 0xfe, 0x2d, 0xee, 0x12, 0x22, 0x0c, 0x8b, 0xdc, 0xe4, 0xd1, 0xc0, 0x52, 0xfb,
	0xef, 0xce, 0x10, 0x1c, 0x9c, 0x20, 0x66, 0x00, 0xbd, 0x0f, 0xc6, 0xea, 0x8b, 0xa9, 0xea, 0x34,
	0xe7, 0xab, 0x1f, 0x0d, 0xe7, 0x97, 0xfc, 0x5d, 0x7c, 0x3b, 0x7b, 0xff, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x90, 0x0a, 0x67, 0x15, 0x01, 0x00, 0x00,
}
