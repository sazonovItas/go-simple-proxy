// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: request.proto

package pb_request

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	ProxyId       string                 `protobuf:"bytes,2,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	ProxyName     string                 `protobuf:"bytes,3,opt,name=proxy_name,json=proxyName,proto3" json:"proxy_name,omitempty"`
	ProxyUserId   string                 `protobuf:"bytes,4,opt,name=proxy_user_id,json=proxyUserId,proto3" json:"proxy_user_id,omitempty"`
	ProxyUserIp   string                 `protobuf:"bytes,5,opt,name=proxy_user_ip,json=proxyUserIp,proto3" json:"proxy_user_ip,omitempty"`
	ProxyUserName string                 `protobuf:"bytes,6,opt,name=proxy_user_name,json=proxyUserName,proto3" json:"proxy_user_name,omitempty"`
	Host          string                 `protobuf:"bytes,7,opt,name=host,proto3" json:"host,omitempty"`
	Upload        int64                  `protobuf:"varint,8,opt,name=upload,proto3" json:"upload,omitempty"`
	Download      int64                  `protobuf:"varint,9,opt,name=download,proto3" json:"download,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ProxyRequest) Reset() {
	*x = ProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRequest) ProtoMessage() {}

func (x *ProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRequest.ProtoReflect.Descriptor instead.
func (*ProxyRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *ProxyRequest) GetProxyId() string {
	if x != nil {
		return x.ProxyId
	}
	return ""
}

func (x *ProxyRequest) GetProxyName() string {
	if x != nil {
		return x.ProxyName
	}
	return ""
}

func (x *ProxyRequest) GetProxyUserId() string {
	if x != nil {
		return x.ProxyUserId
	}
	return ""
}

func (x *ProxyRequest) GetProxyUserIp() string {
	if x != nil {
		return x.ProxyUserIp
	}
	return ""
}

func (x *ProxyRequest) GetProxyUserName() string {
	if x != nil {
		return x.ProxyUserName
	}
	return ""
}

func (x *ProxyRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ProxyRequest) GetUpload() int64 {
	if x != nil {
		return x.Upload
	}
	return 0
}

func (x *ProxyRequest) GetDownload() int64 {
	if x != nil {
		return x.Download
	}
	return 0
}

func (x *ProxyRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type SaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{1}
}

func (x *SaveResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *ProxyRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{2}
}

func (x *SaveRequest) GetRequest() *ProxyRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

var File_request_proto protoreflect.FileDescriptor

var file_request_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x02, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x70, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0x58, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x53, 0x61,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_request_proto_rawDescOnce sync.Once
	file_request_proto_rawDescData = file_request_proto_rawDesc
)

func file_request_proto_rawDescGZIP() []byte {
	file_request_proto_rawDescOnce.Do(func() {
		file_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_proto_rawDescData)
	})
	return file_request_proto_rawDescData
}

var (
	file_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_request_proto_goTypes  = []interface{}{
		(*ProxyRequest)(nil),          // 0: request.ProxyRequest
		(*SaveResponse)(nil),          // 1: request.SaveResponse
		(*SaveRequest)(nil),           // 2: request.SaveRequest
		(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	}
)

var file_request_proto_depIdxs = []int32{
	3, // 0: request.ProxyRequest.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: request.SaveRequest.request:type_name -> request.ProxyRequest
	2, // 2: request.ProxyRequestService.SaveProxyRequest:input_type -> request.SaveRequest
	1, // 3: request.ProxyRequestService.SaveProxyRequest:output_type -> request.SaveResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_request_proto_init() }
func file_request_proto_init() {
	if File_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRequest); i {
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
		file_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResponse); i {
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
		file_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest); i {
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
	file_request_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_request_proto_goTypes,
		DependencyIndexes: file_request_proto_depIdxs,
		MessageInfos:      file_request_proto_msgTypes,
	}.Build()
	File_request_proto = out.File
	file_request_proto_rawDesc = nil
	file_request_proto_goTypes = nil
	file_request_proto_depIdxs = nil
}
