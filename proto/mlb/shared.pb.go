// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/mlb/shared.proto

package mlb

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

// 构建状态
type BuildStatus int32

const (
	BuildStatus_NONE       BuildStatus = 0
	BuildStatus_PENDING    BuildStatus = 1
	BuildStatus_INPROGRESS BuildStatus = 2
	BuildStatus_SUCCESS    BuildStatus = 3
	BuildStatus_FAILED     BuildStatus = 4
	BuildStatus_CANCELED   BuildStatus = 5
)

var BuildStatus_name = map[int32]string{
	0: "NONE",
	1: "PENDING",
	2: "INPROGRESS",
	3: "SUCCESS",
	4: "FAILED",
	5: "CANCELED",
}

var BuildStatus_value = map[string]int32{
	"NONE":       0,
	"PENDING":    1,
	"INPROGRESS": 2,
	"SUCCESS":    3,
	"FAILED":     4,
	"CANCELED":   5,
}

func (x BuildStatus) String() string {
	return proto.EnumName(BuildStatus_name, int32(x))
}

func (BuildStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c46991e1197981a7, []int{0}
}

// 响应状态
type Status struct {
	// 状态码，0代表正常
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 信息, 状态码等于0时为空
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_c46991e1197981a7, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 任务实体
type TaskEntity struct {
	// 唯一标识码
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// 创建时间
	CreatedAt int64 `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// 更新时间
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// 构建状态
	Status BuildStatus `protobuf:"varint,4,opt,name=status,proto3,enum=mlb.BuildStatus" json:"status,omitempty"`
	// 处理器
	Handler string `protobuf:"bytes,5,opt,name=handler,proto3" json:"handler,omitempty"`
	// 元数据
	Meta string `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
	// 日志
	Log                  string   `protobuf:"bytes,7,opt,name=log,proto3" json:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskEntity) Reset()         { *m = TaskEntity{} }
func (m *TaskEntity) String() string { return proto.CompactTextString(m) }
func (*TaskEntity) ProtoMessage()    {}
func (*TaskEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_c46991e1197981a7, []int{1}
}

func (m *TaskEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskEntity.Unmarshal(m, b)
}
func (m *TaskEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskEntity.Marshal(b, m, deterministic)
}
func (m *TaskEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskEntity.Merge(m, src)
}
func (m *TaskEntity) XXX_Size() int {
	return xxx_messageInfo_TaskEntity.Size(m)
}
func (m *TaskEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskEntity.DiscardUnknown(m)
}

var xxx_messageInfo_TaskEntity proto.InternalMessageInfo

func (m *TaskEntity) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *TaskEntity) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *TaskEntity) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *TaskEntity) GetStatus() BuildStatus {
	if m != nil {
		return m.Status
	}
	return BuildStatus_NONE
}

func (m *TaskEntity) GetHandler() string {
	if m != nil {
		return m.Handler
	}
	return ""
}

func (m *TaskEntity) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *TaskEntity) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func init() {
	proto.RegisterEnum("mlb.BuildStatus", BuildStatus_name, BuildStatus_value)
	proto.RegisterType((*Status)(nil), "mlb.Status")
	proto.RegisterType((*TaskEntity)(nil), "mlb.TaskEntity")
}

func init() {
	proto.RegisterFile("proto/mlb/shared.proto", fileDescriptor_c46991e1197981a7)
}

var fileDescriptor_c46991e1197981a7 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x6b, 0xc2, 0x30,
	0x18, 0xc6, 0x57, 0x5b, 0xab, 0xbe, 0x0e, 0x29, 0xef, 0x61, 0xe4, 0xb0, 0x83, 0x78, 0x2a, 0x3b,
	0x28, 0x6c, 0xb0, 0xbb, 0xab, 0x99, 0x08, 0x23, 0x4a, 0xba, 0x9d, 0x76, 0x4a, 0x4d, 0x50, 0x59,
	0x6a, 0xa5, 0x49, 0x0f, 0xfb, 0x82, 0xfb, 0x5c, 0x23, 0x69, 0xf7, 0xe7, 0xf6, 0x3c, 0xcf, 0x8f,
	0x37, 0xfc, 0x08, 0xdc, 0x5c, 0xea, 0xca, 0x56, 0x8b, 0x52, 0x17, 0x0b, 0x73, 0x14, 0xb5, 0x92,
	0x73, 0x3f, 0x60, 0x58, 0xea, 0x62, 0xf6, 0x08, 0x71, 0x6e, 0x85, 0x6d, 0x0c, 0x22, 0x44, 0xfb,
	0x4a, 0x2a, 0x12, 0x4c, 0x83, 0xb4, 0xcf, 0x7d, 0x46, 0x02, 0x83, 0x52, 0x19, 0x23, 0x0e, 0x8a,
	0xf4, 0xa6, 0x41, 0x3a, 0xe2, 0x3f, 0x75, 0xf6, 0x15, 0x00, 0xbc, 0x0a, 0xf3, 0x41, 0xcf, 0xf6,
	0x64, 0x3f, 0xdd, 0x71, 0xd3, 0x9c, 0xa4, 0x3f, 0x1e, 0x71, 0x9f, 0xf1, 0x16, 0x46, 0xfb, 0x5a,
	0x09, 0xab, 0xe4, 0xd2, 0xfa, 0xf3, 0x90, 0xff, 0x0d, 0x8e, 0x36, 0x17, 0xd9, 0xd1, 0xb0, 0xa5,
	0xbf, 0x03, 0xa6, 0x10, 0x1b, 0xaf, 0x45, 0xa2, 0x69, 0x90, 0x4e, 0xee, 0x93, 0x79, 0xa9, 0x8b,
	0xf9, 0x53, 0x73, 0xd2, 0xb2, 0xd5, 0xe5, 0x1d, 0x77, 0x8a, 0x47, 0x71, 0x96, 0x5a, 0xd5, 0xa4,
	0xdf, 0x2a, 0x76, 0xd5, 0x39, 0x95, 0xca, 0x0a, 0x12, 0xb7, 0x4e, 0x2e, 0x63, 0x02, 0xa1, 0xae,
	0x0e, 0x64, 0xe0, 0x27, 0x17, 0xef, 0xde, 0x61, 0xfc, 0xef, 0x59, 0x1c, 0x42, 0xc4, 0xb6, 0x8c,
	0x26, 0x57, 0x38, 0x86, 0xc1, 0x8e, 0xb2, 0xd5, 0x86, 0xad, 0x93, 0x00, 0x27, 0x00, 0x1b, 0xb6,
	0xe3, 0xdb, 0x35, 0xa7, 0x79, 0x9e, 0xf4, 0x1c, 0xcc, 0xdf, 0xb2, 0xcc, 0x95, 0x10, 0x01, 0xe2,
	0xe7, 0xe5, 0xe6, 0x85, 0xae, 0x92, 0x08, 0xaf, 0x61, 0x98, 0x2d, 0x59, 0x46, 0x5d, 0xeb, 0x17,
	0xb1, 0xff, 0xe9, 0x87, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x38, 0x1a, 0x38, 0x83, 0x01,
	0x00, 0x00,
}
