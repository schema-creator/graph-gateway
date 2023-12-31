// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.1
// source: member-service/v1/member-service.proto

package v1

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

type Auth int32

const (
	Auth_read_only  Auth = 0
	Auth_read_write Auth = 1
	Auth_owner      Auth = 2
)

// Enum value maps for Auth.
var (
	Auth_name = map[int32]string{
		0: "read_only",
		1: "read_write",
		2: "owner",
	}
	Auth_value = map[string]int32{
		"read_only":  0,
		"read_write": 1,
		"owner":      2,
	}
)

func (x Auth) Enum() *Auth {
	p := new(Auth)
	*p = x
	return p
}

func (x Auth) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Auth) Descriptor() protoreflect.EnumDescriptor {
	return file_member_service_v1_member_service_proto_enumTypes[0].Descriptor()
}

func (Auth) Type() protoreflect.EnumType {
	return &file_member_service_v1_member_service_proto_enumTypes[0]
}

func (x Auth) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Auth.Descriptor instead.
func (Auth) EnumDescriptor() ([]byte, []int) {
	return file_member_service_v1_member_service_proto_rawDescGZIP(), []int{0}
}

type MemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Authority string `protobuf:"bytes,3,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (x *MemberRequest) Reset() {
	*x = MemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_service_v1_member_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberRequest) ProtoMessage() {}

func (x *MemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_member_service_v1_member_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberRequest.ProtoReflect.Descriptor instead.
func (*MemberRequest) Descriptor() ([]byte, []int) {
	return file_member_service_v1_member_service_proto_rawDescGZIP(), []int{0}
}

func (x *MemberRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MemberRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *MemberRequest) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Authority Auth   `protobuf:"varint,3,opt,name=authority,proto3,enum=member.v1.Auth" json:"authority,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_service_v1_member_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_member_service_v1_member_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_member_service_v1_member_service_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Member) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Member) GetAuthority() Auth {
	if x != nil {
		return x.Authority
	}
	return Auth_read_only
}

type GetMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetMembersRequest) Reset() {
	*x = GetMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_service_v1_member_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMembersRequest) ProtoMessage() {}

func (x *GetMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_member_service_v1_member_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMembersRequest.ProtoReflect.Descriptor instead.
func (*GetMembersRequest) Descriptor() ([]byte, []int) {
	return file_member_service_v1_member_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetMembersRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ListMembers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *ListMembers) Reset() {
	*x = ListMembers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_service_v1_member_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMembers) ProtoMessage() {}

func (x *ListMembers) ProtoReflect() protoreflect.Message {
	mi := &file_member_service_v1_member_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMembers.ProtoReflect.Descriptor instead.
func (*ListMembers) Descriptor() ([]byte, []int) {
	return file_member_service_v1_member_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListMembers) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type DeleteMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DeleteMemberRequest) Reset() {
	*x = DeleteMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_service_v1_member_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemberRequest) ProtoMessage() {}

func (x *DeleteMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_member_service_v1_member_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemberRequest.ProtoReflect.Descriptor instead.
func (*DeleteMemberRequest) Descriptor() ([]byte, []int) {
	return file_member_service_v1_member_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMemberRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteMemberRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type DeleteMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteMemberResponse) Reset() {
	*x = DeleteMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_service_v1_member_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemberResponse) ProtoMessage() {}

func (x *DeleteMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_service_v1_member_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemberResponse.ProtoReflect.Descriptor instead.
func (*DeleteMemberResponse) Descriptor() ([]byte, []int) {
	return file_member_service_v1_member_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteMemberResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_member_service_v1_member_service_proto protoreflect.FileDescriptor

var file_member_service_v1_member_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x65, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x6f, 0x0a, 0x06, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x32, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x3a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2b,
	0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x4d, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x30, 0x0a, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x12, 0x0d, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x10, 0x02, 0x32, 0xa1,
	0x02, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x42, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x3e, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x4f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_member_service_v1_member_service_proto_rawDescOnce sync.Once
	file_member_service_v1_member_service_proto_rawDescData = file_member_service_v1_member_service_proto_rawDesc
)

func file_member_service_v1_member_service_proto_rawDescGZIP() []byte {
	file_member_service_v1_member_service_proto_rawDescOnce.Do(func() {
		file_member_service_v1_member_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_member_service_v1_member_service_proto_rawDescData)
	})
	return file_member_service_v1_member_service_proto_rawDescData
}

var file_member_service_v1_member_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_member_service_v1_member_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_member_service_v1_member_service_proto_goTypes = []interface{}{
	(Auth)(0),                    // 0: member.v1.auth
	(*MemberRequest)(nil),        // 1: member.v1.MemberRequest
	(*Member)(nil),               // 2: member.v1.Member
	(*GetMembersRequest)(nil),    // 3: member.v1.GetMembersRequest
	(*ListMembers)(nil),          // 4: member.v1.ListMembers
	(*DeleteMemberRequest)(nil),  // 5: member.v1.DeleteMemberRequest
	(*DeleteMemberResponse)(nil), // 6: member.v1.DeleteMemberResponse
}
var file_member_service_v1_member_service_proto_depIdxs = []int32{
	0, // 0: member.v1.Member.authority:type_name -> member.v1.auth
	2, // 1: member.v1.ListMembers.members:type_name -> member.v1.Member
	1, // 2: member.v1.MemberService.CreateMember:input_type -> member.v1.MemberRequest
	3, // 3: member.v1.MemberService.GetMembers:input_type -> member.v1.GetMembersRequest
	1, // 4: member.v1.MemberService.UpdateAuthority:input_type -> member.v1.MemberRequest
	5, // 5: member.v1.MemberService.DeleteMember:input_type -> member.v1.DeleteMemberRequest
	2, // 6: member.v1.MemberService.CreateMember:output_type -> member.v1.Member
	4, // 7: member.v1.MemberService.GetMembers:output_type -> member.v1.ListMembers
	2, // 8: member.v1.MemberService.UpdateAuthority:output_type -> member.v1.Member
	6, // 9: member.v1.MemberService.DeleteMember:output_type -> member.v1.DeleteMemberResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_member_service_v1_member_service_proto_init() }
func file_member_service_v1_member_service_proto_init() {
	if File_member_service_v1_member_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_member_service_v1_member_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberRequest); i {
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
		file_member_service_v1_member_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_member_service_v1_member_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMembersRequest); i {
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
		file_member_service_v1_member_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMembers); i {
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
		file_member_service_v1_member_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemberRequest); i {
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
		file_member_service_v1_member_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemberResponse); i {
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
			RawDescriptor: file_member_service_v1_member_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_member_service_v1_member_service_proto_goTypes,
		DependencyIndexes: file_member_service_v1_member_service_proto_depIdxs,
		EnumInfos:         file_member_service_v1_member_service_proto_enumTypes,
		MessageInfos:      file_member_service_v1_member_service_proto_msgTypes,
	}.Build()
	File_member_service_v1_member_service_proto = out.File
	file_member_service_v1_member_service_proto_rawDesc = nil
	file_member_service_v1_member_service_proto_goTypes = nil
	file_member_service_v1_member_service_proto_depIdxs = nil
}
