// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: proto/debug/debug.proto

package debug

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SpanType int32

const (
	SpanType_INBOUND  SpanType = 0
	SpanType_OUTBOUND SpanType = 1
)

// Enum value maps for SpanType.
var (
	SpanType_name = map[int32]string{
		0: "INBOUND",
		1: "OUTBOUND",
	}
	SpanType_value = map[string]int32{
		"INBOUND":  0,
		"OUTBOUND": 1,
	}
)

func (x SpanType) Enum() *SpanType {
	p := new(SpanType)
	*p = x
	return p
}

func (x SpanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpanType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_debug_debug_proto_enumTypes[0].Descriptor()
}

func (SpanType) Type() protoreflect.EnumType {
	return &file_proto_debug_debug_proto_enumTypes[0]
}

func (x SpanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpanType.Descriptor instead.
func (SpanType) EnumDescriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{0}
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{0}
}

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// default: ok
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{1}
}

func (x *HealthResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{2}
}

type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp of recording
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// unix timestamp
	Started uint64 `protobuf:"varint,2,opt,name=started,proto3" json:"started,omitempty"`
	// in seconds
	Uptime uint64 `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// in bytes
	Memory uint64 `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	// num threads
	Threads uint64 `protobuf:"varint,5,opt,name=threads,proto3" json:"threads,omitempty"`
	// total gc in nanoseconds
	Gc uint64 `protobuf:"varint,6,opt,name=gc,proto3" json:"gc,omitempty"`
	// total number of requests
	Requests uint64 `protobuf:"varint,7,opt,name=requests,proto3" json:"requests,omitempty"`
	// total number of errors
	Errors uint64 `protobuf:"varint,8,opt,name=errors,proto3" json:"errors,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{3}
}

func (x *StatsResponse) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *StatsResponse) GetStarted() uint64 {
	if x != nil {
		return x.Started
	}
	return 0
}

func (x *StatsResponse) GetUptime() uint64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *StatsResponse) GetMemory() uint64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *StatsResponse) GetThreads() uint64 {
	if x != nil {
		return x.Threads
	}
	return 0
}

func (x *StatsResponse) GetGc() uint64 {
	if x != nil {
		return x.Gc
	}
	return 0
}

func (x *StatsResponse) GetRequests() uint64 {
	if x != nil {
		return x.Requests
	}
	return 0
}

func (x *StatsResponse) GetErrors() uint64 {
	if x != nil {
		return x.Errors
	}
	return 0
}

// LogRequest requests service logs
type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// count of records to request
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// relative time in seconds
	// before the current time
	// from which to show logs
	Since int64 `protobuf:"varint,2,opt,name=since,proto3" json:"since,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{4}
}

func (x *LogRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *LogRequest) GetSince() int64 {
	if x != nil {
		return x.Since
	}
	return 0
}

// LogResponse returns a list of logs
type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{5}
}

func (x *LogResponse) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

// Record is service log record
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp of log record
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// record metadata
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// message
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{6}
}

func (x *Record) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Record) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Record) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// trace id to retrieve
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TraceRequest) Reset() {
	*x = TraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceRequest) ProtoMessage() {}

func (x *TraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceRequest.ProtoReflect.Descriptor instead.
func (*TraceRequest) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{7}
}

func (x *TraceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spans []*Span `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
}

func (x *TraceResponse) Reset() {
	*x = TraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceResponse) ProtoMessage() {}

func (x *TraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceResponse.ProtoReflect.Descriptor instead.
func (*TraceResponse) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{8}
}

func (x *TraceResponse) GetSpans() []*Span {
	if x != nil {
		return x.Spans
	}
	return nil
}

type Span struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the trace id
	Trace string `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	// id of the span
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// parent span
	Parent string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	// name of the resource
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// time of start in nanoseconds
	Started uint64 `protobuf:"varint,5,opt,name=started,proto3" json:"started,omitempty"`
	// duration of the execution in nanoseconds
	Duration uint64 `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	// associated metadata
	Metadata map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type     SpanType          `protobuf:"varint,8,opt,name=type,proto3,enum=debug.SpanType" json:"type,omitempty"`
}

func (x *Span) Reset() {
	*x = Span{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_debug_debug_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_proto_debug_debug_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_proto_debug_debug_proto_rawDescGZIP(), []int{9}
}

func (x *Span) GetTrace() string {
	if x != nil {
		return x.Trace
	}
	return ""
}

func (x *Span) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Span) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Span) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Span) GetStarted() uint64 {
	if x != nil {
		return x.Started
	}
	return 0
}

func (x *Span) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Span) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Span) GetType() SpanType {
	if x != nil {
		return x.Type
	}
	return SpanType_INBOUND
}

var File_proto_debug_debug_proto protoreflect.FileDescriptor

var file_proto_debug_debug_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x28, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xd5, 0x01, 0x0a, 0x0d,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x67, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x67, 0x63, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x22, 0x38, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x36, 0x0a,
	0x0b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x37,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1e,
	0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32,
	0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x05, 0x73, 0x70, 0x61,
	0x6e, 0x73, 0x22, 0xa7, 0x02, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x70,
	0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x2e, 0x53, 0x70, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x25, 0x0a, 0x08,
	0x53, 0x70, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x42, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x01, 0x32, 0xdc, 0x01, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x2e, 0x0a,
	0x03, 0x4c, 0x6f, 0x67, 0x12, 0x11, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x14, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x13, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x32, 0x36, 0x33, 0x37, 0x33, 0x30, 0x39, 0x39, 0x34, 0x39, 0x2f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_debug_debug_proto_rawDescOnce sync.Once
	file_proto_debug_debug_proto_rawDescData = file_proto_debug_debug_proto_rawDesc
)

func file_proto_debug_debug_proto_rawDescGZIP() []byte {
	file_proto_debug_debug_proto_rawDescOnce.Do(func() {
		file_proto_debug_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_debug_debug_proto_rawDescData)
	})
	return file_proto_debug_debug_proto_rawDescData
}

var file_proto_debug_debug_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_debug_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_debug_debug_proto_goTypes = []interface{}{
	(SpanType)(0),          // 0: debug.SpanType
	(*HealthRequest)(nil),  // 1: debug.HealthRequest
	(*HealthResponse)(nil), // 2: debug.HealthResponse
	(*StatsRequest)(nil),   // 3: debug.StatsRequest
	(*StatsResponse)(nil),  // 4: debug.StatsResponse
	(*LogRequest)(nil),     // 5: debug.LogRequest
	(*LogResponse)(nil),    // 6: debug.LogResponse
	(*Record)(nil),         // 7: debug.Record
	(*TraceRequest)(nil),   // 8: debug.TraceRequest
	(*TraceResponse)(nil),  // 9: debug.TraceResponse
	(*Span)(nil),           // 10: debug.Span
	nil,                    // 11: debug.Record.MetadataEntry
	nil,                    // 12: debug.Span.MetadataEntry
}
var file_proto_debug_debug_proto_depIdxs = []int32{
	7,  // 0: debug.LogResponse.records:type_name -> debug.Record
	11, // 1: debug.Record.metadata:type_name -> debug.Record.MetadataEntry
	10, // 2: debug.TraceResponse.spans:type_name -> debug.Span
	12, // 3: debug.Span.metadata:type_name -> debug.Span.MetadataEntry
	0,  // 4: debug.Span.type:type_name -> debug.SpanType
	5,  // 5: debug.Debug.Log:input_type -> debug.LogRequest
	1,  // 6: debug.Debug.Health:input_type -> debug.HealthRequest
	3,  // 7: debug.Debug.Stats:input_type -> debug.StatsRequest
	8,  // 8: debug.Debug.Trace:input_type -> debug.TraceRequest
	6,  // 9: debug.Debug.Log:output_type -> debug.LogResponse
	2,  // 10: debug.Debug.Health:output_type -> debug.HealthResponse
	4,  // 11: debug.Debug.Stats:output_type -> debug.StatsResponse
	9,  // 12: debug.Debug.Trace:output_type -> debug.TraceResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_debug_debug_proto_init() }
func file_proto_debug_debug_proto_init() {
	if File_proto_debug_debug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_debug_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_debug_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_debug_debug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_debug_debug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_debug_debug_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_debug_debug_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_debug_debug_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_debug_debug_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_debug_debug_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_debug_debug_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Span); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_debug_debug_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_debug_debug_proto_goTypes,
		DependencyIndexes: file_proto_debug_debug_proto_depIdxs,
		EnumInfos:         file_proto_debug_debug_proto_enumTypes,
		MessageInfos:      file_proto_debug_debug_proto_msgTypes,
	}.Build()
	File_proto_debug_debug_proto = out.File
	file_proto_debug_debug_proto_rawDesc = nil
	file_proto_debug_debug_proto_goTypes = nil
	file_proto_debug_debug_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DebugClient interface {
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceResponse, error)
}

type debugClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugClient(cc grpc.ClientConnInterface) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/debug.Debug/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/debug.Debug/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, "/debug.Debug/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceResponse, error) {
	out := new(TraceResponse)
	err := c.cc.Invoke(ctx, "/debug.Debug/Trace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	Log(context.Context, *LogRequest) (*LogResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
	Trace(context.Context, *TraceRequest) (*TraceResponse, error)
}

// UnimplementedDebugServer can be embedded to have forward compatible implementations.
type UnimplementedDebugServer struct {
}

func (*UnimplementedDebugServer) Log(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (*UnimplementedDebugServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedDebugServer) Stats(context.Context, *StatsRequest) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (*UnimplementedDebugServer) Trace(context.Context, *TraceRequest) (*TraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trace not implemented")
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).Stats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_Trace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).Trace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/Trace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).Trace(ctx, req.(*TraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "debug.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _Debug_Log_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Debug_Health_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _Debug_Stats_Handler,
		},
		{
			MethodName: "Trace",
			Handler:    _Debug_Trace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/debug/debug.proto",
}
