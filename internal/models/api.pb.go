// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: api.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TransferCard_Properties_Type int32

const (
	TransferCard_Properties_NONE      TransferCard_Properties_Type = 0
	TransferCard_Properties_INVITE    TransferCard_Properties_Type = 1
	TransferCard_Properties_REPLY     TransferCard_Properties_Type = 2
	TransferCard_Properties_COMPLETED TransferCard_Properties_Type = 3
)

// Enum value maps for TransferCard_Properties_Type.
var (
	TransferCard_Properties_Type_name = map[int32]string{
		0: "NONE",
		1: "INVITE",
		2: "REPLY",
		3: "COMPLETED",
	}
	TransferCard_Properties_Type_value = map[string]int32{
		"NONE":      0,
		"INVITE":    1,
		"REPLY":     2,
		"COMPLETED": 3,
	}
)

func (x TransferCard_Properties_Type) Enum() *TransferCard_Properties_Type {
	p := new(TransferCard_Properties_Type)
	*p = x
	return p
}

func (x TransferCard_Properties_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferCard_Properties_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (TransferCard_Properties_Type) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x TransferCard_Properties_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferCard_Properties_Type.Descriptor instead.
func (TransferCard_Properties_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4, 0, 0}
}

// Define ConnectionRequest: Initial Connection Message
type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude    float64      `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float64      `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Profile     *Profile     `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	Device      *Device      `protobuf:"bytes,4,opt,name=device,proto3" json:"device,omitempty"`
	Directories *Directories `protobuf:"bytes,5,opt,name=directories,proto3" json:"directories,omitempty"`
	Contact     *Contact     `protobuf:"bytes,6,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *ConnectionRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *ConnectionRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *ConnectionRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *ConnectionRequest) GetDirectories() *Directories {
	if x != nil {
		return x.Directories
	}
	return nil
}

func (x *ConnectionRequest) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

// Define ProcessRequest: Processes Given File with Information
type ProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExternal    bool   `protobuf:"varint,1,opt,name=isExternal,proto3" json:"isExternal,omitempty"`
	FilePath      string `protobuf:"bytes,2,opt,name=filePath,proto3" json:"filePath,omitempty"`
	ThumbnailPath string `protobuf:"bytes,3,opt,name=thumbnailPath,proto3" json:"thumbnailPath,omitempty"`
	Duration      int32  `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *ProcessRequest) Reset() {
	*x = ProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessRequest) ProtoMessage() {}

func (x *ProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessRequest.ProtoReflect.Descriptor instead.
func (*ProcessRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessRequest) GetIsExternal() bool {
	if x != nil {
		return x.IsExternal
	}
	return false
}

func (x *ProcessRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *ProcessRequest) GetThumbnailPath() string {
	if x != nil {
		return x.ThumbnailPath
	}
	return ""
}

func (x *ProcessRequest) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

// Define AuthInvite: Invitation Message sent on RPC
type AuthInvite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  Payload       `protobuf:"varint,1,opt,name=payload,proto3,enum=Payload" json:"payload,omitempty"`
	From     *Peer         `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Card     *TransferCard `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`
	IsDirect bool          `protobuf:"varint,4,opt,name=isDirect,proto3" json:"isDirect,omitempty"`
	IsFile   bool          `protobuf:"varint,5,opt,name=isFile,proto3" json:"isFile,omitempty"`
}

func (x *AuthInvite) Reset() {
	*x = AuthInvite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInvite) ProtoMessage() {}

func (x *AuthInvite) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInvite.ProtoReflect.Descriptor instead.
func (*AuthInvite) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *AuthInvite) GetPayload() Payload {
	if x != nil {
		return x.Payload
	}
	return Payload_UNDEFINED
}

func (x *AuthInvite) GetFrom() *Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *AuthInvite) GetCard() *TransferCard {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *AuthInvite) GetIsDirect() bool {
	if x != nil {
		return x.IsDirect
	}
	return false
}

func (x *AuthInvite) GetIsFile() bool {
	if x != nil {
		return x.IsFile
	}
	return false
}

// Define AuthReply: Reply Message sent on RPC
type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  Payload       `protobuf:"varint,1,opt,name=payload,proto3,enum=Payload" json:"payload,omitempty"`
	From     *Peer         `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Decision bool          `protobuf:"varint,3,opt,name=decision,proto3" json:"decision,omitempty"`
	Card     *TransferCard `protobuf:"bytes,4,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *AuthReply) GetPayload() Payload {
	if x != nil {
		return x.Payload
	}
	return Payload_UNDEFINED
}

func (x *AuthReply) GetFrom() *Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *AuthReply) GetDecision() bool {
	if x != nil {
		return x.Decision
	}
	return false
}

func (x *AuthReply) GetCard() *TransferCard {
	if x != nil {
		return x.Card
	}
	return nil
}

// Define TransferCard Type: Transferred Data Card
type TransferCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SQL Properties
	Id         int32                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload    Payload                  `protobuf:"varint,2,opt,name=payload,proto3,enum=Payload" json:"payload,omitempty"`
	Properties *TransferCard_Properties `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
	LastOpened int32                    `protobuf:"varint,4,opt,name=lastOpened,proto3" json:"lastOpened,omitempty"`
	Platform   Platform                 `protobuf:"varint,5,opt,name=platform,proto3,enum=Platform" json:"platform,omitempty"`
	Preview    []byte                   `protobuf:"bytes,6,opt,name=preview,proto3" json:"preview,omitempty"`
	// Owner Properties
	Username  string `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	FirstName string `protobuf:"bytes,8,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,9,opt,name=lastName,proto3" json:"lastName,omitempty"`
	// Data Properties
	Contact  *Contact  `protobuf:"bytes,10,opt,name=contact,proto3" json:"contact,omitempty"`
	Metadata *Metadata `protobuf:"bytes,11,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Url      string    `protobuf:"bytes,12,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *TransferCard) Reset() {
	*x = TransferCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferCard) ProtoMessage() {}

func (x *TransferCard) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferCard.ProtoReflect.Descriptor instead.
func (*TransferCard) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *TransferCard) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TransferCard) GetPayload() Payload {
	if x != nil {
		return x.Payload
	}
	return Payload_UNDEFINED
}

func (x *TransferCard) GetProperties() *TransferCard_Properties {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *TransferCard) GetLastOpened() int32 {
	if x != nil {
		return x.LastOpened
	}
	return 0
}

func (x *TransferCard) GetPlatform() Platform {
	if x != nil {
		return x.Platform
	}
	return Platform_Unknown
}

func (x *TransferCard) GetPreview() []byte {
	if x != nil {
		return x.Preview
	}
	return nil
}

func (x *TransferCard) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TransferCard) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *TransferCard) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *TransferCard) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *TransferCard) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TransferCard) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Define ErrorMessage: Error Message
type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *ErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorMessage) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

// -- Card Properties -- //
type TransferCard_Properties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     TransferCard_Properties_Type `protobuf:"varint,1,opt,name=type,proto3,enum=TransferCard_Properties_Type" json:"type,omitempty"`
	Mime     *MIME                        `protobuf:"bytes,2,opt,name=mime,proto3" json:"mime,omitempty"`
	Size     int32                        `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Duration int32                        `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Name     string                       `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TransferCard_Properties) Reset() {
	*x = TransferCard_Properties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferCard_Properties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferCard_Properties) ProtoMessage() {}

func (x *TransferCard_Properties) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferCard_Properties.ProtoReflect.Descriptor instead.
func (*TransferCard_Properties) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4, 0}
}

func (x *TransferCard_Properties) GetType() TransferCard_Properties_Type {
	if x != nil {
		return x.Type
	}
	return TransferCard_Properties_NONE
}

func (x *TransferCard_Properties) GetMime() *MIME {
	if x != nil {
		return x.Mime
	}
	return nil
}

func (x *TransferCard_Properties) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *TransferCard_Properties) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *TransferCard_Properties) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x0b, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x8e, 0x01, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa2, 0x01,
	0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x19, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x04, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x46, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x22, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x08, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x04, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x22, 0xe9,
	0x04, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x08, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x25, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x09, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x1a, 0xd6, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x4d, 0x49, 0x4d, 0x45, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x36, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x22, 0x40, 0x0a, 0x0c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(TransferCard_Properties_Type)(0), // 0: TransferCard.Properties.Type
	(*ConnectionRequest)(nil),         // 1: ConnectionRequest
	(*ProcessRequest)(nil),            // 2: ProcessRequest
	(*AuthInvite)(nil),                // 3: AuthInvite
	(*AuthReply)(nil),                 // 4: AuthReply
	(*TransferCard)(nil),              // 5: TransferCard
	(*ErrorMessage)(nil),              // 6: ErrorMessage
	(*TransferCard_Properties)(nil),   // 7: TransferCard.Properties
	(*Profile)(nil),                   // 8: Profile
	(*Device)(nil),                    // 9: Device
	(*Directories)(nil),               // 10: Directories
	(*Contact)(nil),                   // 11: Contact
	(Payload)(0),                      // 12: Payload
	(*Peer)(nil),                      // 13: Peer
	(Platform)(0),                     // 14: Platform
	(*Metadata)(nil),                  // 15: Metadata
	(*MIME)(nil),                      // 16: MIME
}
var file_api_proto_depIdxs = []int32{
	8,  // 0: ConnectionRequest.profile:type_name -> Profile
	9,  // 1: ConnectionRequest.device:type_name -> Device
	10, // 2: ConnectionRequest.directories:type_name -> Directories
	11, // 3: ConnectionRequest.contact:type_name -> Contact
	12, // 4: AuthInvite.payload:type_name -> Payload
	13, // 5: AuthInvite.from:type_name -> Peer
	5,  // 6: AuthInvite.card:type_name -> TransferCard
	12, // 7: AuthReply.payload:type_name -> Payload
	13, // 8: AuthReply.from:type_name -> Peer
	5,  // 9: AuthReply.card:type_name -> TransferCard
	12, // 10: TransferCard.payload:type_name -> Payload
	7,  // 11: TransferCard.properties:type_name -> TransferCard.Properties
	14, // 12: TransferCard.platform:type_name -> Platform
	11, // 13: TransferCard.contact:type_name -> Contact
	15, // 14: TransferCard.metadata:type_name -> Metadata
	0,  // 15: TransferCard.Properties.type:type_name -> TransferCard.Properties.Type
	16, // 16: TransferCard.Properties.mime:type_name -> MIME
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_data_proto_init()
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInvite); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferCard); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferCard_Properties); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
