// idl/qwen.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.27.0
// source: qwen.proto

package qwen

import (
	_ "ai_helper/biz/model/api"
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

type QwenApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`                                                              // primary key, qwen task id
	UserId       uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"`                              // user id
	InputModel   string `protobuf:"bytes,3,opt,name=input_model,json=inputModel,proto3" json:"input_model,omitempty" form:"input_model" query:"input_model"`           // input model
	InputRole    string `protobuf:"bytes,4,opt,name=input_role,json=inputRole,proto3" json:"input_role,omitempty" form:"input_role" query:"input_role"`                // input role
	InputContent string `protobuf:"bytes,5,opt,name=input_content,json=inputContent,proto3" json:"input_content,omitempty" form:"input_content" query:"input_content"` // input content
	HistoryNum   int32  `protobuf:"varint,6,opt,name=history_num,json=historyNum,proto3" json:"history_num,omitempty" form:"history_num" query:"history_num"`          // input history num
}

func (x *QwenApiRequest) Reset() {
	*x = QwenApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QwenApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QwenApiRequest) ProtoMessage() {}

func (x *QwenApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qwen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QwenApiRequest.ProtoReflect.Descriptor instead.
func (*QwenApiRequest) Descriptor() ([]byte, []int) {
	return file_qwen_proto_rawDescGZIP(), []int{0}
}

func (x *QwenApiRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QwenApiRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *QwenApiRequest) GetInputModel() string {
	if x != nil {
		return x.InputModel
	}
	return ""
}

func (x *QwenApiRequest) GetInputRole() string {
	if x != nil {
		return x.InputRole
	}
	return ""
}

func (x *QwenApiRequest) GetInputContent() string {
	if x != nil {
		return x.InputContent
	}
	return ""
}

func (x *QwenApiRequest) GetHistoryNum() int32 {
	if x != nil {
		return x.HistoryNum
	}
	return 0
}

type QwenApiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" query:"status_code" form:"status_code" json:"status_code"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`
}

func (x *QwenApiResponse) Reset() {
	*x = QwenApiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QwenApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QwenApiResponse) ProtoMessage() {}

func (x *QwenApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qwen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QwenApiResponse.ProtoReflect.Descriptor instead.
func (*QwenApiResponse) Descriptor() ([]byte, []int) {
	return file_qwen_proto_rawDescGZIP(), []int{1}
}

func (x *QwenApiResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *QwenApiResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

var File_qwen_proto protoreflect.FileDescriptor

var file_qwen_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x71, 0x77, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x71, 0x77,
	0x65, 0x6e, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01,
	0x0a, 0x0e, 0x51, 0x77, 0x65, 0x6e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x22,
	0x69, 0x0a, 0x0f, 0x51, 0x77, 0x65, 0x6e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x32, 0xbd, 0x01, 0x0a, 0x0e, 0x51,
	0x77, 0x65, 0x6e, 0x41, 0x70, 0x69, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x55, 0x0a,
	0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x77, 0x65, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x14, 0x2e, 0x71, 0x77, 0x65, 0x6e, 0x2e, 0x51, 0x77, 0x65, 0x6e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x71, 0x77, 0x65, 0x6e, 0x2e, 0x51, 0x77, 0x65,
	0x6e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0xd2, 0xc1,
	0x18, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x77, 0x65, 0x6e, 0x2f, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x12, 0x54, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x51, 0x77, 0x65,
	0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x71, 0x77, 0x65, 0x6e, 0x2e, 0x51, 0x77, 0x65,
	0x6e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x71, 0x77,
	0x65, 0x6e, 0x2e, 0x51, 0x77, 0x65, 0x6e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0xca, 0xc1, 0x18, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71,
	0x77, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x21, 0x5a, 0x1f, 0x61, 0x69,
	0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x71, 0x77, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qwen_proto_rawDescOnce sync.Once
	file_qwen_proto_rawDescData = file_qwen_proto_rawDesc
)

func file_qwen_proto_rawDescGZIP() []byte {
	file_qwen_proto_rawDescOnce.Do(func() {
		file_qwen_proto_rawDescData = protoimpl.X.CompressGZIP(file_qwen_proto_rawDescData)
	})
	return file_qwen_proto_rawDescData
}

var file_qwen_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_qwen_proto_goTypes = []interface{}{
	(*QwenApiRequest)(nil),  // 0: qwen.QwenApiRequest
	(*QwenApiResponse)(nil), // 1: qwen.QwenApiResponse
}
var file_qwen_proto_depIdxs = []int32{
	0, // 0: qwen.QwenApiHandler.SubmitQwenTask:input_type -> qwen.QwenApiRequest
	0, // 1: qwen.QwenApiHandler.QueryQwenTask:input_type -> qwen.QwenApiRequest
	1, // 2: qwen.QwenApiHandler.SubmitQwenTask:output_type -> qwen.QwenApiResponse
	1, // 3: qwen.QwenApiHandler.QueryQwenTask:output_type -> qwen.QwenApiResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_qwen_proto_init() }
func file_qwen_proto_init() {
	if File_qwen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qwen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QwenApiRequest); i {
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
		file_qwen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QwenApiResponse); i {
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
			RawDescriptor: file_qwen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qwen_proto_goTypes,
		DependencyIndexes: file_qwen_proto_depIdxs,
		MessageInfos:      file_qwen_proto_msgTypes,
	}.Build()
	File_qwen_proto = out.File
	file_qwen_proto_rawDesc = nil
	file_qwen_proto_goTypes = nil
	file_qwen_proto_depIdxs = nil
}
