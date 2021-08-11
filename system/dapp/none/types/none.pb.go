// Code generated by protoc-gen-go. DO NOT EDIT.
// source: none.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type NoneAction struct {
	// Types that are valid to be assigned to Value:
	//	*NoneAction_CommitDelayTx
	Value                isNoneAction_Value `protobuf_oneof:"value"`
	Ty                   int32              `protobuf:"varint,2,opt,name=Ty,proto3" json:"Ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NoneAction) Reset()         { *m = NoneAction{} }
func (m *NoneAction) String() string { return proto.CompactTextString(m) }
func (*NoneAction) ProtoMessage()    {}
func (*NoneAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_784bcb57d038eb7f, []int{0}
}

func (m *NoneAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoneAction.Unmarshal(m, b)
}
func (m *NoneAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoneAction.Marshal(b, m, deterministic)
}
func (m *NoneAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoneAction.Merge(m, src)
}
func (m *NoneAction) XXX_Size() int {
	return xxx_messageInfo_NoneAction.Size(m)
}
func (m *NoneAction) XXX_DiscardUnknown() {
	xxx_messageInfo_NoneAction.DiscardUnknown(m)
}

var xxx_messageInfo_NoneAction proto.InternalMessageInfo

type isNoneAction_Value interface {
	isNoneAction_Value()
}

type NoneAction_CommitDelayTx struct {
	CommitDelayTx *CommitDelayTx `protobuf:"bytes,1,opt,name=commitDelayTx,proto3,oneof"`
}

func (*NoneAction_CommitDelayTx) isNoneAction_Value() {}

func (m *NoneAction) GetValue() isNoneAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *NoneAction) GetCommitDelayTx() *CommitDelayTx {
	if x, ok := m.GetValue().(*NoneAction_CommitDelayTx); ok {
		return x.CommitDelayTx
	}
	return nil
}

func (m *NoneAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NoneAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NoneAction_CommitDelayTx)(nil),
	}
}

// 提交延时交易类型
type CommitDelayTx struct {
	DelayTx                string   `protobuf:"bytes,1,opt,name=delayTx,proto3" json:"delayTx,omitempty"`
	RelativeDelayTime      int64    `protobuf:"varint,2,opt,name=relativeDelayTime,proto3" json:"relativeDelayTime,omitempty"`
	IsBlockHeightDelayTime bool     `protobuf:"varint,3,opt,name=isBlockHeightDelayTime,proto3" json:"isBlockHeightDelayTime,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *CommitDelayTx) Reset()         { *m = CommitDelayTx{} }
func (m *CommitDelayTx) String() string { return proto.CompactTextString(m) }
func (*CommitDelayTx) ProtoMessage()    {}
func (*CommitDelayTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_784bcb57d038eb7f, []int{1}
}

func (m *CommitDelayTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitDelayTx.Unmarshal(m, b)
}
func (m *CommitDelayTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitDelayTx.Marshal(b, m, deterministic)
}
func (m *CommitDelayTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitDelayTx.Merge(m, src)
}
func (m *CommitDelayTx) XXX_Size() int {
	return xxx_messageInfo_CommitDelayTx.Size(m)
}
func (m *CommitDelayTx) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitDelayTx.DiscardUnknown(m)
}

var xxx_messageInfo_CommitDelayTx proto.InternalMessageInfo

func (m *CommitDelayTx) GetDelayTx() string {
	if m != nil {
		return m.DelayTx
	}
	return ""
}

func (m *CommitDelayTx) GetRelativeDelayTime() int64 {
	if m != nil {
		return m.RelativeDelayTime
	}
	return 0
}

func (m *CommitDelayTx) GetIsBlockHeightDelayTime() bool {
	if m != nil {
		return m.IsBlockHeightDelayTime
	}
	return false
}

// 提交延时交易回执
type CommitDelayTxLog struct {
	Submitter            string   `protobuf:"bytes,1,opt,name=submitter,proto3" json:"submitter,omitempty"`
	DelayTxHash          string   `protobuf:"bytes,2,opt,name=delayTxHash,proto3" json:"delayTxHash,omitempty"`
	DelayBeginTime       int64    `protobuf:"varint,3,opt,name=delayBeginTime,proto3" json:"delayBeginTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitDelayTxLog) Reset()         { *m = CommitDelayTxLog{} }
func (m *CommitDelayTxLog) String() string { return proto.CompactTextString(m) }
func (*CommitDelayTxLog) ProtoMessage()    {}
func (*CommitDelayTxLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_784bcb57d038eb7f, []int{2}
}

func (m *CommitDelayTxLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitDelayTxLog.Unmarshal(m, b)
}
func (m *CommitDelayTxLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitDelayTxLog.Marshal(b, m, deterministic)
}
func (m *CommitDelayTxLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitDelayTxLog.Merge(m, src)
}
func (m *CommitDelayTxLog) XXX_Size() int {
	return xxx_messageInfo_CommitDelayTxLog.Size(m)
}
func (m *CommitDelayTxLog) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitDelayTxLog.DiscardUnknown(m)
}

var xxx_messageInfo_CommitDelayTxLog proto.InternalMessageInfo

func (m *CommitDelayTxLog) GetSubmitter() string {
	if m != nil {
		return m.Submitter
	}
	return ""
}

func (m *CommitDelayTxLog) GetDelayTxHash() string {
	if m != nil {
		return m.DelayTxHash
	}
	return ""
}

func (m *CommitDelayTxLog) GetDelayBeginTime() int64 {
	if m != nil {
		return m.DelayBeginTime
	}
	return 0
}

func init() {
	proto.RegisterType((*NoneAction)(nil), "types.NoneAction")
	proto.RegisterType((*CommitDelayTx)(nil), "types.CommitDelayTx")
	proto.RegisterType((*CommitDelayTxLog)(nil), "types.CommitDelayTxLog")
}

func init() {
	proto.RegisterFile("none.proto", fileDescriptor_784bcb57d038eb7f)
}

var fileDescriptor_784bcb57d038eb7f = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x4d, 0x4b, 0xbd, 0xeb, 0x3b, 0xee, 0xd0, 0x87, 0x48, 0x07, 0x87, 0xd2, 0x41, 0x3a,
	0x48, 0x04, 0x05, 0x27, 0x17, 0xab, 0x43, 0x07, 0x71, 0x08, 0x9d, 0xdc, 0x7a, 0xf5, 0xd1, 0x0b,
	0xb6, 0xc9, 0xd1, 0xe6, 0x0e, 0xe3, 0x3f, 0xe1, 0xbf, 0x2c, 0xc4, 0xd3, 0x6b, 0x15, 0xc7, 0xef,
	0xaf, 0xe4, 0x43, 0x02, 0xa0, 0xb4, 0x22, 0xbe, 0xee, 0xb4, 0xd1, 0x18, 0x18, 0xbb, 0xa6, 0x3e,
	0xa9, 0x00, 0x9e, 0xb4, 0xa2, 0xbb, 0xca, 0x48, 0xad, 0xf0, 0x16, 0xe6, 0x95, 0x6e, 0x5b, 0x69,
	0x1e, 0xa8, 0x29, 0x6d, 0xf1, 0x16, 0xb1, 0x98, 0xa5, 0xb3, 0xab, 0x13, 0xee, 0xca, 0xfc, 0x7e,
	0x98, 0xe5, 0x07, 0x62, 0x5c, 0xc6, 0x05, 0x78, 0x85, 0x8d, 0xbc, 0x98, 0xa5, 0x81, 0xf0, 0x0a,
	0x9b, 0x4d, 0x20, 0xd8, 0x96, 0xcd, 0x86, 0x92, 0x0f, 0x06, 0xf3, 0xd1, 0x16, 0x23, 0x98, 0xbc,
	0x0c, 0xae, 0x08, 0xc5, 0xb7, 0xc4, 0x0b, 0x38, 0xee, 0xa8, 0x29, 0x8d, 0xdc, 0xd2, 0x57, 0x59,
	0xb6, 0xe4, 0xce, 0xf4, 0xc5, 0xdf, 0x00, 0x6f, 0xe0, 0x54, 0xf6, 0x59, 0xa3, 0xab, 0xd7, 0x9c,
	0x64, 0xbd, 0x32, 0xfb, 0x89, 0x1f, 0xb3, 0x74, 0x2a, 0xfe, 0x49, 0x93, 0x77, 0x38, 0x1a, 0x01,
	0x3d, 0xea, 0x1a, 0xcf, 0x20, 0xec, 0x37, 0xcb, 0x56, 0x1a, 0x43, 0xdd, 0x8e, 0x6a, 0x6f, 0x60,
	0x0c, 0xb3, 0x1d, 0x62, 0x5e, 0xf6, 0x2b, 0x47, 0x14, 0x8a, 0xa1, 0x85, 0xe7, 0xb0, 0x70, 0x32,
	0xa3, 0x5a, 0xaa, 0x1f, 0x06, 0x5f, 0xfc, 0x72, 0x33, 0x78, 0x9e, 0x72, 0x7e, 0xe9, 0x5e, 0x74,
	0x79, 0xe8, 0x3e, 0xe3, 0xfa, 0x33, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x94, 0x73, 0xbb, 0x9a, 0x01,
	0x00, 0x00,
}
