// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tinyci/ci-agents/grpc/types/ref.proto

package types

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Ref is the encapsulation of a git ref and communicates repository as well as version information.
type Ref struct {
	Id                   int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Repository           *Repository `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	RefName              string      `protobuf:"bytes,3,opt,name=refName,proto3" json:"refName,omitempty"`
	Sha                  string      `protobuf:"bytes,4,opt,name=sha,proto3" json:"sha,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Ref) Reset()         { *m = Ref{} }
func (m *Ref) String() string { return proto.CompactTextString(m) }
func (*Ref) ProtoMessage()    {}
func (*Ref) Descriptor() ([]byte, []int) {
	return fileDescriptor_cacf18990df7113e, []int{0}
}

func (m *Ref) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ref.Unmarshal(m, b)
}
func (m *Ref) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ref.Marshal(b, m, deterministic)
}
func (m *Ref) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ref.Merge(m, src)
}
func (m *Ref) XXX_Size() int {
	return xxx_messageInfo_Ref.Size(m)
}
func (m *Ref) XXX_DiscardUnknown() {
	xxx_messageInfo_Ref.DiscardUnknown(m)
}

var xxx_messageInfo_Ref proto.InternalMessageInfo

func (m *Ref) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Ref) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *Ref) GetRefName() string {
	if m != nil {
		return m.RefName
	}
	return ""
}

func (m *Ref) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func init() {
	proto.RegisterType((*Ref)(nil), "types.Ref")
}

func init() {
	proto.RegisterFile("github.com/tinyci/ci-agents/grpc/types/ref.proto", fileDescriptor_cacf18990df7113e)
}

var fileDescriptor_cacf18990df7113e = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xc9, 0xcc, 0xab, 0x4c, 0xce, 0xd4, 0x4f, 0xce,
	0xd4, 0x4d, 0x4c, 0x4f, 0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0xa9, 0x2c,
	0x48, 0x2d, 0xd6, 0x2f, 0x4a, 0x4d, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x0b,
	0x48, 0x99, 0x13, 0xad, 0xb1, 0x20, 0xbf, 0x38, 0xb3, 0x24, 0xbf, 0xa8, 0x12, 0xa2, 0x5f, 0xa9,
	0x84, 0x8b, 0x39, 0x28, 0x35, 0x4d, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x39, 0x88, 0x29, 0x33, 0x45, 0xc8, 0x90, 0x8b, 0x0b, 0xa1, 0x54, 0x82, 0x49, 0x81, 0x51,
	0x83, 0xdb, 0x48, 0x50, 0x0f, 0x6c, 0x86, 0x5e, 0x10, 0x5c, 0x22, 0x08, 0x49, 0x91, 0x90, 0x04,
	0x17, 0x7b, 0x51, 0x6a, 0x9a, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x8c, 0x2b, 0x24, 0xc0, 0xc5, 0x5c, 0x9c, 0x91, 0x28, 0xc1, 0x02, 0x16, 0x05, 0x31, 0x9d, 0x34,
	0xa2, 0xd4, 0x88, 0x73, 0x70, 0x12, 0x1b, 0xd8, 0x99, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7a, 0x45, 0xc5, 0xac, 0x1a, 0x01, 0x00, 0x00,
}