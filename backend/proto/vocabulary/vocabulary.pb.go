// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: vocabulary.proto

package vocabulary

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

type SetAlreadyLearnedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	WordId           string `protobuf:"bytes,2,opt,name=word_id,json=wordId,proto3" json:"word_id,omitempty"`
	IsAlreadyLearned bool   `protobuf:"varint,3,opt,name=is_already_learned,json=isAlreadyLearned,proto3" json:"is_already_learned,omitempty"`
}

func (x *SetAlreadyLearnedRequest) Reset() {
	*x = SetAlreadyLearnedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAlreadyLearnedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAlreadyLearnedRequest) ProtoMessage() {}

func (x *SetAlreadyLearnedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAlreadyLearnedRequest.ProtoReflect.Descriptor instead.
func (*SetAlreadyLearnedRequest) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{0}
}

func (x *SetAlreadyLearnedRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SetAlreadyLearnedRequest) GetWordId() string {
	if x != nil {
		return x.WordId
	}
	return ""
}

func (x *SetAlreadyLearnedRequest) GetIsAlreadyLearned() bool {
	if x != nil {
		return x.IsAlreadyLearned
	}
	return false
}

type SetAlreadyLearnedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetAlreadyLearnedResponse) Reset() {
	*x = SetAlreadyLearnedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAlreadyLearnedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAlreadyLearnedResponse) ProtoMessage() {}

func (x *SetAlreadyLearnedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAlreadyLearnedResponse.ProtoReflect.Descriptor instead.
func (*SetAlreadyLearnedResponse) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{1}
}

type GetChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetChallengeRequest) Reset() {
	*x = GetChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChallengeRequest) ProtoMessage() {}

func (x *GetChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChallengeRequest.ProtoReflect.Descriptor instead.
func (*GetChallengeRequest) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{2}
}

func (x *GetChallengeRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WordId       string `protobuf:"bytes,1,opt,name=word_id,json=wordId,proto3" json:"word_id,omitempty"`
	LearningStep int32  `protobuf:"varint,2,opt,name=learning_step,json=learningStep,proto3" json:"learning_step,omitempty"`
}

func (x *GetChallengeResponse) Reset() {
	*x = GetChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChallengeResponse) ProtoMessage() {}

func (x *GetChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChallengeResponse.ProtoReflect.Descriptor instead.
func (*GetChallengeResponse) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{3}
}

func (x *GetChallengeResponse) GetWordId() string {
	if x != nil {
		return x.WordId
	}
	return ""
}

func (x *GetChallengeResponse) GetLearningStep() int32 {
	if x != nil {
		return x.LearningStep
	}
	return 0
}

type PromoteWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	WordId string `protobuf:"bytes,2,opt,name=word_id,json=wordId,proto3" json:"word_id,omitempty"`
}

func (x *PromoteWordRequest) Reset() {
	*x = PromoteWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromoteWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromoteWordRequest) ProtoMessage() {}

func (x *PromoteWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromoteWordRequest.ProtoReflect.Descriptor instead.
func (*PromoteWordRequest) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{4}
}

func (x *PromoteWordRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PromoteWordRequest) GetWordId() string {
	if x != nil {
		return x.WordId
	}
	return ""
}

type PromoteWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PromoteWordResponse) Reset() {
	*x = PromoteWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromoteWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromoteWordResponse) ProtoMessage() {}

func (x *PromoteWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromoteWordResponse.ProtoReflect.Descriptor instead.
func (*PromoteWordResponse) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{5}
}

type ResistWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	WordId string `protobuf:"bytes,2,opt,name=word_id,json=wordId,proto3" json:"word_id,omitempty"`
}

func (x *ResistWordRequest) Reset() {
	*x = ResistWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResistWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResistWordRequest) ProtoMessage() {}

func (x *ResistWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResistWordRequest.ProtoReflect.Descriptor instead.
func (*ResistWordRequest) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{6}
}

func (x *ResistWordRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ResistWordRequest) GetWordId() string {
	if x != nil {
		return x.WordId
	}
	return ""
}

type ResistWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResistWordResponse) Reset() {
	*x = ResistWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResistWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResistWordResponse) ProtoMessage() {}

func (x *ResistWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResistWordResponse.ProtoReflect.Descriptor instead.
func (*ResistWordResponse) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{7}
}

var File_vocabulary_proto protoreflect.FileDescriptor

var file_vocabulary_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x22, 0x7a,
	0x0a, 0x18, 0x53, 0x65, 0x74, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x4c, 0x65, 0x61, 0x72,
	0x6e, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x69, 0x73, 0x5f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x6c, 0x65, 0x61, 0x72, 0x6e,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x41, 0x6c, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x65,
	0x74, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x65, 0x70, 0x22, 0x46, 0x0a,
	0x12, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77,
	0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x11,
	0x52, 0x65, 0x73, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x6f, 0x72,
	0x64, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xed, 0x02, 0x0a, 0x11, 0x56, 0x6f,
	0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x53, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12,
	0x1f, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x69, 0x73, 0x74,
	0x57, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72,
	0x79, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79,
	0x2e, 0x52, 0x65, 0x73, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x41, 0x6c, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x24, 0x2e, 0x76, 0x6f, 0x63,
	0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x6c, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x65,
	0x74, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6b, 0x66, 0x63, 0x63, 0x63, 0x63, 0x63,
	0x63, 0x2f, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vocabulary_proto_rawDescOnce sync.Once
	file_vocabulary_proto_rawDescData = file_vocabulary_proto_rawDesc
)

func file_vocabulary_proto_rawDescGZIP() []byte {
	file_vocabulary_proto_rawDescOnce.Do(func() {
		file_vocabulary_proto_rawDescData = protoimpl.X.CompressGZIP(file_vocabulary_proto_rawDescData)
	})
	return file_vocabulary_proto_rawDescData
}

var file_vocabulary_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_vocabulary_proto_goTypes = []interface{}{
	(*SetAlreadyLearnedRequest)(nil),  // 0: vocabulary.SetAlreadyLearnedRequest
	(*SetAlreadyLearnedResponse)(nil), // 1: vocabulary.SetAlreadyLearnedResponse
	(*GetChallengeRequest)(nil),       // 2: vocabulary.GetChallengeRequest
	(*GetChallengeResponse)(nil),      // 3: vocabulary.GetChallengeResponse
	(*PromoteWordRequest)(nil),        // 4: vocabulary.PromoteWordRequest
	(*PromoteWordResponse)(nil),       // 5: vocabulary.PromoteWordResponse
	(*ResistWordRequest)(nil),         // 6: vocabulary.ResistWordRequest
	(*ResistWordResponse)(nil),        // 7: vocabulary.ResistWordResponse
}
var file_vocabulary_proto_depIdxs = []int32{
	2, // 0: vocabulary.VocabularyService.GetChallenge:input_type -> vocabulary.GetChallengeRequest
	4, // 1: vocabulary.VocabularyService.PromoteWord:input_type -> vocabulary.PromoteWordRequest
	6, // 2: vocabulary.VocabularyService.ResistWord:input_type -> vocabulary.ResistWordRequest
	0, // 3: vocabulary.VocabularyService.SetAlreadyLearned:input_type -> vocabulary.SetAlreadyLearnedRequest
	3, // 4: vocabulary.VocabularyService.GetChallenge:output_type -> vocabulary.GetChallengeResponse
	5, // 5: vocabulary.VocabularyService.PromoteWord:output_type -> vocabulary.PromoteWordResponse
	7, // 6: vocabulary.VocabularyService.ResistWord:output_type -> vocabulary.ResistWordResponse
	1, // 7: vocabulary.VocabularyService.SetAlreadyLearned:output_type -> vocabulary.SetAlreadyLearnedResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vocabulary_proto_init() }
func file_vocabulary_proto_init() {
	if File_vocabulary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vocabulary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAlreadyLearnedRequest); i {
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
		file_vocabulary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAlreadyLearnedResponse); i {
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
		file_vocabulary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChallengeRequest); i {
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
		file_vocabulary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChallengeResponse); i {
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
		file_vocabulary_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromoteWordRequest); i {
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
		file_vocabulary_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromoteWordResponse); i {
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
		file_vocabulary_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResistWordRequest); i {
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
		file_vocabulary_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResistWordResponse); i {
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
			RawDescriptor: file_vocabulary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vocabulary_proto_goTypes,
		DependencyIndexes: file_vocabulary_proto_depIdxs,
		MessageInfos:      file_vocabulary_proto_msgTypes,
	}.Build()
	File_vocabulary_proto = out.File
	file_vocabulary_proto_rawDesc = nil
	file_vocabulary_proto_goTypes = nil
	file_vocabulary_proto_depIdxs = nil
}
