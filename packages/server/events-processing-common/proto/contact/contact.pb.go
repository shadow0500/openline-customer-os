// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: proto/contact/contact.proto

package contactGrpcService

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

	TenantID  string `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	UUID      string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contact_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contact_contact_proto_msgTypes[0]
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
	return file_proto_contact_contact_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetTenantID() string {
	if x != nil {
		return x.TenantID
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
		mi := &file_proto_contact_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertContactGrpcRequest) ProtoMessage() {}

func (x *UpsertContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contact_contact_proto_msgTypes[1]
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
	return file_proto_contact_contact_proto_rawDescGZIP(), []int{1}
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

type AddPhoneNumberToContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant        string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	ContactId     string `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId,omitempty"`
	PhoneNumberId string `protobuf:"bytes,3,opt,name=phoneNumberId,proto3" json:"phoneNumberId,omitempty"`
	Primary       bool   `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Label         string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *AddPhoneNumberToContactGrpcRequest) Reset() {
	*x = AddPhoneNumberToContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contact_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPhoneNumberToContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPhoneNumberToContactGrpcRequest) ProtoMessage() {}

func (x *AddPhoneNumberToContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contact_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPhoneNumberToContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*AddPhoneNumberToContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_proto_contact_contact_proto_rawDescGZIP(), []int{2}
}

func (x *AddPhoneNumberToContactGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *AddPhoneNumberToContactGrpcRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *AddPhoneNumberToContactGrpcRequest) GetPhoneNumberId() string {
	if x != nil {
		return x.PhoneNumberId
	}
	return ""
}

func (x *AddPhoneNumberToContactGrpcRequest) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *AddPhoneNumberToContactGrpcRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type CreateContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID      string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	TenantID  string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *CreateContactGrpcRequest) Reset() {
	*x = CreateContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contact_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContactGrpcRequest) ProtoMessage() {}

func (x *CreateContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contact_contact_proto_msgTypes[3]
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
	return file_proto_contact_contact_proto_rawDescGZIP(), []int{3}
}

func (x *CreateContactGrpcRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
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

type CreateContactGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregateID string `protobuf:"bytes,1,opt,name=AggregateID,proto3" json:"AggregateID,omitempty"`
	UUID        string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *CreateContactGrpcResponse) Reset() {
	*x = CreateContactGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contact_contact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContactGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContactGrpcResponse) ProtoMessage() {}

func (x *CreateContactGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contact_contact_proto_msgTypes[4]
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
	return file_proto_contact_contact_proto_rawDescGZIP(), []int{4}
}

func (x *CreateContactGrpcResponse) GetAggregateID() string {
	if x != nil {
		return x.AggregateID
	}
	return ""
}

func (x *CreateContactGrpcResponse) GetUUID() string {
	if x != nil {
		return x.UUID
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
		mi := &file_proto_contact_contact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactIdGrpcResponse) ProtoMessage() {}

func (x *ContactIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contact_contact_proto_msgTypes[5]
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
	return file_proto_contact_contact_proto_rawDescGZIP(), []int{5}
}

func (x *ContactIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_contact_contact_proto protoreflect.FileDescriptor

var file_proto_contact_contact_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x73, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf8, 0x02, 0x0a, 0x18, 0x55, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x22, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x84, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x51, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22,
	0x27, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xea, 0x02, 0x0a, 0x12, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a,
	0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2c,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_contact_contact_proto_rawDescOnce sync.Once
	file_proto_contact_contact_proto_rawDescData = file_proto_contact_contact_proto_rawDesc
)

func file_proto_contact_contact_proto_rawDescGZIP() []byte {
	file_proto_contact_contact_proto_rawDescOnce.Do(func() {
		file_proto_contact_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_contact_contact_proto_rawDescData)
	})
	return file_proto_contact_contact_proto_rawDescData
}

var file_proto_contact_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_contact_contact_proto_goTypes = []interface{}{
	(*Contact)(nil),                            // 0: contactGrpcService.Contact
	(*UpsertContactGrpcRequest)(nil),           // 1: contactGrpcService.UpsertContactGrpcRequest
	(*AddPhoneNumberToContactGrpcRequest)(nil), // 2: contactGrpcService.AddPhoneNumberToContactGrpcRequest
	(*CreateContactGrpcRequest)(nil),           // 3: contactGrpcService.CreateContactGrpcRequest
	(*CreateContactGrpcResponse)(nil),          // 4: contactGrpcService.CreateContactGrpcResponse
	(*ContactIdGrpcResponse)(nil),              // 5: contactGrpcService.ContactIdGrpcResponse
	(*timestamppb.Timestamp)(nil),              // 6: google.protobuf.Timestamp
}
var file_proto_contact_contact_proto_depIdxs = []int32{
	6, // 0: contactGrpcService.UpsertContactGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	6, // 1: contactGrpcService.UpsertContactGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	3, // 2: contactGrpcService.contactGrpcService.CreateContact:input_type -> contactGrpcService.CreateContactGrpcRequest
	1, // 3: contactGrpcService.contactGrpcService.UpsertContact:input_type -> contactGrpcService.UpsertContactGrpcRequest
	2, // 4: contactGrpcService.contactGrpcService.AddPhoneNumberToContact:input_type -> contactGrpcService.AddPhoneNumberToContactGrpcRequest
	4, // 5: contactGrpcService.contactGrpcService.CreateContact:output_type -> contactGrpcService.CreateContactGrpcResponse
	5, // 6: contactGrpcService.contactGrpcService.UpsertContact:output_type -> contactGrpcService.ContactIdGrpcResponse
	5, // 7: contactGrpcService.contactGrpcService.AddPhoneNumberToContact:output_type -> contactGrpcService.ContactIdGrpcResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_contact_contact_proto_init() }
func file_proto_contact_contact_proto_init() {
	if File_proto_contact_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_contact_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_contact_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_contact_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPhoneNumberToContactGrpcRequest); i {
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
		file_proto_contact_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_contact_contact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_contact_contact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_contact_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_contact_contact_proto_goTypes,
		DependencyIndexes: file_proto_contact_contact_proto_depIdxs,
		MessageInfos:      file_proto_contact_contact_proto_msgTypes,
	}.Build()
	File_proto_contact_contact_proto = out.File
	file_proto_contact_contact_proto_rawDesc = nil
	file_proto_contact_contact_proto_goTypes = nil
	file_proto_contact_contact_proto_depIdxs = nil
}
