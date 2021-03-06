// Code generated by protoc-gen-go.
// source: go-micro/server/debug/proto/debug.proto
// DO NOT EDIT!

/*
Package debug is a generated protocol buffer package.

It is generated from these files:
	go-micro/server/debug/proto/debug.proto

It has these top-level messages:
	HealthRequest
	HealthResponse
*/
package debug

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthRequest struct {
}

func (m *HealthRequest) Reset()                    { *m = HealthRequest{} }
func (m *HealthRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()               {}
func (*HealthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HealthResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *HealthResponse) Reset()                    { *m = HealthResponse{} }
func (m *HealthResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()               {}
func (*HealthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HealthRequest)(nil), "HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "HealthResponse")
}

var fileDescriptor0 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcf, 0xd7, 0xcd,
	0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x4f, 0x49, 0x4d, 0x2a,
	0x4d, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0xb0, 0xf5, 0xc0, 0x6c, 0x25, 0x7e, 0x2e, 0x5e,
	0x8f, 0xd4, 0xc4, 0x9c, 0x92, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x05, 0x2e,
	0x3e, 0x98, 0x40, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x10, 0x1f, 0x17, 0x5b, 0x71, 0x49, 0x62,
	0x49, 0x69, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x12, 0x1b, 0x58, 0xa7, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x1c, 0xef, 0x98, 0xac, 0x64, 0x00, 0x00, 0x00,
}
