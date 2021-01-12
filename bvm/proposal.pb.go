// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bvm/proposal.proto

package bvm

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

type ProposalData_Status int32

const (
	ProposalData_CREATE      ProposalData_Status = 0
	ProposalData_VOTING      ProposalData_Status = 1
	ProposalData_REJECT      ProposalData_Status = 2
	ProposalData_WAITING_EXE ProposalData_Status = 3
	ProposalData_CANCEL      ProposalData_Status = 4
	ProposalData_COMPLETED   ProposalData_Status = 5
	ProposalData_TIMEOUT     ProposalData_Status = 6
)

var ProposalData_Status_name = map[int32]string{
	0: "CREATE",
	1: "VOTING",
	2: "REJECT",
	3: "WAITING_EXE",
	4: "CANCEL",
	5: "COMPLETED",
	6: "TIMEOUT",
}

var ProposalData_Status_value = map[string]int32{
	"CREATE":      0,
	"VOTING":      1,
	"REJECT":      2,
	"WAITING_EXE": 3,
	"CANCEL":      4,
	"COMPLETED":   5,
	"TIMEOUT":     6,
}

func (x ProposalData_Status) String() string {
	return proto.EnumName(ProposalData_Status_name, int32(x))
}

func (ProposalData_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12923764dba3f7db, []int{0, 0}
}

type ProposalData_PType int32

const (
	ProposalData_CONFIG     ProposalData_PType = 0
	ProposalData_PERMISSION ProposalData_PType = 1
	ProposalData_NODE       ProposalData_PType = 2
)

var ProposalData_PType_name = map[int32]string{
	0: "CONFIG",
	1: "PERMISSION",
	2: "NODE",
}

var ProposalData_PType_value = map[string]int32{
	"CONFIG":     0,
	"PERMISSION": 1,
	"NODE":       2,
}

func (x ProposalData_PType) String() string {
	return proto.EnumName(ProposalData_PType_name, int32(x))
}

func (ProposalData_PType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12923764dba3f7db, []int{0, 1}
}

type ProposalData struct {
	Id                   uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 []byte              `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Timestamp            int64               `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Timeout              int64               `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Status               ProposalData_Status `protobuf:"varint,5,opt,name=status,proto3,enum=protos.ProposalData_Status" json:"status,omitempty"`
	Assentor             []*VoteInfo         `protobuf:"bytes,6,rep,name=assentor,proto3" json:"assentor,omitempty"`
	Objector             []*VoteInfo         `protobuf:"bytes,7,rep,name=objector,proto3" json:"objector,omitempty"`
	Threshold            uint32              `protobuf:"varint,8,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Score                uint32              `protobuf:"varint,9,opt,name=score,proto3" json:"score,omitempty"`
	Creator              string              `protobuf:"bytes,10,opt,name=creator,proto3" json:"creator,omitempty"`
	Version              string              `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`
	Type                 ProposalData_PType  `protobuf:"varint,12,opt,name=type,proto3,enum=protos.ProposalData_PType" json:"type,omitempty"`
	Completed            []byte              `protobuf:"bytes,13,opt,name=completed,proto3" json:"completed,omitempty"`
	Cancel               []byte              `protobuf:"bytes,14,opt,name=cancel,proto3" json:"cancel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProposalData) Reset()         { *m = ProposalData{} }
func (m *ProposalData) String() string { return proto.CompactTextString(m) }
func (*ProposalData) ProtoMessage()    {}
func (*ProposalData) Descriptor() ([]byte, []int) {
	return fileDescriptor_12923764dba3f7db, []int{0}
}

func (m *ProposalData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalData.Unmarshal(m, b)
}
func (m *ProposalData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalData.Marshal(b, m, deterministic)
}
func (m *ProposalData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalData.Merge(m, src)
}
func (m *ProposalData) XXX_Size() int {
	return xxx_messageInfo_ProposalData.Size(m)
}
func (m *ProposalData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalData.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalData proto.InternalMessageInfo

func (m *ProposalData) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProposalData) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ProposalData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ProposalData) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ProposalData) GetStatus() ProposalData_Status {
	if m != nil {
		return m.Status
	}
	return ProposalData_CREATE
}

func (m *ProposalData) GetAssentor() []*VoteInfo {
	if m != nil {
		return m.Assentor
	}
	return nil
}

func (m *ProposalData) GetObjector() []*VoteInfo {
	if m != nil {
		return m.Objector
	}
	return nil
}

func (m *ProposalData) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *ProposalData) GetScore() uint32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *ProposalData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ProposalData) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ProposalData) GetType() ProposalData_PType {
	if m != nil {
		return m.Type
	}
	return ProposalData_CONFIG
}

func (m *ProposalData) GetCompleted() []byte {
	if m != nil {
		return m.Completed
	}
	return nil
}

func (m *ProposalData) GetCancel() []byte {
	if m != nil {
		return m.Cancel
	}
	return nil
}

type VoteInfo struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	TxHash               string   `protobuf:"bytes,2,opt,name=txHash,proto3" json:"txHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteInfo) Reset()         { *m = VoteInfo{} }
func (m *VoteInfo) String() string { return proto.CompactTextString(m) }
func (*VoteInfo) ProtoMessage()    {}
func (*VoteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_12923764dba3f7db, []int{1}
}

func (m *VoteInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteInfo.Unmarshal(m, b)
}
func (m *VoteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteInfo.Marshal(b, m, deterministic)
}
func (m *VoteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteInfo.Merge(m, src)
}
func (m *VoteInfo) XXX_Size() int {
	return xxx_messageInfo_VoteInfo.Size(m)
}
func (m *VoteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VoteInfo proto.InternalMessageInfo

func (m *VoteInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *VoteInfo) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.ProposalData_Status", ProposalData_Status_name, ProposalData_Status_value)
	proto.RegisterEnum("protos.ProposalData_PType", ProposalData_PType_name, ProposalData_PType_value)
	proto.RegisterType((*ProposalData)(nil), "protos.ProposalData")
	proto.RegisterType((*VoteInfo)(nil), "protos.VoteInfo")
}

func init() { proto.RegisterFile("bvm/proposal.proto", fileDescriptor_12923764dba3f7db) }

var fileDescriptor_12923764dba3f7db = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x5d, 0xa7, 0x69, 0xda, 0x4c, 0x3f, 0xb0, 0x2c, 0x84, 0x2c, 0xe0, 0x10, 0xf5, 0x94, 0x03,
	0x14, 0x69, 0x57, 0xe2, 0x5e, 0xb5, 0x66, 0x09, 0xda, 0x26, 0x95, 0x1b, 0x16, 0x6e, 0xc8, 0x4d,
	0x8c, 0x5a, 0xd4, 0xd6, 0x51, 0xec, 0x5d, 0xb1, 0xbf, 0x96, 0xbf, 0x82, 0xec, 0x24, 0x94, 0x03,
	0x9c, 0x3c, 0x6f, 0xde, 0x1b, 0x79, 0xde, 0xe8, 0x01, 0xd9, 0x3d, 0x9e, 0xde, 0x55, 0xb5, 0xaa,
	0x94, 0x16, 0xc7, 0x79, 0x55, 0x2b, 0xa3, 0x48, 0xe0, 0x1e, 0x3d, 0xfb, 0xe5, 0xc3, 0x78, 0xd3,
	0x52, 0x2b, 0x61, 0x04, 0x99, 0x82, 0x77, 0x28, 0x29, 0x8a, 0x50, 0xec, 0x73, 0xef, 0x50, 0x12,
	0x02, 0x7e, 0xa1, 0x4a, 0x49, 0xbd, 0x08, 0xc5, 0x63, 0xee, 0x6a, 0xf2, 0x1a, 0x42, 0x73, 0x38,
	0x49, 0x6d, 0xc4, 0xa9, 0xa2, 0xbd, 0x08, 0xc5, 0x3d, 0x7e, 0x69, 0x10, 0x0a, 0x03, 0x0b, 0xd4,
	0x83, 0xa1, 0xbe, 0xe3, 0x3a, 0x48, 0x6e, 0x20, 0xd0, 0x46, 0x98, 0x07, 0x4d, 0xfb, 0x11, 0x8a,
	0xa7, 0xd7, 0xaf, 0x9a, 0x65, 0xf4, 0xfc, 0xef, 0x0d, 0xe6, 0x5b, 0x27, 0xe1, 0xad, 0x94, 0xbc,
	0x81, 0xa1, 0xd0, 0x5a, 0x9e, 0x8d, 0xaa, 0x69, 0x10, 0xf5, 0xe2, 0xd1, 0x35, 0xee, 0xc6, 0xee,
	0x95, 0x91, 0xc9, 0xf9, 0xbb, 0xe2, 0x7f, 0x14, 0x56, 0xad, 0x76, 0x3f, 0x64, 0x61, 0xd5, 0x83,
	0xff, 0xa9, 0x3b, 0x85, 0x33, 0xb2, 0xaf, 0xa5, 0xde, 0xab, 0x63, 0x49, 0x87, 0x11, 0x8a, 0x27,
	0xfc, 0xd2, 0x20, 0xcf, 0xa1, 0xaf, 0x0b, 0x55, 0x4b, 0x1a, 0x3a, 0xa6, 0x01, 0xd6, 0x5e, 0x51,
	0x4b, 0x61, 0x3f, 0x80, 0x08, 0xc5, 0x21, 0xef, 0xa0, 0x65, 0x1e, 0x65, 0xad, 0x0f, 0xea, 0x4c,
	0x47, 0x0d, 0xd3, 0x42, 0x32, 0x07, 0xdf, 0x3c, 0x55, 0x92, 0x8e, 0x9d, 0xed, 0x97, 0xff, 0xb4,
	0xbd, 0xc9, 0x9f, 0x2a, 0xc9, 0x9d, 0xce, 0xee, 0x55, 0xa8, 0x53, 0x75, 0x94, 0x46, 0x96, 0x74,
	0xe2, 0x2e, 0x7f, 0x69, 0x90, 0x17, 0x10, 0x14, 0xe2, 0x5c, 0xc8, 0x23, 0x9d, 0x3a, 0xaa, 0x45,
	0x33, 0x09, 0x41, 0x73, 0x3b, 0x02, 0x10, 0x2c, 0x39, 0x5b, 0xe4, 0x0c, 0x5f, 0xd9, 0xfa, 0x3e,
	0xcb, 0x93, 0xf4, 0x16, 0x23, 0x5b, 0x73, 0xf6, 0x89, 0x2d, 0x73, 0xec, 0x91, 0x67, 0x30, 0xfa,
	0xb2, 0x48, 0x2c, 0xf1, 0x8d, 0x7d, 0x65, 0xb8, 0xe7, 0x86, 0x16, 0xe9, 0x92, 0xdd, 0x61, 0x9f,
	0x4c, 0x20, 0x5c, 0x66, 0xeb, 0xcd, 0x1d, 0xcb, 0xd9, 0x0a, 0xf7, 0xc9, 0x08, 0x06, 0x79, 0xb2,
	0x66, 0xd9, 0xe7, 0x1c, 0x07, 0xb3, 0xb7, 0xd0, 0x77, 0xbb, 0xba, 0x81, 0x2c, 0xfd, 0x90, 0xdc,
	0xe2, 0x2b, 0x32, 0x05, 0xd8, 0x30, 0xbe, 0x4e, 0xb6, 0xdb, 0x24, 0x4b, 0x31, 0x22, 0x43, 0xf0,
	0xd3, 0x6c, 0xc5, 0xb0, 0x37, 0x7b, 0x0f, 0xc3, 0xee, 0xf2, 0x36, 0x4c, 0xa2, 0x2c, 0x6b, 0x17,
	0xaf, 0x90, 0xbb, 0xda, 0xba, 0x31, 0x3f, 0x3f, 0x0a, 0xbd, 0x77, 0x11, 0x0b, 0x79, 0x8b, 0x76,
	0x4d, 0x42, 0x6f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x79, 0x16, 0x9b, 0x40, 0xbe, 0x02, 0x00,
	0x00,
}
