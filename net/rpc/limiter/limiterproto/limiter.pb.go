// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: net/rpc/limiter/limiterproto/protos/limiter.proto

package limiterproto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type ErrCodes int32

const (
	ErrCodes_RateLimitExceeded ErrCodes = 0
	ErrCodes_ErrorOffset       ErrCodes = 700
)

var ErrCodes_name = map[int32]string{
	0:   "RateLimitExceeded",
	700: "ErrorOffset",
}

var ErrCodes_value = map[string]int32{
	"RateLimitExceeded": 0,
	"ErrorOffset":       700,
}

func (x ErrCodes) String() string {
	return proto.EnumName(ErrCodes_name, int32(x))
}

func (ErrCodes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_43f29163996a95d8, []int{0}
}

func init() {
	proto.RegisterEnum("limiter.ErrCodes", ErrCodes_name, ErrCodes_value)
}

func init() {
	proto.RegisterFile("net/rpc/limiter/limiterproto/protos/limiter.proto", fileDescriptor_43f29163996a95d8)
}

var fileDescriptor_43f29163996a95d8 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xcc, 0x4b, 0x2d, 0xd1,
	0x2f, 0x2a, 0x48, 0xd6, 0xcf, 0xc9, 0xcc, 0xcd, 0x2c, 0x49, 0x2d, 0x82, 0xd1, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0xfa, 0x60, 0xb2, 0x18, 0x26, 0xa6, 0x07, 0xe6, 0x0a, 0xb1, 0x43, 0xb9, 0x5a, 0xc6,
	0x5c, 0x1c, 0xae, 0x45, 0x45, 0xce, 0xf9, 0x29, 0xa9, 0xc5, 0x42, 0xa2, 0x5c, 0x82, 0x41, 0x89,
	0x25, 0xa9, 0x3e, 0x20, 0x29, 0xd7, 0x8a, 0xe4, 0xd4, 0xd4, 0x94, 0xd4, 0x14, 0x01, 0x06, 0x21,
	0x01, 0x2e, 0x6e, 0xd7, 0xa2, 0xa2, 0xfc, 0x22, 0xff, 0xb4, 0xb4, 0xe2, 0xd4, 0x12, 0x81, 0x3d,
	0xac, 0x4e, 0x66, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x83, 0xcf,
	0x29, 0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x6b, 0xb3, 0x13, 0xb1,
	0x00, 0x00, 0x00,
}
