// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/report.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReportRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	mi := &file_proto_report_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{0}
}

func (x *ReportRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ReportResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReportId      string                 `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	mi := &file_proto_report_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportResponse) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

func (x *ReportResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type HealthRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	mi := &file_proto_report_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{2}
}

type HealthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	mi := &file_proto_report_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{3}
}

func (x *HealthResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_report_proto protoreflect.FileDescriptor

const file_proto_report_proto_rawDesc = "" +
	"\n" +
	"\x12proto/report.proto\x12\x06report\"(\n" +
	"\rReportRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"C\n" +
	"\x0eReportResponse\x12\x1b\n" +
	"\treport_id\x18\x01 \x01(\tR\breportId\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\"\x0f\n" +
	"\rHealthRequest\"(\n" +
	"\x0eHealthResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2\x8e\x01\n" +
	"\rReportService\x12?\n" +
	"\x0eGenerateReport\x12\x15.report.ReportRequest\x1a\x16.report.ReportResponse\x12<\n" +
	"\vHealthCheck\x12\x15.report.HealthRequest\x1a\x16.report.HealthResponseB\tZ\a./protob\x06proto3"

var (
	file_proto_report_proto_rawDescOnce sync.Once
	file_proto_report_proto_rawDescData []byte
)

func file_proto_report_proto_rawDescGZIP() []byte {
	file_proto_report_proto_rawDescOnce.Do(func() {
		file_proto_report_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_report_proto_rawDesc), len(file_proto_report_proto_rawDesc)))
	})
	return file_proto_report_proto_rawDescData
}

var file_proto_report_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_report_proto_goTypes = []any{
	(*ReportRequest)(nil),  // 0: report.ReportRequest
	(*ReportResponse)(nil), // 1: report.ReportResponse
	(*HealthRequest)(nil),  // 2: report.HealthRequest
	(*HealthResponse)(nil), // 3: report.HealthResponse
}
var file_proto_report_proto_depIdxs = []int32{
	0, // 0: report.ReportService.GenerateReport:input_type -> report.ReportRequest
	2, // 1: report.ReportService.HealthCheck:input_type -> report.HealthRequest
	1, // 2: report.ReportService.GenerateReport:output_type -> report.ReportResponse
	3, // 3: report.ReportService.HealthCheck:output_type -> report.HealthResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_report_proto_init() }
func file_proto_report_proto_init() {
	if File_proto_report_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_report_proto_rawDesc), len(file_proto_report_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_report_proto_goTypes,
		DependencyIndexes: file_proto_report_proto_depIdxs,
		MessageInfos:      file_proto_report_proto_msgTypes,
	}.Build()
	File_proto_report_proto = out.File
	file_proto_report_proto_goTypes = nil
	file_proto_report_proto_depIdxs = nil
}
