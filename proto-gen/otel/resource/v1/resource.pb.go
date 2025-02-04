// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource/v1/resource.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/jaegertracing/jaeger/proto-gen/otel/common/v1"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Resource information.
type Resource struct {
	// Set of attributes that describe the resource.
	// Attribute keys MUST be unique (it is not allowed to have more than one
	// attribute with the same key).
	Attributes []*v1.KeyValue `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0, then
	// no attributes were dropped.
	DroppedAttributesCount uint32   `protobuf:"varint,2,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_cebae6241f1ea243, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetAttributes() []*v1.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Resource) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Resource)(nil), "jaeger.resource.v1.Resource")
}

func init() { proto.RegisterFile("resource/v1/resource.proto", fileDescriptor_cebae6241f1ea243) }

var fileDescriptor_cebae6241f1ea243 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x22, 0x21, 0x64, 0xd4, 0x25, 0x43, 0x15, 0x65, 0xa1, 0xea, 0xd4, 0x05, 0x5b,
	0x81, 0x05, 0x75, 0xa3, 0x8c, 0x0c, 0x54, 0x11, 0xea, 0xc0, 0x52, 0x25, 0xe9, 0x29, 0x04, 0x25,
	0x3e, 0xeb, 0x7a, 0x8e, 0x94, 0x8d, 0x77, 0xe0, 0x2d, 0x78, 0x4a, 0x94, 0xc4, 0x09, 0xd9, 0xce,
	0xf7, 0xff, 0xf7, 0xdf, 0xe7, 0x93, 0x11, 0xc1, 0x05, 0x1d, 0xe5, 0xa0, 0x9b, 0x58, 0x8f, 0xb5,
	0xb2, 0x84, 0x8c, 0x41, 0xf0, 0x95, 0x42, 0x01, 0xa4, 0xa6, 0x76, 0x13, 0x47, 0xab, 0x1c, 0xeb,
	0x1a, 0x4d, 0xe7, 0x1e, 0xaa, 0xc1, 0xbb, 0xf9, 0x16, 0xf2, 0x26, 0xf1, 0xbe, 0x60, 0x27, 0x65,
	0xca, 0x4c, 0x65, 0xe6, 0x18, 0x2e, 0xa1, 0x58, 0x5f, 0x6d, 0x6f, 0x1f, 0x22, 0xe5, 0xd3, 0xfc,
	0x58, 0x13, 0xab, 0x57, 0x68, 0x8f, 0x69, 0xe5, 0x20, 0x99, 0xb9, 0x83, 0x27, 0x19, 0x9e, 0x09,
	0xad, 0x85, 0xf3, 0xe9, 0xbf, 0x7b, 0xca, 0xd1, 0x19, 0x0e, 0x17, 0x6b, 0xb1, 0x5d, 0x26, 0x2b,
	0xaf, 0x3f, 0x4f, 0xf2, 0x4b, 0xa7, 0xee, 0x7f, 0x84, 0xdc, 0x94, 0xa8, 0xd0, 0x82, 0x61, 0xa8,
	0xa0, 0x06, 0xa6, 0x76, 0xa0, 0x9b, 0xff, 0x60, 0xbf, 0x1c, 0x31, 0x0f, 0x9d, 0x74, 0x10, 0x1f,
	0xbb, 0xa2, 0xe4, 0x4f, 0x97, 0x75, 0x60, 0x7a, 0x60, 0x64, 0x4a, 0xf3, 0xd2, 0x14, 0xfe, 0xa5,
	0xfb, 0x8c, 0xfb, 0x02, 0x8c, 0x46, 0x86, 0x4a, 0xcf, 0x2e, 0xf6, 0xbb, 0xb8, 0x7b, 0xb3, 0x60,
	0xde, 0xa7, 0x75, 0x7d, 0xa6, 0x1a, 0x37, 0xa8, 0x63, 0x9c, 0x5d, 0xf7, 0xd3, 0x8f, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x55, 0x91, 0xb4, 0x1d, 0x69, 0x01, 0x00, 0x00,
}
