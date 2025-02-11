// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/contact.proto

package contact_grpc_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant    string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	UUID      string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_v1_contact_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *Contact) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *Contact) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Contact) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type UpsertContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant        string                 `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Name          string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Prefix        string                 `protobuf:"bytes,6,opt,name=prefix,proto3" json:"prefix,omitempty"`
	AppSource     string                 `protobuf:"bytes,7,opt,name=appSource,proto3" json:"appSource,omitempty"`
	Source        string                 `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	SourceOfTruth string                 `protobuf:"bytes,9,opt,name=sourceOfTruth,proto3" json:"sourceOfTruth,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *UpsertContactGrpcRequest) Reset() {
	*x = UpsertContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertContactGrpcRequest) ProtoMessage() {}

func (x *UpsertContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpsertContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_v1_contact_proto_rawDescGZIP(), []int{1}
}

func (x *UpsertContactGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetSourceOfTruth() string {
	if x != nil {
		return x.SourceOfTruth
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UpsertContactGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type LinkPhoneNumberToContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant        string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	ContactId     string `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId,omitempty"`
	PhoneNumberId string `protobuf:"bytes,3,opt,name=phoneNumberId,proto3" json:"phoneNumberId,omitempty"`
	Primary       bool   `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Label         string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *LinkPhoneNumberToContactGrpcRequest) Reset() {
	*x = LinkPhoneNumberToContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkPhoneNumberToContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkPhoneNumberToContactGrpcRequest) ProtoMessage() {}

func (x *LinkPhoneNumberToContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkPhoneNumberToContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*LinkPhoneNumberToContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_v1_contact_proto_rawDescGZIP(), []int{2}
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetPhoneNumberId() string {
	if x != nil {
		return x.PhoneNumberId
	}
	return ""
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type LinkEmailToContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant    string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	ContactId string `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId,omitempty"`
	EmailId   string `protobuf:"bytes,3,opt,name=emailId,proto3" json:"emailId,omitempty"`
	Primary   bool   `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Label     string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *LinkEmailToContactGrpcRequest) Reset() {
	*x = LinkEmailToContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkEmailToContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkEmailToContactGrpcRequest) ProtoMessage() {}

func (x *LinkEmailToContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkEmailToContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*LinkEmailToContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_v1_contact_proto_rawDescGZIP(), []int{3}
}

func (x *LinkEmailToContactGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LinkEmailToContactGrpcRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *LinkEmailToContactGrpcRequest) GetEmailId() string {
	if x != nil {
		return x.EmailId
	}
	return ""
}

func (x *LinkEmailToContactGrpcRequest) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *LinkEmailToContactGrpcRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type CreateContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant        string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Prefix        string                 `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	AppSource     string                 `protobuf:"bytes,6,opt,name=appSource,proto3" json:"appSource,omitempty"`
	Source        string                 `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	SourceOfTruth string                 `protobuf:"bytes,8,opt,name=sourceOfTruth,proto3" json:"sourceOfTruth,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=createdAt,proto3,oneof" json:"createdAt,omitempty"`
}

func (x *CreateContactGrpcRequest) Reset() {
	*x = CreateContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContactGrpcRequest) ProtoMessage() {}

func (x *CreateContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*CreateContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_v1_contact_proto_rawDescGZIP(), []int{4}
}

func (x *CreateContactGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetSourceOfTruth() string {
	if x != nil {
		return x.SourceOfTruth
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateContactGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateContactGrpcResponse) Reset() {
	*x = CreateContactGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContactGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContactGrpcResponse) ProtoMessage() {}

func (x *CreateContactGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContactGrpcResponse.ProtoReflect.Descriptor instead.
func (*CreateContactGrpcResponse) Descriptor() ([]byte, []int) {
	return file_v1_contact_proto_rawDescGZIP(), []int{5}
}

func (x *CreateContactGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ContactIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ContactIdGrpcResponse) Reset() {
	*x = ContactIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contact_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactIdGrpcResponse) ProtoMessage() {}

func (x *ContactIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contact_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*ContactIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_v1_contact_proto_rawDescGZIP(), []int{6}
}

func (x *ContactIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_v1_contact_proto protoreflect.FileDescriptor

var file_v1_contact_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf8, 0x02, 0x0a, 0x18, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f,
	0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xb1, 0x01, 0x0a, 0x23, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x22, 0x9f, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xcf, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66,
	0x54, 0x72, 0x75, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49,
	0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc8, 0x02,
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x19, 0x2e,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x58, 0x0a, 0x18, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x4c, 0x69,
	0x6e, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x12, 0x1e, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x42, 0x0c, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x3b,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_contact_proto_rawDescOnce sync.Once
	file_v1_contact_proto_rawDescData = file_v1_contact_proto_rawDesc
)

func file_v1_contact_proto_rawDescGZIP() []byte {
	file_v1_contact_proto_rawDescOnce.Do(func() {
		file_v1_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_contact_proto_rawDescData)
	})
	return file_v1_contact_proto_rawDescData
}

var file_v1_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_contact_proto_goTypes = []interface{}{
	(*Contact)(nil),                             // 0: Contact
	(*UpsertContactGrpcRequest)(nil),            // 1: UpsertContactGrpcRequest
	(*LinkPhoneNumberToContactGrpcRequest)(nil), // 2: LinkPhoneNumberToContactGrpcRequest
	(*LinkEmailToContactGrpcRequest)(nil),       // 3: LinkEmailToContactGrpcRequest
	(*CreateContactGrpcRequest)(nil),            // 4: CreateContactGrpcRequest
	(*CreateContactGrpcResponse)(nil),           // 5: CreateContactGrpcResponse
	(*ContactIdGrpcResponse)(nil),               // 6: ContactIdGrpcResponse
	(*timestamppb.Timestamp)(nil),               // 7: google.protobuf.Timestamp
}
var file_v1_contact_proto_depIdxs = []int32{
	7, // 0: UpsertContactGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	7, // 1: UpsertContactGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	7, // 2: CreateContactGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	4, // 3: contactGrpcService.CreateContact:input_type -> CreateContactGrpcRequest
	1, // 4: contactGrpcService.UpsertContact:input_type -> UpsertContactGrpcRequest
	2, // 5: contactGrpcService.LinkPhoneNumberToContact:input_type -> LinkPhoneNumberToContactGrpcRequest
	3, // 6: contactGrpcService.LinkEmailToContact:input_type -> LinkEmailToContactGrpcRequest
	5, // 7: contactGrpcService.CreateContact:output_type -> CreateContactGrpcResponse
	6, // 8: contactGrpcService.UpsertContact:output_type -> ContactIdGrpcResponse
	6, // 9: contactGrpcService.LinkPhoneNumberToContact:output_type -> ContactIdGrpcResponse
	6, // 10: contactGrpcService.LinkEmailToContact:output_type -> ContactIdGrpcResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_contact_proto_init() }
func file_v1_contact_proto_init() {
	if File_v1_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_v1_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertContactGrpcRequest); i {
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
		file_v1_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkPhoneNumberToContactGrpcRequest); i {
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
		file_v1_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkEmailToContactGrpcRequest); i {
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
		file_v1_contact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContactGrpcRequest); i {
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
		file_v1_contact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContactGrpcResponse); i {
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
		file_v1_contact_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactIdGrpcResponse); i {
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
	file_v1_contact_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_contact_proto_goTypes,
		DependencyIndexes: file_v1_contact_proto_depIdxs,
		MessageInfos:      file_v1_contact_proto_msgTypes,
	}.Build()
	File_v1_contact_proto = out.File
	file_v1_contact_proto_rawDesc = nil
	file_v1_contact_proto_goTypes = nil
	file_v1_contact_proto_depIdxs = nil
}
