// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: covid.proto

package covid

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CovidDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacientCode     string `protobuf:"bytes,1,opt,name=pacient_code,json=pacientCode,proto3" json:"pacient_code,omitempty"`
	PacientAge      string `protobuf:"bytes,2,opt,name=pacient_age,json=pacientAge,proto3" json:"pacient_age,omitempty"`
	PacientGender   string `protobuf:"bytes,3,opt,name=pacient_gender,json=pacientGender,proto3" json:"pacient_gender,omitempty"`
	PacientDistrict string `protobuf:"bytes,4,opt,name=pacient_district,json=pacientDistrict,proto3" json:"pacient_district,omitempty"`
	PacientCity     string `protobuf:"bytes,5,opt,name=pacient_city,json=pacientCity,proto3" json:"pacient_city,omitempty"`
	PacientState    string `protobuf:"bytes,6,opt,name=pacient_state,json=pacientState,proto3" json:"pacient_state,omitempty"`
	CityCode        string `protobuf:"bytes,7,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	Date            string `protobuf:"bytes,8,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *CovidDataResponse) Reset() {
	*x = CovidDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CovidDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CovidDataResponse) ProtoMessage() {}

func (x *CovidDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_covid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CovidDataResponse.ProtoReflect.Descriptor instead.
func (*CovidDataResponse) Descriptor() ([]byte, []int) {
	return file_covid_proto_rawDescGZIP(), []int{0}
}

func (x *CovidDataResponse) GetPacientCode() string {
	if x != nil {
		return x.PacientCode
	}
	return ""
}

func (x *CovidDataResponse) GetPacientAge() string {
	if x != nil {
		return x.PacientAge
	}
	return ""
}

func (x *CovidDataResponse) GetPacientGender() string {
	if x != nil {
		return x.PacientGender
	}
	return ""
}

func (x *CovidDataResponse) GetPacientDistrict() string {
	if x != nil {
		return x.PacientDistrict
	}
	return ""
}

func (x *CovidDataResponse) GetPacientCity() string {
	if x != nil {
		return x.PacientCity
	}
	return ""
}

func (x *CovidDataResponse) GetPacientState() string {
	if x != nil {
		return x.PacientState
	}
	return ""
}

func (x *CovidDataResponse) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

func (x *CovidDataResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type CovidDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CovidDataRequest) Reset() {
	*x = CovidDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CovidDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CovidDataRequest) ProtoMessage() {}

func (x *CovidDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_covid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CovidDataRequest.ProtoReflect.Descriptor instead.
func (*CovidDataRequest) Descriptor() ([]byte, []int) {
	return file_covid_proto_rawDescGZIP(), []int{1}
}

func (x *CovidDataRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_covid_proto protoreflect.FileDescriptor

var file_covid_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02,
	0x0a, 0x11, 0x43, 0x6f, 0x76, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x63,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x61, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x10, 0x70, 0x61, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x69, 0x65, 0x6e,
	0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x61, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x61, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x6f, 0x76, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x4c, 0x0a, 0x10, 0x43, 0x6f,
	0x76, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x11, 0x2e, 0x43, 0x6f, 0x76, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x43, 0x6f, 0x76, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x63, 0x6f,
	0x76, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_covid_proto_rawDescOnce sync.Once
	file_covid_proto_rawDescData = file_covid_proto_rawDesc
)

func file_covid_proto_rawDescGZIP() []byte {
	file_covid_proto_rawDescOnce.Do(func() {
		file_covid_proto_rawDescData = protoimpl.X.CompressGZIP(file_covid_proto_rawDescData)
	})
	return file_covid_proto_rawDescData
}

var file_covid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_covid_proto_goTypes = []interface{}{
	(*CovidDataResponse)(nil), // 0: CovidDataResponse
	(*CovidDataRequest)(nil),  // 1: CovidDataRequest
}
var file_covid_proto_depIdxs = []int32{
	1, // 0: CovidDataService.GetDataStream:input_type -> CovidDataRequest
	0, // 1: CovidDataService.GetDataStream:output_type -> CovidDataResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_covid_proto_init() }
func file_covid_proto_init() {
	if File_covid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_covid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CovidDataResponse); i {
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
		file_covid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CovidDataRequest); i {
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
			RawDescriptor: file_covid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_covid_proto_goTypes,
		DependencyIndexes: file_covid_proto_depIdxs,
		MessageInfos:      file_covid_proto_msgTypes,
	}.Build()
	File_covid_proto = out.File
	file_covid_proto_rawDesc = nil
	file_covid_proto_goTypes = nil
	file_covid_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CovidDataServiceClient is the client API for CovidDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CovidDataServiceClient interface {
	GetDataStream(ctx context.Context, in *CovidDataRequest, opts ...grpc.CallOption) (CovidDataService_GetDataStreamClient, error)
}

type covidDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCovidDataServiceClient(cc grpc.ClientConnInterface) CovidDataServiceClient {
	return &covidDataServiceClient{cc}
}

func (c *covidDataServiceClient) GetDataStream(ctx context.Context, in *CovidDataRequest, opts ...grpc.CallOption) (CovidDataService_GetDataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CovidDataService_serviceDesc.Streams[0], "/CovidDataService/GetDataStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &covidDataServiceGetDataStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CovidDataService_GetDataStreamClient interface {
	Recv() (*CovidDataResponse, error)
	grpc.ClientStream
}

type covidDataServiceGetDataStreamClient struct {
	grpc.ClientStream
}

func (x *covidDataServiceGetDataStreamClient) Recv() (*CovidDataResponse, error) {
	m := new(CovidDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CovidDataServiceServer is the server API for CovidDataService service.
type CovidDataServiceServer interface {
	GetDataStream(*CovidDataRequest, CovidDataService_GetDataStreamServer) error
}

// UnimplementedCovidDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCovidDataServiceServer struct {
}

func (*UnimplementedCovidDataServiceServer) GetDataStream(*CovidDataRequest, CovidDataService_GetDataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDataStream not implemented")
}

func RegisterCovidDataServiceServer(s *grpc.Server, srv CovidDataServiceServer) {
	s.RegisterService(&_CovidDataService_serviceDesc, srv)
}

func _CovidDataService_GetDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CovidDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CovidDataServiceServer).GetDataStream(m, &covidDataServiceGetDataStreamServer{stream})
}

type CovidDataService_GetDataStreamServer interface {
	Send(*CovidDataResponse) error
	grpc.ServerStream
}

type covidDataServiceGetDataStreamServer struct {
	grpc.ServerStream
}

func (x *covidDataServiceGetDataStreamServer) Send(m *CovidDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CovidDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CovidDataService",
	HandlerType: (*CovidDataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDataStream",
			Handler:       _CovidDataService_GetDataStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "covid.proto",
}
