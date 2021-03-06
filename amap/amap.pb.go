// Code generated by protoc-gen-go. DO NOT EDIT.
// source: amap.proto

package amap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Address struct {
	Street               string   `protobuf:"bytes,1,opt,name=Street,proto3" json:"Street,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_amap_9327eab1ff11e142, []int{0}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type Location struct {
	Count                string     `protobuf:"bytes,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Geocodes             []*Geocode `protobuf:"bytes,2,rep,name=Geocodes,proto3" json:"Geocodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_amap_9327eab1ff11e142, []int{1}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetCount() string {
	if m != nil {
		return m.Count
	}
	return ""
}

func (m *Location) GetGeocodes() []*Geocode {
	if m != nil {
		return m.Geocodes
	}
	return nil
}

type Geocode struct {
	Location             string   `protobuf:"bytes,1,opt,name=Location,proto3" json:"Location,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Country              string   `protobuf:"bytes,3,opt,name=Country,proto3" json:"Country,omitempty"`
	Province             string   `protobuf:"bytes,4,opt,name=Province,proto3" json:"Province,omitempty"`
	City                 string   `protobuf:"bytes,5,opt,name=City,proto3" json:"City,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Geocode) Reset()         { *m = Geocode{} }
func (m *Geocode) String() string { return proto.CompactTextString(m) }
func (*Geocode) ProtoMessage()    {}
func (*Geocode) Descriptor() ([]byte, []int) {
	return fileDescriptor_amap_9327eab1ff11e142, []int{2}
}
func (m *Geocode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Geocode.Unmarshal(m, b)
}
func (m *Geocode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Geocode.Marshal(b, m, deterministic)
}
func (dst *Geocode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Geocode.Merge(dst, src)
}
func (m *Geocode) XXX_Size() int {
	return xxx_messageInfo_Geocode.Size(m)
}
func (m *Geocode) XXX_DiscardUnknown() {
	xxx_messageInfo_Geocode.DiscardUnknown(m)
}

var xxx_messageInfo_Geocode proto.InternalMessageInfo

func (m *Geocode) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Geocode) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Geocode) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Geocode) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Geocode) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func init() {
	proto.RegisterType((*Address)(nil), "amap.Address")
	proto.RegisterType((*Location)(nil), "amap.Location")
	proto.RegisterType((*Geocode)(nil), "amap.Geocode")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AmapClient is the client API for Amap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AmapClient interface {
	// 正向地理编码： 将地址描述信息转换成地理坐标（经纬度)
	ForwardGeocode(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Location, error)
}

type amapClient struct {
	cc *grpc.ClientConn
}

func NewAmapClient(cc *grpc.ClientConn) AmapClient {
	return &amapClient{cc}
}

func (c *amapClient) ForwardGeocode(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/amap.Amap/ForwardGeocode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AmapServer is the server API for Amap service.
type AmapServer interface {
	// 正向地理编码： 将地址描述信息转换成地理坐标（经纬度)
	ForwardGeocode(context.Context, *Address) (*Location, error)
}

func RegisterAmapServer(s *grpc.Server, srv AmapServer) {
	s.RegisterService(&_Amap_serviceDesc, srv)
}

func _Amap_ForwardGeocode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmapServer).ForwardGeocode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amap.Amap/ForwardGeocode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmapServer).ForwardGeocode(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

var _Amap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "amap.Amap",
	HandlerType: (*AmapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForwardGeocode",
			Handler:    _Amap_ForwardGeocode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "amap.proto",
}

func init() { proto.RegisterFile("amap.proto", fileDescriptor_amap_9327eab1ff11e142) }

var fileDescriptor_amap_9327eab1ff11e142 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x49, 0xeb, 0x36, 0xe5, 0x10, 0x1d, 0x4e, 0x08, 0x59, 0x9d, 0x2a, 0x4f, 0x65, 0xa9,
	0x44, 0x11, 0x03, 0x63, 0x55, 0x09, 0x06, 0x18, 0x50, 0x79, 0x02, 0x13, 0x7b, 0xc8, 0x90, 0x5c,
	0xe4, 0x18, 0x50, 0x1e, 0x81, 0xb7, 0xae, 0x62, 0x9f, 0x9d, 0xcd, 0xdf, 0x9d, 0xf2, 0xff, 0xdf,
	0x05, 0x40, 0x37, 0xba, 0xdb, 0x77, 0x8e, 0x3c, 0xa1, 0x18, 0xdf, 0xea, 0x19, 0xca, 0xa3, 0x31,
	0xce, 0xf6, 0x3d, 0xde, 0xc3, 0xf2, 0xcb, 0x3b, 0x6b, 0xbd, 0x2c, 0xb6, 0xc5, 0xee, 0xfa, 0xcc,
	0x84, 0x08, 0xe2, 0x54, 0xfb, 0x41, 0xce, 0xc2, 0x34, 0xbc, 0xd5, 0x3b, 0xac, 0x3e, 0xa8, 0xd2,
	0xbe, 0xa6, 0x16, 0xef, 0x60, 0x71, 0xa2, 0x9f, 0x36, 0x7d, 0x16, 0x01, 0x1f, 0x60, 0xf5, 0x66,
	0xa9, 0x22, 0x63, 0x7b, 0x39, 0xdb, 0xce, 0x77, 0x37, 0x87, 0xdb, 0x7d, 0x68, 0xe7, 0xe9, 0x39,
	0xaf, 0xd5, 0x7f, 0x01, 0x25, 0x03, 0x6e, 0xa6, 0x60, 0xce, 0x9b, 0x8a, 0x64, 0x76, 0x65, 0x97,
	0xac, 0x2e, 0xa1, 0x0c, 0xad, 0x6e, 0x90, 0xf3, 0xb8, 0x61, 0x1c, 0xf3, 0x3e, 0x1d, 0xfd, 0xd6,
	0x6d, 0x65, 0xa5, 0x88, 0x79, 0x89, 0xf3, 0x61, 0x8b, 0xe9, 0xb0, 0xc3, 0x0b, 0x88, 0x63, 0xa3,
	0x3b, 0x7c, 0x84, 0xf5, 0x2b, 0xb9, 0x3f, 0xed, 0x4c, 0x32, 0x63, 0x7d, 0xae, 0xdc, 0xac, 0x23,
	0x26, 0x39, 0x75, 0xf5, 0xbd, 0x0c, 0xff, 0xf5, 0xe9, 0x12, 0x00, 0x00, 0xff, 0xff, 0x88, 0x31,
	0x96, 0x63, 0x65, 0x01, 0x00, 0x00,
}
