// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: blogger.proto

package blogger_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type BlogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	WebsiteUrl  string `protobuf:"bytes,3,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
}

func (x *BlogInfo) Reset() {
	*x = BlogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogInfo) ProtoMessage() {}

func (x *BlogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogInfo.ProtoReflect.Descriptor instead.
func (*BlogInfo) Descriptor() ([]byte, []int) {
	return file_blogger_proto_rawDescGZIP(), []int{0}
}

func (x *BlogInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlogInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BlogInfo) GetWebsiteUrl() string {
	if x != nil {
		return x.WebsiteUrl
	}
	return ""
}

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Blog      *BlogInfo            `protobuf:"bytes,2,opt,name=blog,proto3" json:"blog,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[1]
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
	return file_blogger_proto_rawDescGZIP(), []int{1}
}

func (x *Blog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Blog) GetBlog() *BlogInfo {
	if x != nil {
		return x.Blog
	}
	return nil
}

func (x *Blog) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Blog) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UpdateBlogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description *wrappers.StringValue `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	WebsiteUrl  *wrappers.StringValue `protobuf:"bytes,3,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
}

func (x *UpdateBlogInfo) Reset() {
	*x = UpdateBlogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogInfo) ProtoMessage() {}

func (x *UpdateBlogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogInfo.ProtoReflect.Descriptor instead.
func (*UpdateBlogInfo) Descriptor() ([]byte, []int) {
	return file_blogger_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBlogInfo) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdateBlogInfo) GetDescription() *wrappers.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *UpdateBlogInfo) GetWebsiteUrl() *wrappers.StringValue {
	if x != nil {
		return x.WebsiteUrl
	}
	return nil
}

type CreateBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *BlogInfo `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *CreateBlogRequest) Reset() {
	*x = CreateBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogRequest) ProtoMessage() {}

func (x *CreateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogRequest.ProtoReflect.Descriptor instead.
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return file_blogger_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBlogRequest) GetBlog() *BlogInfo {
	if x != nil {
		return x.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *CreateBlogResponse) Reset() {
	*x = CreateBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogResponse) ProtoMessage() {}

func (x *CreateBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogResponse.ProtoReflect.Descriptor instead.
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return file_blogger_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type UpdateBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Blog *UpdateBlogInfo `protobuf:"bytes,2,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *UpdateBlogRequest) Reset() {
	*x = UpdateBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogRequest) ProtoMessage() {}

func (x *UpdateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogRequest.ProtoReflect.Descriptor instead.
func (*UpdateBlogRequest) Descriptor() ([]byte, []int) {
	return file_blogger_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBlogRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBlogRequest) GetBlog() *UpdateBlogInfo {
	if x != nil {
		return x.Blog
	}
	return nil
}

type DeleteBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBlogRequest) Reset() {
	*x = DeleteBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogRequest) ProtoMessage() {}

func (x *DeleteBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlogRequest) Descriptor() ([]byte, []int) {
	return file_blogger_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBlogRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBlogRequest) Reset() {
	*x = GetBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogRequest) ProtoMessage() {}

func (x *GetBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[7]
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
	return file_blogger_proto_rawDescGZIP(), []int{7}
}

func (x *GetBlogRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *GetBlogResponse) Reset() {
	*x = GetBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogResponse) ProtoMessage() {}

func (x *GetBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[8]
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
	return file_blogger_proto_rawDescGZIP(), []int{8}
}

func (x *GetBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type ListBlogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListBlogsRequest) Reset() {
	*x = ListBlogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogsRequest) ProtoMessage() {}

func (x *ListBlogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogsRequest.ProtoReflect.Descriptor instead.
func (*ListBlogsRequest) Descriptor() ([]byte, []int) {
	return file_blogger_proto_rawDescGZIP(), []int{9}
}

func (x *ListBlogsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListBlogsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListBlogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog []*Blog `protobuf:"bytes,1,rep,name=blog,proto3" json:"blog,omitempty"`
}

func (x *ListBlogsResponse) Reset() {
	*x = ListBlogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogger_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogsResponse) ProtoMessage() {}

func (x *ListBlogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogger_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogsResponse.ProtoReflect.Descriptor instead.
func (*ListBlogsResponse) Descriptor() ([]byte, []int) {
	return file_blogger_proto_rawDescGZIP(), []int{10}
}

func (x *ListBlogsResponse) GetBlog() []*Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

var File_blogger_proto protoreflect.FileDescriptor

var file_blogger_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x61, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x55, 0x72, 0x6c, 0x22, 0xb6, 0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04,
	0x62, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc1, 0x01, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x0b, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c,
	0x22, 0x3d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22,
	0x3a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x53, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2e, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67,
	0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x6c,
	0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67,
	0x22, 0x40, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x39, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f,
	0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x32, 0xe4, 0x03,
	0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x56, 0x31, 0x12, 0x61, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x5e,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1d, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x32, 0x0e, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x5b,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1d, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5a, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x12, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x6c, 0x6f, 0x67, 0x73, 0x42, 0xb2, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x65, 0x6f, 0x72, 0x67, 0x79, 0x32, 0x37, 0x2f, 0x62, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x3b, 0x62,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x92, 0x41, 0x6d, 0x12, 0x33, 0x0a, 0x0b,
	0x42, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49, 0x22, 0x1d, 0x0a, 0x08, 0x47,
	0x65, 0x6f, 0x72, 0x67, 0x79, 0x20, 0x50, 0x1a, 0x11, 0x41, 0x65, 0x67, 0x6f, 0x72, 0x61, 0x61,
	0x40, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x72, 0x75, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e,
	0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38,
	0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_blogger_proto_rawDescOnce sync.Once
	file_blogger_proto_rawDescData = file_blogger_proto_rawDesc
)

func file_blogger_proto_rawDescGZIP() []byte {
	file_blogger_proto_rawDescOnce.Do(func() {
		file_blogger_proto_rawDescData = protoimpl.X.CompressGZIP(file_blogger_proto_rawDescData)
	})
	return file_blogger_proto_rawDescData
}

var file_blogger_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_blogger_proto_goTypes = []interface{}{
	(*BlogInfo)(nil),             // 0: blogger_v1.BlogInfo
	(*Blog)(nil),                 // 1: blogger_v1.Blog
	(*UpdateBlogInfo)(nil),       // 2: blogger_v1.UpdateBlogInfo
	(*CreateBlogRequest)(nil),    // 3: blogger_v1.CreateBlogRequest
	(*CreateBlogResponse)(nil),   // 4: blogger_v1.CreateBlogResponse
	(*UpdateBlogRequest)(nil),    // 5: blogger_v1.UpdateBlogRequest
	(*DeleteBlogRequest)(nil),    // 6: blogger_v1.DeleteBlogRequest
	(*GetBlogRequest)(nil),       // 7: blogger_v1.GetBlogRequest
	(*GetBlogResponse)(nil),      // 8: blogger_v1.GetBlogResponse
	(*ListBlogsRequest)(nil),     // 9: blogger_v1.ListBlogsRequest
	(*ListBlogsResponse)(nil),    // 10: blogger_v1.ListBlogsResponse
	(*timestamp.Timestamp)(nil),  // 11: google.protobuf.Timestamp
	(*wrappers.StringValue)(nil), // 12: google.protobuf.StringValue
	(*empty.Empty)(nil),          // 13: google.protobuf.Empty
}
var file_blogger_proto_depIdxs = []int32{
	0,  // 0: blogger_v1.Blog.blog:type_name -> blogger_v1.BlogInfo
	11, // 1: blogger_v1.Blog.created_at:type_name -> google.protobuf.Timestamp
	11, // 2: blogger_v1.Blog.updated_at:type_name -> google.protobuf.Timestamp
	12, // 3: blogger_v1.UpdateBlogInfo.name:type_name -> google.protobuf.StringValue
	12, // 4: blogger_v1.UpdateBlogInfo.description:type_name -> google.protobuf.StringValue
	12, // 5: blogger_v1.UpdateBlogInfo.website_url:type_name -> google.protobuf.StringValue
	0,  // 6: blogger_v1.CreateBlogRequest.blog:type_name -> blogger_v1.BlogInfo
	1,  // 7: blogger_v1.CreateBlogResponse.blog:type_name -> blogger_v1.Blog
	2,  // 8: blogger_v1.UpdateBlogRequest.blog:type_name -> blogger_v1.UpdateBlogInfo
	1,  // 9: blogger_v1.GetBlogResponse.blog:type_name -> blogger_v1.Blog
	1,  // 10: blogger_v1.ListBlogsResponse.blog:type_name -> blogger_v1.Blog
	3,  // 11: blogger_v1.BloggerV1.CreateBlog:input_type -> blogger_v1.CreateBlogRequest
	5,  // 12: blogger_v1.BloggerV1.UpdateBlog:input_type -> blogger_v1.UpdateBlogRequest
	6,  // 13: blogger_v1.BloggerV1.DeleteBlog:input_type -> blogger_v1.DeleteBlogRequest
	7,  // 14: blogger_v1.BloggerV1.GetBlog:input_type -> blogger_v1.GetBlogRequest
	9,  // 15: blogger_v1.BloggerV1.ListBlogs:input_type -> blogger_v1.ListBlogsRequest
	4,  // 16: blogger_v1.BloggerV1.CreateBlog:output_type -> blogger_v1.CreateBlogResponse
	13, // 17: blogger_v1.BloggerV1.UpdateBlog:output_type -> google.protobuf.Empty
	13, // 18: blogger_v1.BloggerV1.DeleteBlog:output_type -> google.protobuf.Empty
	8,  // 19: blogger_v1.BloggerV1.GetBlog:output_type -> blogger_v1.GetBlogResponse
	10, // 20: blogger_v1.BloggerV1.ListBlogs:output_type -> blogger_v1.ListBlogsResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_blogger_proto_init() }
func file_blogger_proto_init() {
	if File_blogger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blogger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogInfo); i {
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
		file_blogger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_blogger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogInfo); i {
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
		file_blogger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogRequest); i {
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
		file_blogger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogResponse); i {
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
		file_blogger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogRequest); i {
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
		file_blogger_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlogRequest); i {
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
		file_blogger_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_blogger_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_blogger_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogsRequest); i {
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
		file_blogger_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogsResponse); i {
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
			RawDescriptor: file_blogger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blogger_proto_goTypes,
		DependencyIndexes: file_blogger_proto_depIdxs,
		MessageInfos:      file_blogger_proto_msgTypes,
	}.Build()
	File_blogger_proto = out.File
	file_blogger_proto_rawDesc = nil
	file_blogger_proto_goTypes = nil
	file_blogger_proto_depIdxs = nil
}