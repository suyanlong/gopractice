// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	example.proto
	wallet.proto

It has these top-level messages:
	Test
	Wallet
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Test struct {
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	Type  int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Reps  int64  `protobuf:"varint,3,opt,name=reps" json:"reps,omitempty"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto1.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Test) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Test) GetReps() int64 {
	if m != nil {
		return m.Reps
	}
	return 0
}

func init() {
	proto1.RegisterType((*Test)(nil), "proto.Test")
}

func init() { proto1.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x2e, 0x5c,
	0x2c, 0x21, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c, 0xac, 0x39, 0x89, 0x49, 0xa9, 0x39, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x65, 0x41, 0xaa, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x98, 0x0d, 0x12, 0x2b, 0x4a, 0x2d, 0x28, 0x96, 0x60, 0x56,
	0x60, 0xd4, 0x60, 0x0e, 0x02, 0xb3, 0x93, 0xd8, 0xc0, 0x86, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xf7, 0x0f, 0x41, 0x20, 0x64, 0x00, 0x00, 0x00,
}
