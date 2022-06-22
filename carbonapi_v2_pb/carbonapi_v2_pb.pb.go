// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: carbonapi_v2_pb/carbonapi_v2_pb.proto

package carbonapi_v2_pb

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

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StartTime int32     `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	StopTime  int32     `protobuf:"varint,3,opt,name=stopTime,proto3" json:"stopTime,omitempty"`
	StepTime  int32     `protobuf:"varint,4,opt,name=stepTime,proto3" json:"stepTime,omitempty"`
	Values    []float64 `protobuf:"fixed64,5,rep,packed,name=values,proto3" json:"values,omitempty"`
	IsAbsent  []bool    `protobuf:"varint,6,rep,packed,name=isAbsent,proto3" json:"isAbsent,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{0}
}

func (x *FetchResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FetchResponse) GetStartTime() int32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *FetchResponse) GetStopTime() int32 {
	if x != nil {
		return x.StopTime
	}
	return 0
}

func (x *FetchResponse) GetStepTime() int32 {
	if x != nil {
		return x.StepTime
	}
	return 0
}

func (x *FetchResponse) GetValues() []float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *FetchResponse) GetIsAbsent() []bool {
	if x != nil {
		return x.IsAbsent
	}
	return nil
}

type MultiFetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*FetchResponse `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *MultiFetchResponse) Reset() {
	*x = MultiFetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiFetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiFetchResponse) ProtoMessage() {}

func (x *MultiFetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiFetchResponse.ProtoReflect.Descriptor instead.
func (*MultiFetchResponse) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{1}
}

func (x *MultiFetchResponse) GetMetrics() []*FetchResponse {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type GlobMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	IsLeaf bool   `protobuf:"varint,2,opt,name=isLeaf,proto3" json:"isLeaf,omitempty"`
}

func (x *GlobMatch) Reset() {
	*x = GlobMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobMatch) ProtoMessage() {}

func (x *GlobMatch) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobMatch.ProtoReflect.Descriptor instead.
func (*GlobMatch) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{2}
}

func (x *GlobMatch) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GlobMatch) GetIsLeaf() bool {
	if x != nil {
		return x.IsLeaf
	}
	return false
}

type GlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Matches []*GlobMatch `protobuf:"bytes,2,rep,name=matches,proto3" json:"matches,omitempty"`
}

func (x *GlobResponse) Reset() {
	*x = GlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobResponse) ProtoMessage() {}

func (x *GlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobResponse.ProtoReflect.Descriptor instead.
func (*GlobResponse) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{3}
}

func (x *GlobResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GlobResponse) GetMatches() []*GlobMatch {
	if x != nil {
		return x.Matches
	}
	return nil
}

type Retention struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecondsPerPoint int32 `protobuf:"varint,1,opt,name=secondsPerPoint,proto3" json:"secondsPerPoint,omitempty"`
	NumberOfPoints  int32 `protobuf:"varint,2,opt,name=numberOfPoints,proto3" json:"numberOfPoints,omitempty"`
}

func (x *Retention) Reset() {
	*x = Retention{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Retention) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retention) ProtoMessage() {}

func (x *Retention) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retention.ProtoReflect.Descriptor instead.
func (*Retention) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{4}
}

func (x *Retention) GetSecondsPerPoint() int32 {
	if x != nil {
		return x.SecondsPerPoint
	}
	return 0
}

func (x *Retention) GetNumberOfPoints() int32 {
	if x != nil {
		return x.NumberOfPoints
	}
	return 0
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AggregationMethod string       `protobuf:"bytes,2,opt,name=aggregationMethod,proto3" json:"aggregationMethod,omitempty"`
	MaxRetention      int32        `protobuf:"varint,3,opt,name=maxRetention,proto3" json:"maxRetention,omitempty"`
	XFilesFactor      float32      `protobuf:"fixed32,4,opt,name=xFilesFactor,proto3" json:"xFilesFactor,omitempty"`
	Retentions        []*Retention `protobuf:"bytes,5,rep,name=retentions,proto3" json:"retentions,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{5}
}

func (x *InfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InfoResponse) GetAggregationMethod() string {
	if x != nil {
		return x.AggregationMethod
	}
	return ""
}

func (x *InfoResponse) GetMaxRetention() int32 {
	if x != nil {
		return x.MaxRetention
	}
	return 0
}

func (x *InfoResponse) GetXFilesFactor() float32 {
	if x != nil {
		return x.XFilesFactor
	}
	return 0
}

func (x *InfoResponse) GetRetentions() []*Retention {
	if x != nil {
		return x.Retentions
	}
	return nil
}

type ServerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server string        `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Info   *InfoResponse `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ServerInfoResponse) Reset() {
	*x = ServerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfoResponse) ProtoMessage() {}

func (x *ServerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfoResponse.ProtoReflect.Descriptor instead.
func (*ServerInfoResponse) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{6}
}

func (x *ServerInfoResponse) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *ServerInfoResponse) GetInfo() *InfoResponse {
	if x != nil {
		return x.Info
	}
	return nil
}

type ZipperInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*ServerInfoResponse `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *ZipperInfoResponse) Reset() {
	*x = ZipperInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZipperInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZipperInfoResponse) ProtoMessage() {}

func (x *ZipperInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZipperInfoResponse.ProtoReflect.Descriptor instead.
func (*ZipperInfoResponse) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{7}
}

func (x *ZipperInfoResponse) GetResponses() []*ServerInfoResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

type ListMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []string `protobuf:"bytes,1,rep,name=Metrics,proto3" json:"Metrics,omitempty"`
}

func (x *ListMetricsResponse) Reset() {
	*x = ListMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMetricsResponse) ProtoMessage() {}

func (x *ListMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMetricsResponse.ProtoReflect.Descriptor instead.
func (*ListMetricsResponse) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{8}
}

func (x *ListMetricsResponse) GetMetrics() []string {
	if x != nil {
		return x.Metrics
	}
	return nil
}

// MetricDetails and MetricDetailsResponse is not guaranteed to stay the same in future releases
// But we'll do our best to make them stable and only to extend them if that's necessary
type MetricDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size    int64 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	ModTime int64 `protobuf:"varint,3,opt,name=ModTime,proto3" json:"ModTime,omitempty"`
	ATime   int64 `protobuf:"varint,4,opt,name=ATime,proto3" json:"ATime,omitempty"`
	RdTime  int64 `protobuf:"varint,5,opt,name=RdTime,proto3" json:"RdTime,omitempty"`
}

func (x *MetricDetails) Reset() {
	*x = MetricDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDetails) ProtoMessage() {}

func (x *MetricDetails) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDetails.ProtoReflect.Descriptor instead.
func (*MetricDetails) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{9}
}

func (x *MetricDetails) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MetricDetails) GetModTime() int64 {
	if x != nil {
		return x.ModTime
	}
	return 0
}

func (x *MetricDetails) GetATime() int64 {
	if x != nil {
		return x.ATime
	}
	return 0
}

func (x *MetricDetails) GetRdTime() int64 {
	if x != nil {
		return x.RdTime
	}
	return 0
}

type MetricDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics    map[string]*MetricDetails `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FreeSpace  uint64                    `protobuf:"varint,2,opt,name=FreeSpace,proto3" json:"FreeSpace,omitempty"`
	TotalSpace uint64                    `protobuf:"varint,3,opt,name=TotalSpace,proto3" json:"TotalSpace,omitempty"`
}

func (x *MetricDetailsResponse) Reset() {
	*x = MetricDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDetailsResponse) ProtoMessage() {}

func (x *MetricDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDetailsResponse.ProtoReflect.Descriptor instead.
func (*MetricDetailsResponse) Descriptor() ([]byte, []int) {
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP(), []int{10}
}

func (x *MetricDetailsResponse) GetMetrics() map[string]*MetricDetails {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *MetricDetailsResponse) GetFreeSpace() uint64 {
	if x != nil {
		return x.FreeSpace
	}
	return 0
}

func (x *MetricDetailsResponse) GetTotalSpace() uint64 {
	if x != nil {
		return x.TotalSpace
	}
	return 0
}

var File_carbonapi_v2_pb_carbonapi_v2_pb_proto protoreflect.FileDescriptor

var file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x70,
	0x62, 0x2f, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x70,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x70, 0x62, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x73, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x65, 0x70,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x01, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x73, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x12, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x70,
	0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x37, 0x0a, 0x09, 0x47, 0x6c, 0x6f, 0x62,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4c,
	0x65, 0x61, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4c, 0x65, 0x61,
	0x66, 0x22, 0x58, 0x0a, 0x0c, 0x47, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x5d, 0x0a, 0x09, 0x52,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x50, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x50, 0x65, 0x72, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x66, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x0c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x62,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x5f, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x70, 0x62, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x57, 0x0a, 0x12, 0x5a, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x61,
	0x72, 0x62, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x6b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x41, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x52, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x15, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x70, 0x69,
	0x5f, 0x76, 0x32, 0x5f, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x72, 0x65, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x46, 0x72, 0x65, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x1a, 0x5a, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32,
	0x5f, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x31, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescOnce sync.Once
	file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescData = file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDesc
)

func file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescGZIP() []byte {
	file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescOnce.Do(func() {
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescData)
	})
	return file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDescData
}

var file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_carbonapi_v2_pb_carbonapi_v2_pb_proto_goTypes = []interface{}{
	(*FetchResponse)(nil),         // 0: carbonapi_v2_pb.FetchResponse
	(*MultiFetchResponse)(nil),    // 1: carbonapi_v2_pb.MultiFetchResponse
	(*GlobMatch)(nil),             // 2: carbonapi_v2_pb.GlobMatch
	(*GlobResponse)(nil),          // 3: carbonapi_v2_pb.GlobResponse
	(*Retention)(nil),             // 4: carbonapi_v2_pb.Retention
	(*InfoResponse)(nil),          // 5: carbonapi_v2_pb.InfoResponse
	(*ServerInfoResponse)(nil),    // 6: carbonapi_v2_pb.ServerInfoResponse
	(*ZipperInfoResponse)(nil),    // 7: carbonapi_v2_pb.ZipperInfoResponse
	(*ListMetricsResponse)(nil),   // 8: carbonapi_v2_pb.ListMetricsResponse
	(*MetricDetails)(nil),         // 9: carbonapi_v2_pb.MetricDetails
	(*MetricDetailsResponse)(nil), // 10: carbonapi_v2_pb.MetricDetailsResponse
	nil,                           // 11: carbonapi_v2_pb.MetricDetailsResponse.MetricsEntry
}
var file_carbonapi_v2_pb_carbonapi_v2_pb_proto_depIdxs = []int32{
	0,  // 0: carbonapi_v2_pb.MultiFetchResponse.metrics:type_name -> carbonapi_v2_pb.FetchResponse
	2,  // 1: carbonapi_v2_pb.GlobResponse.matches:type_name -> carbonapi_v2_pb.GlobMatch
	4,  // 2: carbonapi_v2_pb.InfoResponse.retentions:type_name -> carbonapi_v2_pb.Retention
	5,  // 3: carbonapi_v2_pb.ServerInfoResponse.info:type_name -> carbonapi_v2_pb.InfoResponse
	6,  // 4: carbonapi_v2_pb.ZipperInfoResponse.responses:type_name -> carbonapi_v2_pb.ServerInfoResponse
	11, // 5: carbonapi_v2_pb.MetricDetailsResponse.metrics:type_name -> carbonapi_v2_pb.MetricDetailsResponse.MetricsEntry
	9,  // 6: carbonapi_v2_pb.MetricDetailsResponse.MetricsEntry.value:type_name -> carbonapi_v2_pb.MetricDetails
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_carbonapi_v2_pb_carbonapi_v2_pb_proto_init() }
func file_carbonapi_v2_pb_carbonapi_v2_pb_proto_init() {
	if File_carbonapi_v2_pb_carbonapi_v2_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiFetchResponse); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobMatch); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobResponse); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Retention); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfoResponse); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZipperInfoResponse); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMetricsResponse); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricDetails); i {
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
		file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricDetailsResponse); i {
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
			RawDescriptor: file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_carbonapi_v2_pb_carbonapi_v2_pb_proto_goTypes,
		DependencyIndexes: file_carbonapi_v2_pb_carbonapi_v2_pb_proto_depIdxs,
		MessageInfos:      file_carbonapi_v2_pb_carbonapi_v2_pb_proto_msgTypes,
	}.Build()
	File_carbonapi_v2_pb_carbonapi_v2_pb_proto = out.File
	file_carbonapi_v2_pb_carbonapi_v2_pb_proto_rawDesc = nil
	file_carbonapi_v2_pb_carbonapi_v2_pb_proto_goTypes = nil
	file_carbonapi_v2_pb_carbonapi_v2_pb_proto_depIdxs = nil
}
