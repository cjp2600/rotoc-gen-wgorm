// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin/options/worm.proto

package worm

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type WormFileOptions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WormFileOptions) Reset()         { *m = WormFileOptions{} }
func (m *WormFileOptions) String() string { return proto.CompactTextString(m) }
func (*WormFileOptions) ProtoMessage()    {}
func (*WormFileOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c056d40fbc59afe5, []int{0}
}

func (m *WormFileOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WormFileOptions.Unmarshal(m, b)
}
func (m *WormFileOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WormFileOptions.Marshal(b, m, deterministic)
}
func (m *WormFileOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WormFileOptions.Merge(m, src)
}
func (m *WormFileOptions) XXX_Size() int {
	return xxx_messageInfo_WormFileOptions.Size(m)
}
func (m *WormFileOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_WormFileOptions.DiscardUnknown(m)
}

var xxx_messageInfo_WormFileOptions proto.InternalMessageInfo

type WormMessageOptions struct {
	Model                *bool    `protobuf:"varint,1,req,name=model" json:"model,omitempty"`
	Table                *string  `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
	Merge                *string  `protobuf:"bytes,4,opt,name=merge" json:"merge,omitempty"`
	Migrate              *bool    `protobuf:"varint,3,opt,name=migrate" json:"migrate,omitempty"`
	ConvertTo            *string  `protobuf:"bytes,5,opt,name=convertTo" json:"convertTo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WormMessageOptions) Reset()         { *m = WormMessageOptions{} }
func (m *WormMessageOptions) String() string { return proto.CompactTextString(m) }
func (*WormMessageOptions) ProtoMessage()    {}
func (*WormMessageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c056d40fbc59afe5, []int{1}
}

func (m *WormMessageOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WormMessageOptions.Unmarshal(m, b)
}
func (m *WormMessageOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WormMessageOptions.Marshal(b, m, deterministic)
}
func (m *WormMessageOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WormMessageOptions.Merge(m, src)
}
func (m *WormMessageOptions) XXX_Size() int {
	return xxx_messageInfo_WormMessageOptions.Size(m)
}
func (m *WormMessageOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_WormMessageOptions.DiscardUnknown(m)
}

var xxx_messageInfo_WormMessageOptions proto.InternalMessageInfo

func (m *WormMessageOptions) GetModel() bool {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return false
}

func (m *WormMessageOptions) GetTable() string {
	if m != nil && m.Table != nil {
		return *m.Table
	}
	return ""
}

func (m *WormMessageOptions) GetMerge() string {
	if m != nil && m.Merge != nil {
		return *m.Merge
	}
	return ""
}

func (m *WormMessageOptions) GetMigrate() bool {
	if m != nil && m.Migrate != nil {
		return *m.Migrate
	}
	return false
}

func (m *WormMessageOptions) GetConvertTo() string {
	if m != nil && m.ConvertTo != nil {
		return *m.ConvertTo
	}
	return ""
}

type WormFieldOptions struct {
	Tag                  *WormTag `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WormFieldOptions) Reset()         { *m = WormFieldOptions{} }
func (m *WormFieldOptions) String() string { return proto.CompactTextString(m) }
func (*WormFieldOptions) ProtoMessage()    {}
func (*WormFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c056d40fbc59afe5, []int{2}
}

func (m *WormFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WormFieldOptions.Unmarshal(m, b)
}
func (m *WormFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WormFieldOptions.Marshal(b, m, deterministic)
}
func (m *WormFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WormFieldOptions.Merge(m, src)
}
func (m *WormFieldOptions) XXX_Size() int {
	return xxx_messageInfo_WormFieldOptions.Size(m)
}
func (m *WormFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_WormFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_WormFieldOptions proto.InternalMessageInfo

func (m *WormFieldOptions) GetTag() *WormTag {
	if m != nil {
		return m.Tag
	}
	return nil
}

type WormTag struct {
	Gorm                 *string  `protobuf:"bytes,3,opt,name=gorm" json:"gorm,omitempty"`
	Validator            *string  `protobuf:"bytes,4,opt,name=validator" json:"validator,omitempty"`
	Jsonb                *bool    `protobuf:"varint,5,opt,name=jsonb" json:"jsonb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WormTag) Reset()         { *m = WormTag{} }
func (m *WormTag) String() string { return proto.CompactTextString(m) }
func (*WormTag) ProtoMessage()    {}
func (*WormTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c056d40fbc59afe5, []int{3}
}

func (m *WormTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WormTag.Unmarshal(m, b)
}
func (m *WormTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WormTag.Marshal(b, m, deterministic)
}
func (m *WormTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WormTag.Merge(m, src)
}
func (m *WormTag) XXX_Size() int {
	return xxx_messageInfo_WormTag.Size(m)
}
func (m *WormTag) XXX_DiscardUnknown() {
	xxx_messageInfo_WormTag.DiscardUnknown(m)
}

var xxx_messageInfo_WormTag proto.InternalMessageInfo

func (m *WormTag) GetGorm() string {
	if m != nil && m.Gorm != nil {
		return *m.Gorm
	}
	return ""
}

func (m *WormTag) GetValidator() string {
	if m != nil && m.Validator != nil {
		return *m.Validator
	}
	return ""
}

func (m *WormTag) GetJsonb() bool {
	if m != nil && m.Jsonb != nil {
		return *m.Jsonb
	}
	return false
}

type Pagination struct {
	TotalCount           *int32   `protobuf:"varint,1,req,name=totalCount" json:"totalCount,omitempty"`
	TotalPages           *int32   `protobuf:"varint,2,req,name=totalPages" json:"totalPages,omitempty"`
	CurrentPage          *int32   `protobuf:"varint,3,req,name=currentPage" json:"currentPage,omitempty"`
	Size                 *int32   `protobuf:"varint,4,req,name=size" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_c056d40fbc59afe5, []int{4}
}

func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

func (m *Pagination) GetTotalPages() int32 {
	if m != nil && m.TotalPages != nil {
		return *m.TotalPages
	}
	return 0
}

func (m *Pagination) GetCurrentPage() int32 {
	if m != nil && m.CurrentPage != nil {
		return *m.CurrentPage
	}
	return 0
}

func (m *Pagination) GetSize() int32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

type AutoServerOptions struct {
	Autogen              *bool    `protobuf:"varint,1,opt,name=autogen" json:"autogen,omitempty"`
	TxnMiddleware        *bool    `protobuf:"varint,2,opt,name=txn_middleware,json=txnMiddleware" json:"txn_middleware,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoServerOptions) Reset()         { *m = AutoServerOptions{} }
func (m *AutoServerOptions) String() string { return proto.CompactTextString(m) }
func (*AutoServerOptions) ProtoMessage()    {}
func (*AutoServerOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c056d40fbc59afe5, []int{5}
}

func (m *AutoServerOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoServerOptions.Unmarshal(m, b)
}
func (m *AutoServerOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoServerOptions.Marshal(b, m, deterministic)
}
func (m *AutoServerOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoServerOptions.Merge(m, src)
}
func (m *AutoServerOptions) XXX_Size() int {
	return xxx_messageInfo_AutoServerOptions.Size(m)
}
func (m *AutoServerOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoServerOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AutoServerOptions proto.InternalMessageInfo

func (m *AutoServerOptions) GetAutogen() bool {
	if m != nil && m.Autogen != nil {
		return *m.Autogen
	}
	return false
}

func (m *AutoServerOptions) GetTxnMiddleware() bool {
	if m != nil && m.TxnMiddleware != nil {
		return *m.TxnMiddleware
	}
	return false
}

type MethodOptions struct {
	ObjectType           *string  `protobuf:"bytes,1,opt,name=object_type,json=objectType" json:"object_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MethodOptions) Reset()         { *m = MethodOptions{} }
func (m *MethodOptions) String() string { return proto.CompactTextString(m) }
func (*MethodOptions) ProtoMessage()    {}
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c056d40fbc59afe5, []int{6}
}

func (m *MethodOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodOptions.Unmarshal(m, b)
}
func (m *MethodOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodOptions.Marshal(b, m, deterministic)
}
func (m *MethodOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodOptions.Merge(m, src)
}
func (m *MethodOptions) XXX_Size() int {
	return xxx_messageInfo_MethodOptions.Size(m)
}
func (m *MethodOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MethodOptions proto.InternalMessageInfo

func (m *MethodOptions) GetObjectType() string {
	if m != nil && m.ObjectType != nil {
		return *m.ObjectType
	}
	return ""
}

var E_FileOpts = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*WormFileOptions)(nil),
	Field:         99432,
	Name:          "worm.file_opts",
	Tag:           "bytes,99432,opt,name=file_opts",
	Filename:      "plugin/options/worm.proto",
}

var E_Opts = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*WormMessageOptions)(nil),
	Field:         99432,
	Name:          "worm.opts",
	Tag:           "bytes,99432,opt,name=opts",
	Filename:      "plugin/options/worm.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*WormFieldOptions)(nil),
	Field:         99432,
	Name:          "worm.field",
	Tag:           "bytes,99432,opt,name=field",
	Filename:      "plugin/options/worm.proto",
}

var E_Server = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*AutoServerOptions)(nil),
	Field:         99432,
	Name:          "worm.server",
	Tag:           "bytes,99432,opt,name=server",
	Filename:      "plugin/options/worm.proto",
}

var E_Method = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*MethodOptions)(nil),
	Field:         99432,
	Name:          "worm.method",
	Tag:           "bytes,99432,opt,name=method",
	Filename:      "plugin/options/worm.proto",
}

func init() {
	proto.RegisterType((*WormFileOptions)(nil), "worm.WormFileOptions")
	proto.RegisterType((*WormMessageOptions)(nil), "worm.WormMessageOptions")
	proto.RegisterType((*WormFieldOptions)(nil), "worm.WormFieldOptions")
	proto.RegisterType((*WormTag)(nil), "worm.WormTag")
	proto.RegisterType((*Pagination)(nil), "worm.Pagination")
	proto.RegisterType((*AutoServerOptions)(nil), "worm.AutoServerOptions")
	proto.RegisterType((*MethodOptions)(nil), "worm.MethodOptions")
	proto.RegisterExtension(E_FileOpts)
	proto.RegisterExtension(E_Opts)
	proto.RegisterExtension(E_Field)
	proto.RegisterExtension(E_Server)
	proto.RegisterExtension(E_Method)
}

func init() { proto.RegisterFile("plugin/options/worm.proto", fileDescriptor_c056d40fbc59afe5) }

var fileDescriptor_c056d40fbc59afe5 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0x76, 0xfb, 0x91, 0x4e, 0x55, 0x60, 0xcd, 0x97, 0x41, 0xcb, 0x6e, 0x14, 0x09, 0xa9,
	0xa7, 0x16, 0xc1, 0xad, 0x37, 0x84, 0xc4, 0xad, 0xda, 0xc5, 0x54, 0xe2, 0x58, 0xb9, 0x8d, 0x6b,
	0xbc, 0x4a, 0x32, 0x91, 0xe3, 0x74, 0x77, 0x39, 0x72, 0xe2, 0xc4, 0xef, 0xe4, 0x67, 0x20, 0xdb,
	0x4d, 0x93, 0xaa, 0xbd, 0xcd, 0x7b, 0x33, 0x79, 0x79, 0x7e, 0x33, 0xf0, 0xa6, 0x48, 0x2b, 0xa9,
	0xf2, 0x29, 0x16, 0x46, 0x61, 0x5e, 0x4e, 0xef, 0x51, 0x67, 0x93, 0x42, 0xa3, 0x41, 0xd2, 0xb1,
	0xf5, 0xdb, 0x48, 0x22, 0xca, 0x54, 0x4c, 0x1d, 0xb7, 0xaa, 0x36, 0xd3, 0x44, 0x94, 0x6b, 0xad,
	0x0a, 0x83, 0xda, 0xcf, 0xc5, 0x17, 0xf0, 0xf4, 0x07, 0xea, 0xec, 0xab, 0x4a, 0xc5, 0x8d, 0x57,
	0x89, 0xff, 0x06, 0x40, 0x2c, 0x37, 0x17, 0x65, 0xc9, 0x65, 0x4d, 0x93, 0x17, 0xd0, 0xcd, 0x30,
	0x11, 0x29, 0x0d, 0xa2, 0xb3, 0x71, 0xc8, 0x3c, 0xb0, 0xac, 0xe1, 0xab, 0x54, 0xd0, 0xb3, 0x28,
	0x18, 0x0f, 0x98, 0x07, 0x6e, 0x56, 0x68, 0x29, 0x68, 0xc7, 0xb3, 0x0e, 0x10, 0x0a, 0xfd, 0x4c,
	0x49, 0xcd, 0x8d, 0xa0, 0xe7, 0x51, 0x30, 0x0e, 0x59, 0x0d, 0xc9, 0x25, 0x0c, 0xd6, 0x98, 0x6f,
	0x85, 0x36, 0x0b, 0xa4, 0x5d, 0xf7, 0x4d, 0x43, 0xc4, 0x9f, 0xe0, 0x99, 0xf7, 0x28, 0xd2, 0xa4,
	0x76, 0x73, 0x0d, 0xe7, 0x86, 0x4b, 0x1a, 0x44, 0xc1, 0x78, 0xf8, 0x71, 0x34, 0x71, 0x2f, 0xb7,
	0x43, 0x0b, 0x2e, 0x99, 0xed, 0xc4, 0xdf, 0xa0, 0xbf, 0xc3, 0x84, 0x40, 0x47, 0xa2, 0xce, 0xdc,
	0x4f, 0x07, 0xcc, 0xd5, 0xf6, 0x8f, 0x5b, 0x9e, 0xaa, 0x84, 0x1b, 0xd4, 0x3b, 0x97, 0x0d, 0x61,
	0xfd, 0xdf, 0x95, 0x98, 0xaf, 0x9c, 0x97, 0x90, 0x79, 0x10, 0xff, 0x0e, 0x00, 0x6e, 0xb9, 0x54,
	0x39, 0xb7, 0x1e, 0xc8, 0x15, 0x80, 0x41, 0xc3, 0xd3, 0x2f, 0x58, 0xe5, 0xc6, 0xa5, 0xd2, 0x65,
	0x2d, 0x66, 0xdf, 0xbf, 0xe5, 0x52, 0x94, 0xf4, 0xac, 0xd5, 0x77, 0x0c, 0x89, 0x60, 0xb8, 0xae,
	0xb4, 0x16, 0xb9, 0xb1, 0x98, 0x9e, 0xbb, 0x81, 0x36, 0x65, 0x8d, 0x97, 0xea, 0x97, 0x4d, 0xd1,
	0xb6, 0x5c, 0x1d, 0x2f, 0xe0, 0xe2, 0x73, 0x65, 0xf0, 0xbb, 0xd0, 0x5b, 0xa1, 0xeb, 0x34, 0x28,
	0xf4, 0x79, 0x65, 0x50, 0x8a, 0xdc, 0x25, 0x12, 0xb2, 0x1a, 0x92, 0xf7, 0xf0, 0xc4, 0x3c, 0xe4,
	0xcb, 0x4c, 0x25, 0x49, 0x2a, 0xee, 0xb9, 0xf6, 0x8b, 0x0a, 0xd9, 0xc8, 0x3c, 0xe4, 0xf3, 0x3d,
	0x19, 0x7f, 0x80, 0xd1, 0x5c, 0x98, 0x9f, 0xd8, 0xca, 0x77, 0x88, 0xab, 0x3b, 0xb1, 0x36, 0x4b,
	0xf3, 0x58, 0x08, 0xa7, 0x3a, 0x60, 0xe0, 0xa9, 0xc5, 0x63, 0x21, 0x66, 0x0c, 0x06, 0x1b, 0x95,
	0x8a, 0x25, 0x16, 0xa6, 0x24, 0x97, 0x13, 0x7f, 0x68, 0x93, 0xfa, 0xd0, 0x26, 0xad, 0x83, 0xa2,
	0xff, 0xfe, 0xf4, 0xdc, 0x9a, 0x5e, 0x36, 0x6b, 0x6a, 0xb5, 0x59, 0xb8, 0xf1, 0xa0, 0x9c, 0xdd,
	0x40, 0xc7, 0xc9, 0x5d, 0x1f, 0xc9, 0x1d, 0xde, 0xe2, 0x5e, 0x91, 0x36, 0x8a, 0x87, 0x13, 0xcc,
	0x09, 0xcd, 0xe6, 0xd0, 0xdd, 0xd8, 0xab, 0x21, 0xef, 0x4e, 0x18, 0x6c, 0xae, 0x69, 0xaf, 0xf7,
	0xaa, 0xed, 0xb0, 0xe9, 0x33, 0xaf, 0x32, 0x63, 0xd0, 0x2b, 0x5d, 0xee, 0x27, 0x1c, 0xda, 0x85,
	0xa8, 0xf5, 0x91, 0xc3, 0xd7, 0x5e, 0xf1, 0x68, 0x65, 0x6c, 0xa7, 0x34, 0x9b, 0x43, 0x2f, 0x73,
	0xc9, 0x93, 0xab, 0x13, 0xaf, 0x6e, 0xad, 0x64, 0x2f, 0xf9, 0xdc, 0x4b, 0x1e, 0x34, 0xd9, 0x4e,
	0xe4, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x5c, 0xf3, 0x89, 0x13, 0x04, 0x00, 0x00,
}
