// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/app_store.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// App store type in an app extension.
type AppStoreEnum_AppStore int32

const (
	// Not specified.
	AppStoreEnum_UNSPECIFIED AppStoreEnum_AppStore = 0
	// Used for return value only. Represents value unknown in this version.
	AppStoreEnum_UNKNOWN AppStoreEnum_AppStore = 1
	// Apple iTunes.
	AppStoreEnum_APPLE_ITUNES AppStoreEnum_AppStore = 2
	// Google Play.
	AppStoreEnum_GOOGLE_PLAY AppStoreEnum_AppStore = 3
)

var AppStoreEnum_AppStore_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "APPLE_ITUNES",
	3: "GOOGLE_PLAY",
}

var AppStoreEnum_AppStore_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"APPLE_ITUNES": 2,
	"GOOGLE_PLAY":  3,
}

func (x AppStoreEnum_AppStore) String() string {
	return proto.EnumName(AppStoreEnum_AppStore_name, int32(x))
}

func (AppStoreEnum_AppStore) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e712fee7ba2d1fd3, []int{0, 0}
}

// Container for enum describing app store type in an app extension.
type AppStoreEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppStoreEnum) Reset()         { *m = AppStoreEnum{} }
func (m *AppStoreEnum) String() string { return proto.CompactTextString(m) }
func (*AppStoreEnum) ProtoMessage()    {}
func (*AppStoreEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e712fee7ba2d1fd3, []int{0}
}

func (m *AppStoreEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppStoreEnum.Unmarshal(m, b)
}
func (m *AppStoreEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppStoreEnum.Marshal(b, m, deterministic)
}
func (m *AppStoreEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppStoreEnum.Merge(m, src)
}
func (m *AppStoreEnum) XXX_Size() int {
	return xxx_messageInfo_AppStoreEnum.Size(m)
}
func (m *AppStoreEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppStoreEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppStoreEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AppStoreEnum_AppStore", AppStoreEnum_AppStore_name, AppStoreEnum_AppStore_value)
	proto.RegisterType((*AppStoreEnum)(nil), "google.ads.googleads.v2.enums.AppStoreEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/app_store.proto", fileDescriptor_e712fee7ba2d1fd3)
}

var fileDescriptor_e712fee7ba2d1fd3 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x5d, 0x07, 0x2a, 0xd9, 0xc4, 0xd2, 0xa3, 0xb8, 0xc3, 0x76, 0x37, 0x81, 0x7a, 0x8b,
	0xa7, 0x4c, 0x63, 0x19, 0x1b, 0x5d, 0x60, 0x6e, 0xa2, 0x16, 0x46, 0xb4, 0x25, 0x14, 0xd6, 0x24,
	0x34, 0xdd, 0x1e, 0xc8, 0xa3, 0x8f, 0xe2, 0x63, 0x78, 0xf4, 0x29, 0x24, 0x89, 0xed, 0x4d, 0x2f,
	0xe1, 0x4b, 0xbe, 0x5f, 0xbe, 0x7c, 0xf9, 0x83, 0x2b, 0xa1, 0x94, 0xd8, 0x15, 0x88, 0xe7, 0x06,
	0x79, 0x69, 0xd5, 0x21, 0x46, 0x85, 0xdc, 0x57, 0x06, 0x71, 0xad, 0xb7, 0xa6, 0x51, 0x75, 0x01,
	0x75, 0xad, 0x1a, 0x15, 0x8d, 0x3c, 0x03, 0x79, 0x6e, 0x60, 0x87, 0xc3, 0x43, 0x0c, 0x1d, 0x7e,
	0x71, 0xd9, 0xa6, 0xe9, 0x12, 0x71, 0x29, 0x55, 0xc3, 0x9b, 0x52, 0x49, 0xe3, 0x2f, 0x4f, 0x5e,
	0xc0, 0x90, 0x68, 0xbd, 0xb2, 0x71, 0x54, 0xee, 0xab, 0xc9, 0x1c, 0x9c, 0xb6, 0xfb, 0xe8, 0x1c,
	0x0c, 0xd6, 0xe9, 0x8a, 0xd1, 0xdb, 0xd9, 0xfd, 0x8c, 0xde, 0x85, 0x47, 0xd1, 0x00, 0x9c, 0xac,
	0xd3, 0x79, 0xba, 0x7c, 0x4c, 0xc3, 0x5e, 0x14, 0x82, 0x21, 0x61, 0x6c, 0x41, 0xb7, 0xb3, 0x87,
	0x75, 0x4a, 0x57, 0x61, 0x60, 0xf9, 0x64, 0xb9, 0x4c, 0x16, 0x74, 0xcb, 0x16, 0xe4, 0x29, 0xec,
	0x4f, 0xbf, 0x7a, 0x60, 0xfc, 0xa6, 0x2a, 0xf8, 0x6f, 0xc1, 0xe9, 0x59, 0xfb, 0x20, 0xb3, 0x8d,
	0x58, 0xef, 0x79, 0xfa, 0xcb, 0x0b, 0xb5, 0xe3, 0x52, 0x40, 0x55, 0x0b, 0x24, 0x0a, 0xe9, 0xfa,
	0xb6, 0xf3, 0xd0, 0xa5, 0xf9, 0x63, 0x3c, 0x37, 0x6e, 0x7d, 0x0f, 0xfa, 0x09, 0x21, 0x1f, 0xc1,
	0x28, 0xf1, 0x51, 0x24, 0x37, 0xd0, 0x4b, 0xab, 0x36, 0x31, 0xb4, 0x9f, 0x35, 0x9f, 0xad, 0x9f,
	0x91, 0xdc, 0x64, 0x9d, 0x9f, 0x6d, 0xe2, 0xcc, 0xf9, 0xdf, 0xc1, 0xd8, 0x1f, 0x62, 0x4c, 0x72,
	0x83, 0x71, 0x47, 0x60, 0xbc, 0x89, 0x31, 0x76, 0xcc, 0xeb, 0xb1, 0x2b, 0x76, 0xfd, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xef, 0x21, 0x1f, 0x60, 0xb6, 0x01, 0x00, 0x00,
}
