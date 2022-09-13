// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/v1/info.proto

// Package common defines commonly used types agnostic to the node role on the Sonr network.

package common

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EntityKind int32

const (
	EntityKind_ADDRESS EntityKind = 0
	EntityKind_DID     EntityKind = 1
	EntityKind_CID     EntityKind = 2
)

var EntityKind_name = map[int32]string{
	0: "ADDRESS",
	1: "DID",
	2: "CID",
}

var EntityKind_value = map[string]int32{
	"ADDRESS": 0,
	"DID":     1,
	"CID":     2,
}

func (x EntityKind) String() string {
	return proto.EnumName(EntityKind_name, int32(x))
}

func (EntityKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{0}
}

type BlockchainModule int32

const (
	// Query x/registry module
	BlockchainModule_REGISTRY BlockchainModule = 0
	// Query x/schema module
	BlockchainModule_SCHEMA BlockchainModule = 1
	// Query x/bucket module
	BlockchainModule_BUCKET BlockchainModule = 2
)

var BlockchainModule_name = map[int32]string{
	0: "REGISTRY",
	1: "SCHEMA",
	2: "BUCKET",
}

var BlockchainModule_value = map[string]int32{
	"REGISTRY": 0,
	"SCHEMA":   1,
	"BUCKET":   2,
}

func (x BlockchainModule) String() string {
	return proto.EnumName(BlockchainModule_name, int32(x))
}

func (BlockchainModule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{1}
}

type MotorCallbackMessage int32

const (
	MotorCallbackMessage_MTR_INIT            MotorCallbackMessage = 0
	MotorCallbackMessage_MTR_FAUCET_RECEIVED MotorCallbackMessage = 1
	MotorCallbackMessage_MTR_SHARDS_CREATED  MotorCallbackMessage = 2
	MotorCallbackMessage_MTR_VAULT_CREATED   MotorCallbackMessage = 3
	MotorCallbackMessage_MTR_ACCOUNT_CREATED MotorCallbackMessage = 4
	MotorCallbackMessage_MTR_LOGGED_IN       MotorCallbackMessage = 5
)

var MotorCallbackMessage_name = map[int32]string{
	0: "MTR_INIT",
	1: "MTR_FAUCET_RECEIVED",
	2: "MTR_SHARDS_CREATED",
	3: "MTR_VAULT_CREATED",
	4: "MTR_ACCOUNT_CREATED",
	5: "MTR_LOGGED_IN",
}

var MotorCallbackMessage_value = map[string]int32{
	"MTR_INIT":            0,
	"MTR_FAUCET_RECEIVED": 1,
	"MTR_SHARDS_CREATED":  2,
	"MTR_VAULT_CREATED":   3,
	"MTR_ACCOUNT_CREATED": 4,
	"MTR_LOGGED_IN":       5,
}

func (x MotorCallbackMessage) String() string {
	return proto.EnumName(MotorCallbackMessage_name, int32(x))
}

func (MotorCallbackMessage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{2}
}

// Peers Active Status
type Peer_Status int32

const (
	Peer_STATUS_UNSPECIFIED Peer_Status = 0
	Peer_STATUS_ONLINE      Peer_Status = 1
	Peer_STATUS_AWAY        Peer_Status = 2
	Peer_STATUS_BUSY        Peer_Status = 3
)

var Peer_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ONLINE",
	2: "STATUS_AWAY",
	3: "STATUS_BUSY",
}

var Peer_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_ONLINE":      1,
	"STATUS_AWAY":        2,
	"STATUS_BUSY":        3,
}

func (x Peer_Status) String() string {
	return proto.EnumName(Peer_Status_name, int32(x))
}

func (Peer_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{0, 0}
}

// Basic Info Sent to Peers to Establish Connections
type Peer struct {
	PeerId string      `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Did    string      `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	Status Peer_Status `protobuf:"varint,3,opt,name=status,proto3,enum=sonrio.common.v1.Peer_Status" json:"status,omitempty"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{0}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return m.Size()
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *Peer) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *Peer) GetStatus() Peer_Status {
	if m != nil {
		return m.Status
	}
	return Peer_STATUS_UNSPECIFIED
}

// AuthInfo is a object used by Motor clients to store authentication information in Biometric storage
type AuthInfo struct {
	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Did       string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	AesDscKey []byte `protobuf:"bytes,3,opt,name=aes_dsc_key,json=aesDscKey,proto3" json:"aes_dsc_key,omitempty"`
	AesPskKey []byte `protobuf:"bytes,4,opt,name=aes_psk_key,json=aesPskKey,proto3" json:"aes_psk_key,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *AuthInfo) Reset()         { *m = AuthInfo{} }
func (m *AuthInfo) String() string { return proto.CompactTextString(m) }
func (*AuthInfo) ProtoMessage()    {}
func (*AuthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{1}
}
func (m *AuthInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthInfo.Merge(m, src)
}
func (m *AuthInfo) XXX_Size() int {
	return m.Size()
}
func (m *AuthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AuthInfo proto.InternalMessageInfo

func (m *AuthInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AuthInfo) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *AuthInfo) GetAesDscKey() []byte {
	if m != nil {
		return m.AesDscKey
	}
	return nil
}

func (m *AuthInfo) GetAesPskKey() []byte {
	if m != nil {
		return m.AesPskKey
	}
	return nil
}

func (m *AuthInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthInfo) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("sonrio.common.v1.EntityKind", EntityKind_name, EntityKind_value)
	proto.RegisterEnum("sonrio.common.v1.BlockchainModule", BlockchainModule_name, BlockchainModule_value)
	proto.RegisterEnum("sonrio.common.v1.MotorCallbackMessage", MotorCallbackMessage_name, MotorCallbackMessage_value)
	proto.RegisterEnum("sonrio.common.v1.Peer_Status", Peer_Status_name, Peer_Status_value)
	proto.RegisterType((*Peer)(nil), "sonrio.common.v1.Peer")
	proto.RegisterType((*AuthInfo)(nil), "sonrio.common.v1.AuthInfo")
}

func init() { proto.RegisterFile("common/v1/info.proto", fileDescriptor_a6ffb5b3e6a498f4) }

var fileDescriptor_a6ffb5b3e6a498f4 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x4e, 0xdb, 0x4c,
	0x10, 0xc0, 0xbd, 0x09, 0x24, 0x30, 0xf0, 0x7d, 0x5d, 0xb6, 0x14, 0xa2, 0xaa, 0xb5, 0x10, 0x27,
	0x44, 0x55, 0xa7, 0xb4, 0xaa, 0xd4, 0xab, 0x63, 0x2f, 0x60, 0x41, 0x42, 0xe4, 0x3f, 0x54, 0xf4,
	0x62, 0x39, 0xf6, 0x42, 0x56, 0x49, 0xbc, 0x96, 0x77, 0x43, 0x95, 0xb7, 0xe8, 0xb1, 0x8f, 0xd1,
	0x37, 0xe8, 0xb5, 0x47, 0x8e, 0x3d, 0x56, 0xf0, 0x22, 0xd5, 0x9a, 0x10, 0xaa, 0xaa, 0x27, 0xcf,
	0xfc, 0x7e, 0x33, 0x23, 0x8f, 0x76, 0x60, 0x33, 0x15, 0x93, 0x89, 0xc8, 0xdb, 0xd7, 0x07, 0x6d,
	0x9e, 0x5f, 0x0a, 0xab, 0x28, 0x85, 0x12, 0x04, 0x4b, 0x91, 0x97, 0x5c, 0x58, 0xf7, 0xd2, 0xba,
	0x3e, 0xd8, 0xfd, 0x8e, 0x60, 0xa9, 0xcf, 0x58, 0x49, 0xb6, 0xa1, 0x59, 0x30, 0x56, 0xc6, 0x3c,
	0x6b, 0xa1, 0x1d, 0xb4, 0xb7, 0xea, 0x37, 0x74, 0xea, 0x65, 0x04, 0x43, 0x3d, 0xe3, 0x59, 0xab,
	0x56, 0x41, 0x1d, 0x92, 0xf7, 0xd0, 0x90, 0x2a, 0x51, 0x53, 0xd9, 0xaa, 0xef, 0xa0, 0xbd, 0xff,
	0xdf, 0xbe, 0xb4, 0xfe, 0x1e, 0x6b, 0xe9, 0x91, 0x56, 0x50, 0x15, 0xf9, 0xf3, 0xe2, 0xdd, 0x08,
	0x1a, 0xf7, 0x84, 0x6c, 0x01, 0x09, 0x42, 0x3b, 0x8c, 0x82, 0x38, 0xea, 0x05, 0x7d, 0xea, 0x78,
	0x87, 0x1e, 0x75, 0xb1, 0x41, 0x36, 0xe0, 0xbf, 0x39, 0x3f, 0xeb, 0x9d, 0x7a, 0x3d, 0x8a, 0x11,
	0x79, 0x02, 0x6b, 0x73, 0x64, 0x7f, 0xb4, 0x2f, 0x70, 0xed, 0x0f, 0xd0, 0x89, 0x82, 0x0b, 0x5c,
	0xdf, 0xfd, 0x86, 0x60, 0xc5, 0x9e, 0xaa, 0xa1, 0x97, 0x5f, 0x0a, 0xd2, 0x82, 0x66, 0x92, 0x65,
	0x25, 0x93, 0x72, 0xbe, 0xc5, 0x43, 0xfa, 0x8f, 0x35, 0x4c, 0x58, 0x4b, 0x98, 0x8c, 0x33, 0x99,
	0xc6, 0x23, 0x36, 0xab, 0x76, 0x59, 0xf7, 0x57, 0x13, 0x26, 0x5d, 0x99, 0x9e, 0xb0, 0xd9, 0x83,
	0x2f, 0xe4, 0xa8, 0xf2, 0x4b, 0x0b, 0xdf, 0x97, 0x23, 0xed, 0x9f, 0xc3, 0x4a, 0x91, 0x48, 0xf9,
	0x59, 0x94, 0x59, 0x6b, 0xb9, 0x1a, 0xbb, 0xc8, 0xc9, 0x0b, 0x58, 0x55, 0x7c, 0xc2, 0xa4, 0x4a,
	0x26, 0x45, 0xab, 0xb1, 0x83, 0xf6, 0xea, 0xfe, 0x23, 0xd8, 0x7f, 0x05, 0x40, 0x73, 0xc5, 0xd5,
	0xec, 0x84, 0xe7, 0x19, 0x59, 0x83, 0xa6, 0xed, 0xba, 0x3e, 0x0d, 0x02, 0x6c, 0x90, 0x26, 0xd4,
	0x5d, 0xcf, 0xc5, 0x48, 0x07, 0x8e, 0xe7, 0xe2, 0xda, 0xfe, 0x07, 0xc0, 0x9d, 0xb1, 0x48, 0x47,
	0xe9, 0x30, 0xe1, 0x79, 0x57, 0x64, 0xd3, 0x31, 0x23, 0xeb, 0xb0, 0xe2, 0xd3, 0x23, 0x2f, 0x08,
	0xfd, 0x0b, 0x6c, 0x10, 0x80, 0x46, 0xe0, 0x1c, 0xd3, 0xae, 0x8d, 0x91, 0x8e, 0x3b, 0x91, 0x73,
	0x42, 0x43, 0x5c, 0xdb, 0xff, 0x8a, 0x60, 0xb3, 0x2b, 0x94, 0x28, 0x9d, 0x64, 0x3c, 0x1e, 0x24,
	0xe9, 0xa8, 0xcb, 0xa4, 0x4c, 0xae, 0xaa, 0xf6, 0x6e, 0xe8, 0xc7, 0x5e, 0xcf, 0x0b, 0xb1, 0x41,
	0xb6, 0xe1, 0xa9, 0xce, 0x0e, 0xed, 0xc8, 0xa1, 0x61, 0xec, 0x53, 0x87, 0x7a, 0xe7, 0x54, 0xff,
	0xc2, 0x16, 0x10, 0x2d, 0x82, 0x63, 0xdb, 0x77, 0x83, 0xd8, 0xf1, 0xa9, 0x1d, 0x52, 0x17, 0xd7,
	0xc8, 0x33, 0xd8, 0xd0, 0xfc, 0xdc, 0x8e, 0x4e, 0xc3, 0x05, 0xae, 0x3f, 0xcc, 0xb1, 0x1d, 0xe7,
	0x2c, 0xea, 0x3d, 0x8a, 0x25, 0xfd, 0xac, 0x5a, 0x9c, 0x9e, 0x1d, 0x1d, 0x51, 0x37, 0xf6, 0x7a,
	0x78, 0xb9, 0x33, 0xf8, 0x71, 0x6b, 0xa2, 0x9b, 0x5b, 0x13, 0xfd, 0xba, 0x35, 0xd1, 0x97, 0x3b,
	0xd3, 0xb8, 0xb9, 0x33, 0x8d, 0x9f, 0x77, 0xa6, 0x01, 0x9b, 0x5c, 0x54, 0xe7, 0x64, 0xa9, 0x59,
	0xc1, 0xe4, 0xfc, 0xa4, 0xfa, 0xe8, 0xd3, 0x9b, 0x2b, 0xae, 0x86, 0xd3, 0x81, 0x06, 0x6d, 0xed,
	0x5f, 0x73, 0x51, 0x7d, 0xdb, 0x6a, 0xc8, 0xcb, 0x2c, 0x2e, 0x92, 0x52, 0xcd, 0xda, 0x55, 0x4f,
	0xfb, 0xbe, 0x67, 0xd0, 0xa8, 0x6e, 0xfe, 0xdd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x65,
	0x20, 0x82, 0x0b, 0x03, 0x00, 0x00,
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PeerId) > 0 {
		i -= len(m.PeerId)
		copy(dAtA[i:], m.PeerId)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.PeerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AesPskKey) > 0 {
		i -= len(m.AesPskKey)
		copy(dAtA[i:], m.AesPskKey)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.AesPskKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AesDscKey) > 0 {
		i -= len(m.AesDscKey)
		copy(dAtA[i:], m.AesDscKey)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.AesDscKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovInfo(uint64(m.Status))
	}
	return n
}

func (m *AuthInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.AesDscKey)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.AesPskKey)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovInfo(uint64(m.Timestamp))
	}
	return n
}

func sovInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Peer_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: AuthInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AesDscKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AesDscKey = append(m.AesDscKey[:0], dAtA[iNdEx:postIndex]...)
			if m.AesDscKey == nil {
				m.AesDscKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AesPskKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AesPskKey = append(m.AesPskKey[:0], dAtA[iNdEx:postIndex]...)
			if m.AesPskKey == nil {
				m.AesPskKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
				return 0, ErrInvalidLengthInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInfo = fmt.Errorf("proto: unexpected end of group")
)
