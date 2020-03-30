// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/mlb/task.proto

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

// 查询请求
type QueryRequest struct {
	// 任务ID
	TaskID               string   `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff78b02eed61a09d, []int{0}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

// 查询回复
type QueryResponse struct {
	// 状态
	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 构建状态
	BuildStatus BuildStatus `protobuf:"varint,2,opt,name=buildStatus,proto3,enum=mlb.BuildStatus" json:"buildStatus,omitempty"`
	// 视频的存放地址，仅当构建成功时有值
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// 封面的存放地址，仅当构建成功时有值
	Cover string `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	// 任务创建时间
	// 任务创建时间
	CreatedAt int64 `protobuf:"varint,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// 任务更新时间
	UpdatedAt int64 `protobuf:"varint,11,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// 构建日志
	BuildLog             string   `protobuf:"bytes,12,opt,name=buildLog,proto3" json:"buildLog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff78b02eed61a09d, []int{1}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QueryResponse) GetBuildStatus() BuildStatus {
	if m != nil {
		return m.BuildStatus
	}
	return BuildStatus_PENDING
}

func (m *QueryResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *QueryResponse) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *QueryResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *QueryResponse) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *QueryResponse) GetBuildLog() string {
	if m != nil {
		return m.BuildLog
	}
	return ""
}

// 取消请求
type CancelRequest struct {
	// 任务ID
	TaskID               string   `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelRequest) Reset()         { *m = CancelRequest{} }
func (m *CancelRequest) String() string { return proto.CompactTextString(m) }
func (*CancelRequest) ProtoMessage()    {}
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff78b02eed61a09d, []int{2}
}

func (m *CancelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelRequest.Unmarshal(m, b)
}
func (m *CancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelRequest.Marshal(b, m, deterministic)
}
func (m *CancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelRequest.Merge(m, src)
}
func (m *CancelRequest) XXX_Size() int {
	return xxx_messageInfo_CancelRequest.Size(m)
}
func (m *CancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelRequest proto.InternalMessageInfo

func (m *CancelRequest) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

// 取消回复
type CancelResponse struct {
	// 状态
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelResponse) Reset()         { *m = CancelResponse{} }
func (m *CancelResponse) String() string { return proto.CompactTextString(m) }
func (*CancelResponse) ProtoMessage()    {}
func (*CancelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff78b02eed61a09d, []int{3}
}

func (m *CancelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelResponse.Unmarshal(m, b)
}
func (m *CancelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelResponse.Marshal(b, m, deterministic)
}
func (m *CancelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelResponse.Merge(m, src)
}
func (m *CancelResponse) XXX_Size() int {
	return xxx_messageInfo_CancelResponse.Size(m)
}
func (m *CancelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelResponse proto.InternalMessageInfo

func (m *CancelResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// 获取请求
type FetchRequest struct {
	// 开始时间
	StartTime int64 `protobuf:"varint,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	// 结束时间
	EndTime              int64    `protobuf:"varint,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRequest) Reset()         { *m = FetchRequest{} }
func (m *FetchRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()    {}
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff78b02eed61a09d, []int{4}
}

func (m *FetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRequest.Unmarshal(m, b)
}
func (m *FetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRequest.Marshal(b, m, deterministic)
}
func (m *FetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRequest.Merge(m, src)
}
func (m *FetchRequest) XXX_Size() int {
	return xxx_messageInfo_FetchRequest.Size(m)
}
func (m *FetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRequest proto.InternalMessageInfo

func (m *FetchRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *FetchRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

// 获取回复
type FetchResponse struct {
	// 状态
	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 任务列表
	Task                 []*TaskEntity `protobuf:"bytes,2,rep,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FetchResponse) Reset()         { *m = FetchResponse{} }
func (m *FetchResponse) String() string { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()    {}
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff78b02eed61a09d, []int{5}
}

func (m *FetchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchResponse.Unmarshal(m, b)
}
func (m *FetchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchResponse.Marshal(b, m, deterministic)
}
func (m *FetchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchResponse.Merge(m, src)
}
func (m *FetchResponse) XXX_Size() int {
	return xxx_messageInfo_FetchResponse.Size(m)
}
func (m *FetchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchResponse proto.InternalMessageInfo

func (m *FetchResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *FetchResponse) GetTask() []*TaskEntity {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "mlb.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "mlb.QueryResponse")
	proto.RegisterType((*CancelRequest)(nil), "mlb.CancelRequest")
	proto.RegisterType((*CancelResponse)(nil), "mlb.CancelResponse")
	proto.RegisterType((*FetchRequest)(nil), "mlb.FetchRequest")
	proto.RegisterType((*FetchResponse)(nil), "mlb.FetchResponse")
}

func init() {
	proto.RegisterFile("proto/mlb/task.proto", fileDescriptor_ff78b02eed61a09d)
}

var fileDescriptor_ff78b02eed61a09d = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x49, 0xcf, 0xd2, 0x40,
	0x18, 0xfe, 0x4a, 0xa1, 0xca, 0x5b, 0x40, 0x1c, 0x09, 0x99, 0x34, 0x1c, 0x9a, 0x92, 0x68, 0x4f,
	0x60, 0x4a, 0x88, 0x67, 0x37, 0xa2, 0x89, 0x17, 0x2a, 0x17, 0x8f, 0x5d, 0x26, 0xd2, 0xd0, 0x05,
	0x3b, 0x33, 0x26, 0xfc, 0x58, 0xef, 0xfe, 0x8c, 0x2f, 0x7d, 0x3b, 0xb4, 0xe5, 0x44, 0xb8, 0xcd,
	0xb3, 0xbc, 0xcb, 0x3c, 0x33, 0x30, 0x3b, 0x97, 0x85, 0x28, 0xd6, 0x59, 0x1a, 0xae, 0x45, 0xc0,
	0x4f, 0x2b, 0x84, 0x44, 0xcf, 0xd2, 0xd0, 0x9a, 0xb7, 0x12, 0x3f, 0x06, 0x25, 0x8b, 0x6b, 0xd1,
	0x79, 0x0b, 0xa3, 0xbd, 0x64, 0xe5, 0xc5, 0x67, 0x7f, 0x24, 0xe3, 0x82, 0xcc, 0xc1, 0xa8, 0x4a,
	0xbf, 0x7f, 0xa1, 0x9a, 0xad, 0xb9, 0x43, 0x5f, 0x21, 0xe7, 0xbf, 0x06, 0x63, 0x65, 0xe4, 0xe7,
	0x22, 0xe7, 0x8c, 0x2c, 0xc1, 0xe0, 0x22, 0x10, 0x92, 0xa3, 0xd3, 0xf4, 0xcc, 0x55, 0x96, 0x86,
	0xab, 0x9f, 0x48, 0xf9, 0x4a, 0x22, 0x1e, 0x98, 0xa1, 0x4c, 0xd2, 0xb8, 0xa6, 0x69, 0xcf, 0xd6,
	0xdc, 0x89, 0x37, 0x45, 0xe7, 0xa7, 0x96, 0xf7, 0xbb, 0x26, 0x32, 0x05, 0x5d, 0x96, 0x29, 0xd5,
	0x71, 0x7e, 0x75, 0x24, 0x33, 0x18, 0x44, 0xc5, 0x5f, 0x56, 0xd2, 0x3e, 0x72, 0x35, 0x20, 0x0b,
	0x18, 0x46, 0x25, 0x0b, 0x04, 0x8b, 0x3f, 0x0a, 0x0a, 0xb6, 0xe6, 0xea, 0x7e, 0x4b, 0x54, 0xaa,
	0x3c, 0xc7, 0x4a, 0x35, 0x6b, 0xb5, 0x21, 0x88, 0x05, 0x2f, 0x71, 0xe4, 0x8f, 0xe2, 0x37, 0x1d,
	0x61, 0xd3, 0x06, 0x3b, 0xef, 0x60, 0xfc, 0x39, 0xc8, 0x23, 0x96, 0xde, 0xcb, 0x64, 0x0b, 0x93,
	0xab, 0xf1, 0x81, 0x4c, 0x9c, 0x1d, 0x8c, 0x76, 0x4c, 0x44, 0xc7, 0x6b, 0xfb, 0x05, 0x0c, 0xb9,
	0x08, 0x4a, 0x71, 0x48, 0x32, 0x86, 0x75, 0xba, 0xdf, 0x12, 0x84, 0xc2, 0x0b, 0x96, 0xc7, 0xa8,
	0xf5, 0x50, 0xbb, 0x42, 0xe7, 0x17, 0x8c, 0x55, 0x9f, 0x47, 0x5e, 0x64, 0x09, 0xfd, 0x6a, 0x7d,
	0xda, 0xb3, 0x75, 0xd7, 0xf4, 0x5e, 0xa1, 0xe5, 0x10, 0xf0, 0xd3, 0xd7, 0x5c, 0x24, 0xe2, 0xe2,
	0xa3, 0xe8, 0xfd, 0xd3, 0xa0, 0x5f, 0x91, 0xe4, 0x3d, 0x0c, 0xf0, 0xd5, 0xc9, 0x6b, 0x34, 0x76,
	0xbf, 0x8a, 0x45, 0xba, 0x54, 0xbd, 0x82, 0xf3, 0x44, 0x36, 0x60, 0xd4, 0xa1, 0x90, 0x5a, 0xbf,
	0x89, 0xd2, 0x7a, 0x73, 0xc3, 0x35, 0x45, 0x5b, 0x00, 0xbc, 0xca, 0x5e, 0x32, 0xc9, 0xd4, 0xac,
	0x6e, 0x46, 0x6a, 0xd6, 0xcd, 0x75, 0x9d, 0x27, 0xf2, 0x41, 0x25, 0xf9, 0x2d, 0xe1, 0xa2, 0x68,
	0x96, 0xbc, 0x5f, 0x18, 0x1a, 0xf8, 0xf9, 0x37, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x68, 0x9a,
	0x6b, 0x38, 0x31, 0x03, 0x00, 0x00,
}
