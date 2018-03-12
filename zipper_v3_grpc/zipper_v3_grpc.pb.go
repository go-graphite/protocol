// Code generated by protoc-gen-gogo.
// source: zipper_v3_grpc.proto
// DO NOT EDIT!

/*
	Package zipper_v3_grpc is a generated protocol buffer package.

	It is generated from these files:
		zipper_v3_grpc.proto

	It has these top-level messages:
		ProtocolVersionResponse
*/
package zipper_v3_grpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import carbonzipperpbng "zipper_v3_pb"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import strings "strings"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Version
type ProtocolVersionResponse struct {
	Version int64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ProtocolVersionResponse) Reset()      { *m = ProtocolVersionResponse{} }
func (*ProtocolVersionResponse) ProtoMessage() {}
func (*ProtocolVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorZipperV3Grpc, []int{0}
}

func (m *ProtocolVersionResponse) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtocolVersionResponse)(nil), "zipper_v3_grpc.ProtocolVersionResponse")
}
func (this *ProtocolVersionResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ProtocolVersionResponse)
	if !ok {
		that2, ok := that.(ProtocolVersionResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (this *ProtocolVersionResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&zipper_v3_grpc.ProtocolVersionResponse{")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringZipperV3Grpc(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CarbonV1 service

type CarbonV1Client interface {
	GetVersion(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ProtocolVersionResponse, error)
	FetchMetrics(ctx context.Context, in *carbonzipperpbng.MultiFetchRequest, opts ...grpc.CallOption) (*carbonzipperpbng.MultiFetchResponse, error)
	FindMetrics(ctx context.Context, in *carbonzipperpbng.MultiGlobRequest, opts ...grpc.CallOption) (*carbonzipperpbng.MultiGlobResponse, error)
	MetricsInfo(ctx context.Context, in *carbonzipperpbng.MultiMetricsInfoRequest, opts ...grpc.CallOption) (*carbonzipperpbng.MultiMetricsInfoResponse, error)
	ListMetrics(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*carbonzipperpbng.ListMetricsResponse, error)
	Stats(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*carbonzipperpbng.MetricDetailsResponse, error)
}

type carbonV1Client struct {
	cc *grpc.ClientConn
}

func NewCarbonV1Client(cc *grpc.ClientConn) CarbonV1Client {
	return &carbonV1Client{cc}
}

func (c *carbonV1Client) GetVersion(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ProtocolVersionResponse, error) {
	out := new(ProtocolVersionResponse)
	err := grpc.Invoke(ctx, "/zipper_v3_grpc.CarbonV1/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carbonV1Client) FetchMetrics(ctx context.Context, in *carbonzipperpbng.MultiFetchRequest, opts ...grpc.CallOption) (*carbonzipperpbng.MultiFetchResponse, error) {
	out := new(carbonzipperpbng.MultiFetchResponse)
	err := grpc.Invoke(ctx, "/zipper_v3_grpc.CarbonV1/FetchMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carbonV1Client) FindMetrics(ctx context.Context, in *carbonzipperpbng.MultiGlobRequest, opts ...grpc.CallOption) (*carbonzipperpbng.MultiGlobResponse, error) {
	out := new(carbonzipperpbng.MultiGlobResponse)
	err := grpc.Invoke(ctx, "/zipper_v3_grpc.CarbonV1/FindMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carbonV1Client) MetricsInfo(ctx context.Context, in *carbonzipperpbng.MultiMetricsInfoRequest, opts ...grpc.CallOption) (*carbonzipperpbng.MultiMetricsInfoResponse, error) {
	out := new(carbonzipperpbng.MultiMetricsInfoResponse)
	err := grpc.Invoke(ctx, "/zipper_v3_grpc.CarbonV1/MetricsInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carbonV1Client) ListMetrics(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*carbonzipperpbng.ListMetricsResponse, error) {
	out := new(carbonzipperpbng.ListMetricsResponse)
	err := grpc.Invoke(ctx, "/zipper_v3_grpc.CarbonV1/ListMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carbonV1Client) Stats(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*carbonzipperpbng.MetricDetailsResponse, error) {
	out := new(carbonzipperpbng.MetricDetailsResponse)
	err := grpc.Invoke(ctx, "/zipper_v3_grpc.CarbonV1/Stats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CarbonV1 service

type CarbonV1Server interface {
	GetVersion(context.Context, *google_protobuf1.Empty) (*ProtocolVersionResponse, error)
	FetchMetrics(context.Context, *carbonzipperpbng.MultiFetchRequest) (*carbonzipperpbng.MultiFetchResponse, error)
	FindMetrics(context.Context, *carbonzipperpbng.MultiGlobRequest) (*carbonzipperpbng.MultiGlobResponse, error)
	MetricsInfo(context.Context, *carbonzipperpbng.MultiMetricsInfoRequest) (*carbonzipperpbng.MultiMetricsInfoResponse, error)
	ListMetrics(context.Context, *google_protobuf1.Empty) (*carbonzipperpbng.ListMetricsResponse, error)
	Stats(context.Context, *google_protobuf1.Empty) (*carbonzipperpbng.MetricDetailsResponse, error)
}

func RegisterCarbonV1Server(s *grpc.Server, srv CarbonV1Server) {
	s.RegisterService(&_CarbonV1_serviceDesc, srv)
}

func _CarbonV1_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV1Server).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zipper_v3_grpc.CarbonV1/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV1Server).GetVersion(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarbonV1_FetchMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(carbonzipperpbng.MultiFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV1Server).FetchMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zipper_v3_grpc.CarbonV1/FetchMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV1Server).FetchMetrics(ctx, req.(*carbonzipperpbng.MultiFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarbonV1_FindMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(carbonzipperpbng.MultiGlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV1Server).FindMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zipper_v3_grpc.CarbonV1/FindMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV1Server).FindMetrics(ctx, req.(*carbonzipperpbng.MultiGlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarbonV1_MetricsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(carbonzipperpbng.MultiMetricsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV1Server).MetricsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zipper_v3_grpc.CarbonV1/MetricsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV1Server).MetricsInfo(ctx, req.(*carbonzipperpbng.MultiMetricsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarbonV1_ListMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV1Server).ListMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zipper_v3_grpc.CarbonV1/ListMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV1Server).ListMetrics(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarbonV1_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV1Server).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zipper_v3_grpc.CarbonV1/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV1Server).Stats(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarbonV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zipper_v3_grpc.CarbonV1",
	HandlerType: (*CarbonV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _CarbonV1_GetVersion_Handler,
		},
		{
			MethodName: "FetchMetrics",
			Handler:    _CarbonV1_FetchMetrics_Handler,
		},
		{
			MethodName: "FindMetrics",
			Handler:    _CarbonV1_FindMetrics_Handler,
		},
		{
			MethodName: "MetricsInfo",
			Handler:    _CarbonV1_MetricsInfo_Handler,
		},
		{
			MethodName: "ListMetrics",
			Handler:    _CarbonV1_ListMetrics_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _CarbonV1_Stats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zipper_v3_grpc.proto",
}

func (m *ProtocolVersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolVersionResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintZipperV3Grpc(dAtA, i, uint64(m.Version))
	}
	return i, nil
}

func encodeFixed64ZipperV3Grpc(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ZipperV3Grpc(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintZipperV3Grpc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ProtocolVersionResponse) Size() (n int) {
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovZipperV3Grpc(uint64(m.Version))
	}
	return n
}

func sovZipperV3Grpc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozZipperV3Grpc(x uint64) (n int) {
	return sovZipperV3Grpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ProtocolVersionResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProtocolVersionResponse{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringZipperV3Grpc(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ProtocolVersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZipperV3Grpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtocolVersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolVersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZipperV3Grpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipZipperV3Grpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthZipperV3Grpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipZipperV3Grpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowZipperV3Grpc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowZipperV3Grpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowZipperV3Grpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthZipperV3Grpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowZipperV3Grpc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipZipperV3Grpc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthZipperV3Grpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowZipperV3Grpc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("zipper_v3_grpc.proto", fileDescriptorZipperV3Grpc) }

var fileDescriptorZipperV3Grpc = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcb, 0x4a, 0xfb, 0x40,
	0x18, 0xc5, 0x33, 0xfc, 0x2f, 0xca, 0x54, 0x5c, 0x04, 0xd1, 0x12, 0x61, 0x94, 0x56, 0xf1, 0x82,
	0x4e, 0xd1, 0xbe, 0x81, 0x97, 0x16, 0xc5, 0xaa, 0x54, 0x28, 0x82, 0x8b, 0xd2, 0xc4, 0x69, 0x3a,
	0x90, 0x66, 0xc6, 0xcc, 0xa4, 0xa0, 0x2b, 0x1f, 0xc1, 0x95, 0xcf, 0xe0, 0xa3, 0xb8, 0xec, 0xd2,
	0xa5, 0x1d, 0x37, 0x2e, 0xfb, 0x08, 0xd2, 0x49, 0xd3, 0x46, 0x4b, 0x8a, 0xbb, 0xef, 0x7c, 0x39,
	0xe7, 0x77, 0xf2, 0x25, 0x70, 0xe1, 0x81, 0x72, 0x4e, 0x82, 0x7a, 0xa7, 0x58, 0x77, 0x03, 0xee,
	0x60, 0x1e, 0x30, 0xc9, 0xcc, 0xf9, 0xef, 0x5b, 0x6b, 0x65, 0xac, 0xb9, 0x5d, 0x48, 0x8a, 0x28,
	0x60, 0xed, 0xba, 0x54, 0xb6, 0x42, 0x1b, 0x3b, 0xac, 0x5d, 0x70, 0x99, 0xcb, 0x0a, 0x7a, 0x6d,
	0x87, 0x4d, 0xad, 0xb4, 0xd0, 0xd3, 0xd0, 0xbe, 0xec, 0x32, 0xe6, 0x7a, 0x64, 0xec, 0x22, 0x6d,
	0x2e, 0xef, 0xa3, 0x87, 0xb9, 0x22, 0x5c, 0xba, 0x1c, 0x0c, 0x0e, 0xf3, 0x6a, 0x24, 0x10, 0x94,
	0xf9, 0x55, 0x22, 0x38, 0xf3, 0x05, 0x31, 0xb3, 0x70, 0xa6, 0x13, 0xad, 0xb2, 0x60, 0x15, 0x6c,
	0xfe, 0xa9, 0xc6, 0x72, 0xff, 0xf9, 0x2f, 0x9c, 0x3d, 0x6c, 0x04, 0x36, 0xf3, 0x6b, 0x7b, 0xe6,
	0x05, 0x84, 0x65, 0x22, 0x87, 0x61, 0x73, 0x11, 0x47, 0x6d, 0x38, 0x6e, 0xc3, 0xc7, 0x83, 0x36,
	0x6b, 0x03, 0xff, 0xb8, 0x3d, 0xa5, 0x35, 0x67, 0x98, 0x37, 0x70, 0xae, 0x44, 0xa4, 0xd3, 0xaa,
	0x10, 0x19, 0x50, 0x47, 0x98, 0x79, 0xec, 0xe8, 0xae, 0x08, 0xc0, 0x6d, 0xdf, 0xc5, 0x95, 0xd0,
	0x93, 0x54, 0x9b, 0xaa, 0xe4, 0x2e, 0x24, 0x42, 0x5a, 0x6b, 0xd3, 0x4d, 0x23, 0xf8, 0x35, 0xcc,
	0x94, 0xa8, 0x7f, 0x1b, 0xb3, 0x73, 0x29, 0xb1, 0xb2, 0xc7, 0xec, 0x18, 0x9d, 0x9f, 0xea, 0x19,
	0x91, 0x9b, 0x30, 0x33, 0xa4, 0x9e, 0xf8, 0x4d, 0x66, 0x6e, 0xa5, 0xa4, 0x12, 0x9e, 0xb8, 0x60,
	0xfb, 0x37, 0xd6, 0x51, 0xcf, 0x39, 0xcc, 0x9c, 0x51, 0x21, 0xe3, 0x0b, 0xd2, 0x3e, 0xf8, 0xfa,
	0x24, 0x34, 0x11, 0x4b, 0xf0, 0x4e, 0xe1, 0xbf, 0x2b, 0xd9, 0x90, 0x62, 0xca, 0xaf, 0x9b, 0x7c,
	0x3d, 0x4d, 0x39, 0x22, 0xb2, 0x41, 0xbd, 0x04, 0xeb, 0x60, 0xa7, 0xdb, 0x43, 0xc6, 0x5b, 0x0f,
	0x19, 0xfd, 0x1e, 0x02, 0x8f, 0x0a, 0x81, 0x17, 0x85, 0xc0, 0xab, 0x42, 0xa0, 0xab, 0x10, 0x78,
	0x57, 0x08, 0x7c, 0x2a, 0x64, 0xf4, 0x15, 0x02, 0x4f, 0x1f, 0xc8, 0xb0, 0xff, 0xeb, 0xa2, 0xe2,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xbf, 0x86, 0x28, 0x17, 0x03, 0x00, 0x00,
}
