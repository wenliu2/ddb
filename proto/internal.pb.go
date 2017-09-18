// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/internal.proto

package ddb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Mutation_Type int32

const (
	Mutation_UNSPECIFIED Mutation_Type = 0
	Mutation_PUT         Mutation_Type = 1
	Mutation_DELETE      Mutation_Type = 2
)

var Mutation_Type_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "PUT",
	2: "DELETE",
}
var Mutation_Type_value = map[string]int32{
	"UNSPECIFIED": 0,
	"PUT":         1,
	"DELETE":      2,
}

func (x Mutation_Type) String() string {
	return proto.EnumName(Mutation_Type_name, int32(x))
}
func (Mutation_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

type LogRecord struct {
	Mutation *Mutation `protobuf:"bytes,2,opt,name=mutation" json:"mutation,omitempty"`
}

func (m *LogRecord) Reset()                    { *m = LogRecord{} }
func (m *LogRecord) String() string            { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()               {}
func (*LogRecord) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *LogRecord) GetMutation() *Mutation {
	if m != nil {
		return m.Mutation
	}
	return nil
}

type Mutation struct {
	// The key associated with this mutation.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The type of this mutation.
	Type Mutation_Type `protobuf:"varint,2,opt,name=type,enum=ddb.internal.Mutation_Type" json:"type,omitempty"`
	// For mutations that update, this is the value.
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Mutation) Reset()                    { *m = Mutation{} }
func (m *Mutation) String() string            { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()               {}
func (*Mutation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Mutation) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Mutation) GetType() Mutation_Type {
	if m != nil {
		return m.Type
	}
	return Mutation_UNSPECIFIED
}

func (m *Mutation) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*LogRecord)(nil), "ddb.internal.LogRecord")
	proto.RegisterType((*Mutation)(nil), "ddb.internal.Mutation")
	proto.RegisterEnum("ddb.internal.Mutation_Type", Mutation_Type_name, Mutation_Type_value)
}

func init() { proto.RegisterFile("proto/internal.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x03, 0x73, 0x85, 0x78, 0x52,
	0x52, 0x92, 0xf4, 0x60, 0x62, 0x4a, 0xf6, 0x5c, 0x9c, 0x3e, 0xf9, 0xe9, 0x41, 0xa9, 0xc9, 0xf9,
	0x45, 0x29, 0x42, 0x46, 0x5c, 0x1c, 0xb9, 0xa5, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x62, 0x7a, 0xc8, 0xaa, 0xf5, 0x7c, 0xa1, 0xb2, 0x41, 0x70, 0x75,
	0x4a, 0x13, 0x19, 0xb9, 0x38, 0x60, 0xc2, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x3e, 0x17, 0x4b, 0x49, 0x65, 0x41, 0x2a, 0xd8,
	0x38, 0x3e, 0x23, 0x69, 0xec, 0xc6, 0xe9, 0x85, 0x54, 0x16, 0xa4, 0x06, 0x81, 0x15, 0x0a, 0x89,
	0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x04, 0x41, 0x38,
	0x4a, 0x3a, 0x5c, 0x2c, 0x20, 0x35, 0x42, 0xfc, 0x5c, 0xdc, 0xa1, 0x7e, 0xc1, 0x01, 0xae, 0xce,
	0x9e, 0x6e, 0x9e, 0xae, 0x2e, 0x02, 0x0c, 0x42, 0xec, 0x5c, 0xcc, 0x01, 0xa1, 0x21, 0x02, 0x8c,
	0x42, 0x5c, 0x5c, 0x6c, 0x2e, 0xae, 0x3e, 0xae, 0x21, 0xae, 0x02, 0x4c, 0x49, 0x6c, 0x60, 0x9f,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x52, 0xc9, 0x49, 0x01, 0x01, 0x00, 0x00,
}
