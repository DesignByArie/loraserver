// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gw.proto

package gw // import "github.com/brocaar/loraserver/api/gw"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/brocaar/loraserver/api/common"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UplinkTXInfo struct {
	// Frequency (Hz).
	Frequency uint32 `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Modulation.
	Modulation common.Modulation `protobuf:"varint,2,opt,name=modulation,proto3,enum=common.Modulation" json:"modulation,omitempty"`
	// Types that are valid to be assigned to ModulationInfo:
	//	*UplinkTXInfo_LoraModulationInfo
	//	*UplinkTXInfo_FskModulationInfo
	ModulationInfo       isUplinkTXInfo_ModulationInfo `protobuf_oneof:"modulation_info"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *UplinkTXInfo) Reset()         { *m = UplinkTXInfo{} }
func (m *UplinkTXInfo) String() string { return proto.CompactTextString(m) }
func (*UplinkTXInfo) ProtoMessage()    {}
func (*UplinkTXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_f7b5aebffa35cc66, []int{0}
}
func (m *UplinkTXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkTXInfo.Unmarshal(m, b)
}
func (m *UplinkTXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkTXInfo.Marshal(b, m, deterministic)
}
func (dst *UplinkTXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkTXInfo.Merge(dst, src)
}
func (m *UplinkTXInfo) XXX_Size() int {
	return xxx_messageInfo_UplinkTXInfo.Size(m)
}
func (m *UplinkTXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkTXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkTXInfo proto.InternalMessageInfo

type isUplinkTXInfo_ModulationInfo interface {
	isUplinkTXInfo_ModulationInfo()
}

type UplinkTXInfo_LoraModulationInfo struct {
	LoraModulationInfo *LoRaModulationInfo `protobuf:"bytes,3,opt,name=lora_modulation_info,json=loraModulationInfo,proto3,oneof"`
}
type UplinkTXInfo_FskModulationInfo struct {
	FskModulationInfo *FSKModulationInfo `protobuf:"bytes,4,opt,name=fsk_modulation_info,json=fskModulationInfo,proto3,oneof"`
}

func (*UplinkTXInfo_LoraModulationInfo) isUplinkTXInfo_ModulationInfo() {}
func (*UplinkTXInfo_FskModulationInfo) isUplinkTXInfo_ModulationInfo()  {}

func (m *UplinkTXInfo) GetModulationInfo() isUplinkTXInfo_ModulationInfo {
	if m != nil {
		return m.ModulationInfo
	}
	return nil
}

func (m *UplinkTXInfo) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *UplinkTXInfo) GetModulation() common.Modulation {
	if m != nil {
		return m.Modulation
	}
	return common.Modulation_LORA
}

func (m *UplinkTXInfo) GetLoraModulationInfo() *LoRaModulationInfo {
	if x, ok := m.GetModulationInfo().(*UplinkTXInfo_LoraModulationInfo); ok {
		return x.LoraModulationInfo
	}
	return nil
}

func (m *UplinkTXInfo) GetFskModulationInfo() *FSKModulationInfo {
	if x, ok := m.GetModulationInfo().(*UplinkTXInfo_FskModulationInfo); ok {
		return x.FskModulationInfo
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UplinkTXInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UplinkTXInfo_OneofMarshaler, _UplinkTXInfo_OneofUnmarshaler, _UplinkTXInfo_OneofSizer, []interface{}{
		(*UplinkTXInfo_LoraModulationInfo)(nil),
		(*UplinkTXInfo_FskModulationInfo)(nil),
	}
}

func _UplinkTXInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UplinkTXInfo)
	// modulation_info
	switch x := m.ModulationInfo.(type) {
	case *UplinkTXInfo_LoraModulationInfo:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LoraModulationInfo); err != nil {
			return err
		}
	case *UplinkTXInfo_FskModulationInfo:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FskModulationInfo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UplinkTXInfo.ModulationInfo has unexpected type %T", x)
	}
	return nil
}

func _UplinkTXInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UplinkTXInfo)
	switch tag {
	case 3: // modulation_info.lora_modulation_info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoRaModulationInfo)
		err := b.DecodeMessage(msg)
		m.ModulationInfo = &UplinkTXInfo_LoraModulationInfo{msg}
		return true, err
	case 4: // modulation_info.fsk_modulation_info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FSKModulationInfo)
		err := b.DecodeMessage(msg)
		m.ModulationInfo = &UplinkTXInfo_FskModulationInfo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UplinkTXInfo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UplinkTXInfo)
	// modulation_info
	switch x := m.ModulationInfo.(type) {
	case *UplinkTXInfo_LoraModulationInfo:
		s := proto.Size(x.LoraModulationInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UplinkTXInfo_FskModulationInfo:
		s := proto.Size(x.FskModulationInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LoRaModulationInfo struct {
	// Bandwidth.
	Bandwidth uint32 `protobuf:"varint,1,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// Speading-factor.
	SpreadingFactor uint32 `protobuf:"varint,2,opt,name=spreading_factor,json=spreadingFactor,proto3" json:"spreading_factor,omitempty"`
	// Code-rate.
	CodeRate             string   `protobuf:"bytes,3,opt,name=code_rate,json=codeRate,proto3" json:"code_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoRaModulationInfo) Reset()         { *m = LoRaModulationInfo{} }
func (m *LoRaModulationInfo) String() string { return proto.CompactTextString(m) }
func (*LoRaModulationInfo) ProtoMessage()    {}
func (*LoRaModulationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_f7b5aebffa35cc66, []int{1}
}
func (m *LoRaModulationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoRaModulationInfo.Unmarshal(m, b)
}
func (m *LoRaModulationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoRaModulationInfo.Marshal(b, m, deterministic)
}
func (dst *LoRaModulationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoRaModulationInfo.Merge(dst, src)
}
func (m *LoRaModulationInfo) XXX_Size() int {
	return xxx_messageInfo_LoRaModulationInfo.Size(m)
}
func (m *LoRaModulationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LoRaModulationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LoRaModulationInfo proto.InternalMessageInfo

func (m *LoRaModulationInfo) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *LoRaModulationInfo) GetSpreadingFactor() uint32 {
	if m != nil {
		return m.SpreadingFactor
	}
	return 0
}

func (m *LoRaModulationInfo) GetCodeRate() string {
	if m != nil {
		return m.CodeRate
	}
	return ""
}

type FSKModulationInfo struct {
	// Bandwidth.
	Bandwidth uint32 `protobuf:"varint,1,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// Bitrate.
	Bitrate              uint32   `protobuf:"varint,2,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FSKModulationInfo) Reset()         { *m = FSKModulationInfo{} }
func (m *FSKModulationInfo) String() string { return proto.CompactTextString(m) }
func (*FSKModulationInfo) ProtoMessage()    {}
func (*FSKModulationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_f7b5aebffa35cc66, []int{2}
}
func (m *FSKModulationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FSKModulationInfo.Unmarshal(m, b)
}
func (m *FSKModulationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FSKModulationInfo.Marshal(b, m, deterministic)
}
func (dst *FSKModulationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FSKModulationInfo.Merge(dst, src)
}
func (m *FSKModulationInfo) XXX_Size() int {
	return xxx_messageInfo_FSKModulationInfo.Size(m)
}
func (m *FSKModulationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FSKModulationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FSKModulationInfo proto.InternalMessageInfo

func (m *FSKModulationInfo) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *FSKModulationInfo) GetBitrate() uint32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

type UplinkRXInfo struct {
	// Gateway ID.
	GatewayId []byte `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// RX time (only set when the gateway has a GPS module).
	Time *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// RX time since GPS epoch (only set when the gateway has a GPS module).
	TimeSinceGpsEpoch *duration.Duration `protobuf:"bytes,3,opt,name=time_since_gps_epoch,json=timeSinceGpsEpoch,proto3" json:"time_since_gps_epoch,omitempty"`
	// RSSI.
	Rssi int32 `protobuf:"varint,4,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// LoRa SNR.
	LoraSnr float64 `protobuf:"fixed64,5,opt,name=lora_snr,json=loraSnr,proto3" json:"lora_snr,omitempty"`
	// Channel.
	Channel uint32 `protobuf:"varint,6,opt,name=channel,proto3" json:"channel,omitempty"`
	// RF Chain.
	RfChain uint32 `protobuf:"varint,7,opt,name=rf_chain,json=rfChain,proto3" json:"rf_chain,omitempty"`
	// Board.
	Board uint32 `protobuf:"varint,8,opt,name=board,proto3" json:"board,omitempty"`
	// Antenna.
	Antenna uint32 `protobuf:"varint,9,opt,name=antenna,proto3" json:"antenna,omitempty"`
	// Latitude.
	Latitude float64 `protobuf:"fixed64,10,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,11,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude             float64  `protobuf:"fixed64,12,opt,name=altitude,proto3" json:"altitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UplinkRXInfo) Reset()         { *m = UplinkRXInfo{} }
func (m *UplinkRXInfo) String() string { return proto.CompactTextString(m) }
func (*UplinkRXInfo) ProtoMessage()    {}
func (*UplinkRXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gw_f7b5aebffa35cc66, []int{3}
}
func (m *UplinkRXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkRXInfo.Unmarshal(m, b)
}
func (m *UplinkRXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkRXInfo.Marshal(b, m, deterministic)
}
func (dst *UplinkRXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkRXInfo.Merge(dst, src)
}
func (m *UplinkRXInfo) XXX_Size() int {
	return xxx_messageInfo_UplinkRXInfo.Size(m)
}
func (m *UplinkRXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkRXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkRXInfo proto.InternalMessageInfo

func (m *UplinkRXInfo) GetGatewayId() []byte {
	if m != nil {
		return m.GatewayId
	}
	return nil
}

func (m *UplinkRXInfo) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *UplinkRXInfo) GetTimeSinceGpsEpoch() *duration.Duration {
	if m != nil {
		return m.TimeSinceGpsEpoch
	}
	return nil
}

func (m *UplinkRXInfo) GetRssi() int32 {
	if m != nil {
		return m.Rssi
	}
	return 0
}

func (m *UplinkRXInfo) GetLoraSnr() float64 {
	if m != nil {
		return m.LoraSnr
	}
	return 0
}

func (m *UplinkRXInfo) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *UplinkRXInfo) GetRfChain() uint32 {
	if m != nil {
		return m.RfChain
	}
	return 0
}

func (m *UplinkRXInfo) GetBoard() uint32 {
	if m != nil {
		return m.Board
	}
	return 0
}

func (m *UplinkRXInfo) GetAntenna() uint32 {
	if m != nil {
		return m.Antenna
	}
	return 0
}

func (m *UplinkRXInfo) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *UplinkRXInfo) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *UplinkRXInfo) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func init() {
	proto.RegisterType((*UplinkTXInfo)(nil), "gw.UplinkTXInfo")
	proto.RegisterType((*LoRaModulationInfo)(nil), "gw.LoRaModulationInfo")
	proto.RegisterType((*FSKModulationInfo)(nil), "gw.FSKModulationInfo")
	proto.RegisterType((*UplinkRXInfo)(nil), "gw.UplinkRXInfo")
}

func init() { proto.RegisterFile("gw.proto", fileDescriptor_gw_f7b5aebffa35cc66) }

var fileDescriptor_gw_f7b5aebffa35cc66 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0xa5, 0x6b, 0xd7, 0xf6, 0xb6, 0xfb, 0x8d, 0x9a, 0x81, 0xb2, 0xf2, 0xaf, 0xaa, 0x10,
	0x2a, 0x2f, 0x89, 0x54, 0xc4, 0x17, 0x18, 0xb0, 0x31, 0x06, 0x2f, 0xee, 0x90, 0x10, 0x2f, 0x91,
	0x93, 0x38, 0xa9, 0xd5, 0xc4, 0x0e, 0xb6, 0x43, 0x34, 0x3e, 0x01, 0xef, 0x7c, 0x61, 0x64, 0x27,
	0x69, 0xd9, 0xfa, 0xb0, 0xa7, 0xf4, 0x9c, 0x7b, 0xcf, 0xf1, 0xfd, 0x57, 0x18, 0xa4, 0x95, 0x57,
	0x48, 0xa1, 0x05, 0xea, 0xa4, 0xd5, 0xf4, 0x6d, 0xca, 0xf4, 0xba, 0x0c, 0xbd, 0x48, 0xe4, 0x7e,
	0x28, 0x45, 0x44, 0x88, 0xf4, 0x33, 0x21, 0x89, 0xa2, 0xf2, 0x27, 0x95, 0x3e, 0x29, 0x98, 0x1f,
	0x89, 0x3c, 0x17, 0xbc, 0xf9, 0xd4, 0xd2, 0xe9, 0x8b, 0x54, 0x88, 0x34, 0xa3, 0xbe, 0x45, 0x61,
	0x99, 0xf8, 0x9a, 0xe5, 0x54, 0x69, 0x92, 0x17, 0x4d, 0xc2, 0xf3, 0xbb, 0x09, 0x71, 0x29, 0x89,
	0x66, 0xad, 0xc1, 0xfc, 0x77, 0x07, 0xc6, 0x5f, 0x8b, 0x8c, 0xf1, 0xcd, 0xf5, 0xb7, 0x4b, 0x9e,
	0x08, 0xf4, 0x14, 0x86, 0x89, 0xa4, 0x3f, 0x4a, 0xca, 0xa3, 0x1b, 0xd7, 0x99, 0x39, 0x8b, 0x23,
	0xbc, 0x23, 0xd0, 0x12, 0x20, 0x17, 0x71, 0x99, 0x59, 0x0b, 0xb7, 0x33, 0x73, 0x16, 0xff, 0x2f,
	0x91, 0xd7, 0x94, 0xf4, 0x65, 0x1b, 0xc1, 0xff, 0x64, 0xa1, 0x4f, 0x70, 0x62, 0x3a, 0x09, 0x76,
	0x54, 0xc0, 0x78, 0x22, 0xdc, 0x83, 0x99, 0xb3, 0x18, 0x2d, 0x1f, 0x7b, 0x69, 0xe5, 0x7d, 0x16,
	0x98, 0xec, 0xd4, 0xa6, 0x8e, 0x8f, 0xff, 0x61, 0x64, 0x54, 0xb7, 0x59, 0x74, 0x01, 0x0f, 0x13,
	0xb5, 0xd9, 0xb3, 0xea, 0x5a, 0xab, 0x47, 0xc6, 0xea, 0x7c, 0x75, 0xb5, 0xe7, 0x34, 0x49, 0xd4,
	0xe6, 0x36, 0x79, 0x36, 0x81, 0xe3, 0x3b, 0x26, 0xf3, 0x5f, 0x80, 0xf6, 0xeb, 0x30, 0xf3, 0x08,
	0x09, 0x8f, 0x2b, 0x16, 0xeb, 0x75, 0x3b, 0x8f, 0x2d, 0x81, 0x5e, 0xc3, 0x03, 0x55, 0x48, 0x4a,
	0x62, 0xc6, 0xd3, 0x20, 0x21, 0x91, 0x16, 0xd2, 0x4e, 0xe5, 0x08, 0x1f, 0x6f, 0xf9, 0x73, 0x4b,
	0xa3, 0x27, 0x30, 0x8c, 0x44, 0x4c, 0x03, 0x49, 0x34, 0xb5, 0xbd, 0x0f, 0xf1, 0xc0, 0x10, 0x98,
	0x68, 0x3a, 0xbf, 0x82, 0xc9, 0x5e, 0xe1, 0xf7, 0x3c, 0xed, 0x42, 0x3f, 0x64, 0xda, 0xba, 0xd5,
	0x2f, 0xb6, 0x70, 0xfe, 0xe7, 0xa0, 0xdd, 0x29, 0xae, 0x77, 0xfa, 0x0c, 0x20, 0x25, 0x9a, 0x56,
	0xe4, 0x26, 0x60, 0xb1, 0x75, 0x1a, 0xe3, 0x61, 0xc3, 0x5c, 0xc6, 0xc8, 0x83, 0xae, 0x39, 0x1b,
	0x6b, 0x33, 0x5a, 0x4e, 0xbd, 0xfa, 0x64, 0xbc, 0xf6, 0x64, 0xbc, 0xeb, 0xf6, 0xa6, 0xb0, 0xcd,
	0x33, 0x0b, 0x35, 0xdf, 0x40, 0x31, 0x1e, 0xd1, 0x20, 0x2d, 0x54, 0x40, 0x0b, 0x11, 0xad, 0x9b,
	0x85, 0x9e, 0xee, 0xe9, 0xdf, 0x37, 0x27, 0x87, 0x27, 0x46, 0xb6, 0x32, 0xaa, 0x8b, 0x42, 0x7d,
	0x30, 0x1a, 0x84, 0xa0, 0x2b, 0x95, 0x62, 0x76, 0x83, 0x3d, 0x6c, 0x7f, 0xa3, 0x53, 0x18, 0xd8,
	0x83, 0x51, 0x5c, 0xba, 0xbd, 0x99, 0xb3, 0x70, 0x70, 0xdf, 0xe0, 0x15, 0x97, 0xa6, 0xe9, 0x68,
	0x4d, 0x38, 0xa7, 0x99, 0x7b, 0x58, 0x37, 0xdd, 0x40, 0x23, 0x92, 0x49, 0x10, 0xad, 0x09, 0xe3,
	0x6e, 0xbf, 0x0e, 0xc9, 0xe4, 0x9d, 0x81, 0xe8, 0x04, 0x7a, 0xa1, 0x20, 0x32, 0x76, 0x07, 0x96,
	0xaf, 0x81, 0xb1, 0x22, 0x5c, 0x53, 0xce, 0x89, 0x3b, 0xac, 0xf3, 0x1b, 0x88, 0xa6, 0x30, 0x30,
	0x5b, 0xd0, 0x65, 0x4c, 0x5d, 0xb0, 0xef, 0x6f, 0xb1, 0xd9, 0x49, 0x26, 0x78, 0x5a, 0x07, 0x47,
	0x36, 0xb8, 0x23, 0x8c, 0x92, 0x64, 0x8d, 0x72, 0x5c, 0x2b, 0x5b, 0x7c, 0xf6, 0xea, 0xfb, 0xcb,
	0xfb, 0xff, 0xe3, 0x69, 0x15, 0x1e, 0xda, 0xb9, 0xbd, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x91,
	0x2f, 0xbd, 0xf8, 0x20, 0x04, 0x00, 0x00,
}
