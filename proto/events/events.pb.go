// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: proto/events/events.proto

package events

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

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string            `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload   []byte            `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp int64             `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_proto_events_events_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PublishRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PublishRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *PublishRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_proto_events_events_proto_rawDescGZIP(), []int{1}
}

type ConsumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group   string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Topic   string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Offset  int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	AutoAck bool   `protobuf:"varint,4,opt,name=auto_ack,json=autoAck,proto3" json:"auto_ack,omitempty"`
	// duration in nanoseconds
	AckWait    int64 `protobuf:"varint,5,opt,name=ack_wait,json=ackWait,proto3" json:"ack_wait,omitempty"`
	RetryLimit int64 `protobuf:"varint,6,opt,name=retry_limit,json=retryLimit,proto3" json:"retry_limit,omitempty"`
}

func (x *ConsumeRequest) Reset() {
	*x = ConsumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeRequest) ProtoMessage() {}

func (x *ConsumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeRequest.ProtoReflect.Descriptor instead.
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return file_proto_events_events_proto_rawDescGZIP(), []int{2}
}

func (x *ConsumeRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ConsumeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ConsumeRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ConsumeRequest) GetAutoAck() bool {
	if x != nil {
		return x.AutoAck
	}
	return false
}

func (x *ConsumeRequest) GetAckWait() int64 {
	if x != nil {
		return x.AckWait
	}
	return 0
}

func (x *ConsumeRequest) GetRetryLimit() int64 {
	if x != nil {
		return x.RetryLimit
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic     string            `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload   []byte            `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp int64             `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_proto_events_events_proto_rawDescGZIP(), []int{3}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Event) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Event) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Event) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic  string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_proto_events_events_proto_rawDescGZIP(), []int{4}
}

func (x *ReadRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ReadRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ReadRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_proto_events_events_proto_rawDescGZIP(), []int{5}
}

func (x *ReadResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Ttl   int64  `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_proto_events_events_proto_rawDescGZIP(), []int{6}
}

func (x *WriteRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *WriteRequest) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_events_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_events_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_proto_events_events_proto_rawDescGZIP(), []int{7}
}

type AckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AckRequest) Reset() {
	*x = AckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_events_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckRequest) ProtoMessage() {}

func (x *AckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_events_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckRequest.ProtoReflect.Descriptor instead.
func (*AckRequest) Descriptor() ([]byte, []int) {
	return file_proto_events_events_proto_rawDescGZIP(), []int{8}
}

func (x *AckRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AckRequest) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_events_events_proto protoreflect.FileDescriptor

var file_proto_events_events_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x40, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x61, 0x75, 0x74, 0x6f, 0x41, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x6b, 0x5f,
	0x77, 0x61, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x63, 0x6b, 0x57,
	0x61, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0xdb, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x51, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x35, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x0c,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x74, 0x74, 0x6c, 0x22, 0x0f, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x78, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3a, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x12, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x16, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x32, 0x70, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x32, 0x36, 0x33, 0x37, 0x33, 0x30, 0x39, 0x39, 0x34,
	0x39, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_events_events_proto_rawDescOnce sync.Once
	file_proto_events_events_proto_rawDescData = file_proto_events_events_proto_rawDesc
)

func file_proto_events_events_proto_rawDescGZIP() []byte {
	file_proto_events_events_proto_rawDescOnce.Do(func() {
		file_proto_events_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_events_events_proto_rawDescData)
	})
	return file_proto_events_events_proto_rawDescData
}

var file_proto_events_events_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_events_events_proto_goTypes = []interface{}{
	(*PublishRequest)(nil),  // 0: events.PublishRequest
	(*PublishResponse)(nil), // 1: events.PublishResponse
	(*ConsumeRequest)(nil),  // 2: events.ConsumeRequest
	(*Event)(nil),           // 3: events.Event
	(*ReadRequest)(nil),     // 4: events.ReadRequest
	(*ReadResponse)(nil),    // 5: events.ReadResponse
	(*WriteRequest)(nil),    // 6: events.WriteRequest
	(*WriteResponse)(nil),   // 7: events.WriteResponse
	(*AckRequest)(nil),      // 8: events.AckRequest
	nil,                     // 9: events.PublishRequest.MetadataEntry
	nil,                     // 10: events.Event.MetadataEntry
}
var file_proto_events_events_proto_depIdxs = []int32{
	9,  // 0: events.PublishRequest.metadata:type_name -> events.PublishRequest.MetadataEntry
	10, // 1: events.Event.metadata:type_name -> events.Event.MetadataEntry
	3,  // 2: events.ReadResponse.events:type_name -> events.Event
	3,  // 3: events.WriteRequest.event:type_name -> events.Event
	0,  // 4: events.Stream.Publish:input_type -> events.PublishRequest
	2,  // 5: events.Stream.Consume:input_type -> events.ConsumeRequest
	4,  // 6: events.Store.Read:input_type -> events.ReadRequest
	6,  // 7: events.Store.Write:input_type -> events.WriteRequest
	1,  // 8: events.Stream.Publish:output_type -> events.PublishResponse
	3,  // 9: events.Stream.Consume:output_type -> events.Event
	5,  // 10: events.Store.Read:output_type -> events.ReadResponse
	7,  // 11: events.Store.Write:output_type -> events.WriteResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_events_events_proto_init() }
func file_proto_events_events_proto_init() {
	if File_proto_events_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_events_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_proto_events_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
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
		file_proto_events_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeRequest); i {
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
		file_proto_events_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_proto_events_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_proto_events_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_proto_events_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_proto_events_events_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
		file_proto_events_events_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckRequest); i {
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
			RawDescriptor: file_proto_events_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_events_events_proto_goTypes,
		DependencyIndexes: file_proto_events_events_proto_depIdxs,
		MessageInfos:      file_proto_events_events_proto_msgTypes,
	}.Build()
	File_proto_events_events_proto = out.File
	file_proto_events_events_proto_rawDesc = nil
	file_proto_events_events_proto_goTypes = nil
	file_proto_events_events_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Stream_ConsumeClient, error)
}

type streamClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClient(cc grpc.ClientConnInterface) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/events.Stream/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Stream_ConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[0], "/events.Stream/Consume", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamConsumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stream_ConsumeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type streamConsumeClient struct {
	grpc.ClientStream
}

func (x *streamConsumeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServer is the server API for Stream service.
type StreamServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	Consume(*ConsumeRequest, Stream_ConsumeServer) error
}

// UnimplementedStreamServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (*UnimplementedStreamServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedStreamServer) Consume(*ConsumeRequest, Stream_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
	s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.Stream/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stream_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServer).Consume(m, &streamConsumeServer{stream})
}

type Stream_ConsumeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type streamConsumeServer struct {
	grpc.ServerStream
}

func (x *streamConsumeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Stream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Stream_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Consume",
			Handler:       _Stream_Consume_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/events/events.proto",
}

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
}

type storeClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreClient(cc grpc.ClientConnInterface) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/events.Store/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/events.Store/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
}

// UnimplementedStoreServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (*UnimplementedStoreServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedStoreServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.Store/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.Store/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Store_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Store_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/events/events.proto",
}
