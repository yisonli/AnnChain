// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gemmill/protos/blockchain/message.proto

package blockchain

import (
	fmt "fmt"
	types "github.com/dappledger/AnnChain/gemmill/protos/types"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type MsgType int32

const (
	MsgType_None      MsgType = 0
	MsgType_BlockReq  MsgType = 1
	MsgType_BlockRsp  MsgType = 2
	MsgType_StatusReq MsgType = 3
	MsgType_StatusRsp MsgType = 4
	MsgType_HeaderReq MsgType = 5
	MsgType_HeaderRsp MsgType = 6
)

var MsgType_name = map[int32]string{
	0: "None",
	1: "BlockReq",
	2: "BlockRsp",
	3: "StatusReq",
	4: "StatusRsp",
	5: "HeaderReq",
	6: "HeaderRsp",
}

var MsgType_value = map[string]int32{
	"None":      0,
	"BlockReq":  1,
	"BlockRsp":  2,
	"StatusReq": 3,
	"StatusRsp": 4,
	"HeaderReq": 5,
	"HeaderRsp": 6,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1bdd4fb1b2575be, []int{0}
}

type BlockRequestMessage struct {
	Height               int64    `protobuf:"varint,1,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockRequestMessage) Reset()         { *m = BlockRequestMessage{} }
func (m *BlockRequestMessage) String() string { return proto.CompactTextString(m) }
func (*BlockRequestMessage) ProtoMessage()    {}
func (*BlockRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bdd4fb1b2575be, []int{0}
}
func (m *BlockRequestMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockRequestMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRequestMessage.Merge(m, src)
}
func (m *BlockRequestMessage) XXX_Size() int {
	return m.Size()
}
func (m *BlockRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRequestMessage proto.InternalMessageInfo

func (m *BlockRequestMessage) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type BlockResponseMessage struct {
	Block                *types.Block `protobuf:"bytes,1,opt,name=Block,proto3" json:"Block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BlockResponseMessage) Reset()         { *m = BlockResponseMessage{} }
func (m *BlockResponseMessage) String() string { return proto.CompactTextString(m) }
func (*BlockResponseMessage) ProtoMessage()    {}
func (*BlockResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bdd4fb1b2575be, []int{1}
}
func (m *BlockResponseMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockResponseMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockResponseMessage.Merge(m, src)
}
func (m *BlockResponseMessage) XXX_Size() int {
	return m.Size()
}
func (m *BlockResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BlockResponseMessage proto.InternalMessageInfo

func (m *BlockResponseMessage) GetBlock() *types.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type StatusRequestMessage struct {
	Height               int64    `protobuf:"varint,1,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequestMessage) Reset()         { *m = StatusRequestMessage{} }
func (m *StatusRequestMessage) String() string { return proto.CompactTextString(m) }
func (*StatusRequestMessage) ProtoMessage()    {}
func (*StatusRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bdd4fb1b2575be, []int{2}
}
func (m *StatusRequestMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusRequestMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequestMessage.Merge(m, src)
}
func (m *StatusRequestMessage) XXX_Size() int {
	return m.Size()
}
func (m *StatusRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequestMessage proto.InternalMessageInfo

func (m *StatusRequestMessage) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type StatusResponseMessage struct {
	Height               int64    `protobuf:"varint,1,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponseMessage) Reset()         { *m = StatusResponseMessage{} }
func (m *StatusResponseMessage) String() string { return proto.CompactTextString(m) }
func (*StatusResponseMessage) ProtoMessage()    {}
func (*StatusResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bdd4fb1b2575be, []int{3}
}
func (m *StatusResponseMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusResponseMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponseMessage.Merge(m, src)
}
func (m *StatusResponseMessage) XXX_Size() int {
	return m.Size()
}
func (m *StatusResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponseMessage proto.InternalMessageInfo

func (m *StatusResponseMessage) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type BlockHeaderRequestMessage struct {
	Height               int64    `protobuf:"varint,1,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeaderRequestMessage) Reset()         { *m = BlockHeaderRequestMessage{} }
func (m *BlockHeaderRequestMessage) String() string { return proto.CompactTextString(m) }
func (*BlockHeaderRequestMessage) ProtoMessage()    {}
func (*BlockHeaderRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bdd4fb1b2575be, []int{4}
}
func (m *BlockHeaderRequestMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeaderRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeaderRequestMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeaderRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeaderRequestMessage.Merge(m, src)
}
func (m *BlockHeaderRequestMessage) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeaderRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeaderRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeaderRequestMessage proto.InternalMessageInfo

func (m *BlockHeaderRequestMessage) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type BlockHeaderResponseMessage struct {
	Header               *types.Header `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BlockHeaderResponseMessage) Reset()         { *m = BlockHeaderResponseMessage{} }
func (m *BlockHeaderResponseMessage) String() string { return proto.CompactTextString(m) }
func (*BlockHeaderResponseMessage) ProtoMessage()    {}
func (*BlockHeaderResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bdd4fb1b2575be, []int{5}
}
func (m *BlockHeaderResponseMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeaderResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeaderResponseMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeaderResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeaderResponseMessage.Merge(m, src)
}
func (m *BlockHeaderResponseMessage) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeaderResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeaderResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeaderResponseMessage proto.InternalMessageInfo

func (m *BlockHeaderResponseMessage) GetHeader() *types.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

type BlockMessage struct {
	Type                 MsgType  `protobuf:"varint,1,opt,name=Type,proto3,enum=blockchain.MsgType" json:"Type,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockMessage) Reset()         { *m = BlockMessage{} }
func (m *BlockMessage) String() string { return proto.CompactTextString(m) }
func (*BlockMessage) ProtoMessage()    {}
func (*BlockMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bdd4fb1b2575be, []int{6}
}
func (m *BlockMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMessage.Merge(m, src)
}
func (m *BlockMessage) XXX_Size() int {
	return m.Size()
}
func (m *BlockMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMessage proto.InternalMessageInfo

func (m *BlockMessage) GetType() MsgType {
	if m != nil {
		return m.Type
	}
	return MsgType_None
}

func (m *BlockMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("blockchain.MsgType", MsgType_name, MsgType_value)
	proto.RegisterType((*BlockRequestMessage)(nil), "blockchain.BlockRequestMessage")
	proto.RegisterType((*BlockResponseMessage)(nil), "blockchain.BlockResponseMessage")
	proto.RegisterType((*StatusRequestMessage)(nil), "blockchain.StatusRequestMessage")
	proto.RegisterType((*StatusResponseMessage)(nil), "blockchain.StatusResponseMessage")
	proto.RegisterType((*BlockHeaderRequestMessage)(nil), "blockchain.BlockHeaderRequestMessage")
	proto.RegisterType((*BlockHeaderResponseMessage)(nil), "blockchain.BlockHeaderResponseMessage")
	proto.RegisterType((*BlockMessage)(nil), "blockchain.BlockMessage")
}

func init() {
	proto.RegisterFile("gemmill/protos/blockchain/message.proto", fileDescriptor_e1bdd4fb1b2575be)
}

var fileDescriptor_e1bdd4fb1b2575be = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0xea, 0x40,
	0x18, 0xc5, 0x8d, 0x46, 0xaf, 0xf7, 0xbb, 0xf1, 0x12, 0x46, 0x5b, 0xac, 0x8b, 0x20, 0x81, 0xa2,
	0x14, 0x9a, 0x80, 0x6e, 0x4a, 0x77, 0xd5, 0x2e, 0xa4, 0xc5, 0x2e, 0xd2, 0xae, 0xba, 0x1b, 0xf5,
	0x23, 0x06, 0xf3, 0x67, 0xea, 0x8c, 0x0b, 0xdf, 0xa4, 0x8f, 0xd4, 0x65, 0x1f, 0xa1, 0xd8, 0x17,
	0x29, 0x99, 0x4c, 0xfc, 0x53, 0x10, 0xdc, 0x84, 0x9c, 0x39, 0xbf, 0x33, 0x87, 0x99, 0xf9, 0xa0,
	0xe3, 0x63, 0x14, 0x05, 0x61, 0xe8, 0xb2, 0x65, 0x22, 0x12, 0xee, 0x4e, 0xc2, 0x64, 0xba, 0x98,
	0xce, 0x69, 0x10, 0xbb, 0x11, 0x72, 0x4e, 0x7d, 0x74, 0xa4, 0x43, 0x60, 0xe7, 0xb4, 0xda, 0xbf,
	0x42, 0x62, 0xcd, 0x50, 0x7d, 0x33, 0xda, 0xbe, 0x86, 0xfa, 0x20, 0xe5, 0x3d, 0x7c, 0x5b, 0x21,
	0x17, 0xe3, 0x6c, 0x2b, 0x72, 0x0e, 0x95, 0x11, 0x06, 0xfe, 0x5c, 0x34, 0xb5, 0xb6, 0xd6, 0x2d,
	0x79, 0x4a, 0xd9, 0xb7, 0xd0, 0x50, 0x38, 0x67, 0x49, 0xcc, 0x31, 0xe7, 0x6d, 0x28, 0xcb, 0x75,
	0x89, 0xff, 0xeb, 0x19, 0x4e, 0xd6, 0x91, 0xb1, 0x99, 0x65, 0x3b, 0xd0, 0x78, 0x16, 0x54, 0xac,
	0xf8, 0x89, 0x5d, 0x2e, 0x9c, 0xe5, 0xfc, 0x61, 0xd9, 0xb1, 0x40, 0x1f, 0x2e, 0x64, 0xd3, 0x08,
	0xe9, 0x0c, 0x97, 0x27, 0xb6, 0x0c, 0xa1, 0x75, 0x10, 0x3a, 0xac, 0xba, 0x4c, 0x53, 0xa9, 0xa1,
	0x0e, 0x56, 0x53, 0x07, 0x53, 0xb4, 0x32, 0xed, 0x47, 0x30, 0xe4, 0x26, 0x79, 0xac, 0x03, 0xfa,
	0xcb, 0x9a, 0xa1, 0x0c, 0xfd, 0xef, 0xd5, 0x9d, 0xdd, 0x93, 0x38, 0x63, 0xee, 0xa7, 0x96, 0x27,
	0x01, 0x42, 0x40, 0xbf, 0xa7, 0x82, 0x36, 0x8b, 0x6d, 0xad, 0x6b, 0x78, 0xf2, 0xff, 0x6a, 0x01,
	0x7f, 0x14, 0x44, 0xaa, 0xa0, 0x3f, 0x25, 0x31, 0x9a, 0x05, 0x62, 0x40, 0x35, 0x7f, 0x27, 0x53,
	0xdb, 0x29, 0xce, 0xcc, 0x22, 0xa9, 0xc1, 0xdf, 0xed, 0xc5, 0x9a, 0xa5, 0x3d, 0xc9, 0x99, 0xa9,
	0xa7, 0x72, 0x7b, 0x21, 0x66, 0x79, 0x4f, 0x72, 0x66, 0x56, 0x06, 0x0f, 0x1f, 0x1b, 0x4b, 0xfb,
	0xdc, 0x58, 0xda, 0xd7, 0xc6, 0xd2, 0xde, 0xbf, 0xad, 0xc2, 0xeb, 0x8d, 0x1f, 0x88, 0xf9, 0x6a,
	0xe2, 0x4c, 0x93, 0xc8, 0x9d, 0x51, 0xc6, 0x42, 0x9c, 0xf9, 0xb8, 0x74, 0xef, 0xe2, 0x78, 0x28,
	0xa7, 0xed, 0xe8, 0x1c, 0x4e, 0x2a, 0x72, 0xa9, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xa5,
	0xfc, 0x90, 0xab, 0x02, 0x00, 0x00,
}

func (m *BlockRequestMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockRequestMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockRequestMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Height != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockResponseMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockResponseMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockResponseMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StatusRequestMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusRequestMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusRequestMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Height != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StatusResponseMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusResponseMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusResponseMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Height != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeaderRequestMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeaderRequestMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeaderRequestMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Height != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeaderResponseMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeaderResponseMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeaderResponseMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockRequestMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockResponseMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatusRequestMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatusResponseMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockHeaderRequestMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockHeaderResponseMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMessage(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockRequestMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockRequestMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockRequestMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockResponseMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockResponseMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockResponseMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &types.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusRequestMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusRequestMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusRequestMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusResponseMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusResponseMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusResponseMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockHeaderRequestMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHeaderRequestMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeaderRequestMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockHeaderResponseMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHeaderResponseMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeaderResponseMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &types.Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= MsgType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
