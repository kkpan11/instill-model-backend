// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: ray.proto

package rayserver

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

// ModelReadyRequest represents a request to check if a model is ready
type ModelReadyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// model id
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// model tag verion
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ModelReadyRequest) Reset() {
	*x = ModelReadyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ray_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelReadyRequest) ProtoMessage() {}

func (x *ModelReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ray_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelReadyRequest.ProtoReflect.Descriptor instead.
func (*ModelReadyRequest) Descriptor() ([]byte, []int) {
	return file_ray_proto_rawDescGZIP(), []int{0}
}

func (x *ModelReadyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelReadyRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// ModelReadyResponse represents a response to check if a model is ready
type ModelReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// whether the model is ready or not
	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *ModelReadyResponse) Reset() {
	*x = ModelReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ray_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelReadyResponse) ProtoMessage() {}

func (x *ModelReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ray_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelReadyResponse.ProtoReflect.Descriptor instead.
func (*ModelReadyResponse) Descriptor() ([]byte, []int) {
	return file_ray_proto_rawDescGZIP(), []int{1}
}

func (x *ModelReadyResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

// ModelMetadataRequest represents a request to get the model metadata
type ModelMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// model id
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// model tag verion
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ModelMetadataRequest) Reset() {
	*x = ModelMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ray_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelMetadataRequest) ProtoMessage() {}

func (x *ModelMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ray_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelMetadataRequest.ProtoReflect.Descriptor instead.
func (*ModelMetadataRequest) Descriptor() ([]byte, []int) {
	return file_ray_proto_rawDescGZIP(), []int{2}
}

func (x *ModelMetadataRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelMetadataRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// tensor for inference
type InferTensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tensor name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// tensor data type.
	Datatype string `protobuf:"bytes,2,opt,name=datatype,proto3" json:"datatype,omitempty"`
	// tensor shape.
	Shape []int64 `protobuf:"varint,3,rep,packed,name=shape,proto3" json:"shape,omitempty"`
}

func (x *InferTensor) Reset() {
	*x = InferTensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ray_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferTensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferTensor) ProtoMessage() {}

func (x *InferTensor) ProtoReflect() protoreflect.Message {
	mi := &file_ray_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferTensor.ProtoReflect.Descriptor instead.
func (*InferTensor) Descriptor() ([]byte, []int) {
	return file_ray_proto_rawDescGZIP(), []int{3}
}

func (x *InferTensor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InferTensor) GetDatatype() string {
	if x != nil {
		return x.Datatype
	}
	return ""
}

func (x *InferTensor) GetShape() []int64 {
	if x != nil {
		return x.Shape
	}
	return nil
}

// ModelMetadataResponse represents a response to get the model metadata
type ModelMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// model name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// model tag version
	Versions []string `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
	// model inference framework
	Framework string `protobuf:"bytes,3,opt,name=framework,proto3" json:"framework,omitempty"`
	// model inputs
	Inputs []*ModelMetadataResponse_TensorMetadata `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// model outputs
	Outputs []*ModelMetadataResponse_TensorMetadata `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *ModelMetadataResponse) Reset() {
	*x = ModelMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ray_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelMetadataResponse) ProtoMessage() {}

func (x *ModelMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ray_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelMetadataResponse.ProtoReflect.Descriptor instead.
func (*ModelMetadataResponse) Descriptor() ([]byte, []int) {
	return file_ray_proto_rawDescGZIP(), []int{4}
}

func (x *ModelMetadataResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelMetadataResponse) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *ModelMetadataResponse) GetFramework() string {
	if x != nil {
		return x.Framework
	}
	return ""
}

func (x *ModelMetadataResponse) GetInputs() []*ModelMetadataResponse_TensorMetadata {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *ModelMetadataResponse) GetOutputs() []*ModelMetadataResponse_TensorMetadata {
	if x != nil {
		return x.Outputs
	}
	return nil
}

// ModelInferRequest represents a request for model inference
type ModelInferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the model to use for inferencing.
	ModelName string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	// model tag version
	ModelVersion string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	// input tensors for the inference.
	Inputs []*InferTensor `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// The requested output tensors for the inference. Optional, if not
	// specified all outputs specified in the model config will be
	// returned.
	Outputs []*ModelInferRequest_InferRequestedOutputTensor `protobuf:"bytes,6,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// raw input contents
	RawInputContents [][]byte `protobuf:"bytes,7,rep,name=raw_input_contents,json=rawInputContents,proto3" json:"raw_input_contents,omitempty"`
}

func (x *ModelInferRequest) Reset() {
	*x = ModelInferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ray_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelInferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelInferRequest) ProtoMessage() {}

func (x *ModelInferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ray_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelInferRequest.ProtoReflect.Descriptor instead.
func (*ModelInferRequest) Descriptor() ([]byte, []int) {
	return file_ray_proto_rawDescGZIP(), []int{5}
}

func (x *ModelInferRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *ModelInferRequest) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

func (x *ModelInferRequest) GetInputs() []*InferTensor {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *ModelInferRequest) GetOutputs() []*ModelInferRequest_InferRequestedOutputTensor {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *ModelInferRequest) GetRawInputContents() [][]byte {
	if x != nil {
		return x.RawInputContents
	}
	return nil
}

// ModelInferResponse represents a response for model inference
type ModelInferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the model to use for inferencing.
	ModelName string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	// model tag version
	ModelVersion string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	// output tensors
	Outputs []*InferTensor `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// raw output contents
	RawOutputContents [][]byte `protobuf:"bytes,6,rep,name=raw_output_contents,json=rawOutputContents,proto3" json:"raw_output_contents,omitempty"`
}

func (x *ModelInferResponse) Reset() {
	*x = ModelInferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ray_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelInferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelInferResponse) ProtoMessage() {}

func (x *ModelInferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ray_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelInferResponse.ProtoReflect.Descriptor instead.
func (*ModelInferResponse) Descriptor() ([]byte, []int) {
	return file_ray_proto_rawDescGZIP(), []int{6}
}

func (x *ModelInferResponse) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *ModelInferResponse) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

func (x *ModelInferResponse) GetOutputs() []*InferTensor {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *ModelInferResponse) GetRawOutputContents() [][]byte {
	if x != nil {
		return x.RawOutputContents
	}
	return nil
}

// metadata for a tensor
type ModelMetadataResponse_TensorMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tensor name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// tensor data type
	Datatype string `protobuf:"bytes,2,opt,name=datatype,proto3" json:"datatype,omitempty"`
	// tensor shape
	Shape []int64 `protobuf:"varint,3,rep,packed,name=shape,proto3" json:"shape,omitempty"`
}

func (x *ModelMetadataResponse_TensorMetadata) Reset() {
	*x = ModelMetadataResponse_TensorMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ray_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelMetadataResponse_TensorMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelMetadataResponse_TensorMetadata) ProtoMessage() {}

func (x *ModelMetadataResponse_TensorMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_ray_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelMetadataResponse_TensorMetadata.ProtoReflect.Descriptor instead.
func (*ModelMetadataResponse_TensorMetadata) Descriptor() ([]byte, []int) {
	return file_ray_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ModelMetadataResponse_TensorMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelMetadataResponse_TensorMetadata) GetDatatype() string {
	if x != nil {
		return x.Datatype
	}
	return ""
}

func (x *ModelMetadataResponse_TensorMetadata) GetShape() []int64 {
	if x != nil {
		return x.Shape
	}
	return nil
}

// An output tensor requested for an inference request.
type ModelInferRequest_InferRequestedOutputTensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tensor name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ModelInferRequest_InferRequestedOutputTensor) Reset() {
	*x = ModelInferRequest_InferRequestedOutputTensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ray_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelInferRequest_InferRequestedOutputTensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelInferRequest_InferRequestedOutputTensor) ProtoMessage() {}

func (x *ModelInferRequest_InferRequestedOutputTensor) ProtoReflect() protoreflect.Message {
	mi := &file_ray_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelInferRequest_InferRequestedOutputTensor.ProtoReflect.Descriptor instead.
func (*ModelInferRequest_InferRequestedOutputTensor) Descriptor() ([]byte, []int) {
	return file_ray_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ModelInferRequest_InferRequestedOutputTensor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_ray_proto protoreflect.FileDescriptor

var file_ray_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x61, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x12, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x22, 0x44, 0x0a, 0x14, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x0b, 0x49,
	0x6e, 0x66, 0x65, 0x72, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65,
	0x22, 0xd1, 0x02, 0x0a, 0x15, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x47, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x73, 0x12, 0x49, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x1a, 0x56, 0x0a, 0x0e,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x68, 0x61, 0x70, 0x65, 0x22, 0xba, 0x02, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e,
	0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x72, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x51,
	0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x72, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x61, 0x77, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x10, 0x72,
	0x61, 0x77, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0x30, 0x0a, 0x1a, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xba, 0x01, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x07,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x72, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x2e,
	0x0a, 0x13, 0x72, 0x61, 0x77, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x11, 0x72, 0x61, 0x77,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xfc,
	0x01, 0x0a, 0x0a, 0x52, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x1c, 0x2e, 0x72, 0x61,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x61, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x72, 0x61,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72,
	0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x72, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x49, 0x6e, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72,
	0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x2f, 0x72, 0x61, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ray_proto_rawDescOnce sync.Once
	file_ray_proto_rawDescData = file_ray_proto_rawDesc
)

func file_ray_proto_rawDescGZIP() []byte {
	file_ray_proto_rawDescOnce.Do(func() {
		file_ray_proto_rawDescData = protoimpl.X.CompressGZIP(file_ray_proto_rawDescData)
	})
	return file_ray_proto_rawDescData
}

var file_ray_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ray_proto_goTypes = []interface{}{
	(*ModelReadyRequest)(nil),                            // 0: ray.serve.ModelReadyRequest
	(*ModelReadyResponse)(nil),                           // 1: ray.serve.ModelReadyResponse
	(*ModelMetadataRequest)(nil),                         // 2: ray.serve.ModelMetadataRequest
	(*InferTensor)(nil),                                  // 3: ray.serve.InferTensor
	(*ModelMetadataResponse)(nil),                        // 4: ray.serve.ModelMetadataResponse
	(*ModelInferRequest)(nil),                            // 5: ray.serve.ModelInferRequest
	(*ModelInferResponse)(nil),                           // 6: ray.serve.ModelInferResponse
	(*ModelMetadataResponse_TensorMetadata)(nil),         // 7: ray.serve.ModelMetadataResponse.TensorMetadata
	(*ModelInferRequest_InferRequestedOutputTensor)(nil), // 8: ray.serve.ModelInferRequest.InferRequestedOutputTensor
}
var file_ray_proto_depIdxs = []int32{
	7, // 0: ray.serve.ModelMetadataResponse.inputs:type_name -> ray.serve.ModelMetadataResponse.TensorMetadata
	7, // 1: ray.serve.ModelMetadataResponse.outputs:type_name -> ray.serve.ModelMetadataResponse.TensorMetadata
	3, // 2: ray.serve.ModelInferRequest.inputs:type_name -> ray.serve.InferTensor
	8, // 3: ray.serve.ModelInferRequest.outputs:type_name -> ray.serve.ModelInferRequest.InferRequestedOutputTensor
	3, // 4: ray.serve.ModelInferResponse.outputs:type_name -> ray.serve.InferTensor
	0, // 5: ray.serve.RayService.ModelReady:input_type -> ray.serve.ModelReadyRequest
	2, // 6: ray.serve.RayService.ModelMetadata:input_type -> ray.serve.ModelMetadataRequest
	5, // 7: ray.serve.RayService.ModelInfer:input_type -> ray.serve.ModelInferRequest
	1, // 8: ray.serve.RayService.ModelReady:output_type -> ray.serve.ModelReadyResponse
	4, // 9: ray.serve.RayService.ModelMetadata:output_type -> ray.serve.ModelMetadataResponse
	6, // 10: ray.serve.RayService.ModelInfer:output_type -> ray.serve.ModelInferResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ray_proto_init() }
func file_ray_proto_init() {
	if File_ray_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ray_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelReadyRequest); i {
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
		file_ray_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelReadyResponse); i {
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
		file_ray_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelMetadataRequest); i {
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
		file_ray_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferTensor); i {
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
		file_ray_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelMetadataResponse); i {
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
		file_ray_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelInferRequest); i {
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
		file_ray_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelInferResponse); i {
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
		file_ray_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelMetadataResponse_TensorMetadata); i {
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
		file_ray_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelInferRequest_InferRequestedOutputTensor); i {
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
			RawDescriptor: file_ray_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ray_proto_goTypes,
		DependencyIndexes: file_ray_proto_depIdxs,
		MessageInfos:      file_ray_proto_msgTypes,
	}.Build()
	File_ray_proto = out.File
	file_ray_proto_rawDesc = nil
	file_ray_proto_goTypes = nil
	file_ray_proto_depIdxs = nil
}
