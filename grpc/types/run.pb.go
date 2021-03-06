// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tinyci/ci-agents/grpc/types/run.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Run is a single CI run, intended to be sent to a runner.
type Run struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	FinishedAt           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"`
	Status               bool                 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	StatusSet            bool                 `protobuf:"varint,7,opt,name=statusSet,proto3" json:"statusSet,omitempty"`
	Settings             *RunSettings         `protobuf:"bytes,8,opt,name=settings,proto3" json:"settings,omitempty"`
	Task                 *Task                `protobuf:"bytes,9,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Run) Reset()         { *m = Run{} }
func (m *Run) String() string { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()    {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_dadf5638dd3f7587, []int{0}
}

func (m *Run) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Run.Unmarshal(m, b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Run.Marshal(b, m, deterministic)
}
func (m *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(m, src)
}
func (m *Run) XXX_Size() int {
	return xxx_messageInfo_Run.Size(m)
}
func (m *Run) XXX_DiscardUnknown() {
	xxx_messageInfo_Run.DiscardUnknown(m)
}

var xxx_messageInfo_Run proto.InternalMessageInfo

func (m *Run) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Run) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Run) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Run) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *Run) GetFinishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.FinishedAt
	}
	return nil
}

func (m *Run) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Run) GetStatusSet() bool {
	if m != nil {
		return m.StatusSet
	}
	return false
}

func (m *Run) GetSettings() *RunSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Run) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

// RunList is just an array of runs
type RunList struct {
	List                 []*Run   `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunList) Reset()         { *m = RunList{} }
func (m *RunList) String() string { return proto.CompactTextString(m) }
func (*RunList) ProtoMessage()    {}
func (*RunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_dadf5638dd3f7587, []int{1}
}

func (m *RunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunList.Unmarshal(m, b)
}
func (m *RunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunList.Marshal(b, m, deterministic)
}
func (m *RunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunList.Merge(m, src)
}
func (m *RunList) XXX_Size() int {
	return xxx_messageInfo_RunList.Size(m)
}
func (m *RunList) XXX_DiscardUnknown() {
	xxx_messageInfo_RunList.DiscardUnknown(m)
}

var xxx_messageInfo_RunList proto.InternalMessageInfo

func (m *RunList) GetList() []*Run {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Run)(nil), "types.Run")
	proto.RegisterType((*RunList)(nil), "types.RunList")
}

func init() {
	proto.RegisterFile("github.com/tinyci/ci-agents/grpc/types/run.proto", fileDescriptor_dadf5638dd3f7587)
}

var fileDescriptor_dadf5638dd3f7587 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3b, 0x6f, 0xea, 0x30,
	0x14, 0xc7, 0x95, 0x07, 0x2f, 0x23, 0xdd, 0xc1, 0xc3, 0x95, 0x85, 0xae, 0x6e, 0x23, 0x86, 0x2a,
	0x1d, 0xea, 0xb4, 0x74, 0xe9, 0x63, 0x6a, 0xe7, 0x4e, 0x86, 0xa9, 0x4b, 0x65, 0x82, 0x31, 0x16,
	0xc4, 0x89, 0x72, 0x8e, 0x07, 0xbe, 0x6d, 0x3f, 0x4a, 0x85, 0x13, 0x08, 0x1b, 0x6c, 0xf6, 0xf9,
	0xff, 0x7f, 0xe7, 0x49, 0x1e, 0xb4, 0xc1, 0x8d, 0x5b, 0xf2, 0xbc, 0x2c, 0x32, 0x34, 0x76, 0x9f,
	0x9b, 0x2c, 0x37, 0xf7, 0x52, 0x2b, 0x8b, 0x90, 0xe9, 0xba, 0xca, 0x33, 0xdc, 0x57, 0x0a, 0xb2,
	0xda, 0x59, 0x5e, 0xd5, 0x25, 0x96, 0xb4, 0xe7, 0x03, 0x93, 0xb7, 0x33, 0x50, 0x97, 0x3b, 0x69,
	0x75, 0xe6, 0xf5, 0xa5, 0x5b, 0x67, 0x55, 0xc3, 0xa0, 0x29, 0x14, 0xa0, 0x2c, 0xaa, 0xee, 0xd5,
	0xe4, 0x98, 0x3c, 0x5e, 0x59, 0x15, 0x25, 0x6c, 0x5b, 0xe4, 0xe5, 0xfa, 0x46, 0xbf, 0x41, 0x21,
	0x1a, 0xab, 0xa1, 0x41, 0xa7, 0x3f, 0x21, 0x89, 0x84, 0xb3, 0xf4, 0x0f, 0x09, 0xcd, 0x8a, 0x05,
	0x49, 0x90, 0x46, 0x22, 0x34, 0x2b, 0x4a, 0x49, 0x6c, 0x65, 0xa1, 0x58, 0x98, 0x04, 0xe9, 0x48,
	0xf8, 0x37, 0x7d, 0x26, 0xa3, 0xbc, 0x56, 0x12, 0xd5, 0xea, 0x1d, 0x59, 0x94, 0x04, 0xe9, 0x78,
	0x36, 0xe1, 0xba, 0x2c, 0xf5, 0x4e, 0xf1, 0xe3, 0x7c, 0x7c, 0x71, 0x1c, 0x47, 0x74, 0xe6, 0x03,
	0x09, 0x28, 0xeb, 0x86, 0x8c, 0x2f, 0x93, 0x27, 0x33, 0x7d, 0x25, 0x64, 0x6d, 0xac, 0x81, 0x8d,
	0x47, 0x7b, 0x17, 0xd1, 0x33, 0x37, 0xfd, 0x4b, 0xfa, 0x80, 0x12, 0x1d, 0xb0, 0x7e, 0x12, 0xa4,
	0x43, 0xd1, 0xfe, 0xe8, 0x3f, 0xdf, 0x0d, 0x3a, 0x98, 0x2b, 0x64, 0x03, 0x2f, 0x75, 0x01, 0xca,
	0xc9, 0xf0, 0xb8, 0x23, 0x36, 0xf4, 0xf5, 0x28, 0xf7, 0xeb, 0xe3, 0xc2, 0xd9, 0x79, 0xab, 0x88,
	0x93, 0x87, 0xde, 0x90, 0xf8, 0x70, 0x0a, 0x36, 0xf2, 0xde, 0x71, 0xeb, 0x5d, 0x48, 0xd8, 0x0a,
	0x2f, 0x4c, 0xef, 0xc8, 0x40, 0x38, 0xfb, 0x69, 0x00, 0xe9, 0x7f, 0x12, 0xef, 0x0c, 0x20, 0x0b,
	0x92, 0x28, 0x1d, 0xcf, 0x48, 0x97, 0x57, 0xf8, 0xf8, 0x47, 0xfa, 0x75, 0x7b, 0xdd, 0x29, 0x97,
	0x7d, 0x3f, 0xfb, 0xd3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0xbf, 0x58, 0x4e, 0xa4, 0x02,
	0x00, 0x00,
}
