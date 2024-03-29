// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: internal/pb/callcenter.proto

package pb

import (
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

type CallCenterEmployee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone       string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Verified    bool   `protobuf:"varint,3,opt,name=verified,proto3" json:"verified,omitempty"`
	Role        string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	DateOfBirth string `protobuf:"bytes,5,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty"`
	UrlImage    string `protobuf:"bytes,6,opt,name=urlImage,proto3" json:"urlImage,omitempty"`
}

func (x *CallCenterEmployee) Reset() {
	*x = CallCenterEmployee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_callcenter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallCenterEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallCenterEmployee) ProtoMessage() {}

func (x *CallCenterEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_callcenter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallCenterEmployee.ProtoReflect.Descriptor instead.
func (*CallCenterEmployee) Descriptor() ([]byte, []int) {
	return file_internal_pb_callcenter_proto_rawDescGZIP(), []int{0}
}

func (x *CallCenterEmployee) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CallCenterEmployee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CallCenterEmployee) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *CallCenterEmployee) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *CallCenterEmployee) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *CallCenterEmployee) GetUrlImage() string {
	if x != nil {
		return x.UrlImage
	}
	return ""
}

type CreateCallCenterEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Role  string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateCallCenterEmployeeRequest) Reset() {
	*x = CreateCallCenterEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_callcenter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCallCenterEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCallCenterEmployeeRequest) ProtoMessage() {}

func (x *CreateCallCenterEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_callcenter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCallCenterEmployeeRequest.ProtoReflect.Descriptor instead.
func (*CreateCallCenterEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_callcenter_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCallCenterEmployeeRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateCallCenterEmployeeRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type UpdateCallCenterEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *CallCenterEmployee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *UpdateCallCenterEmployeeRequest) Reset() {
	*x = UpdateCallCenterEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_callcenter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCallCenterEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCallCenterEmployeeRequest) ProtoMessage() {}

func (x *UpdateCallCenterEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_callcenter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCallCenterEmployeeRequest.ProtoReflect.Descriptor instead.
func (*UpdateCallCenterEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_callcenter_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCallCenterEmployeeRequest) GetEmployee() *CallCenterEmployee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type GetEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetEmployeeRequest) Reset() {
	*x = GetEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_callcenter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmployeeRequest) ProtoMessage() {}

func (x *GetEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_callcenter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmployeeRequest.ProtoReflect.Descriptor instead.
func (*GetEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_callcenter_proto_rawDescGZIP(), []int{3}
}

func (x *GetEmployeeRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64               `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Item   *CallCenterEmployee `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetEmployeeResponse) Reset() {
	*x = GetEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_callcenter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmployeeResponse) ProtoMessage() {}

func (x *GetEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_callcenter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmployeeResponse.ProtoReflect.Descriptor instead.
func (*GetEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_callcenter_proto_rawDescGZIP(), []int{4}
}

func (x *GetEmployeeResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetEmployeeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetEmployeeResponse) GetItem() *CallCenterEmployee {
	if x != nil {
		return x.Item
	}
	return nil
}

type GetListEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetListEmployeeRequest) Reset() {
	*x = GetListEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_callcenter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListEmployeeRequest) ProtoMessage() {}

func (x *GetListEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_callcenter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListEmployeeRequest.ProtoReflect.Descriptor instead.
func (*GetListEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_callcenter_proto_rawDescGZIP(), []int{5}
}

func (x *GetListEmployeeRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListEmployeeRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetListEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64                 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string                `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Items  []*CallCenterEmployee `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Total  int64                 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetListEmployeeResponse) Reset() {
	*x = GetListEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_callcenter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListEmployeeResponse) ProtoMessage() {}

func (x *GetListEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_callcenter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListEmployeeResponse.ProtoReflect.Descriptor instead.
func (*GetListEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_callcenter_proto_rawDescGZIP(), []int{6}
}

func (x *GetListEmployeeResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetListEmployeeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetListEmployeeResponse) GetItems() []*CallCenterEmployee {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GetListEmployeeResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_internal_pb_callcenter_proto protoreflect.FileDescriptor

var file_internal_pb_callcenter_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61,
	0x6c, 0x6c, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x72, 0x6c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0x4b, 0x0a, 0x1f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x65,
	0x0a, 0x1f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x42, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x22, 0x7f, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x22, 0x46, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x17, 0x67,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xbf, 0x03, 0x0a, 0x11, 0x43, 0x61, 0x6c,
	0x6c, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e,
	0x0a, 0x0b, 0x67, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x26, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a,
	0x0a, 0x0f, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x12, 0x2a, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x0e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x33, 0x2e, 0x74,
	0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x0e, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x33, 0x2e, 0x74,
	0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x50, 0x01, 0x5a, 0x0d,
	0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_callcenter_proto_rawDescOnce sync.Once
	file_internal_pb_callcenter_proto_rawDescData = file_internal_pb_callcenter_proto_rawDesc
)

func file_internal_pb_callcenter_proto_rawDescGZIP() []byte {
	file_internal_pb_callcenter_proto_rawDescOnce.Do(func() {
		file_internal_pb_callcenter_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_callcenter_proto_rawDescData)
	})
	return file_internal_pb_callcenter_proto_rawDescData
}

var file_internal_pb_callcenter_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_pb_callcenter_proto_goTypes = []interface{}{
	(*CallCenterEmployee)(nil),              // 0: tech2.microservice.CallCenterEmployee
	(*CreateCallCenterEmployeeRequest)(nil), // 1: tech2.microservice.createCallCenterEmployeeRequest
	(*UpdateCallCenterEmployeeRequest)(nil), // 2: tech2.microservice.updateCallCenterEmployeeRequest
	(*GetEmployeeRequest)(nil),              // 3: tech2.microservice.getEmployeeRequest
	(*GetEmployeeResponse)(nil),             // 4: tech2.microservice.getEmployeeResponse
	(*GetListEmployeeRequest)(nil),          // 5: tech2.microservice.getListEmployeeRequest
	(*GetListEmployeeResponse)(nil),         // 6: tech2.microservice.getListEmployeeResponse
}
var file_internal_pb_callcenter_proto_depIdxs = []int32{
	0, // 0: tech2.microservice.updateCallCenterEmployeeRequest.employee:type_name -> tech2.microservice.CallCenterEmployee
	0, // 1: tech2.microservice.getEmployeeResponse.item:type_name -> tech2.microservice.CallCenterEmployee
	0, // 2: tech2.microservice.getListEmployeeResponse.items:type_name -> tech2.microservice.CallCenterEmployee
	3, // 3: tech2.microservice.CallCenterService.getEmployee:input_type -> tech2.microservice.getEmployeeRequest
	5, // 4: tech2.microservice.CallCenterService.getListEmployee:input_type -> tech2.microservice.getListEmployeeRequest
	1, // 5: tech2.microservice.CallCenterService.createEmployee:input_type -> tech2.microservice.createCallCenterEmployeeRequest
	2, // 6: tech2.microservice.CallCenterService.updateEmployee:input_type -> tech2.microservice.updateCallCenterEmployeeRequest
	4, // 7: tech2.microservice.CallCenterService.getEmployee:output_type -> tech2.microservice.getEmployeeResponse
	6, // 8: tech2.microservice.CallCenterService.getListEmployee:output_type -> tech2.microservice.getListEmployeeResponse
	4, // 9: tech2.microservice.CallCenterService.createEmployee:output_type -> tech2.microservice.getEmployeeResponse
	4, // 10: tech2.microservice.CallCenterService.updateEmployee:output_type -> tech2.microservice.getEmployeeResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_pb_callcenter_proto_init() }
func file_internal_pb_callcenter_proto_init() {
	if File_internal_pb_callcenter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_callcenter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallCenterEmployee); i {
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
		file_internal_pb_callcenter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCallCenterEmployeeRequest); i {
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
		file_internal_pb_callcenter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCallCenterEmployeeRequest); i {
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
		file_internal_pb_callcenter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmployeeRequest); i {
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
		file_internal_pb_callcenter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmployeeResponse); i {
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
		file_internal_pb_callcenter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListEmployeeRequest); i {
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
		file_internal_pb_callcenter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListEmployeeResponse); i {
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
			RawDescriptor: file_internal_pb_callcenter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pb_callcenter_proto_goTypes,
		DependencyIndexes: file_internal_pb_callcenter_proto_depIdxs,
		MessageInfos:      file_internal_pb_callcenter_proto_msgTypes,
	}.Build()
	File_internal_pb_callcenter_proto = out.File
	file_internal_pb_callcenter_proto_rawDesc = nil
	file_internal_pb_callcenter_proto_goTypes = nil
	file_internal_pb_callcenter_proto_depIdxs = nil
}
