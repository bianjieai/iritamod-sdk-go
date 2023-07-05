// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: perm/perm.proto

package perm

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Role represents a role
type Role int32

const (
	// ROOT_ADMIN defines the root admin role index.
	RoleRootAdmin Role = 0
	// PERM_ADMIN defines the permission admin role index.
	RolePermAdmin Role = 1
	// BLACKLIST_ADMIN defines the blacklist admin role index.
	RoleBlacklistAdmin Role = 2
	// NODE_ADMIN defines the node admin role index.
	RoleNodeAdmin Role = 3
	// PARAM_ADMIN defines the param admin role index.
	RoleParamAdmin Role = 4
	// POWER_USER defines the power user role index.
	RolePowerUser Role = 5
	// RELAYER_USER defines the relayer user role index.
	RoleRelayerUser Role = 6
	// ID_ADMIN defines the identity admin role index.
	RoleIDAdmin Role = 7
	// BASE_M1_ADMIN defines the base M1 admin role index.
	RoleBaseM1Admin Role = 8
	// Chain_Account_Role defines the platform admin role index.
	RolePlatformUser Role = 9
	// POWER_USER_ADMIN defines the power admin role index.
	RolePowerUserAdmin Role = 10
	//LAYER2_USER defines the layer2 user role index.
	RoleLayer2User Role = 11
)

var Role_name = map[int32]string{
	0:  "ROOT_ADMIN",
	1:  "PERM_ADMIN",
	2:  "BLACKLIST_ADMIN",
	3:  "NODE_ADMIN",
	4:  "PARAM_ADMIN",
	5:  "POWER_USER",
	6:  "RELAYER_USER",
	7:  "ID_ADMIN",
	8:  "BASE_M1_ADMIN",
	9:  "PLATFORM_USER",
	10: "POWER_USER_ADMIN",
	11: "LAYER2_USER",
}

var Role_value = map[string]int32{
	"ROOT_ADMIN":       0,
	"PERM_ADMIN":       1,
	"BLACKLIST_ADMIN":  2,
	"NODE_ADMIN":       3,
	"PARAM_ADMIN":      4,
	"POWER_USER":       5,
	"RELAYER_USER":     6,
	"ID_ADMIN":         7,
	"BASE_M1_ADMIN":    8,
	"PLATFORM_USER":    9,
	"POWER_USER_ADMIN": 10,
	"LAYER2_USER":      11,
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb77ba30a3a45e51, []int{0}
}

func init() {
	proto.RegisterEnum("iritamod.perm.Role", Role_name, Role_value)
	golang_proto.RegisterEnum("iritamod.perm.Role", Role_name, Role_value)
}

func init() { proto.RegisterFile("perm/perm.proto", fileDescriptor_bb77ba30a3a45e51) }
func init() { golang_proto.RegisterFile("perm/perm.proto", fileDescriptor_bb77ba30a3a45e51) }

var fileDescriptor_bb77ba30a3a45e51 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x13, 0x36, 0xc6, 0x70, 0x57, 0x5a, 0xc2, 0xc4, 0xc1, 0x12, 0x56, 0x10, 0x02, 0xc4,
	0x9f, 0xb5, 0xda, 0xf8, 0x04, 0x0e, 0x0d, 0x52, 0xb5, 0xa4, 0xa9, 0xdc, 0x4d, 0x08, 0x2e, 0x95,
	0xbb, 0x9a, 0x60, 0x96, 0xcc, 0x93, 0x13, 0x84, 0xf8, 0x06, 0x28, 0x27, 0xbe, 0x40, 0x4e, 0xec,
	0xc0, 0x91, 0x8f, 0xc0, 0x71, 0xc7, 0x1d, 0x39, 0x70, 0x80, 0xe6, 0x8b, 0x20, 0x3b, 0x49, 0xcd,
	0x25, 0x8a, 0xfc, 0xfe, 0xf4, 0x3c, 0xcf, 0xfb, 0xea, 0x01, 0xbd, 0x73, 0x26, 0xd3, 0xa1, 0xfa,
	0x0c, 0xce, 0xa5, 0xc8, 0x85, 0xd3, 0xe5, 0x92, 0xe7, 0x34, 0x15, 0xcb, 0x81, 0x7a, 0x84, 0xbb,
	0xb1, 0x88, 0x85, 0x9e, 0x0c, 0xd5, 0x5f, 0x0d, 0x3d, 0xfd, 0xbd, 0x01, 0x36, 0x89, 0x48, 0x98,
	0x73, 0x1f, 0x00, 0x12, 0x45, 0x47, 0x73, 0x3c, 0x0a, 0xc7, 0x93, 0xbe, 0x05, 0x6f, 0x17, 0xa5,
	0xdb, 0x55, 0x13, 0x22, 0x44, 0x8e, 0x97, 0x29, 0x3f, 0x53, 0xc8, 0xd4, 0x27, 0x61, 0x83, 0xd8,
	0x06, 0x99, 0x32, 0x99, 0xd6, 0xc8, 0x33, 0xd0, 0xf3, 0x02, 0xfc, 0xf2, 0x30, 0x18, 0xcf, 0x5a,
	0xa9, 0x6b, 0xf0, 0x6e, 0x51, 0xba, 0x8e, 0xe2, 0xbc, 0x84, 0x9e, 0x9c, 0x26, 0x3c, 0x33, 0x7a,
	0x93, 0x68, 0xe4, 0x37, 0xdc, 0x86, 0xd1, 0x9b, 0x88, 0x25, 0xab, 0x91, 0x07, 0xa0, 0x33, 0xc5,
	0x04, 0xb7, 0x9e, 0x9b, 0xd0, 0x29, 0x4a, 0xf7, 0x96, 0xf6, 0xa4, 0x92, 0xa6, 0x26, 0x57, 0xf4,
	0xda, 0x27, 0xf3, 0xe3, 0x99, 0x4f, 0xfa, 0xd7, 0xff, 0xcb, 0x25, 0x3e, 0x31, 0x79, 0x9c, 0x31,
	0xe9, 0x3c, 0x04, 0x3b, 0xc4, 0x0f, 0xf0, 0x9b, 0x16, 0xda, 0x82, 0x77, 0x8a, 0xd2, 0xed, 0xe9,
	0xfd, 0x58, 0x42, 0x3f, 0x37, 0xd8, 0x3d, 0xb0, 0x3d, 0x1e, 0x35, 0x5e, 0x37, 0x60, 0xaf, 0x28,
	0xdd, 0x8e, 0x42, 0xc6, 0xa3, 0xda, 0xe8, 0x11, 0xe8, 0x7a, 0x78, 0xe6, 0xcf, 0xc3, 0xfd, 0x86,
	0xd9, 0x36, 0x32, 0x1e, 0xcd, 0x58, 0xb8, 0x5f, 0x73, 0x8f, 0x41, 0x77, 0x1a, 0xe0, 0xa3, 0x57,
	0x11, 0x09, 0x6b, 0xbb, 0x9b, 0x70, 0xb7, 0x28, 0xdd, 0xbe, 0xce, 0x94, 0xd0, 0xfc, 0x9d, 0x90,
	0xa9, 0xf6, 0x7b, 0x0e, 0xfa, 0x26, 0x79, 0xa3, 0x09, 0xcc, 0xbd, 0xd6, 0xf9, 0xd7, 0xc7, 0xd0,
	0x2b, 0x1c, 0xd4, 0xa2, 0x1d, 0x73, 0x8c, 0x40, 0x6d, 0x70, 0xa0, 0x48, 0xb8, 0xf3, 0xe5, 0x1b,
	0xb2, 0xbe, 0x5f, 0x20, 0xeb, 0xc7, 0x05, 0xb2, 0xbd, 0xc3, 0xcb, 0xbf, 0xc8, 0xba, 0x5c, 0x21,
	0xfb, 0x6a, 0x85, 0xec, 0x3f, 0x2b, 0x64, 0x7f, 0xad, 0x90, 0xf5, 0xb3, 0x42, 0xf6, 0x55, 0x85,
	0xac, 0x5f, 0x15, 0xb2, 0xde, 0x3e, 0x89, 0x79, 0xfe, 0xfe, 0xe3, 0x62, 0x70, 0x22, 0xd2, 0xe1,
	0x82, 0xd3, 0xb3, 0x0f, 0x9c, 0x51, 0x3e, 0x6c, 0xab, 0xb3, 0x97, 0x2d, 0x4f, 0xf7, 0x62, 0xa1,
	0x6b, 0xb5, 0xd8, 0xd2, 0x95, 0x79, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xd2, 0xd0, 0x6d,
	0x6a, 0x02, 0x00, 0x00,
}

func (x Role) String() string {
	s, ok := Role_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
