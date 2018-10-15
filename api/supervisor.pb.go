// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supervisor.proto

package api

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Program contains information necessary to execute program.
type Program struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arg                  []string `protobuf:"bytes,2,rep,name=arg,proto3" json:"arg,omitempty"`
	Env                  []string `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Program) Reset()         { *m = Program{} }
func (m *Program) String() string { return proto.CompactTextString(m) }
func (*Program) ProtoMessage()    {}
func (*Program) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{0}
}

func (m *Program) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Program.Unmarshal(m, b)
}
func (m *Program) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Program.Marshal(b, m, deterministic)
}
func (m *Program) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Program.Merge(m, src)
}
func (m *Program) XXX_Size() int {
	return xxx_messageInfo_Program.Size(m)
}
func (m *Program) XXX_DiscardUnknown() {
	xxx_messageInfo_Program.DiscardUnknown(m)
}

var xxx_messageInfo_Program proto.InternalMessageInfo

func (m *Program) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Program) GetArg() []string {
	if m != nil {
		return m.Arg
	}
	return nil
}

func (m *Program) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

// Status of the Program.
type Status struct {
	Program              *Program `protobuf:"bytes,1,opt,name=program,proto3" json:"program,omitempty"`
	Pid                  int64    `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Running              bool     `protobuf:"varint,3,opt,name=running,proto3" json:"running,omitempty"`
	Out                  string   `protobuf:"bytes,4,opt,name=out,proto3" json:"out,omitempty"`
	Err                  string   `protobuf:"bytes,5,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{1}
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

func (m *Status) GetProgram() *Program {
	if m != nil {
		return m.Program
	}
	return nil
}

func (m *Status) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Status) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *Status) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func (m *Status) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// ListRequest
type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

// ListResponse
type ListResponse struct {
	Statuses             map[string]*Status `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{3}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetStatuses() map[string]*Status {
	if m != nil {
		return m.Statuses
	}
	return nil
}

// AddRequest
type AddRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Program              *Program `protobuf:"bytes,2,opt,name=program,proto3" json:"program,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{4}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddRequest) GetProgram() *Program {
	if m != nil {
		return m.Program
	}
	return nil
}

// AddResponse
type AddResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{5}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

// RemoveRequest
type RemoveRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{6}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// RemoveResponse
type RemoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{7}
}

func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

// StartRequest
type StartRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{8}
}

func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (m *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(m, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// StartResponse
type StartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{9}
}

func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (m *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(m, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

// StopRequest
type StopRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{10}
}

func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (m *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(m, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

func (m *StopRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// StopResponse
type StopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{11}
}

func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (m *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(m, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

// StopAllRequest
type StopAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAllRequest) Reset()         { *m = StopAllRequest{} }
func (m *StopAllRequest) String() string { return proto.CompactTextString(m) }
func (*StopAllRequest) ProtoMessage()    {}
func (*StopAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{12}
}

func (m *StopAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAllRequest.Unmarshal(m, b)
}
func (m *StopAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAllRequest.Marshal(b, m, deterministic)
}
func (m *StopAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAllRequest.Merge(m, src)
}
func (m *StopAllRequest) XXX_Size() int {
	return xxx_messageInfo_StopAllRequest.Size(m)
}
func (m *StopAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopAllRequest proto.InternalMessageInfo

// StopAllResponse
type StopAllResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAllResponse) Reset()         { *m = StopAllResponse{} }
func (m *StopAllResponse) String() string { return proto.CompactTextString(m) }
func (*StopAllResponse) ProtoMessage()    {}
func (*StopAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{13}
}

func (m *StopAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAllResponse.Unmarshal(m, b)
}
func (m *StopAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAllResponse.Marshal(b, m, deterministic)
}
func (m *StopAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAllResponse.Merge(m, src)
}
func (m *StopAllResponse) XXX_Size() int {
	return xxx_messageInfo_StopAllResponse.Size(m)
}
func (m *StopAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopAllResponse proto.InternalMessageInfo

// StartAllRequest
type StartAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartAllRequest) Reset()         { *m = StartAllRequest{} }
func (m *StartAllRequest) String() string { return proto.CompactTextString(m) }
func (*StartAllRequest) ProtoMessage()    {}
func (*StartAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{14}
}

func (m *StartAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartAllRequest.Unmarshal(m, b)
}
func (m *StartAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartAllRequest.Marshal(b, m, deterministic)
}
func (m *StartAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartAllRequest.Merge(m, src)
}
func (m *StartAllRequest) XXX_Size() int {
	return xxx_messageInfo_StartAllRequest.Size(m)
}
func (m *StartAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartAllRequest proto.InternalMessageInfo

// StartAllResponse
type StartAllResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartAllResponse) Reset()         { *m = StartAllResponse{} }
func (m *StartAllResponse) String() string { return proto.CompactTextString(m) }
func (*StartAllResponse) ProtoMessage()    {}
func (*StartAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{15}
}

func (m *StartAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartAllResponse.Unmarshal(m, b)
}
func (m *StartAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartAllResponse.Marshal(b, m, deterministic)
}
func (m *StartAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartAllResponse.Merge(m, src)
}
func (m *StartAllResponse) XXX_Size() int {
	return xxx_messageInfo_StartAllResponse.Size(m)
}
func (m *StartAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartAllResponse proto.InternalMessageInfo

// RemoveAllRequest
type RemoveAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAllRequest) Reset()         { *m = RemoveAllRequest{} }
func (m *RemoveAllRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveAllRequest) ProtoMessage()    {}
func (*RemoveAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{16}
}

func (m *RemoveAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAllRequest.Unmarshal(m, b)
}
func (m *RemoveAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAllRequest.Marshal(b, m, deterministic)
}
func (m *RemoveAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAllRequest.Merge(m, src)
}
func (m *RemoveAllRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveAllRequest.Size(m)
}
func (m *RemoveAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAllRequest proto.InternalMessageInfo

// RemoveAllResponse
type RemoveAllResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAllResponse) Reset()         { *m = RemoveAllResponse{} }
func (m *RemoveAllResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveAllResponse) ProtoMessage()    {}
func (*RemoveAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{17}
}

func (m *RemoveAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAllResponse.Unmarshal(m, b)
}
func (m *RemoveAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAllResponse.Marshal(b, m, deterministic)
}
func (m *RemoveAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAllResponse.Merge(m, src)
}
func (m *RemoveAllResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveAllResponse.Size(m)
}
func (m *RemoveAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAllResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Program)(nil), "api.Program")
	proto.RegisterType((*Status)(nil), "api.Status")
	proto.RegisterType((*ListRequest)(nil), "api.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "api.ListResponse")
	proto.RegisterMapType((map[string]*Status)(nil), "api.ListResponse.StatusesEntry")
	proto.RegisterType((*AddRequest)(nil), "api.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "api.AddResponse")
	proto.RegisterType((*RemoveRequest)(nil), "api.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "api.RemoveResponse")
	proto.RegisterType((*StartRequest)(nil), "api.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "api.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "api.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "api.StopResponse")
	proto.RegisterType((*StopAllRequest)(nil), "api.StopAllRequest")
	proto.RegisterType((*StopAllResponse)(nil), "api.StopAllResponse")
	proto.RegisterType((*StartAllRequest)(nil), "api.StartAllRequest")
	proto.RegisterType((*StartAllResponse)(nil), "api.StartAllResponse")
	proto.RegisterType((*RemoveAllRequest)(nil), "api.RemoveAllRequest")
	proto.RegisterType((*RemoveAllResponse)(nil), "api.RemoveAllResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SupervisorClient is the client API for Supervisor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SupervisorClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	RemoveAll(ctx context.Context, in *RemoveAllRequest, opts ...grpc.CallOption) (*RemoveAllResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	StartAll(ctx context.Context, in *StartAllRequest, opts ...grpc.CallOption) (*StartAllResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	StopAll(ctx context.Context, in *StopAllRequest, opts ...grpc.CallOption) (*StopAllResponse, error)
}

type supervisorClient struct {
	cc *grpc.ClientConn
}

func NewSupervisorClient(cc *grpc.ClientConn) SupervisorClient {
	return &supervisorClient{cc}
}

func (c *supervisorClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/api.Supervisor/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/api.Supervisor/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/api.Supervisor/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) RemoveAll(ctx context.Context, in *RemoveAllRequest, opts ...grpc.CallOption) (*RemoveAllResponse, error) {
	out := new(RemoveAllResponse)
	err := c.cc.Invoke(ctx, "/api.Supervisor/RemoveAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/api.Supervisor/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) StartAll(ctx context.Context, in *StartAllRequest, opts ...grpc.CallOption) (*StartAllResponse, error) {
	out := new(StartAllResponse)
	err := c.cc.Invoke(ctx, "/api.Supervisor/StartAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/api.Supervisor/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) StopAll(ctx context.Context, in *StopAllRequest, opts ...grpc.CallOption) (*StopAllResponse, error) {
	out := new(StopAllResponse)
	err := c.cc.Invoke(ctx, "/api.Supervisor/StopAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupervisorServer is the server API for Supervisor service.
type SupervisorServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	RemoveAll(context.Context, *RemoveAllRequest) (*RemoveAllResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	StartAll(context.Context, *StartAllRequest) (*StartAllResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	StopAll(context.Context, *StopAllRequest) (*StopAllResponse, error)
}

func RegisterSupervisorServer(s *grpc.Server, srv SupervisorServer) {
	s.RegisterService(&_Supervisor_serviceDesc, srv)
}

func _Supervisor_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Supervisor/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Supervisor/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Supervisor/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_RemoveAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).RemoveAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Supervisor/RemoveAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).RemoveAll(ctx, req.(*RemoveAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Supervisor/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_StartAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).StartAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Supervisor/StartAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).StartAll(ctx, req.(*StartAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Supervisor/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_StopAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).StopAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Supervisor/StopAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).StopAll(ctx, req.(*StopAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Supervisor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Supervisor",
	HandlerType: (*SupervisorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Supervisor_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Supervisor_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Supervisor_Remove_Handler,
		},
		{
			MethodName: "RemoveAll",
			Handler:    _Supervisor_RemoveAll_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Supervisor_Start_Handler,
		},
		{
			MethodName: "StartAll",
			Handler:    _Supervisor_StartAll_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Supervisor_Stop_Handler,
		},
		{
			MethodName: "StopAll",
			Handler:    _Supervisor_StopAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor.proto",
}

func init() { proto.RegisterFile("supervisor.proto", fileDescriptor_b8b9452d77b1c7d2) }

var fileDescriptor_b8b9452d77b1c7d2 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc9, 0x8a, 0xdb, 0x4a,
	0x14, 0x7d, 0xb2, 0x3c, 0x5e, 0x79, 0x90, 0xaf, 0xbb, 0x1f, 0x42, 0x9b, 0xb8, 0x2b, 0x10, 0x44,
	0xa0, 0x4d, 0x70, 0x43, 0xc8, 0xb4, 0xf1, 0x22, 0xd0, 0x8b, 0x2c, 0x82, 0xf4, 0x05, 0x0a, 0x2a,
	0x8c, 0x88, 0x2d, 0x29, 0xa5, 0x92, 0xa0, 0x77, 0xf9, 0x89, 0xfc, 0x46, 0xbe, 0x31, 0xd4, 0xa0,
	0x52, 0xb9, 0x13, 0xbc, 0xbb, 0x3a, 0xf7, 0x9c, 0x3b, 0x9d, 0x42, 0xe0, 0xd7, 0x4d, 0x45, 0x59,
	0x9b, 0xd7, 0x25, 0xdb, 0x55, 0xac, 0xe4, 0x25, 0xba, 0x69, 0x95, 0x93, 0x03, 0x4c, 0xbe, 0xb2,
	0xf2, 0xc8, 0xd2, 0x33, 0x22, 0x0c, 0x8b, 0xf4, 0x4c, 0x03, 0x67, 0xeb, 0x44, 0xb3, 0x58, 0xc6,
	0xe8, 0x83, 0x9b, 0xb2, 0x63, 0x30, 0xd8, 0xba, 0xd1, 0x2c, 0x16, 0xa1, 0x40, 0x68, 0xd1, 0x06,
	0xae, 0x42, 0x68, 0xd1, 0x92, 0x9f, 0x0e, 0x8c, 0x13, 0x9e, 0xf2, 0xa6, 0xc6, 0x57, 0x30, 0xa9,
	0x54, 0x35, 0x59, 0xc5, 0xdb, 0xcf, 0x77, 0x69, 0x95, 0xef, 0x74, 0x87, 0xb8, 0x4b, 0x8a, 0x22,
	0x55, 0x9e, 0x05, 0x83, 0xad, 0x13, 0xb9, 0xb1, 0x08, 0x31, 0x80, 0x09, 0x6b, 0x8a, 0x22, 0x2f,
	0x8e, 0x81, 0xbb, 0x75, 0xa2, 0x69, 0xdc, 0x7d, 0x0a, 0x6e, 0xd9, 0xf0, 0x60, 0x28, 0xa7, 0x12,
	0xa1, 0x1c, 0x81, 0xb1, 0x60, 0xa4, 0x10, 0xca, 0x18, 0x59, 0x80, 0xf7, 0x25, 0xaf, 0x79, 0x4c,
	0x7f, 0x34, 0xb4, 0xe6, 0xe4, 0x97, 0x03, 0x73, 0xf5, 0x5d, 0x57, 0x65, 0x51, 0x53, 0xfc, 0x08,
	0xd3, 0x5a, 0x4e, 0x48, 0xeb, 0xc0, 0xd9, 0xba, 0x91, 0xb7, 0x7f, 0x21, 0x07, 0xb3, 0x49, 0xbb,
	0x44, 0x33, 0x3e, 0x17, 0x9c, 0x3d, 0xc5, 0x46, 0x10, 0x3e, 0xc2, 0xe2, 0x22, 0x25, 0xfa, 0x7f,
	0xa7, 0x4f, 0xfa, 0x4e, 0x22, 0xc4, 0x3b, 0x18, 0xb5, 0xe9, 0xa9, 0xa1, 0x72, 0x23, 0x6f, 0xef,
	0xc9, 0xe2, 0x4a, 0x14, 0xab, 0xcc, 0x87, 0xc1, 0x3b, 0x87, 0x3c, 0x02, 0x1c, 0xb2, 0x4c, 0x4f,
	0xf9, 0xcf, 0x7b, 0x5b, 0x07, 0x1c, 0x5c, 0x39, 0xa0, 0x58, 0x58, 0x56, 0x52, 0xa3, 0x93, 0x97,
	0xb0, 0x88, 0xe9, 0xb9, 0x6c, 0xe9, 0x95, 0xda, 0xc4, 0x87, 0x65, 0x47, 0xd2, 0x32, 0x02, 0xf3,
	0x84, 0xa7, 0x8c, 0x5f, 0x53, 0xad, 0xe4, 0xf6, 0xcc, 0x9c, 0x89, 0xdc, 0x81, 0x97, 0xf0, 0xb2,
	0xba, 0xa6, 0x59, 0x8a, 0xba, 0x82, 0xa2, 0x25, 0x3e, 0x2c, 0xc5, 0xf7, 0xe1, 0x74, 0xea, 0x1c,
	0x5a, 0xc3, 0xca, 0x20, 0x9a, 0x24, 0xa1, 0x94, 0x71, 0x8b, 0x85, 0xe0, 0xf7, 0x90, 0xa6, 0x21,
	0xf8, 0x6a, 0x0b, 0x8b, 0xb7, 0x81, 0xb5, 0x85, 0x29, 0xe2, 0xfe, 0xb7, 0x0b, 0x90, 0x98, 0x37,
	0x8f, 0xf7, 0x30, 0x14, 0x6e, 0xa3, 0x6f, 0x19, 0x2f, 0xd5, 0xe1, 0xfa, 0xaf, 0xa7, 0x40, 0xfe,
	0xc3, 0xd7, 0xe0, 0x1e, 0xb2, 0x0c, 0x57, 0x32, 0xd7, 0x9b, 0x16, 0xfa, 0x3d, 0x60, 0xb8, 0x0f,
	0x30, 0x56, 0xed, 0x11, 0x65, 0xf6, 0xc2, 0x8a, 0x70, 0x73, 0x81, 0x19, 0xd1, 0x27, 0x98, 0x99,
	0x99, 0xf1, 0xd6, 0xe2, 0xf4, 0x7b, 0x85, 0xff, 0x3f, 0x87, 0x8d, 0xfa, 0x0d, 0x8c, 0xe4, 0x65,
	0x70, 0xdd, 0x3d, 0x35, 0xe3, 0x62, 0x88, 0x36, 0x64, 0x14, 0xef, 0x61, 0xda, 0xdd, 0x12, 0x6f,
	0x7a, 0x86, 0xd5, 0xed, 0xf6, 0x19, 0x6a, 0xa4, 0xf7, 0x30, 0x14, 0x66, 0xe9, 0xd3, 0x59, 0xe6,
	0x87, 0x6b, 0x0b, 0x31, 0xf4, 0xb7, 0x30, 0xd1, 0xde, 0xe2, 0xc6, 0xe4, 0xad, 0x3e, 0x37, 0x97,
	0x60, 0xa7, 0xfb, 0x36, 0x96, 0xbf, 0xa5, 0x87, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0x8f,
	0x3d, 0x30, 0xaa, 0x04, 0x00, 0x00,
}