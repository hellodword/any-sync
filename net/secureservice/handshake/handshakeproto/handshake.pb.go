// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: net/secureservice/handshake/handshakeproto/protos/handshake.proto

package handshakeproto

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

type CredentialsType int32

const (
	// SkipVerify using when identity is not required, for example in p2p cases
	CredentialsType_SkipVerify CredentialsType = 0
	// SignedPeerIds using a payload containing PayloadSignedPeerIds message
	CredentialsType_SignedPeerIds CredentialsType = 1
)

var CredentialsType_name = map[int32]string{
	0: "SkipVerify",
	1: "SignedPeerIds",
}

var CredentialsType_value = map[string]int32{
	"SkipVerify":    0,
	"SignedPeerIds": 1,
}

func (x CredentialsType) String() string {
	return proto.EnumName(CredentialsType_name, int32(x))
}

func (CredentialsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60283fc75f020893, []int{0}
}

type Error int32

const (
	Error_Null                 Error = 0
	Error_Unexpected           Error = 1
	Error_InvalidCredentials   Error = 2
	Error_UnexpectedPayload    Error = 3
	Error_SkipVerifyNotAllowed Error = 4
	Error_DeadlineExceeded     Error = 5
	Error_IncompatibleVersion  Error = 6
	Error_IncompatibleProto    Error = 7
)

var Error_name = map[int32]string{
	0: "Null",
	1: "Unexpected",
	2: "InvalidCredentials",
	3: "UnexpectedPayload",
	4: "SkipVerifyNotAllowed",
	5: "DeadlineExceeded",
	6: "IncompatibleVersion",
	7: "IncompatibleProto",
}

var Error_value = map[string]int32{
	"Null":                 0,
	"Unexpected":           1,
	"InvalidCredentials":   2,
	"UnexpectedPayload":    3,
	"SkipVerifyNotAllowed": 4,
	"DeadlineExceeded":     5,
	"IncompatibleVersion":  6,
	"IncompatibleProto":    7,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}

func (Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60283fc75f020893, []int{1}
}

type ProtoType int32

const (
	ProtoType_DRPC ProtoType = 0
)

var ProtoType_name = map[int32]string{
	0: "DRPC",
}

var ProtoType_value = map[string]int32{
	"DRPC": 0,
}

func (x ProtoType) String() string {
	return proto.EnumName(ProtoType_name, int32(x))
}

func (ProtoType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60283fc75f020893, []int{2}
}

type Credentials struct {
	Type          CredentialsType `protobuf:"varint,1,opt,name=type,proto3,enum=anyHandshake.CredentialsType" json:"type,omitempty"`
	Payload       []byte          `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Version       uint32          `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	ClientVersion string          `protobuf:"bytes,4,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_60283fc75f020893, []int{0}
}
func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(m, src)
}
func (m *Credentials) XXX_Size() int {
	return m.Size()
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetType() CredentialsType {
	if m != nil {
		return m.Type
	}
	return CredentialsType_SkipVerify
}

func (m *Credentials) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Credentials) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Credentials) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

type PayloadSignedPeerIds struct {
	// account identity
	Identity []byte `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// sign of (localPeerId + remotePeerId)
	Sign []byte `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *PayloadSignedPeerIds) Reset()         { *m = PayloadSignedPeerIds{} }
func (m *PayloadSignedPeerIds) String() string { return proto.CompactTextString(m) }
func (*PayloadSignedPeerIds) ProtoMessage()    {}
func (*PayloadSignedPeerIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_60283fc75f020893, []int{1}
}
func (m *PayloadSignedPeerIds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PayloadSignedPeerIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PayloadSignedPeerIds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PayloadSignedPeerIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadSignedPeerIds.Merge(m, src)
}
func (m *PayloadSignedPeerIds) XXX_Size() int {
	return m.Size()
}
func (m *PayloadSignedPeerIds) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadSignedPeerIds.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadSignedPeerIds proto.InternalMessageInfo

func (m *PayloadSignedPeerIds) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *PayloadSignedPeerIds) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type Ack struct {
	Error Error `protobuf:"varint,1,opt,name=error,proto3,enum=anyHandshake.Error" json:"error,omitempty"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_60283fc75f020893, []int{2}
}
func (m *Ack) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(m, src)
}
func (m *Ack) XXX_Size() int {
	return m.Size()
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_Null
}

type Proto struct {
	Proto ProtoType `protobuf:"varint,1,opt,name=proto,proto3,enum=anyHandshake.ProtoType" json:"proto,omitempty"`
}

func (m *Proto) Reset()         { *m = Proto{} }
func (m *Proto) String() string { return proto.CompactTextString(m) }
func (*Proto) ProtoMessage()    {}
func (*Proto) Descriptor() ([]byte, []int) {
	return fileDescriptor_60283fc75f020893, []int{3}
}
func (m *Proto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proto.Merge(m, src)
}
func (m *Proto) XXX_Size() int {
	return m.Size()
}
func (m *Proto) XXX_DiscardUnknown() {
	xxx_messageInfo_Proto.DiscardUnknown(m)
}

var xxx_messageInfo_Proto proto.InternalMessageInfo

func (m *Proto) GetProto() ProtoType {
	if m != nil {
		return m.Proto
	}
	return ProtoType_DRPC
}

func init() {
	proto.RegisterEnum("anyHandshake.CredentialsType", CredentialsType_name, CredentialsType_value)
	proto.RegisterEnum("anyHandshake.Error", Error_name, Error_value)
	proto.RegisterEnum("anyHandshake.ProtoType", ProtoType_name, ProtoType_value)
	proto.RegisterType((*Credentials)(nil), "anyHandshake.Credentials")
	proto.RegisterType((*PayloadSignedPeerIds)(nil), "anyHandshake.PayloadSignedPeerIds")
	proto.RegisterType((*Ack)(nil), "anyHandshake.Ack")
	proto.RegisterType((*Proto)(nil), "anyHandshake.Proto")
}

func init() {
	proto.RegisterFile("net/secureservice/handshake/handshakeproto/protos/handshake.proto", fileDescriptor_60283fc75f020893)
}

var fileDescriptor_60283fc75f020893 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xf5, 0x36, 0x71, 0x3f, 0x86, 0xa4, 0x6c, 0xb7, 0x29, 0xb5, 0x90, 0xb0, 0xa2, 0x88, 0x43,
	0x88, 0x44, 0xc2, 0x97, 0xb8, 0x87, 0xa6, 0x88, 0x5c, 0xaa, 0xc8, 0x85, 0x1e, 0xb8, 0x6d, 0xbd,
	0x43, 0xbb, 0xca, 0xb2, 0xb6, 0xd6, 0xdb, 0x50, 0xff, 0x0b, 0xce, 0xfc, 0x0a, 0x7e, 0x06, 0xc7,
	0x1e, 0x39, 0xa2, 0xe4, 0x8f, 0x20, 0x6f, 0x9c, 0xc6, 0xe1, 0xc4, 0xc5, 0xde, 0x99, 0xf7, 0x66,
	0xe7, 0xbd, 0x67, 0xc3, 0x50, 0xa3, 0x1d, 0x64, 0x18, 0xdf, 0x18, 0xcc, 0xd0, 0xcc, 0x64, 0x8c,
	0x83, 0x6b, 0xae, 0x45, 0x76, 0xcd, 0xa7, 0x95, 0x53, 0x6a, 0x12, 0x9b, 0x0c, 0xdc, 0x33, 0x5b,
	0x77, 0xfb, 0xae, 0xc1, 0x1a, 0x5c, 0xe7, 0x1f, 0x56, 0xbd, 0xce, 0x0f, 0x02, 0x0f, 0x4e, 0x0c,
	0x0a, 0xd4, 0x56, 0x72, 0x95, 0xb1, 0x97, 0x50, 0xb7, 0x79, 0x8a, 0x01, 0x69, 0x93, 0xee, 0xfe,
	0xab, 0x27, 0xfd, 0x2a, 0xb9, 0x5f, 0x21, 0x7e, 0xcc, 0x53, 0x8c, 0x1c, 0x95, 0x05, 0xb0, 0x93,
	0xf2, 0x5c, 0x25, 0x5c, 0x04, 0x5b, 0x6d, 0xd2, 0x6d, 0x44, 0xab, 0xb2, 0x40, 0x66, 0x68, 0x32,
	0x99, 0xe8, 0xa0, 0xd6, 0x26, 0xdd, 0x66, 0xb4, 0x2a, 0xd9, 0x53, 0x68, 0xc6, 0x4a, 0xa2, 0xb6,
	0x17, 0x25, 0x5e, 0x6f, 0x93, 0xee, 0x5e, 0xb4, 0xd9, 0xec, 0xbc, 0x87, 0xd6, 0x64, 0x79, 0xd5,
	0xb9, 0xbc, 0xd2, 0x28, 0x26, 0x88, 0x66, 0x2c, 0x32, 0xf6, 0x18, 0x76, 0xa5, 0x13, 0x62, 0x73,
	0x27, 0xb4, 0x11, 0xdd, 0xd7, 0x8c, 0x41, 0x3d, 0x93, 0x57, 0xba, 0x94, 0xe2, 0xce, 0x9d, 0x17,
	0x50, 0x1b, 0xc6, 0x53, 0xf6, 0x0c, 0x7c, 0x34, 0x26, 0x31, 0xa5, 0xb9, 0xc3, 0x4d, 0x73, 0xa7,
	0x05, 0x14, 0x2d, 0x19, 0x9d, 0xb7, 0xe0, 0x4f, 0x5c, 0x5a, 0xcf, 0xc1, 0x77, 0xb1, 0x95, 0x33,
	0xc7, 0x9b, 0x33, 0x8e, 0xe3, 0xa2, 0x58, 0xb2, 0x7a, 0x6f, 0xe0, 0xe1, 0x3f, 0x21, 0xb1, 0x7d,
	0x80, 0xf3, 0xa9, 0x4c, 0x2f, 0xd0, 0xc8, 0x2f, 0x39, 0xf5, 0xd8, 0x01, 0x34, 0x37, 0xdc, 0x50,
	0xd2, 0xfb, 0x49, 0xc0, 0x77, 0xeb, 0xd9, 0x2e, 0xd4, 0xcf, 0x6e, 0x94, 0xa2, 0x5e, 0x31, 0xf6,
	0x49, 0xe3, 0x6d, 0x8a, 0xb1, 0x45, 0x41, 0x09, 0x7b, 0x04, 0x6c, 0xac, 0x67, 0x5c, 0x49, 0x51,
	0x59, 0x40, 0xb7, 0xd8, 0x11, 0x1c, 0xac, 0x79, 0x65, 0x5a, 0xb4, 0xc6, 0x02, 0x68, 0xad, 0xb7,
	0x9e, 0x25, 0x76, 0xa8, 0x54, 0xf2, 0x0d, 0x05, 0xad, 0xb3, 0x16, 0xd0, 0x11, 0x72, 0xa1, 0xa4,
	0xc6, 0xd3, 0xdb, 0x18, 0x51, 0xa0, 0xa0, 0x3e, 0x3b, 0x86, 0xc3, 0xb1, 0x8e, 0x93, 0xaf, 0x29,
	0xb7, 0xf2, 0x52, 0x61, 0xf9, 0x05, 0xe8, 0x76, 0x71, 0x7f, 0x15, 0x70, 0x8e, 0xe9, 0x4e, 0xef,
	0x08, 0xf6, 0xee, 0xcd, 0x17, 0xaa, 0x47, 0xd1, 0xe4, 0x84, 0x7a, 0xef, 0x46, 0xbf, 0xe6, 0x21,
	0xb9, 0x9b, 0x87, 0xe4, 0xcf, 0x3c, 0x24, 0xdf, 0x17, 0xa1, 0x77, 0xb7, 0x08, 0xbd, 0xdf, 0x8b,
	0xd0, 0xfb, 0xdc, 0xfb, 0xff, 0x3f, 0xf7, 0x72, 0xdb, 0xbd, 0x5e, 0xff, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xe2, 0xfa, 0x40, 0x67, 0xee, 0x02, 0x00, 0x00,
}

func (m *Credentials) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credentials) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Credentials) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientVersion) > 0 {
		i -= len(m.ClientVersion)
		copy(dAtA[i:], m.ClientVersion)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.ClientVersion)))
		i--
		dAtA[i] = 0x22
	}
	if m.Version != 0 {
		i = encodeVarintHandshake(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintHandshake(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PayloadSignedPeerIds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PayloadSignedPeerIds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PayloadSignedPeerIds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sign) > 0 {
		i -= len(m.Sign)
		copy(dAtA[i:], m.Sign)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.Sign)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Ack) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ack) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ack) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Error != 0 {
		i = encodeVarintHandshake(dAtA, i, uint64(m.Error))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Proto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proto != 0 {
		i = encodeVarintHandshake(dAtA, i, uint64(m.Proto))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHandshake(dAtA []byte, offset int, v uint64) int {
	offset -= sovHandshake(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Credentials) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovHandshake(uint64(m.Type))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovHandshake(uint64(m.Version))
	}
	l = len(m.ClientVersion)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	return n
}

func (m *PayloadSignedPeerIds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	return n
}

func (m *Ack) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Error != 0 {
		n += 1 + sovHandshake(uint64(m.Error))
	}
	return n
}

func (m *Proto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proto != 0 {
		n += 1 + sovHandshake(uint64(m.Proto))
	}
	return n
}

func sovHandshake(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHandshake(x uint64) (n int) {
	return sovHandshake(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Credentials) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: Credentials: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credentials: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= CredentialsType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
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
func (m *PayloadSignedPeerIds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: PayloadSignedPeerIds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PayloadSignedPeerIds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = append(m.Identity[:0], dAtA[iNdEx:postIndex]...)
			if m.Identity == nil {
				m.Identity = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = append(m.Sign[:0], dAtA[iNdEx:postIndex]...)
			if m.Sign == nil {
				m.Sign = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
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
func (m *Ack) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: Ack: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ack: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Error |= Error(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
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
func (m *Proto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: Proto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proto", wireType)
			}
			m.Proto = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Proto |= ProtoType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
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
func skipHandshake(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHandshake
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
					return 0, ErrIntOverflowHandshake
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
					return 0, ErrIntOverflowHandshake
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
				return 0, ErrInvalidLengthHandshake
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHandshake
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHandshake
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHandshake        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHandshake          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHandshake = fmt.Errorf("proto: unexpected end of group")
)
