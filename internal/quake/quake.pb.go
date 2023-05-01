// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: quake.proto

package quake

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

// Quake is for earthquake information.
type Quake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All the Phase used to locate the quake.
	Phases []*Phase `protobuf:"bytes,1,rep,name=phases,proto3" json:"phases,omitempty"`
	// Magnitude all the magnitudes associated with the quake.
	Magnitudes            []*Magnitude `protobuf:"bytes,2,rep,name=magnitudes,proto3" json:"magnitudes,omitempty"`
	PublicID              string       `protobuf:"bytes,10,opt,name=public_iD,json=publicID,proto3" json:"public_iD,omitempty"`
	Type                  string       `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	Time                  *Timestamp   `protobuf:"bytes,12,opt,name=time,proto3" json:"time,omitempty"`
	Latitude              float64      `protobuf:"fixed64,13,opt,name=latitude,proto3" json:"latitude,omitempty"`
	LatitudeUncertainty   float64      `protobuf:"fixed64,14,opt,name=latitude_uncertainty,json=latitudeUncertainty,proto3" json:"latitude_uncertainty,omitempty"`
	Longitude             float64      `protobuf:"fixed64,15,opt,name=longitude,proto3" json:"longitude,omitempty"`
	LongitudeUncertainty  float64      `protobuf:"fixed64,16,opt,name=longitude_uncertainty,json=longitudeUncertainty,proto3" json:"longitude_uncertainty,omitempty"`
	Depth                 float64      `protobuf:"fixed64,17,opt,name=depth,proto3" json:"depth,omitempty"`
	DepthUncertainty      float64      `protobuf:"fixed64,18,opt,name=depth_uncertainty,json=depthUncertainty,proto3" json:"depth_uncertainty,omitempty"`
	DepthType             string       `protobuf:"bytes,19,opt,name=depth_type,json=depthType,proto3" json:"depth_type,omitempty"`
	Magnitude             float64      `protobuf:"fixed64,20,opt,name=magnitude,proto3" json:"magnitude,omitempty"`
	MagnitudeUncertainty  float64      `protobuf:"fixed64,21,opt,name=magnitude_uncertainty,json=magnitudeUncertainty,proto3" json:"magnitude_uncertainty,omitempty"`
	MagnitudeType         string       `protobuf:"bytes,22,opt,name=magnitude_type,json=magnitudeType,proto3" json:"magnitude_type,omitempty"`
	MagnitudeStationCount int64        `protobuf:"varint,23,opt,name=magnitude_station_count,json=magnitudeStationCount,proto3" json:"magnitude_station_count,omitempty"`
	Method                string       `protobuf:"bytes,24,opt,name=method,proto3" json:"method,omitempty"`
	EarthModel            string       `protobuf:"bytes,25,opt,name=earth_model,json=earthModel,proto3" json:"earth_model,omitempty"`
	EvaluationMode        string       `protobuf:"bytes,26,opt,name=evaluation_mode,json=evaluationMode,proto3" json:"evaluation_mode,omitempty"`
	EvaluationStatus      string       `protobuf:"bytes,27,opt,name=evaluation_status,json=evaluationStatus,proto3" json:"evaluation_status,omitempty"`
	UsedPhaseCount        int64        `protobuf:"varint,28,opt,name=used_phase_count,json=usedPhaseCount,proto3" json:"used_phase_count,omitempty"`
	UsedStationCount      int64        `protobuf:"varint,29,opt,name=used_station_count,json=usedStationCount,proto3" json:"used_station_count,omitempty"`
	StandardError         float64      `protobuf:"fixed64,30,opt,name=standard_error,json=standardError,proto3" json:"standard_error,omitempty"`
	AzimuthalGap          float64      `protobuf:"fixed64,31,opt,name=azimuthal_gap,json=azimuthalGap,proto3" json:"azimuthal_gap,omitempty"`
	MinimumDistance       float64      `protobuf:"fixed64,32,opt,name=minimum_distance,json=minimumDistance,proto3" json:"minimum_distance,omitempty"`
	Agency                string       `protobuf:"bytes,33,opt,name=agency,proto3" json:"agency,omitempty"`
	ModificationTime      *Timestamp   `protobuf:"bytes,34,opt,name=modification_time,json=modificationTime,proto3" json:"modification_time,omitempty"`
}

func (x *Quake) Reset() {
	*x = Quake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quake) ProtoMessage() {}

func (x *Quake) ProtoReflect() protoreflect.Message {
	mi := &file_quake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quake.ProtoReflect.Descriptor instead.
func (*Quake) Descriptor() ([]byte, []int) {
	return file_quake_proto_rawDescGZIP(), []int{0}
}

func (x *Quake) GetPhases() []*Phase {
	if x != nil {
		return x.Phases
	}
	return nil
}

func (x *Quake) GetMagnitudes() []*Magnitude {
	if x != nil {
		return x.Magnitudes
	}
	return nil
}

func (x *Quake) GetPublicID() string {
	if x != nil {
		return x.PublicID
	}
	return ""
}

func (x *Quake) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Quake) GetTime() *Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Quake) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Quake) GetLatitudeUncertainty() float64 {
	if x != nil {
		return x.LatitudeUncertainty
	}
	return 0
}

func (x *Quake) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Quake) GetLongitudeUncertainty() float64 {
	if x != nil {
		return x.LongitudeUncertainty
	}
	return 0
}

func (x *Quake) GetDepth() float64 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *Quake) GetDepthUncertainty() float64 {
	if x != nil {
		return x.DepthUncertainty
	}
	return 0
}

func (x *Quake) GetDepthType() string {
	if x != nil {
		return x.DepthType
	}
	return ""
}

func (x *Quake) GetMagnitude() float64 {
	if x != nil {
		return x.Magnitude
	}
	return 0
}

func (x *Quake) GetMagnitudeUncertainty() float64 {
	if x != nil {
		return x.MagnitudeUncertainty
	}
	return 0
}

func (x *Quake) GetMagnitudeType() string {
	if x != nil {
		return x.MagnitudeType
	}
	return ""
}

func (x *Quake) GetMagnitudeStationCount() int64 {
	if x != nil {
		return x.MagnitudeStationCount
	}
	return 0
}

func (x *Quake) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Quake) GetEarthModel() string {
	if x != nil {
		return x.EarthModel
	}
	return ""
}

func (x *Quake) GetEvaluationMode() string {
	if x != nil {
		return x.EvaluationMode
	}
	return ""
}

func (x *Quake) GetEvaluationStatus() string {
	if x != nil {
		return x.EvaluationStatus
	}
	return ""
}

func (x *Quake) GetUsedPhaseCount() int64 {
	if x != nil {
		return x.UsedPhaseCount
	}
	return 0
}

func (x *Quake) GetUsedStationCount() int64 {
	if x != nil {
		return x.UsedStationCount
	}
	return 0
}

func (x *Quake) GetStandardError() float64 {
	if x != nil {
		return x.StandardError
	}
	return 0
}

func (x *Quake) GetAzimuthalGap() float64 {
	if x != nil {
		return x.AzimuthalGap
	}
	return 0
}

func (x *Quake) GetMinimumDistance() float64 {
	if x != nil {
		return x.MinimumDistance
	}
	return 0
}

func (x *Quake) GetAgency() string {
	if x != nil {
		return x.Agency
	}
	return ""
}

func (x *Quake) GetModificationTime() *Timestamp {
	if x != nil {
		return x.ModificationTime
	}
	return nil
}

// Phase represents a seismic phase.
type Phase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkCode      string     `protobuf:"bytes,1,opt,name=network_code,json=networkCode,proto3" json:"network_code,omitempty"`
	StationCode      string     `protobuf:"bytes,2,opt,name=station_code,json=stationCode,proto3" json:"station_code,omitempty"`
	LocationCode     string     `protobuf:"bytes,3,opt,name=location_code,json=locationCode,proto3" json:"location_code,omitempty"`
	ChannelCode      string     `protobuf:"bytes,4,opt,name=channel_code,json=channelCode,proto3" json:"channel_code,omitempty"`
	Phase            string     `protobuf:"bytes,5,opt,name=phase,proto3" json:"phase,omitempty"`
	Time             *Timestamp `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	Residual         float64    `protobuf:"fixed64,7,opt,name=residual,proto3" json:"residual,omitempty"`
	Weight           float64    `protobuf:"fixed64,8,opt,name=weight,proto3" json:"weight,omitempty"`
	Azimuth          float64    `protobuf:"fixed64,9,opt,name=azimuth,proto3" json:"azimuth,omitempty"`
	Distance         float64    `protobuf:"fixed64,10,opt,name=distance,proto3" json:"distance,omitempty"`
	EvaluationMode   string     `protobuf:"bytes,11,opt,name=evaluation_mode,json=evaluationMode,proto3" json:"evaluation_mode,omitempty"`
	EvaluationStatus string     `protobuf:"bytes,12,opt,name=evaluation_status,json=evaluationStatus,proto3" json:"evaluation_status,omitempty"`
}

func (x *Phase) Reset() {
	*x = Phase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Phase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phase) ProtoMessage() {}

func (x *Phase) ProtoReflect() protoreflect.Message {
	mi := &file_quake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phase.ProtoReflect.Descriptor instead.
func (*Phase) Descriptor() ([]byte, []int) {
	return file_quake_proto_rawDescGZIP(), []int{1}
}

func (x *Phase) GetNetworkCode() string {
	if x != nil {
		return x.NetworkCode
	}
	return ""
}

func (x *Phase) GetStationCode() string {
	if x != nil {
		return x.StationCode
	}
	return ""
}

func (x *Phase) GetLocationCode() string {
	if x != nil {
		return x.LocationCode
	}
	return ""
}

func (x *Phase) GetChannelCode() string {
	if x != nil {
		return x.ChannelCode
	}
	return ""
}

func (x *Phase) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *Phase) GetTime() *Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Phase) GetResidual() float64 {
	if x != nil {
		return x.Residual
	}
	return 0
}

func (x *Phase) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Phase) GetAzimuth() float64 {
	if x != nil {
		return x.Azimuth
	}
	return 0
}

func (x *Phase) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Phase) GetEvaluationMode() string {
	if x != nil {
		return x.EvaluationMode
	}
	return ""
}

func (x *Phase) GetEvaluationStatus() string {
	if x != nil {
		return x.EvaluationStatus
	}
	return ""
}

// Magnitude represents a quake magnitude.
type Magnitude struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// StationMagnitude that have contributed to the Magnitude.
	StationMagnitude     []*StationMagnitude `protobuf:"bytes,1,rep,name=station_magnitude,json=stationMagnitude,proto3" json:"station_magnitude,omitempty"`
	Magnitude            float64             `protobuf:"fixed64,2,opt,name=magnitude,proto3" json:"magnitude,omitempty"`
	MagnitudeUncertainty float64             `protobuf:"fixed64,3,opt,name=magnitude_uncertainty,json=magnitudeUncertainty,proto3" json:"magnitude_uncertainty,omitempty"`
	Type                 string              `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Method               string              `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	StationCount         int64               `protobuf:"varint,6,opt,name=station_count,json=stationCount,proto3" json:"station_count,omitempty"`
}

func (x *Magnitude) Reset() {
	*x = Magnitude{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quake_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Magnitude) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Magnitude) ProtoMessage() {}

func (x *Magnitude) ProtoReflect() protoreflect.Message {
	mi := &file_quake_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Magnitude.ProtoReflect.Descriptor instead.
func (*Magnitude) Descriptor() ([]byte, []int) {
	return file_quake_proto_rawDescGZIP(), []int{2}
}

func (x *Magnitude) GetStationMagnitude() []*StationMagnitude {
	if x != nil {
		return x.StationMagnitude
	}
	return nil
}

func (x *Magnitude) GetMagnitude() float64 {
	if x != nil {
		return x.Magnitude
	}
	return 0
}

func (x *Magnitude) GetMagnitudeUncertainty() float64 {
	if x != nil {
		return x.MagnitudeUncertainty
	}
	return 0
}

func (x *Magnitude) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Magnitude) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Magnitude) GetStationCount() int64 {
	if x != nil {
		return x.StationCount
	}
	return 0
}

// StationMagnitude the magnitude calculated at a single station.
type StationMagnitude struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkCode  string  `protobuf:"bytes,1,opt,name=network_code,json=networkCode,proto3" json:"network_code,omitempty"`
	StationCode  string  `protobuf:"bytes,2,opt,name=station_code,json=stationCode,proto3" json:"station_code,omitempty"`
	LocationCode string  `protobuf:"bytes,3,opt,name=location_code,json=locationCode,proto3" json:"location_code,omitempty"`
	ChannelCode  string  `protobuf:"bytes,4,opt,name=channel_code,json=channelCode,proto3" json:"channel_code,omitempty"`
	Magnitude    float64 `protobuf:"fixed64,5,opt,name=magnitude,proto3" json:"magnitude,omitempty"`
	Type         string  `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Residual     float64 `protobuf:"fixed64,7,opt,name=residual,proto3" json:"residual,omitempty"`
	Weight       float64 `protobuf:"fixed64,8,opt,name=weight,proto3" json:"weight,omitempty"`
	Amplitude    float64 `protobuf:"fixed64,9,opt,name=amplitude,proto3" json:"amplitude,omitempty"`
}

func (x *StationMagnitude) Reset() {
	*x = StationMagnitude{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quake_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationMagnitude) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationMagnitude) ProtoMessage() {}

func (x *StationMagnitude) ProtoReflect() protoreflect.Message {
	mi := &file_quake_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationMagnitude.ProtoReflect.Descriptor instead.
func (*StationMagnitude) Descriptor() ([]byte, []int) {
	return file_quake_proto_rawDescGZIP(), []int{3}
}

func (x *StationMagnitude) GetNetworkCode() string {
	if x != nil {
		return x.NetworkCode
	}
	return ""
}

func (x *StationMagnitude) GetStationCode() string {
	if x != nil {
		return x.StationCode
	}
	return ""
}

func (x *StationMagnitude) GetLocationCode() string {
	if x != nil {
		return x.LocationCode
	}
	return ""
}

func (x *StationMagnitude) GetChannelCode() string {
	if x != nil {
		return x.ChannelCode
	}
	return ""
}

func (x *StationMagnitude) GetMagnitude() float64 {
	if x != nil {
		return x.Magnitude
	}
	return 0
}

func (x *StationMagnitude) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StationMagnitude) GetResidual() float64 {
	if x != nil {
		return x.Residual
	}
	return 0
}

func (x *StationMagnitude) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *StationMagnitude) GetAmplitude() float64 {
	if x != nil {
		return x.Amplitude
	}
	return 0
}

// Timestamp for encoding time stamps.
type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix time in seconds
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Frational part of time in nanoseconds.
	Nanos int64 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quake_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_quake_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_quake_proto_rawDescGZIP(), []int{4}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int64 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_quake_proto protoreflect.FileDescriptor

var file_quake_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x71,
	0x75, 0x61, 0x6b, 0x65, 0x22, 0xa1, 0x08, 0x0a, 0x05, 0x51, 0x75, 0x61, 0x6b, 0x65, 0x12, 0x24,
	0x0a, 0x06, 0x70, 0x68, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x71, 0x75, 0x61, 0x6b, 0x65, 0x2e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x06, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x75, 0x61, 0x6b, 0x65,
	0x2e, 0x4d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x52, 0x0a, 0x6d, 0x61, 0x67, 0x6e,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x69, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x75, 0x61, 0x6b, 0x65, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x5f, 0x75, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x74,
	0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x55, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x5f, 0x75, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69,
	0x6e, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x55, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x75,
	0x6e, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x10, 0x64, 0x65, 0x70, 0x74, 0x68, 0x55, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e,
	0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x70, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x33, 0x0a, 0x15, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x5f, 0x75, 0x6e, 0x63,
	0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14,
	0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x55, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x61,
	0x69, 0x6e, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61,
	0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x6d,
	0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x6d, 0x61,
	0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x61, 0x72, 0x74, 0x68, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x61, 0x72, 0x74, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x75, 0x73,
	0x65, 0x64, 0x50, 0x68, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x75, 0x73, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x7a, 0x69, 0x6d, 0x75, 0x74, 0x68, 0x61, 0x6c, 0x5f, 0x67,
	0x61, 0x70, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x61, 0x7a, 0x69, 0x6d, 0x75, 0x74,
	0x68, 0x61, 0x6c, 0x47, 0x61, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75,
	0x6d, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x21, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3d, 0x0a, 0x11, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x22,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x75, 0x61, 0x6b, 0x65, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x91, 0x03, 0x0a, 0x05, 0x50, 0x68, 0x61,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x75, 0x61, 0x6b, 0x65, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x7a, 0x69, 0x6d, 0x75, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x61, 0x7a, 0x69, 0x6d, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xf5, 0x01, 0x0a,
	0x09, 0x4d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x44, 0x0a, 0x11, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x71, 0x75, 0x61, 0x6b, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x52, 0x10,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x33,
	0x0a, 0x15, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x5f, 0x75, 0x6e, 0x63, 0x65,
	0x72, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x6d,
	0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x55, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69,
	0x6e, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa4, 0x02, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x67, 0x6e, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6d, 0x61, 0x67, 0x6e,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x69, 0x64, 0x75, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x69, 0x64, 0x75, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x3b, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x71, 0x75,
	0x61, 0x6b, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quake_proto_rawDescOnce sync.Once
	file_quake_proto_rawDescData = file_quake_proto_rawDesc
)

func file_quake_proto_rawDescGZIP() []byte {
	file_quake_proto_rawDescOnce.Do(func() {
		file_quake_proto_rawDescData = protoimpl.X.CompressGZIP(file_quake_proto_rawDescData)
	})
	return file_quake_proto_rawDescData
}

var file_quake_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_quake_proto_goTypes = []interface{}{
	(*Quake)(nil),            // 0: quake.Quake
	(*Phase)(nil),            // 1: quake.Phase
	(*Magnitude)(nil),        // 2: quake.Magnitude
	(*StationMagnitude)(nil), // 3: quake.StationMagnitude
	(*Timestamp)(nil),        // 4: quake.Timestamp
}
var file_quake_proto_depIdxs = []int32{
	1, // 0: quake.Quake.phases:type_name -> quake.Phase
	2, // 1: quake.Quake.magnitudes:type_name -> quake.Magnitude
	4, // 2: quake.Quake.time:type_name -> quake.Timestamp
	4, // 3: quake.Quake.modification_time:type_name -> quake.Timestamp
	4, // 4: quake.Phase.time:type_name -> quake.Timestamp
	3, // 5: quake.Magnitude.station_magnitude:type_name -> quake.StationMagnitude
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_quake_proto_init() }
func file_quake_proto_init() {
	if File_quake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quake); i {
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
		file_quake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Phase); i {
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
		file_quake_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Magnitude); i {
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
		file_quake_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationMagnitude); i {
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
		file_quake_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
			RawDescriptor: file_quake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_quake_proto_goTypes,
		DependencyIndexes: file_quake_proto_depIdxs,
		MessageInfos:      file_quake_proto_msgTypes,
	}.Build()
	File_quake_proto = out.File
	file_quake_proto_rawDesc = nil
	file_quake_proto_goTypes = nil
	file_quake_proto_depIdxs = nil
}
