// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: proto/filesystem.proto

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

type RegisterNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeAddress string `protobuf:"bytes,1,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
}

func (x *RegisterNodeRequest) Reset() {
	*x = RegisterNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filesystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeRequest) ProtoMessage() {}

func (x *RegisterNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filesystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeRequest.ProtoReflect.Descriptor instead.
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return file_proto_filesystem_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterNodeRequest) GetNodeAddress() string {
	if x != nil {
		return x.NodeAddress
	}
	return ""
}

type RegisterNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterNodeResponse) Reset() {
	*x = RegisterNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filesystem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeResponse) ProtoMessage() {}

func (x *RegisterNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filesystem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeResponse.ProtoReflect.Descriptor instead.
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return file_proto_filesystem_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterNodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetNodesForChunksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId      string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`                 // Unique identifier for the file
	TotalChunks int32  `protobuf:"varint,2,opt,name=total_chunks,json=totalChunks,proto3" json:"total_chunks,omitempty"` // Total number of chunks to be uploaded
	FileType    string `protobuf:"bytes,3,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`           // File Type
}

func (x *GetNodesForChunksRequest) Reset() {
	*x = GetNodesForChunksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filesystem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesForChunksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesForChunksRequest) ProtoMessage() {}

func (x *GetNodesForChunksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filesystem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesForChunksRequest.ProtoReflect.Descriptor instead.
func (*GetNodesForChunksRequest) Descriptor() ([]byte, []int) {
	return file_proto_filesystem_proto_rawDescGZIP(), []int{2}
}

func (x *GetNodesForChunksRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *GetNodesForChunksRequest) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

func (x *GetNodesForChunksRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

type ChunkNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId     int32  `protobuf:"varint,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`            // ID of the chunk
	NodeAddress string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"` // Address of the node where this chunk should be uploaded
}

func (x *ChunkNodeInfo) Reset() {
	*x = ChunkNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filesystem_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkNodeInfo) ProtoMessage() {}

func (x *ChunkNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filesystem_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkNodeInfo.ProtoReflect.Descriptor instead.
func (*ChunkNodeInfo) Descriptor() ([]byte, []int) {
	return file_proto_filesystem_proto_rawDescGZIP(), []int{3}
}

func (x *ChunkNodeInfo) GetChunkId() int32 {
	if x != nil {
		return x.ChunkId
	}
	return 0
}

func (x *ChunkNodeInfo) GetNodeAddress() string {
	if x != nil {
		return x.NodeAddress
	}
	return ""
}

type GetNodesForChunksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*ChunkNodeInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"` // List of node addresses for each chunk
}

func (x *GetNodesForChunksResponse) Reset() {
	*x = GetNodesForChunksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filesystem_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesForChunksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesForChunksResponse) ProtoMessage() {}

func (x *GetNodesForChunksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filesystem_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesForChunksResponse.ProtoReflect.Descriptor instead.
func (*GetNodesForChunksResponse) Descriptor() ([]byte, []int) {
	return file_proto_filesystem_proto_rawDescGZIP(), []int{4}
}

func (x *GetNodesForChunksResponse) GetNodes() []*ChunkNodeInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type GetChunkLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *GetChunkLocationsRequest) Reset() {
	*x = GetChunkLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filesystem_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunkLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkLocationsRequest) ProtoMessage() {}

func (x *GetChunkLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filesystem_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkLocationsRequest.ProtoReflect.Descriptor instead.
func (*GetChunkLocationsRequest) Descriptor() ([]byte, []int) {
	return file_proto_filesystem_proto_rawDescGZIP(), []int{5}
}

func (x *GetChunkLocationsRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type GetChunkLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunks   []*ChunkLocationInfo `protobuf:"bytes,1,rep,name=chunks,proto3" json:"chunks,omitempty"`
	FileType string               `protobuf:"bytes,2,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
}

func (x *GetChunkLocationsResponse) Reset() {
	*x = GetChunkLocationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filesystem_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunkLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkLocationsResponse) ProtoMessage() {}

func (x *GetChunkLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filesystem_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkLocationsResponse.ProtoReflect.Descriptor instead.
func (*GetChunkLocationsResponse) Descriptor() ([]byte, []int) {
	return file_proto_filesystem_proto_rawDescGZIP(), []int{6}
}

func (x *GetChunkLocationsResponse) GetChunks() []*ChunkLocationInfo {
	if x != nil {
		return x.Chunks
	}
	return nil
}

func (x *GetChunkLocationsResponse) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

type ChunkLocationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId int32    `protobuf:"varint,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Nodes   []string `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ChunkLocationInfo) Reset() {
	*x = ChunkLocationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filesystem_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkLocationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkLocationInfo) ProtoMessage() {}

func (x *ChunkLocationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filesystem_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkLocationInfo.ProtoReflect.Descriptor instead.
func (*ChunkLocationInfo) Descriptor() ([]byte, []int) {
	return file_proto_filesystem_proto_rawDescGZIP(), []int{7}
}

func (x *ChunkLocationInfo) GetChunkId() int32 {
	if x != nil {
		return x.ChunkId
	}
	return 0
}

func (x *ChunkLocationInfo) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_proto_filesystem_proto protoreflect.FileDescriptor

var file_proto_filesystem_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x22, 0x38, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x30,
	0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x73, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x0d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x4c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x46, 0x6f, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x22, 0x33, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x11, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x32, 0xa7,
	0x02, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x51, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x1f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x46, 0x6f, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x24, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x46,
	0x6f, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x62, 0x72, 0x65, 0x65,
	0x7a, 0x65, 0x46, 0x53, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_filesystem_proto_rawDescOnce sync.Once
	file_proto_filesystem_proto_rawDescData = file_proto_filesystem_proto_rawDesc
)

func file_proto_filesystem_proto_rawDescGZIP() []byte {
	file_proto_filesystem_proto_rawDescOnce.Do(func() {
		file_proto_filesystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_filesystem_proto_rawDescData)
	})
	return file_proto_filesystem_proto_rawDescData
}

var file_proto_filesystem_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_filesystem_proto_goTypes = []any{
	(*RegisterNodeRequest)(nil),       // 0: filesystem.RegisterNodeRequest
	(*RegisterNodeResponse)(nil),      // 1: filesystem.RegisterNodeResponse
	(*GetNodesForChunksRequest)(nil),  // 2: filesystem.GetNodesForChunksRequest
	(*ChunkNodeInfo)(nil),             // 3: filesystem.ChunkNodeInfo
	(*GetNodesForChunksResponse)(nil), // 4: filesystem.GetNodesForChunksResponse
	(*GetChunkLocationsRequest)(nil),  // 5: filesystem.GetChunkLocationsRequest
	(*GetChunkLocationsResponse)(nil), // 6: filesystem.GetChunkLocationsResponse
	(*ChunkLocationInfo)(nil),         // 7: filesystem.ChunkLocationInfo
}
var file_proto_filesystem_proto_depIdxs = []int32{
	3, // 0: filesystem.GetNodesForChunksResponse.nodes:type_name -> filesystem.ChunkNodeInfo
	7, // 1: filesystem.GetChunkLocationsResponse.chunks:type_name -> filesystem.ChunkLocationInfo
	0, // 2: filesystem.ManagerService.RegisterNode:input_type -> filesystem.RegisterNodeRequest
	2, // 3: filesystem.ManagerService.GetNodesForChunks:input_type -> filesystem.GetNodesForChunksRequest
	5, // 4: filesystem.ManagerService.GetChunkLocations:input_type -> filesystem.GetChunkLocationsRequest
	1, // 5: filesystem.ManagerService.RegisterNode:output_type -> filesystem.RegisterNodeResponse
	4, // 6: filesystem.ManagerService.GetNodesForChunks:output_type -> filesystem.GetNodesForChunksResponse
	6, // 7: filesystem.ManagerService.GetChunkLocations:output_type -> filesystem.GetChunkLocationsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_filesystem_proto_init() }
func file_proto_filesystem_proto_init() {
	if File_proto_filesystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_filesystem_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterNodeRequest); i {
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
		file_proto_filesystem_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterNodeResponse); i {
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
		file_proto_filesystem_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetNodesForChunksRequest); i {
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
		file_proto_filesystem_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ChunkNodeInfo); i {
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
		file_proto_filesystem_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetNodesForChunksResponse); i {
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
		file_proto_filesystem_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetChunkLocationsRequest); i {
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
		file_proto_filesystem_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetChunkLocationsResponse); i {
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
		file_proto_filesystem_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ChunkLocationInfo); i {
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
			RawDescriptor: file_proto_filesystem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_filesystem_proto_goTypes,
		DependencyIndexes: file_proto_filesystem_proto_depIdxs,
		MessageInfos:      file_proto_filesystem_proto_msgTypes,
	}.Build()
	File_proto_filesystem_proto = out.File
	file_proto_filesystem_proto_rawDesc = nil
	file_proto_filesystem_proto_goTypes = nil
	file_proto_filesystem_proto_depIdxs = nil
}
