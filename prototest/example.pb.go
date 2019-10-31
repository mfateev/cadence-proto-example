// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prototest/example.proto

package prototest

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

type ExampleMsg struct {
	SimpleString string `protobuf:"bytes,1,opt,name=simple_string,json=simpleString,proto3" json:"simple_string,omitempty"`
	// Types that are valid to be assigned to ComplexField:
	//	*ExampleMsg_SomeString
	//	*ExampleMsg_SomeNumber
	//	*ExampleMsg_SomeBool
	ComplexField         isExampleMsg_ComplexField `protobuf_oneof:"complex_field"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExampleMsg) Reset()         { *m = ExampleMsg{} }
func (m *ExampleMsg) String() string { return proto.CompactTextString(m) }
func (*ExampleMsg) ProtoMessage()    {}
func (*ExampleMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc265aa5adcb83f7, []int{0}
}

func (m *ExampleMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleMsg.Unmarshal(m, b)
}
func (m *ExampleMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleMsg.Marshal(b, m, deterministic)
}
func (m *ExampleMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleMsg.Merge(m, src)
}
func (m *ExampleMsg) XXX_Size() int {
	return xxx_messageInfo_ExampleMsg.Size(m)
}
func (m *ExampleMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleMsg proto.InternalMessageInfo

func (m *ExampleMsg) GetSimpleString() string {
	if m != nil {
		return m.SimpleString
	}
	return ""
}

type isExampleMsg_ComplexField interface {
	isExampleMsg_ComplexField()
}

type ExampleMsg_SomeString struct {
	SomeString string `protobuf:"bytes,2,opt,name=some_string,json=someString,proto3,oneof"`
}

type ExampleMsg_SomeNumber struct {
	SomeNumber int64 `protobuf:"varint,3,opt,name=some_number,json=someNumber,proto3,oneof"`
}

type ExampleMsg_SomeBool struct {
	SomeBool bool `protobuf:"varint,4,opt,name=some_bool,json=someBool,proto3,oneof"`
}

func (*ExampleMsg_SomeString) isExampleMsg_ComplexField() {}

func (*ExampleMsg_SomeNumber) isExampleMsg_ComplexField() {}

func (*ExampleMsg_SomeBool) isExampleMsg_ComplexField() {}

func (m *ExampleMsg) GetComplexField() isExampleMsg_ComplexField {
	if m != nil {
		return m.ComplexField
	}
	return nil
}

func (m *ExampleMsg) GetSomeString() string {
	if x, ok := m.GetComplexField().(*ExampleMsg_SomeString); ok {
		return x.SomeString
	}
	return ""
}

func (m *ExampleMsg) GetSomeNumber() int64 {
	if x, ok := m.GetComplexField().(*ExampleMsg_SomeNumber); ok {
		return x.SomeNumber
	}
	return 0
}

func (m *ExampleMsg) GetSomeBool() bool {
	if x, ok := m.GetComplexField().(*ExampleMsg_SomeBool); ok {
		return x.SomeBool
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExampleMsg) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExampleMsg_SomeString)(nil),
		(*ExampleMsg_SomeNumber)(nil),
		(*ExampleMsg_SomeBool)(nil),
	}
}

func init() {
	proto.RegisterType((*ExampleMsg)(nil), "ExampleMsg")
}

func init() { proto.RegisterFile("prototest/example.proto", fileDescriptor_bc265aa5adcb83f7) }

var fileDescriptor_bc265aa5adcb83f7 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x03, 0x8b,
	0x28, 0x2d, 0x67, 0xe4, 0xe2, 0x72, 0x85, 0x88, 0xf8, 0x16, 0xa7, 0x0b, 0x29, 0x73, 0xf1, 0x16,
	0x67, 0x82, 0x38, 0xf1, 0xc5, 0x25, 0x45, 0x99, 0x79, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x3c, 0x10, 0xc1, 0x60, 0xb0, 0x98, 0x90, 0x22, 0x17, 0x77, 0x71, 0x7e, 0x2e, 0x5c, 0x09,
	0x13, 0x48, 0x89, 0x07, 0x43, 0x10, 0x17, 0x48, 0x10, 0x4d, 0x49, 0x5e, 0x69, 0x6e, 0x52, 0x6a,
	0x91, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x33, 0x4c, 0x89, 0x1f, 0x58, 0x4c, 0x48, 0x96, 0x8b, 0x13,
	0xac, 0x24, 0x29, 0x3f, 0x3f, 0x47, 0x82, 0x45, 0x81, 0x51, 0x83, 0xc3, 0x83, 0x21, 0x88, 0x03,
	0x24, 0xe4, 0x94, 0x9f, 0x9f, 0xe3, 0xc4, 0xcf, 0xc5, 0x9b, 0x9c, 0x0f, 0xb2, 0xb4, 0x22, 0x3e,
	0x2d, 0x33, 0x35, 0x27, 0xc5, 0x89, 0x3b, 0x8a, 0x13, 0xee, 0x89, 0x24, 0x36, 0x30, 0xd3, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x54, 0x9c, 0xbf, 0x0d, 0xd8, 0x00, 0x00, 0x00,
}