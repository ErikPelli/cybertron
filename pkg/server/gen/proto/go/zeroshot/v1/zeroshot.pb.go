// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: zeroshot/v1/zeroshot.proto

package zeroshotv1

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

type ClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input      string              `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Parameters *ZeroShotParameters `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *ClassifyRequest) Reset() {
	*x = ClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeroshot_v1_zeroshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifyRequest) ProtoMessage() {}

func (x *ClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zeroshot_v1_zeroshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifyRequest.ProtoReflect.Descriptor instead.
func (*ClassifyRequest) Descriptor() ([]byte, []int) {
	return file_zeroshot_v1_zeroshot_proto_rawDescGZIP(), []int{0}
}

func (x *ClassifyRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *ClassifyRequest) GetParameters() *ZeroShotParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type ZeroShotParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HypothesisTemplate string   `protobuf:"bytes,1,opt,name=hypothesis_template,json=hypothesisTemplate,proto3" json:"hypothesis_template,omitempty"`
	CandidateLabels    []string `protobuf:"bytes,2,rep,name=candidate_labels,json=candidateLabels,proto3" json:"candidate_labels,omitempty"`
	MultiLabel         bool     `protobuf:"varint,3,opt,name=multi_label,json=multiLabel,proto3" json:"multi_label,omitempty"`
}

func (x *ZeroShotParameters) Reset() {
	*x = ZeroShotParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeroshot_v1_zeroshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZeroShotParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZeroShotParameters) ProtoMessage() {}

func (x *ZeroShotParameters) ProtoReflect() protoreflect.Message {
	mi := &file_zeroshot_v1_zeroshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZeroShotParameters.ProtoReflect.Descriptor instead.
func (*ZeroShotParameters) Descriptor() ([]byte, []int) {
	return file_zeroshot_v1_zeroshot_proto_rawDescGZIP(), []int{1}
}

func (x *ZeroShotParameters) GetHypothesisTemplate() string {
	if x != nil {
		return x.HypothesisTemplate
	}
	return ""
}

func (x *ZeroShotParameters) GetCandidateLabels() []string {
	if x != nil {
		return x.CandidateLabels
	}
	return nil
}

func (x *ZeroShotParameters) GetMultiLabel() bool {
	if x != nil {
		return x.MultiLabel
	}
	return false
}

type ClassifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: string sequence = ...; ?
	Labels []string  `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Scores []float64 `protobuf:"fixed64,2,rep,packed,name=scores,proto3" json:"scores,omitempty"`
}

func (x *ClassifyResponse) Reset() {
	*x = ClassifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeroshot_v1_zeroshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifyResponse) ProtoMessage() {}

func (x *ClassifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zeroshot_v1_zeroshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifyResponse.ProtoReflect.Descriptor instead.
func (*ClassifyResponse) Descriptor() ([]byte, []int) {
	return file_zeroshot_v1_zeroshot_proto_rawDescGZIP(), []int{2}
}

func (x *ClassifyResponse) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ClassifyResponse) GetScores() []float64 {
	if x != nil {
		return x.Scores
	}
	return nil
}

var File_zeroshot_v1_zeroshot_proto protoreflect.FileDescriptor

var file_zeroshot_v1_zeroshot_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x7a, 0x65, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x65,
	0x72, 0x6f, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x7a, 0x65,
	0x72, 0x6f, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x3f, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x5a, 0x65, 0x72, 0x6f, 0x53, 0x68, 0x6f, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x22, 0x91, 0x01, 0x0a, 0x12, 0x5a, 0x65, 0x72, 0x6f, 0x53, 0x68, 0x6f, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x68, 0x79, 0x70, 0x6f,
	0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x68, 0x79, 0x70, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x69,
	0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x42, 0x0a, 0x10, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x32, 0x73, 0x0a, 0x0f, 0x5a, 0x65, 0x72,
	0x6f, 0x53, 0x68, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x08,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x1c, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x73,
	0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x73, 0x68, 0x6f,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a,
	0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x42, 0x48,
	0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6c, 0x70,
	0x6f, 0x64, 0x79, 0x73, 0x73, 0x65, 0x79, 0x2f, 0x63, 0x79, 0x62, 0x65, 0x72, 0x74, 0x72, 0x6f,
	0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x7a, 0x65, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x7a, 0x65,
	0x72, 0x6f, 0x73, 0x68, 0x6f, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zeroshot_v1_zeroshot_proto_rawDescOnce sync.Once
	file_zeroshot_v1_zeroshot_proto_rawDescData = file_zeroshot_v1_zeroshot_proto_rawDesc
)

func file_zeroshot_v1_zeroshot_proto_rawDescGZIP() []byte {
	file_zeroshot_v1_zeroshot_proto_rawDescOnce.Do(func() {
		file_zeroshot_v1_zeroshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_zeroshot_v1_zeroshot_proto_rawDescData)
	})
	return file_zeroshot_v1_zeroshot_proto_rawDescData
}

var file_zeroshot_v1_zeroshot_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_zeroshot_v1_zeroshot_proto_goTypes = []interface{}{
	(*ClassifyRequest)(nil),    // 0: zeroshot.v1.ClassifyRequest
	(*ZeroShotParameters)(nil), // 1: zeroshot.v1.ZeroShotParameters
	(*ClassifyResponse)(nil),   // 2: zeroshot.v1.ClassifyResponse
}
var file_zeroshot_v1_zeroshot_proto_depIdxs = []int32{
	1, // 0: zeroshot.v1.ClassifyRequest.parameters:type_name -> zeroshot.v1.ZeroShotParameters
	0, // 1: zeroshot.v1.ZeroShotService.Classify:input_type -> zeroshot.v1.ClassifyRequest
	2, // 2: zeroshot.v1.ZeroShotService.Classify:output_type -> zeroshot.v1.ClassifyResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_zeroshot_v1_zeroshot_proto_init() }
func file_zeroshot_v1_zeroshot_proto_init() {
	if File_zeroshot_v1_zeroshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zeroshot_v1_zeroshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassifyRequest); i {
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
		file_zeroshot_v1_zeroshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZeroShotParameters); i {
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
		file_zeroshot_v1_zeroshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassifyResponse); i {
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
			RawDescriptor: file_zeroshot_v1_zeroshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zeroshot_v1_zeroshot_proto_goTypes,
		DependencyIndexes: file_zeroshot_v1_zeroshot_proto_depIdxs,
		MessageInfos:      file_zeroshot_v1_zeroshot_proto_msgTypes,
	}.Build()
	File_zeroshot_v1_zeroshot_proto = out.File
	file_zeroshot_v1_zeroshot_proto_rawDesc = nil
	file_zeroshot_v1_zeroshot_proto_goTypes = nil
	file_zeroshot_v1_zeroshot_proto_depIdxs = nil
}
