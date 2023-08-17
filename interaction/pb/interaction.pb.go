// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: pb/interaction.proto

package interaction

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

type ActionType int32

const (
	ActionType_PublishComment ActionType = 0
	ActionType_DelComment     ActionType = 1
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "PublishComment",
		1: "DelComment",
	}
	ActionType_value = map[string]int32{
		"PublishComment": 0,
		"DelComment":     1,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_interaction_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_pb_interaction_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{0}
}

type FavoriteActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId    int64 `protobuf:"varint,1,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
	ActionType int64 `protobuf:"varint,2,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (x *FavoriteActionReq) Reset() {
	*x = FavoriteActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_interaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteActionReq) ProtoMessage() {}

func (x *FavoriteActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_interaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteActionReq.ProtoReflect.Descriptor instead.
func (*FavoriteActionReq) Descriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{0}
}

func (x *FavoriteActionReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *FavoriteActionReq) GetActionType() int64 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_interaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pb_interaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{1}
}

type FavoriteListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *FavoriteListReq) Reset() {
	*x = FavoriteListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_interaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListReq) ProtoMessage() {}

func (x *FavoriteListReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_interaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListReq.ProtoReflect.Descriptor instead.
func (*FavoriteListReq) Descriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{2}
}

func (x *FavoriteListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FavoriteListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author        *User  `protobuf:"bytes,1,opt,name=Author,proto3" json:"Author,omitempty"`                // 视频作者信息
	CommentCount  int64  `protobuf:"varint,2,opt,name=CommentCount,proto3" json:"CommentCount,omitempty"`   // 视频的评论总数
	CoverURL      string `protobuf:"bytes,3,opt,name=CoverURL,proto3" json:"CoverURL,omitempty"`            // 视频封面地址
	FavoriteCount int64  `protobuf:"varint,4,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"` // 视频的点赞总数
	ID            int64  `protobuf:"varint,5,opt,name=ID,proto3" json:"ID,omitempty"`                       // 视频唯一标识
	IsFavorite    bool   `protobuf:"varint,6,opt,name=IsFavorite,proto3" json:"IsFavorite,omitempty"`       // true-已点赞，false-未点赞
	PlayURL       string `protobuf:"bytes,7,opt,name=PlayURL,proto3" json:"PlayURL,omitempty"`              // 视频播放地址
	Title         string `protobuf:"bytes,8,opt,name=Title,proto3" json:"Title,omitempty"`                  // 视频标题
}

func (x *FavoriteListResp) Reset() {
	*x = FavoriteListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_interaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListResp) ProtoMessage() {}

func (x *FavoriteListResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_interaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListResp.ProtoReflect.Descriptor instead.
func (*FavoriteListResp) Descriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{3}
}

func (x *FavoriteListResp) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *FavoriteListResp) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *FavoriteListResp) GetCoverURL() string {
	if x != nil {
		return x.CoverURL
	}
	return ""
}

func (x *FavoriteListResp) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *FavoriteListResp) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *FavoriteListResp) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *FavoriteListResp) GetPlayURL() string {
	if x != nil {
		return x.PlayURL
	}
	return ""
}

func (x *FavoriteListResp) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar          string `protobuf:"bytes,1,opt,name=Avatar,proto3" json:"Avatar,omitempty"`                   // 用户头像
	BackgroundImage string `protobuf:"bytes,2,opt,name=BackgroundImage,proto3" json:"BackgroundImage,omitempty"` // 用户个人页顶部大图
	FavoriteCount   int64  `protobuf:"varint,3,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`    // 喜欢数
	FollowCount     int64  `protobuf:"varint,4,opt,name=FollowCount,proto3" json:"FollowCount,omitempty"`        // 关注总数
	FollowerCount   int64  `protobuf:"varint,5,opt,name=FollowerCount,proto3" json:"FollowerCount,omitempty"`    // 粉丝总数
	ID              int64  `protobuf:"varint,6,opt,name=ID,proto3" json:"ID,omitempty"`                          // 用户id
	IsFollow        bool   `protobuf:"varint,7,opt,name=IsFollow,proto3" json:"IsFollow,omitempty"`              // true-已关注，false-未关注
	Name            string `protobuf:"bytes,8,opt,name=Name,proto3" json:"Name,omitempty"`                       // 用户名称
	Signature       string `protobuf:"bytes,9,opt,name=Signature,proto3" json:"Signature,omitempty"`             // 个人简介
	TotalFavorited  string `protobuf:"bytes,10,opt,name=TotalFavorited,proto3" json:"TotalFavorited,omitempty"`  // 获赞数量
	WorkCount       int64  `protobuf:"varint,11,opt,name=WorkCount,proto3" json:"WorkCount,omitempty"`           // 作品数
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_interaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pb_interaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *User) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *User) GetTotalFavorited() string {
	if x != nil {
		return x.TotalFavorited
	}
	return ""
}

func (x *User) GetWorkCount() int64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

type CommentActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId     int64      `protobuf:"varint,1,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
	ActionType  ActionType `protobuf:"varint,2,opt,name=ActionType,proto3,enum=interaction.ActionType" json:"ActionType,omitempty"` // 0 发布评论， 1 删除评论
	CommentText string     `protobuf:"bytes,3,opt,name=CommentText,proto3" json:"CommentText,omitempty"`                            // 发布评论带上字段
	CommentId   int64      `protobuf:"varint,4,opt,name=CommentId,proto3" json:"CommentId,omitempty"`                               // 删除评论带上字段
}

func (x *CommentActionReq) Reset() {
	*x = CommentActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_interaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentActionReq) ProtoMessage() {}

func (x *CommentActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_interaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentActionReq.ProtoReflect.Descriptor instead.
func (*CommentActionReq) Descriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{5}
}

func (x *CommentActionReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *CommentActionReq) GetActionType() ActionType {
	if x != nil {
		return x.ActionType
	}
	return ActionType_PublishComment
}

func (x *CommentActionReq) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

func (x *CommentActionReq) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

type CommentActionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId int64  `protobuf:"varint,1,opt,name=CommentId,proto3" json:"CommentId,omitempty"` // 评论的id
	Author    *User  `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`        // 发表评论的用户
	Content   string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`      // 评论内容
	CreateAt  string `protobuf:"bytes,4,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`    // 发布时间
}

func (x *CommentActionResp) Reset() {
	*x = CommentActionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_interaction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentActionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentActionResp) ProtoMessage() {}

func (x *CommentActionResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_interaction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentActionResp.ProtoReflect.Descriptor instead.
func (*CommentActionResp) Descriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{6}
}

func (x *CommentActionResp) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *CommentActionResp) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *CommentActionResp) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentActionResp) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

type CommentListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=VideoId,proto3" json:"VideoId,omitempty"` // 视频id
}

func (x *CommentListReq) Reset() {
	*x = CommentListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_interaction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListReq) ProtoMessage() {}

func (x *CommentListReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_interaction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListReq.ProtoReflect.Descriptor instead.
func (*CommentListReq) Descriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{7}
}

func (x *CommentListReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type CommentListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentList []*CommentActionResp `protobuf:"bytes,1,rep,name=CommentList,proto3" json:"CommentList,omitempty"` // 评论列表
}

func (x *CommentListResp) Reset() {
	*x = CommentListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_interaction_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListResp) ProtoMessage() {}

func (x *CommentListResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_interaction_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListResp.ProtoReflect.Descriptor instead.
func (*CommentListResp) Descriptor() ([]byte, []int) {
	return file_pb_interaction_proto_rawDescGZIP(), []int{8}
}

func (x *CommentListResp) GetCommentList() []*CommentActionResp {
	if x != nil {
		return x.CommentList
	}
	return nil
}

var File_pb_interaction_proto protoreflect.FileDescriptor

var file_pb_interaction_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x4d, 0x0a, 0x11, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x29, 0x0a, 0x0f, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x83, 0x02, 0x0a, 0x10, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x06, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x50, 0x6c, 0x61, 0x79, 0x55, 0x52, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50,
	0x6c, 0x61, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xda, 0x02, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x28, 0x0a,
	0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x57,
	0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x57, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x10, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x92, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x49, 0x64, 0x22, 0x53, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x30, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x32, 0xba, 0x02, 0x0a, 0x0b, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x4b, 0x0a, 0x0c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4e, 0x0a, 0x0d,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_interaction_proto_rawDescOnce sync.Once
	file_pb_interaction_proto_rawDescData = file_pb_interaction_proto_rawDesc
)

func file_pb_interaction_proto_rawDescGZIP() []byte {
	file_pb_interaction_proto_rawDescOnce.Do(func() {
		file_pb_interaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_interaction_proto_rawDescData)
	})
	return file_pb_interaction_proto_rawDescData
}

var file_pb_interaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_interaction_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_interaction_proto_goTypes = []interface{}{
	(ActionType)(0),           // 0: interaction.ActionType
	(*FavoriteActionReq)(nil), // 1: interaction.FavoriteActionReq
	(*Empty)(nil),             // 2: interaction.Empty
	(*FavoriteListReq)(nil),   // 3: interaction.FavoriteListReq
	(*FavoriteListResp)(nil),  // 4: interaction.FavoriteListResp
	(*User)(nil),              // 5: interaction.User
	(*CommentActionReq)(nil),  // 6: interaction.CommentActionReq
	(*CommentActionResp)(nil), // 7: interaction.CommentActionResp
	(*CommentListReq)(nil),    // 8: interaction.CommentListReq
	(*CommentListResp)(nil),   // 9: interaction.CommentListResp
}
var file_pb_interaction_proto_depIdxs = []int32{
	5, // 0: interaction.FavoriteListResp.Author:type_name -> interaction.User
	0, // 1: interaction.CommentActionReq.ActionType:type_name -> interaction.ActionType
	5, // 2: interaction.CommentActionResp.Author:type_name -> interaction.User
	7, // 3: interaction.CommentListResp.CommentList:type_name -> interaction.CommentActionResp
	1, // 4: interaction.Interaction.FavoriteAction:input_type -> interaction.FavoriteActionReq
	3, // 5: interaction.Interaction.FavoriteList:input_type -> interaction.FavoriteListReq
	6, // 6: interaction.Interaction.CommentAction:input_type -> interaction.CommentActionReq
	8, // 7: interaction.Interaction.CommentList:input_type -> interaction.CommentListReq
	2, // 8: interaction.Interaction.FavoriteAction:output_type -> interaction.Empty
	4, // 9: interaction.Interaction.FavoriteList:output_type -> interaction.FavoriteListResp
	7, // 10: interaction.Interaction.CommentAction:output_type -> interaction.CommentActionResp
	9, // 11: interaction.Interaction.CommentList:output_type -> interaction.CommentListResp
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pb_interaction_proto_init() }
func file_pb_interaction_proto_init() {
	if File_pb_interaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_interaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteActionReq); i {
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
		file_pb_interaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_pb_interaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListReq); i {
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
		file_pb_interaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListResp); i {
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
		file_pb_interaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_pb_interaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentActionReq); i {
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
		file_pb_interaction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentActionResp); i {
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
		file_pb_interaction_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListReq); i {
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
		file_pb_interaction_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListResp); i {
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
			RawDescriptor: file_pb_interaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_interaction_proto_goTypes,
		DependencyIndexes: file_pb_interaction_proto_depIdxs,
		EnumInfos:         file_pb_interaction_proto_enumTypes,
		MessageInfos:      file_pb_interaction_proto_msgTypes,
	}.Build()
	File_pb_interaction_proto = out.File
	file_pb_interaction_proto_rawDesc = nil
	file_pb_interaction_proto_goTypes = nil
	file_pb_interaction_proto_depIdxs = nil
}