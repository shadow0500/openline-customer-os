// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: v1/location.proto

package location_grpc_service

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

type UpsertLocationGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant        string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AppSource     string                 `protobuf:"bytes,4,opt,name=appSource,proto3" json:"appSource,omitempty"`
	Source        string                 `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	SourceOfTruth string                 `protobuf:"bytes,6,opt,name=sourceOfTruth,proto3" json:"sourceOfTruth,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	RawAddress    string                 `protobuf:"bytes,9,opt,name=rawAddress,proto3" json:"rawAddress,omitempty"`
	Country       string                 `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`
	Region        string                 `protobuf:"bytes,11,opt,name=region,proto3" json:"region,omitempty"`
	Locality      string                 `protobuf:"bytes,12,opt,name=locality,proto3" json:"locality,omitempty"`
	AddressLine1  string                 `protobuf:"bytes,13,opt,name=addressLine1,proto3" json:"addressLine1,omitempty"`
	AddressLine2  string                 `protobuf:"bytes,14,opt,name=addressLine2,proto3" json:"addressLine2,omitempty"`
	ZipCode       string                 `protobuf:"bytes,15,opt,name=zipCode,proto3" json:"zipCode,omitempty"`
	AddressType   string                 `protobuf:"bytes,16,opt,name=addressType,proto3" json:"addressType,omitempty"`
	HouseNumber   string                 `protobuf:"bytes,17,opt,name=houseNumber,proto3" json:"houseNumber,omitempty"`
	PostalCode    string                 `protobuf:"bytes,18,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	Commercial    bool                   `protobuf:"varint,19,opt,name=commercial,proto3" json:"commercial,omitempty"`
	Predirection  string                 `protobuf:"bytes,20,opt,name=predirection,proto3" json:"predirection,omitempty"`
	District      string                 `protobuf:"bytes,21,opt,name=district,proto3" json:"district,omitempty"`
	Street        string                 `protobuf:"bytes,22,opt,name=street,proto3" json:"street,omitempty"`
	Latitude      string                 `protobuf:"bytes,23,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     string                 `protobuf:"bytes,24,opt,name=longitude,proto3" json:"longitude,omitempty"`
	PlusFour      string                 `protobuf:"bytes,25,opt,name=plusFour,proto3" json:"plusFour,omitempty"`
}

func (x *UpsertLocationGrpcRequest) Reset() {
	*x = UpsertLocationGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertLocationGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertLocationGrpcRequest) ProtoMessage() {}

func (x *UpsertLocationGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertLocationGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpsertLocationGrpcRequest) Descriptor() ([]byte, []int) {
	return file_v1_location_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertLocationGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetSourceOfTruth() string {
	if x != nil {
		return x.SourceOfTruth
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UpsertLocationGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpsertLocationGrpcRequest) GetRawAddress() string {
	if x != nil {
		return x.RawAddress
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetAddressLine1() string {
	if x != nil {
		return x.AddressLine1
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetAddressLine2() string {
	if x != nil {
		return x.AddressLine2
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetAddressType() string {
	if x != nil {
		return x.AddressType
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetHouseNumber() string {
	if x != nil {
		return x.HouseNumber
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetCommercial() bool {
	if x != nil {
		return x.Commercial
	}
	return false
}

func (x *UpsertLocationGrpcRequest) GetPredirection() string {
	if x != nil {
		return x.Predirection
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

func (x *UpsertLocationGrpcRequest) GetPlusFour() string {
	if x != nil {
		return x.PlusFour
	}
	return ""
}

type LocationIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LocationIdGrpcResponse) Reset() {
	*x = LocationIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationIdGrpcResponse) ProtoMessage() {}

func (x *LocationIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*LocationIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_v1_location_proto_rawDescGZIP(), []int{1}
}

func (x *LocationIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_v1_location_proto protoreflect.FileDescriptor

var file_v1_location_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x06, 0x0a, 0x19, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66,
	0x54, 0x72, 0x75, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x61, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x31,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c,
	0x69, 0x6e, 0x65, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69,
	0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x75, 0x73, 0x46, 0x6f, 0x75, 0x72,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x73, 0x46, 0x6f, 0x75, 0x72,
	0x22, 0x28, 0x0a, 0x16, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x5c, 0x0a, 0x13, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x45, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x42, 0x0d, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x3b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_location_proto_rawDescOnce sync.Once
	file_v1_location_proto_rawDescData = file_v1_location_proto_rawDesc
)

func file_v1_location_proto_rawDescGZIP() []byte {
	file_v1_location_proto_rawDescOnce.Do(func() {
		file_v1_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_location_proto_rawDescData)
	})
	return file_v1_location_proto_rawDescData
}

var file_v1_location_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_location_proto_goTypes = []interface{}{
	(*UpsertLocationGrpcRequest)(nil), // 0: UpsertLocationGrpcRequest
	(*LocationIdGrpcResponse)(nil),    // 1: LocationIdGrpcResponse
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
}
var file_v1_location_proto_depIdxs = []int32{
	2, // 0: UpsertLocationGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	2, // 1: UpsertLocationGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: LocationGrpcService.UpsertLocation:input_type -> UpsertLocationGrpcRequest
	1, // 3: LocationGrpcService.UpsertLocation:output_type -> LocationIdGrpcResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_location_proto_init() }
func file_v1_location_proto_init() {
	if File_v1_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertLocationGrpcRequest); i {
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
		file_v1_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationIdGrpcResponse); i {
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
			RawDescriptor: file_v1_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_location_proto_goTypes,
		DependencyIndexes: file_v1_location_proto_depIdxs,
		MessageInfos:      file_v1_location_proto_msgTypes,
	}.Build()
	File_v1_location_proto = out.File
	file_v1_location_proto_rawDesc = nil
	file_v1_location_proto_goTypes = nil
	file_v1_location_proto_depIdxs = nil
}
