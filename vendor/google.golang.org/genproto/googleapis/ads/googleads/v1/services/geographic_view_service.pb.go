// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/geographic_view_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [GeographicViewService.GetGeographicView][google.ads.googleads.v1.services.GeographicViewService.GetGeographicView].
type GetGeographicViewRequest struct {
	// The resource name of the geographic view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGeographicViewRequest) Reset()         { *m = GetGeographicViewRequest{} }
func (m *GetGeographicViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetGeographicViewRequest) ProtoMessage()    {}
func (*GetGeographicViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_geographic_view_service_3209ea15797ce738, []int{0}
}
func (m *GetGeographicViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGeographicViewRequest.Unmarshal(m, b)
}
func (m *GetGeographicViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGeographicViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetGeographicViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGeographicViewRequest.Merge(dst, src)
}
func (m *GetGeographicViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetGeographicViewRequest.Size(m)
}
func (m *GetGeographicViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGeographicViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGeographicViewRequest proto.InternalMessageInfo

func (m *GetGeographicViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGeographicViewRequest)(nil), "google.ads.googleads.v1.services.GetGeographicViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeographicViewServiceClient is the client API for GeographicViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeographicViewServiceClient interface {
	// Returns the requested geographic view in full detail.
	GetGeographicView(ctx context.Context, in *GetGeographicViewRequest, opts ...grpc.CallOption) (*resources.GeographicView, error)
}

type geographicViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewGeographicViewServiceClient(cc *grpc.ClientConn) GeographicViewServiceClient {
	return &geographicViewServiceClient{cc}
}

func (c *geographicViewServiceClient) GetGeographicView(ctx context.Context, in *GetGeographicViewRequest, opts ...grpc.CallOption) (*resources.GeographicView, error) {
	out := new(resources.GeographicView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.GeographicViewService/GetGeographicView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeographicViewServiceServer is the server API for GeographicViewService service.
type GeographicViewServiceServer interface {
	// Returns the requested geographic view in full detail.
	GetGeographicView(context.Context, *GetGeographicViewRequest) (*resources.GeographicView, error)
}

func RegisterGeographicViewServiceServer(s *grpc.Server, srv GeographicViewServiceServer) {
	s.RegisterService(&_GeographicViewService_serviceDesc, srv)
}

func _GeographicViewService_GetGeographicView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGeographicViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographicViewServiceServer).GetGeographicView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.GeographicViewService/GetGeographicView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographicViewServiceServer).GetGeographicView(ctx, req.(*GetGeographicViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeographicViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.GeographicViewService",
	HandlerType: (*GeographicViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGeographicView",
			Handler:    _GeographicViewService_GetGeographicView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/geographic_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/geographic_view_service.proto", fileDescriptor_geographic_view_service_3209ea15797ce738)
}

var fileDescriptor_geographic_view_service_3209ea15797ce738 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0x4a, 0xc3, 0x40,
	0x1c, 0xc7, 0x49, 0x04, 0xc1, 0xa0, 0x83, 0x01, 0xa1, 0x04, 0x87, 0x52, 0x3b, 0x48, 0x87, 0x3b,
	0x62, 0x07, 0xf1, 0x44, 0x25, 0x5d, 0xe2, 0x24, 0xa5, 0x42, 0x06, 0x09, 0x94, 0x33, 0x39, 0xce,
	0x40, 0x93, 0x8b, 0xf7, 0x4b, 0xd3, 0x41, 0x1c, 0xf4, 0x15, 0x7c, 0x03, 0x47, 0xdf, 0xc1, 0x17,
	0x70, 0x75, 0xf0, 0x05, 0x9c, 0x7c, 0x0a, 0x49, 0xaf, 0x17, 0x2c, 0x36, 0x74, 0xfb, 0x72, 0xbf,
	0xef, 0xe7, 0xf7, 0xe7, 0x9b, 0x58, 0xe7, 0x5c, 0x08, 0x3e, 0x61, 0x98, 0xc6, 0x80, 0x95, 0xac,
	0x54, 0xe9, 0x62, 0x60, 0xb2, 0x4c, 0x22, 0x06, 0x98, 0x33, 0xc1, 0x25, 0xcd, 0xef, 0x92, 0x68,
	0x5c, 0x26, 0x6c, 0x36, 0x5e, 0x14, 0x50, 0x2e, 0x45, 0x21, 0xec, 0xb6, 0x82, 0x10, 0x8d, 0x01,
	0xd5, 0x3c, 0x2a, 0x5d, 0xa4, 0x79, 0xe7, 0xb8, 0x69, 0x82, 0x64, 0x20, 0xa6, 0x72, 0xc5, 0x08,
	0xd5, 0xda, 0xd9, 0xd7, 0x60, 0x9e, 0x60, 0x9a, 0x65, 0xa2, 0xa0, 0x45, 0x22, 0x32, 0x50, 0xd5,
	0xce, 0x85, 0xd5, 0xf2, 0x59, 0xe1, 0xd7, 0x64, 0x90, 0xb0, 0xd9, 0x88, 0xdd, 0x4f, 0x19, 0x14,
	0xf6, 0x81, 0xb5, 0xa3, 0x9b, 0x8f, 0x33, 0x9a, 0xb2, 0x96, 0xd1, 0x36, 0x0e, 0xb7, 0x46, 0xdb,
	0xfa, 0xf1, 0x8a, 0xa6, 0xec, 0xe8, 0xcb, 0xb0, 0xf6, 0x96, 0xf1, 0x6b, 0xb5, 0xb2, 0xfd, 0x6e,
	0x58, 0xbb, 0xff, 0x7a, 0xdb, 0x04, 0xad, 0x3b, 0x15, 0x35, 0x2d, 0xe4, 0xb8, 0x8d, 0x6c, 0x1d,
	0x02, 0x5a, 0x26, 0x3b, 0x27, 0xcf, 0x9f, 0xdf, 0x2f, 0x66, 0xdf, 0x76, 0xab, 0xa8, 0x1e, 0x96,
	0xce, 0x39, 0x8b, 0xa6, 0x50, 0x88, 0x94, 0x49, 0xc0, 0xbd, 0x3f, 0xd9, 0x55, 0x18, 0xe0, 0xde,
	0xe3, 0xe0, 0xc9, 0xb4, 0xba, 0x91, 0x48, 0xd7, 0xee, 0x3b, 0x70, 0x56, 0xde, 0x3f, 0xac, 0xf2,
	0x1d, 0x1a, 0x37, 0x97, 0x0b, 0x9e, 0x8b, 0x09, 0xcd, 0x38, 0x12, 0x92, 0x63, 0xce, 0xb2, 0x79,
	0xfa, 0xfa, 0x43, 0xe6, 0x09, 0x34, 0xff, 0x39, 0xa7, 0x5a, 0xbc, 0x9a, 0x1b, 0xbe, 0xe7, 0xbd,
	0x99, 0x6d, 0x5f, 0x35, 0xf4, 0x62, 0x40, 0x4a, 0x56, 0x2a, 0x70, 0xd1, 0x62, 0x30, 0x7c, 0x68,
	0x4b, 0xe8, 0xc5, 0x10, 0xd6, 0x96, 0x30, 0x70, 0x43, 0x6d, 0xf9, 0x31, 0xbb, 0xea, 0x9d, 0x10,
	0x2f, 0x06, 0x42, 0x6a, 0x13, 0x21, 0x81, 0x4b, 0x88, 0xb6, 0xdd, 0x6e, 0xce, 0xf7, 0xec, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x20, 0xad, 0x3a, 0xfd, 0xe0, 0x02, 0x00, 0x00,
}
