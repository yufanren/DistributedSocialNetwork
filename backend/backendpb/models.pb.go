// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: backend/backendpb/models.proto

package backendpb

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

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{1}
}

func (x *AuthResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Author  string `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{2}
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Blog) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type AddBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *AddBlogRequest) Reset() {
	*x = AddBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlogRequest) ProtoMessage() {}

func (x *AddBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlogRequest.ProtoReflect.Descriptor instead.
func (*AddBlogRequest) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{3}
}

func (x *AddBlogRequest) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type AddBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId int32 `protobuf:"varint,1,opt,name=BlogId,proto3" json:"BlogId,omitempty"`
}

func (x *AddBlogResponse) Reset() {
	*x = AddBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlogResponse) ProtoMessage() {}

func (x *AddBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlogResponse.ProtoReflect.Descriptor instead.
func (*AddBlogResponse) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{4}
}

func (x *AddBlogResponse) GetBlogId() int32 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{5}
}

type ListBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int32 `protobuf:"varint,1,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListBlogRequest) Reset() {
	*x = ListBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogRequest) ProtoMessage() {}

func (x *ListBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogRequest.ProtoReflect.Descriptor instead.
func (*ListBlogRequest) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{6}
}

func (x *ListBlogRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListBlogRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32   `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Blogs []*Blog `protobuf:"bytes,2,rep,name=Blogs,proto3" json:"Blogs,omitempty"`
}

func (x *ListBlogResponse) Reset() {
	*x = ListBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogResponse) ProtoMessage() {}

func (x *ListBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogResponse.ProtoReflect.Descriptor instead.
func (*ListBlogResponse) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{7}
}

func (x *ListBlogResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListBlogResponse) GetBlogs() []*Blog {
	if x != nil {
		return x.Blogs
	}
	return nil
}

type GetBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId int32 `protobuf:"varint,1,opt,name=BlogId,proto3" json:"BlogId,omitempty"`
}

func (x *GetBlogRequest) Reset() {
	*x = GetBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogRequest) ProtoMessage() {}

func (x *GetBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogRequest.ProtoReflect.Descriptor instead.
func (*GetBlogRequest) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{8}
}

func (x *GetBlogRequest) GetBlogId() int32 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

type GetBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=Blog,proto3" json:"Blog,omitempty"`
}

func (x *GetBlogResponse) Reset() {
	*x = GetBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogResponse) ProtoMessage() {}

func (x *GetBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogResponse.ProtoReflect.Descriptor instead.
func (*GetBlogResponse) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{9}
}

func (x *GetBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent   string `protobuf:"bytes,1,opt,name=Parent,proto3" json:"Parent,omitempty"`
	Follower string `protobuf:"bytes,2,opt,name=Follower,proto3" json:"Follower,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{10}
}

func (x *FollowRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *FollowRequest) GetFollower() string {
	if x != nil {
		return x.Follower
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res bool `protobuf:"varint,1,opt,name=Res,proto3" json:"Res,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{11}
}

func (x *CheckResponse) GetRes() bool {
	if x != nil {
		return x.Res
	}
	return false
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{12}
}

func (x *UserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	PageNum  int32  `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *UserBlogRequest) Reset() {
	*x = UserBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBlogRequest) ProtoMessage() {}

func (x *UserBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBlogRequest.ProtoReflect.Descriptor instead.
func (*UserBlogRequest) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{13}
}

func (x *UserBlogRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserBlogRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *UserBlogRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   int32    `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Parents []string `protobuf:"bytes,2,rep,name=Parents,proto3" json:"Parents,omitempty"`
}

func (x *ListUserResponse) Reset() {
	*x = ListUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_backendpb_models_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponse) ProtoMessage() {}

func (x *ListUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_backendpb_models_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponse.ProtoReflect.Descriptor instead.
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return file_backend_backendpb_models_proto_rawDescGZIP(), []int{14}
}

func (x *ListUserResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListUserResponse) GetParents() []string {
	if x != nil {
		return x.Parents
	}
	return nil
}

var File_backend_backendpb_models_proto protoreflect.FileDescriptor

var file_backend_backendpb_models_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x22, 0x45, 0x0a, 0x0b, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4e, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x6c,
	0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22,
	0x29, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x4f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x05, 0x42, 0x6c, 0x6f,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x42, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x42, 0x6c,
	0x6f, 0x67, 0x22, 0x43, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x52, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xfa,
	0x05, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x3a, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x19, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64,
	0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x19, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x18,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x18, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70,
	0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x16,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x3b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_backend_backendpb_models_proto_rawDescOnce sync.Once
	file_backend_backendpb_models_proto_rawDescData = file_backend_backendpb_models_proto_rawDesc
)

func file_backend_backendpb_models_proto_rawDescGZIP() []byte {
	file_backend_backendpb_models_proto_rawDescOnce.Do(func() {
		file_backend_backendpb_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_backendpb_models_proto_rawDescData)
	})
	return file_backend_backendpb_models_proto_rawDescData
}

var file_backend_backendpb_models_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_backend_backendpb_models_proto_goTypes = []interface{}{
	(*AuthRequest)(nil),      // 0: backendpb.AuthRequest
	(*AuthResponse)(nil),     // 1: backendpb.AuthResponse
	(*Blog)(nil),             // 2: backendpb.Blog
	(*AddBlogRequest)(nil),   // 3: backendpb.AddBlogRequest
	(*AddBlogResponse)(nil),  // 4: backendpb.AddBlogResponse
	(*Response)(nil),         // 5: backendpb.Response
	(*ListBlogRequest)(nil),  // 6: backendpb.ListBlogRequest
	(*ListBlogResponse)(nil), // 7: backendpb.ListBlogResponse
	(*GetBlogRequest)(nil),   // 8: backendpb.GetBlogRequest
	(*GetBlogResponse)(nil),  // 9: backendpb.GetBlogResponse
	(*FollowRequest)(nil),    // 10: backendpb.FollowRequest
	(*CheckResponse)(nil),    // 11: backendpb.CheckResponse
	(*UserRequest)(nil),      // 12: backendpb.UserRequest
	(*UserBlogRequest)(nil),  // 13: backendpb.UserBlogRequest
	(*ListUserResponse)(nil), // 14: backendpb.ListUserResponse
}
var file_backend_backendpb_models_proto_depIdxs = []int32{
	2,  // 0: backendpb.AddBlogRequest.blog:type_name -> backendpb.Blog
	2,  // 1: backendpb.ListBlogResponse.Blogs:type_name -> backendpb.Blog
	2,  // 2: backendpb.GetBlogResponse.Blog:type_name -> backendpb.Blog
	0,  // 3: backendpb.Backend.Login:input_type -> backendpb.AuthRequest
	0,  // 4: backendpb.Backend.Register:input_type -> backendpb.AuthRequest
	3,  // 5: backendpb.Backend.AddBlog:input_type -> backendpb.AddBlogRequest
	6,  // 6: backendpb.Backend.ListBlog:input_type -> backendpb.ListBlogRequest
	8,  // 7: backendpb.Backend.GetBlog:input_type -> backendpb.GetBlogRequest
	10, // 8: backendpb.Backend.Follow:input_type -> backendpb.FollowRequest
	10, // 9: backendpb.Backend.IsFollow:input_type -> backendpb.FollowRequest
	13, // 10: backendpb.Backend.GetUserBlog:input_type -> backendpb.UserBlogRequest
	12, // 11: backendpb.Backend.GetUserFollowed:input_type -> backendpb.UserRequest
	12, // 12: backendpb.Backend.GetUserFollower:input_type -> backendpb.UserRequest
	13, // 13: backendpb.Backend.GetFollowBlog:input_type -> backendpb.UserBlogRequest
	1,  // 14: backendpb.Backend.Login:output_type -> backendpb.AuthResponse
	1,  // 15: backendpb.Backend.Register:output_type -> backendpb.AuthResponse
	4,  // 16: backendpb.Backend.AddBlog:output_type -> backendpb.AddBlogResponse
	7,  // 17: backendpb.Backend.ListBlog:output_type -> backendpb.ListBlogResponse
	9,  // 18: backendpb.Backend.GetBlog:output_type -> backendpb.GetBlogResponse
	5,  // 19: backendpb.Backend.Follow:output_type -> backendpb.Response
	11, // 20: backendpb.Backend.IsFollow:output_type -> backendpb.CheckResponse
	7,  // 21: backendpb.Backend.GetUserBlog:output_type -> backendpb.ListBlogResponse
	14, // 22: backendpb.Backend.GetUserFollowed:output_type -> backendpb.ListUserResponse
	14, // 23: backendpb.Backend.GetUserFollower:output_type -> backendpb.ListUserResponse
	7,  // 24: backendpb.Backend.GetFollowBlog:output_type -> backendpb.ListBlogResponse
	14, // [14:25] is the sub-list for method output_type
	3,  // [3:14] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_backend_backendpb_models_proto_init() }
func file_backend_backendpb_models_proto_init() {
	if File_backend_backendpb_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_backendpb_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_backend_backendpb_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_backend_backendpb_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
		file_backend_backendpb_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBlogRequest); i {
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
		file_backend_backendpb_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBlogResponse); i {
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
		file_backend_backendpb_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_backend_backendpb_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogRequest); i {
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
		file_backend_backendpb_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogResponse); i {
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
		file_backend_backendpb_models_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogRequest); i {
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
		file_backend_backendpb_models_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogResponse); i {
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
		file_backend_backendpb_models_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_backend_backendpb_models_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
		file_backend_backendpb_models_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_backend_backendpb_models_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBlogRequest); i {
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
		file_backend_backendpb_models_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserResponse); i {
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
			RawDescriptor: file_backend_backendpb_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_backendpb_models_proto_goTypes,
		DependencyIndexes: file_backend_backendpb_models_proto_depIdxs,
		MessageInfos:      file_backend_backendpb_models_proto_msgTypes,
	}.Build()
	File_backend_backendpb_models_proto = out.File
	file_backend_backendpb_models_proto_rawDesc = nil
	file_backend_backendpb_models_proto_goTypes = nil
	file_backend_backendpb_models_proto_depIdxs = nil
}