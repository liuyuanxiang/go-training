// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: biz/user.proto

package biz

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_biz_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_biz_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CreateUserRsp) Reset() {
	*x = CreateUserRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRsp) ProtoMessage() {}

func (x *CreateUserRsp) ProtoReflect() protoreflect.Message {
	mi := &file_biz_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRsp.ProtoReflect.Descriptor instead.
func (*CreateUserRsp) Descriptor() ([]byte, []int) {
	return file_biz_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRsp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateUserRsp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserReq) Reset() {
	*x = GetUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReq) ProtoMessage() {}

func (x *GetUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_biz_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReq.ProtoReflect.Descriptor instead.
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return file_biz_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetUserRsp) Reset() {
	*x = GetUserRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRsp) ProtoMessage() {}

func (x *GetUserRsp) ProtoReflect() protoreflect.Message {
	mi := &file_biz_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRsp.ProtoReflect.Descriptor instead.
func (*GetUserRsp) Descriptor() ([]byte, []int) {
	return file_biz_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRsp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetUserRsp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ListUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int64 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int64 `protobuf:"varint,10,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListUserReq) Reset() {
	*x = ListUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReq) ProtoMessage() {}

func (x *ListUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_biz_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReq.ProtoReflect.Descriptor instead.
func (*ListUserReq) Descriptor() ([]byte, []int) {
	return file_biz_user_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListUserReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListUserRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPage int64               `protobuf:"varint,2,opt,name=currentPage,proto3" json:"currentPage,omitempty"`
	PageSize    int64               `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	TotalCount  int64               `protobuf:"varint,4,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	TotalPage   int64               `protobuf:"varint,5,opt,name=totalPage,proto3" json:"totalPage,omitempty"`
	List        []*ListUserRsp_User `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListUserRsp) Reset() {
	*x = ListUserRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRsp) ProtoMessage() {}

func (x *ListUserRsp) ProtoReflect() protoreflect.Message {
	mi := &file_biz_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRsp.ProtoReflect.Descriptor instead.
func (*ListUserRsp) Descriptor() ([]byte, []int) {
	return file_biz_user_proto_rawDescGZIP(), []int{5}
}

func (x *ListUserRsp) GetCurrentPage() int64 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *ListUserRsp) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListUserRsp) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListUserRsp) GetTotalPage() int64 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *ListUserRsp) GetList() []*ListUserRsp_User {
	if x != nil {
		return x.List
	}
	return nil
}

type ListUserRsp_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ListUserRsp_User) Reset() {
	*x = ListUserRsp_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biz_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRsp_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRsp_User) ProtoMessage() {}

func (x *ListUserRsp_User) ProtoReflect() protoreflect.Message {
	mi := &file_biz_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRsp_User.ProtoReflect.Descriptor instead.
func (*ListUserRsp_User) Descriptor() ([]byte, []int) {
	return file_biz_user_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListUserRsp_User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListUserRsp_User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_biz_user_proto protoreflect.FileDescriptor

var file_biz_user_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x69, 0x7a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x62, 0x69, 0x7a, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3b, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x45, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xe8, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x73, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a,
	0x32, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0xe9, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x64,
	0x3a, 0x01, 0x2a, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x73, 0x70,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x6f, 0x6e, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12,
	0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42,
	0x1f, 0x5a, 0x1d, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x33, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x3b, 0x62, 0x69, 0x7a,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_biz_user_proto_rawDescOnce sync.Once
	file_biz_user_proto_rawDescData = file_biz_user_proto_rawDesc
)

func file_biz_user_proto_rawDescGZIP() []byte {
	file_biz_user_proto_rawDescOnce.Do(func() {
		file_biz_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_biz_user_proto_rawDescData)
	})
	return file_biz_user_proto_rawDescData
}

var file_biz_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_biz_user_proto_goTypes = []interface{}{
	(*CreateUserReq)(nil),    // 0: biz.CreateUserReq
	(*CreateUserRsp)(nil),    // 1: biz.CreateUserRsp
	(*GetUserReq)(nil),       // 2: biz.GetUserReq
	(*GetUserRsp)(nil),       // 3: biz.GetUserRsp
	(*ListUserReq)(nil),      // 4: biz.ListUserReq
	(*ListUserRsp)(nil),      // 5: biz.ListUserRsp
	(*ListUserRsp_User)(nil), // 6: biz.ListUserRsp.User
}
var file_biz_user_proto_depIdxs = []int32{
	6, // 0: biz.ListUserRsp.list:type_name -> biz.ListUserRsp.User
	0, // 1: biz.UserService.CreateUser:input_type -> biz.CreateUserReq
	2, // 2: biz.UserService.GetUser:input_type -> biz.GetUserReq
	4, // 3: biz.UserService.ListUser:input_type -> biz.ListUserReq
	1, // 4: biz.UserService.CreateUser:output_type -> biz.CreateUserRsp
	3, // 5: biz.UserService.GetUser:output_type -> biz.GetUserRsp
	5, // 6: biz.UserService.ListUser:output_type -> biz.ListUserRsp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_biz_user_proto_init() }
func file_biz_user_proto_init() {
	if File_biz_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biz_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReq); i {
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
		file_biz_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRsp); i {
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
		file_biz_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReq); i {
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
		file_biz_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRsp); i {
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
		file_biz_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserReq); i {
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
		file_biz_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRsp); i {
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
		file_biz_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRsp_User); i {
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
			RawDescriptor: file_biz_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biz_user_proto_goTypes,
		DependencyIndexes: file_biz_user_proto_depIdxs,
		MessageInfos:      file_biz_user_proto_msgTypes,
	}.Build()
	File_biz_user_proto = out.File
	file_biz_user_proto_rawDesc = nil
	file_biz_user_proto_goTypes = nil
	file_biz_user_proto_depIdxs = nil
}
