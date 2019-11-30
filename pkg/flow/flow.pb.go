// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bio-routing/tflow3/pkg/flow/flow.proto

package flow

import (
	fmt "fmt"
	api "github.com/bio-routing/bio-rd/net/api"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IPAddress struct {
	Addr                 *api.IP           `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IPAddress) Reset()         { *m = IPAddress{} }
func (m *IPAddress) String() string { return proto.CompactTextString(m) }
func (*IPAddress) ProtoMessage()    {}
func (*IPAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{0}
}

func (m *IPAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddress.Unmarshal(m, b)
}
func (m *IPAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddress.Marshal(b, m, deterministic)
}
func (m *IPAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddress.Merge(m, src)
}
func (m *IPAddress) XXX_Size() int {
	return xxx_messageInfo_IPAddress.Size(m)
}
func (m *IPAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddress.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddress proto.InternalMessageInfo

func (m *IPAddress) GetAddr() *api.IP {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *IPAddress) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Prefix struct {
	Pfx                  *api.Prefix       `protobuf:"bytes,1,opt,name=pfx,proto3" json:"pfx,omitempty"`
	OriginAsn            uint32            `protobuf:"varint,2,opt,name=origin_asn,json=originAsn,proto3" json:"origin_asn,omitempty"`
	NextAsn              uint32            `protobuf:"varint,3,opt,name=next_asn,json=nextAsn,proto3" json:"next_asn,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Prefix) Reset()         { *m = Prefix{} }
func (m *Prefix) String() string { return proto.CompactTextString(m) }
func (*Prefix) ProtoMessage()    {}
func (*Prefix) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{1}
}

func (m *Prefix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prefix.Unmarshal(m, b)
}
func (m *Prefix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prefix.Marshal(b, m, deterministic)
}
func (m *Prefix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prefix.Merge(m, src)
}
func (m *Prefix) XXX_Size() int {
	return xxx_messageInfo_Prefix.Size(m)
}
func (m *Prefix) XXX_DiscardUnknown() {
	xxx_messageInfo_Prefix.DiscardUnknown(m)
}

var xxx_messageInfo_Prefix proto.InternalMessageInfo

func (m *Prefix) GetPfx() *api.Prefix {
	if m != nil {
		return m.Pfx
	}
	return nil
}

func (m *Prefix) GetOriginAsn() uint32 {
	if m != nil {
		return m.OriginAsn
	}
	return 0
}

func (m *Prefix) GetNextAsn() uint32 {
	if m != nil {
		return m.NextAsn
	}
	return 0
}

func (m *Prefix) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ASN struct {
	Asn                  uint32            `protobuf:"varint,1,opt,name=asn,proto3" json:"asn,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ASN) Reset()         { *m = ASN{} }
func (m *ASN) String() string { return proto.CompactTextString(m) }
func (*ASN) ProtoMessage()    {}
func (*ASN) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{2}
}

func (m *ASN) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ASN.Unmarshal(m, b)
}
func (m *ASN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ASN.Marshal(b, m, deterministic)
}
func (m *ASN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ASN.Merge(m, src)
}
func (m *ASN) XXX_Size() int {
	return xxx_messageInfo_ASN.Size(m)
}
func (m *ASN) XXX_DiscardUnknown() {
	xxx_messageInfo_ASN.DiscardUnknown(m)
}

var xxx_messageInfo_ASN proto.InternalMessageInfo

func (m *ASN) GetAsn() uint32 {
	if m != nil {
		return m.Asn
	}
	return 0
}

func (m *ASN) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type VRF struct {
	VrfId                uint64            `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *VRF) Reset()         { *m = VRF{} }
func (m *VRF) String() string { return proto.CompactTextString(m) }
func (*VRF) ProtoMessage()    {}
func (*VRF) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{3}
}

func (m *VRF) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VRF.Unmarshal(m, b)
}
func (m *VRF) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VRF.Marshal(b, m, deterministic)
}
func (m *VRF) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VRF.Merge(m, src)
}
func (m *VRF) XXX_Size() int {
	return xxx_messageInfo_VRF.Size(m)
}
func (m *VRF) XXX_DiscardUnknown() {
	xxx_messageInfo_VRF.DiscardUnknown(m)
}

var xxx_messageInfo_VRF proto.InternalMessageInfo

func (m *VRF) GetVrfId() uint64 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *VRF) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Port struct {
	Port                 uint32            `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{4}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Port) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Protocol struct {
	Protocol             uint32            `protobuf:"varint,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Protocol) Reset()         { *m = Protocol{} }
func (m *Protocol) String() string { return proto.CompactTextString(m) }
func (*Protocol) ProtoMessage()    {}
func (*Protocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{5}
}

func (m *Protocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Protocol.Unmarshal(m, b)
}
func (m *Protocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Protocol.Marshal(b, m, deterministic)
}
func (m *Protocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Protocol.Merge(m, src)
}
func (m *Protocol) XXX_Size() int {
	return xxx_messageInfo_Protocol.Size(m)
}
func (m *Protocol) XXX_DiscardUnknown() {
	xxx_messageInfo_Protocol.DiscardUnknown(m)
}

var xxx_messageInfo_Protocol proto.InternalMessageInfo

func (m *Protocol) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *Protocol) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Interface struct {
	Id                   uint32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Interface) Reset()         { *m = Interface{} }
func (m *Interface) String() string { return proto.CompactTextString(m) }
func (*Interface) ProtoMessage()    {}
func (*Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{6}
}

func (m *Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Interface.Unmarshal(m, b)
}
func (m *Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Interface.Marshal(b, m, deterministic)
}
func (m *Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interface.Merge(m, src)
}
func (m *Interface) XXX_Size() int {
	return xxx_messageInfo_Interface.Size(m)
}
func (m *Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_Interface proto.InternalMessageInfo

func (m *Interface) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Interface) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Flow struct {
	Agent                *IPAddress `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
	Vrf                  *VRF       `protobuf:"bytes,2,opt,name=vrf,proto3" json:"vrf,omitempty"`
	IntIn                *Interface `protobuf:"bytes,3,opt,name=int_in,json=intIn,proto3" json:"int_in,omitempty"`
	IntOut               *Interface `protobuf:"bytes,4,opt,name=int_out,json=intOut,proto3" json:"int_out,omitempty"`
	IpFlow               *IPFlow    `protobuf:"bytes,5,opt,name=ip_flow,json=ipFlow,proto3" json:"ip_flow,omitempty"`
	NextHopIp            *IPAddress `protobuf:"bytes,7,opt,name=next_hop_ip,json=nextHopIp,proto3" json:"next_hop_ip,omitempty"`
	SrcPrefix            *Prefix    `protobuf:"bytes,8,opt,name=src_prefix,json=srcPrefix,proto3" json:"src_prefix,omitempty"`
	DstPrefix            *Prefix    `protobuf:"bytes,9,opt,name=dst_prefix,json=dstPrefix,proto3" json:"dst_prefix,omitempty"`
	PacketsCount         uint64     `protobuf:"varint,13,opt,name=packets_count,json=packetsCount,proto3" json:"packets_count,omitempty"`
	BytesCount           uint64     `protobuf:"varint,14,opt,name=bytes_count,json=bytesCount,proto3" json:"bytes_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Flow) Reset()         { *m = Flow{} }
func (m *Flow) String() string { return proto.CompactTextString(m) }
func (*Flow) ProtoMessage()    {}
func (*Flow) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{7}
}

func (m *Flow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Flow.Unmarshal(m, b)
}
func (m *Flow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Flow.Marshal(b, m, deterministic)
}
func (m *Flow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flow.Merge(m, src)
}
func (m *Flow) XXX_Size() int {
	return xxx_messageInfo_Flow.Size(m)
}
func (m *Flow) XXX_DiscardUnknown() {
	xxx_messageInfo_Flow.DiscardUnknown(m)
}

var xxx_messageInfo_Flow proto.InternalMessageInfo

func (m *Flow) GetAgent() *IPAddress {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *Flow) GetVrf() *VRF {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *Flow) GetIntIn() *Interface {
	if m != nil {
		return m.IntIn
	}
	return nil
}

func (m *Flow) GetIntOut() *Interface {
	if m != nil {
		return m.IntOut
	}
	return nil
}

func (m *Flow) GetIpFlow() *IPFlow {
	if m != nil {
		return m.IpFlow
	}
	return nil
}

func (m *Flow) GetNextHopIp() *IPAddress {
	if m != nil {
		return m.NextHopIp
	}
	return nil
}

func (m *Flow) GetSrcPrefix() *Prefix {
	if m != nil {
		return m.SrcPrefix
	}
	return nil
}

func (m *Flow) GetDstPrefix() *Prefix {
	if m != nil {
		return m.DstPrefix
	}
	return nil
}

func (m *Flow) GetPacketsCount() uint64 {
	if m != nil {
		return m.PacketsCount
	}
	return 0
}

func (m *Flow) GetBytesCount() uint64 {
	if m != nil {
		return m.BytesCount
	}
	return 0
}

type IPFlow struct {
	SrcIp                *IPAddress `protobuf:"bytes,5,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	DstIp                *IPAddress `protobuf:"bytes,6,opt,name=dst_ip,json=dstIp,proto3" json:"dst_ip,omitempty"`
	Protocol             *Protocol  `protobuf:"bytes,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SrcPort              *Port      `protobuf:"bytes,11,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstPort              *Port      `protobuf:"bytes,12,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IPFlow) Reset()         { *m = IPFlow{} }
func (m *IPFlow) String() string { return proto.CompactTextString(m) }
func (*IPFlow) ProtoMessage()    {}
func (*IPFlow) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{8}
}

func (m *IPFlow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPFlow.Unmarshal(m, b)
}
func (m *IPFlow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPFlow.Marshal(b, m, deterministic)
}
func (m *IPFlow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPFlow.Merge(m, src)
}
func (m *IPFlow) XXX_Size() int {
	return xxx_messageInfo_IPFlow.Size(m)
}
func (m *IPFlow) XXX_DiscardUnknown() {
	xxx_messageInfo_IPFlow.DiscardUnknown(m)
}

var xxx_messageInfo_IPFlow proto.InternalMessageInfo

func (m *IPFlow) GetSrcIp() *IPAddress {
	if m != nil {
		return m.SrcIp
	}
	return nil
}

func (m *IPFlow) GetDstIp() *IPAddress {
	if m != nil {
		return m.DstIp
	}
	return nil
}

func (m *IPFlow) GetProtocol() *Protocol {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (m *IPFlow) GetSrcPort() *Port {
	if m != nil {
		return m.SrcPort
	}
	return nil
}

func (m *IPFlow) GetDstPort() *Port {
	if m != nil {
		return m.DstPort
	}
	return nil
}

type FlowRecord struct {
	Agent                uint32   `protobuf:"varint,1,opt,name=agent,proto3" json:"agent,omitempty"`
	Vrf                  uint32   `protobuf:"varint,2,opt,name=vrf,proto3" json:"vrf,omitempty"`
	IntIn                uint32   `protobuf:"varint,3,opt,name=int_in,json=intIn,proto3" json:"int_in,omitempty"`
	IntOut               uint32   `protobuf:"varint,4,opt,name=int_out,json=intOut,proto3" json:"int_out,omitempty"`
	SrcIp                uint32   `protobuf:"varint,5,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	DstIp                uint32   `protobuf:"varint,6,opt,name=dst_ip,json=dstIp,proto3" json:"dst_ip,omitempty"`
	NextHopIp            uint32   `protobuf:"varint,7,opt,name=next_hop_ip,json=nextHopIp,proto3" json:"next_hop_ip,omitempty"`
	SrcPrefix            uint32   `protobuf:"varint,8,opt,name=src_prefix,json=srcPrefix,proto3" json:"src_prefix,omitempty"`
	DstPrefix            uint32   `protobuf:"varint,9,opt,name=dst_prefix,json=dstPrefix,proto3" json:"dst_prefix,omitempty"`
	Protocol             uint32   `protobuf:"varint,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SrcPort              uint32   `protobuf:"varint,11,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstPort              uint32   `protobuf:"varint,12,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	PacketsCount         uint64   `protobuf:"varint,13,opt,name=packets_count,json=packetsCount,proto3" json:"packets_count,omitempty"`
	BytesCount           uint64   `protobuf:"varint,14,opt,name=bytes_count,json=bytesCount,proto3" json:"bytes_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowRecord) Reset()         { *m = FlowRecord{} }
func (m *FlowRecord) String() string { return proto.CompactTextString(m) }
func (*FlowRecord) ProtoMessage()    {}
func (*FlowRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{9}
}

func (m *FlowRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRecord.Unmarshal(m, b)
}
func (m *FlowRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRecord.Marshal(b, m, deterministic)
}
func (m *FlowRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRecord.Merge(m, src)
}
func (m *FlowRecord) XXX_Size() int {
	return xxx_messageInfo_FlowRecord.Size(m)
}
func (m *FlowRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRecord.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRecord proto.InternalMessageInfo

func (m *FlowRecord) GetAgent() uint32 {
	if m != nil {
		return m.Agent
	}
	return 0
}

func (m *FlowRecord) GetVrf() uint32 {
	if m != nil {
		return m.Vrf
	}
	return 0
}

func (m *FlowRecord) GetIntIn() uint32 {
	if m != nil {
		return m.IntIn
	}
	return 0
}

func (m *FlowRecord) GetIntOut() uint32 {
	if m != nil {
		return m.IntOut
	}
	return 0
}

func (m *FlowRecord) GetSrcIp() uint32 {
	if m != nil {
		return m.SrcIp
	}
	return 0
}

func (m *FlowRecord) GetDstIp() uint32 {
	if m != nil {
		return m.DstIp
	}
	return 0
}

func (m *FlowRecord) GetNextHopIp() uint32 {
	if m != nil {
		return m.NextHopIp
	}
	return 0
}

func (m *FlowRecord) GetSrcPrefix() uint32 {
	if m != nil {
		return m.SrcPrefix
	}
	return 0
}

func (m *FlowRecord) GetDstPrefix() uint32 {
	if m != nil {
		return m.DstPrefix
	}
	return 0
}

func (m *FlowRecord) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *FlowRecord) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *FlowRecord) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *FlowRecord) GetPacketsCount() uint64 {
	if m != nil {
		return m.PacketsCount
	}
	return 0
}

func (m *FlowRecord) GetBytesCount() uint64 {
	if m != nil {
		return m.BytesCount
	}
	return 0
}

type FlowFile struct {
	IpAddresses          []*IPAddress  `protobuf:"bytes,1,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
	Prefixes             []*Prefix     `protobuf:"bytes,2,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	Asns                 []*ASN        `protobuf:"bytes,3,rep,name=asns,proto3" json:"asns,omitempty"`
	Vrfs                 []*VRF        `protobuf:"bytes,4,rep,name=vrfs,proto3" json:"vrfs,omitempty"`
	Ports                []*Port       `protobuf:"bytes,5,rep,name=ports,proto3" json:"ports,omitempty"`
	Protocols            []*Protocol   `protobuf:"bytes,6,rep,name=protocols,proto3" json:"protocols,omitempty"`
	Interfaces           []*Interface  `protobuf:"bytes,7,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	FlowRecords          []*FlowRecord `protobuf:"bytes,8,rep,name=flow_records,json=flowRecords,proto3" json:"flow_records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FlowFile) Reset()         { *m = FlowFile{} }
func (m *FlowFile) String() string { return proto.CompactTextString(m) }
func (*FlowFile) ProtoMessage()    {}
func (*FlowFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_827e3cf7c8bdc54e, []int{10}
}

func (m *FlowFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowFile.Unmarshal(m, b)
}
func (m *FlowFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowFile.Marshal(b, m, deterministic)
}
func (m *FlowFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowFile.Merge(m, src)
}
func (m *FlowFile) XXX_Size() int {
	return xxx_messageInfo_FlowFile.Size(m)
}
func (m *FlowFile) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowFile.DiscardUnknown(m)
}

var xxx_messageInfo_FlowFile proto.InternalMessageInfo

func (m *FlowFile) GetIpAddresses() []*IPAddress {
	if m != nil {
		return m.IpAddresses
	}
	return nil
}

func (m *FlowFile) GetPrefixes() []*Prefix {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

func (m *FlowFile) GetAsns() []*ASN {
	if m != nil {
		return m.Asns
	}
	return nil
}

func (m *FlowFile) GetVrfs() []*VRF {
	if m != nil {
		return m.Vrfs
	}
	return nil
}

func (m *FlowFile) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *FlowFile) GetProtocols() []*Protocol {
	if m != nil {
		return m.Protocols
	}
	return nil
}

func (m *FlowFile) GetInterfaces() []*Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *FlowFile) GetFlowRecords() []*FlowRecord {
	if m != nil {
		return m.FlowRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*IPAddress)(nil), "bio.tflow3.IPAddress")
	proto.RegisterMapType((map[string]string)(nil), "bio.tflow3.IPAddress.MetadataEntry")
	proto.RegisterType((*Prefix)(nil), "bio.tflow3.Prefix")
	proto.RegisterMapType((map[string]string)(nil), "bio.tflow3.Prefix.MetadataEntry")
	proto.RegisterType((*ASN)(nil), "bio.tflow3.ASN")
	proto.RegisterMapType((map[string]string)(nil), "bio.tflow3.ASN.MetadataEntry")
	proto.RegisterType((*VRF)(nil), "bio.tflow3.VRF")
	proto.RegisterMapType((map[string]string)(nil), "bio.tflow3.VRF.MetadataEntry")
	proto.RegisterType((*Port)(nil), "bio.tflow3.Port")
	proto.RegisterMapType((map[string]string)(nil), "bio.tflow3.Port.MetadataEntry")
	proto.RegisterType((*Protocol)(nil), "bio.tflow3.Protocol")
	proto.RegisterMapType((map[string]string)(nil), "bio.tflow3.Protocol.MetadataEntry")
	proto.RegisterType((*Interface)(nil), "bio.tflow3.Interface")
	proto.RegisterMapType((map[string]string)(nil), "bio.tflow3.Interface.MetadataEntry")
	proto.RegisterType((*Flow)(nil), "bio.tflow3.Flow")
	proto.RegisterType((*IPFlow)(nil), "bio.tflow3.IPFlow")
	proto.RegisterType((*FlowRecord)(nil), "bio.tflow3.FlowRecord")
	proto.RegisterType((*FlowFile)(nil), "bio.tflow3.FlowFile")
}

func init() {
	proto.RegisterFile("github.com/bio-routing/tflow3/pkg/flow/flow.proto", fileDescriptor_827e3cf7c8bdc54e)
}

var fileDescriptor_827e3cf7c8bdc54e = []byte{
	// 959 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0x66, 0xfd, 0x7b, 0x9c, 0x6d, 0xa3, 0x51, 0x0d, 0x9b, 0x48, 0x69, 0x83, 0x23, 0x55,
	0x91, 0x02, 0x36, 0x49, 0x55, 0xa9, 0x2d, 0x08, 0x64, 0x10, 0x11, 0xbe, 0xa0, 0x58, 0x5b, 0xa9,
	0x17, 0xdc, 0x58, 0x6b, 0xef, 0xae, 0x3b, 0x8a, 0xb3, 0x33, 0x9a, 0x19, 0xbb, 0xc9, 0x1b, 0x70,
	0x81, 0xc4, 0x05, 0xbc, 0x01, 0x77, 0x48, 0xbc, 0x05, 0x0f, 0xc2, 0x23, 0xf0, 0x08, 0xe8, 0x9c,
	0xfd, 0xf1, 0xae, 0x7f, 0x02, 0x17, 0xf8, 0x26, 0x99, 0x9d, 0xf3, 0x9d, 0xd9, 0xef, 0x7c, 0xf3,
	0x9d, 0xe3, 0x85, 0x8b, 0x29, 0x37, 0xef, 0xe6, 0xe3, 0xee, 0x44, 0xdc, 0xf4, 0xc6, 0x5c, 0x7c,
	0xa2, 0xc4, 0xdc, 0xf0, 0x78, 0xda, 0x33, 0xd1, 0x4c, 0xbc, 0x7f, 0xd6, 0x93, 0xd7, 0xd3, 0x1e,
	0xae, 0xe8, 0x4f, 0x57, 0x2a, 0x61, 0x04, 0x83, 0x31, 0x17, 0xdd, 0x24, 0x7e, 0xd4, 0xdb, 0x92,
	0x4e, 0xeb, 0xa0, 0x17, 0x87, 0xa6, 0xe7, 0x4b, 0x8e, 0xff, 0x93, 0xe4, 0xce, 0xef, 0x16, 0x34,
	0x07, 0xc3, 0x7e, 0x10, 0xa8, 0x50, 0x6b, 0xf6, 0x04, 0x2a, 0x7e, 0x10, 0x28, 0xd7, 0x3a, 0xb1,
	0xce, 0x5a, 0x97, 0xad, 0x2e, 0x9e, 0x8c, 0xd8, 0xc1, 0xd0, 0xa3, 0x00, 0xfb, 0x12, 0x1a, 0x37,
	0xa1, 0xf1, 0x03, 0xdf, 0xf8, 0xae, 0x7d, 0x62, 0x9f, 0xb5, 0x2e, 0x4f, 0xbb, 0xcb, 0xd7, 0x77,
	0xf3, 0x93, 0xba, 0xdf, 0xa5, 0xa8, 0x6f, 0x62, 0xa3, 0xee, 0xbc, 0x3c, 0xe9, 0xe8, 0x33, 0x70,
	0x4a, 0x21, 0x76, 0x00, 0xf6, 0x75, 0x78, 0x47, 0x6f, 0x6c, 0x7a, 0xb8, 0x64, 0x8f, 0xa0, 0xba,
	0xf0, 0x67, 0xf3, 0xd0, 0xdd, 0xa3, 0xbd, 0xe4, 0xe1, 0xd5, 0xde, 0x0b, 0xab, 0xf3, 0x97, 0x05,
	0xb5, 0xa1, 0x0a, 0x23, 0x7e, 0xcb, 0x3e, 0x02, 0x5b, 0x46, 0xb7, 0x29, 0xd1, 0x87, 0x39, 0xd1,
	0x24, 0xea, 0x61, 0x8c, 0x1d, 0x03, 0x08, 0xc5, 0xa7, 0x3c, 0x1e, 0xf9, 0x3a, 0xa6, 0xc3, 0x1c,
	0xaf, 0x99, 0xec, 0xf4, 0x75, 0xcc, 0x0e, 0xa1, 0x11, 0x87, 0xb7, 0x86, 0x82, 0x36, 0x05, 0xeb,
	0xf8, 0x8c, 0xa1, 0xcf, 0x0b, 0x55, 0x56, 0xa8, 0xca, 0x93, 0x62, 0x95, 0xc9, 0x4b, 0x76, 0x53,
	0xe2, 0xcf, 0x16, 0xd8, 0xfd, 0x37, 0xaf, 0x31, 0x07, 0x89, 0x59, 0x44, 0x0c, 0x97, 0xec, 0x65,
	0x81, 0xd4, 0x1e, 0x91, 0x3a, 0x2e, 0x92, 0xea, 0xbf, 0x79, 0xbd, 0x1b, 0x46, 0xbf, 0x5a, 0x60,
	0xbf, 0xf5, 0xae, 0x58, 0x1b, 0x6a, 0x0b, 0x15, 0x8d, 0x78, 0x40, 0x69, 0x15, 0xaf, 0xba, 0x50,
	0xd1, 0x20, 0xf8, 0x37, 0x5a, 0x6f, 0xbd, 0xab, 0xdd, 0xd0, 0xfa, 0xc5, 0x82, 0xca, 0x50, 0x28,
	0xc3, 0x18, 0x54, 0xa4, 0x50, 0x26, 0x95, 0x8a, 0xd6, 0xec, 0xd5, 0x1a, 0xa9, 0xc7, 0xa5, 0x0b,
	0x14, 0xca, 0xec, 0x86, 0xd5, 0x6f, 0x16, 0x34, 0x86, 0xd8, 0x58, 0x13, 0x31, 0x63, 0x47, 0xd0,
	0x90, 0xe9, 0x3a, 0x65, 0x97, 0x3f, 0xb3, 0x2f, 0xd6, 0x18, 0x76, 0xca, 0x16, 0x4b, 0x70, 0xbb,
	0x61, 0xf9, 0x07, 0x36, 0x7d, 0x6c, 0x42, 0x15, 0xf9, 0x93, 0x90, 0x3d, 0x80, 0xbd, 0xf4, 0x52,
	0x1d, 0x6f, 0x8f, 0x07, 0x28, 0x68, 0xec, 0xdf, 0x64, 0x69, 0xb4, 0x2e, 0xf5, 0x7d, 0x65, 0x43,
	0xdf, 0x67, 0x87, 0xed, 0x86, 0xef, 0x9f, 0x36, 0x54, 0xae, 0x66, 0xe2, 0x3d, 0x3b, 0x87, 0xaa,
	0x3f, 0x0d, 0x63, 0x93, 0xf6, 0x7d, 0x7b, 0xe3, 0xec, 0xf1, 0x12, 0x0c, 0x8e, 0x88, 0x85, 0x8a,
	0xe8, 0xb4, 0x6c, 0x44, 0x2c, 0x4d, 0xe9, 0x61, 0x8c, 0x7d, 0x0c, 0x35, 0x1e, 0x9b, 0x11, 0x4f,
	0x26, 0xc0, 0xea, 0x81, 0x59, 0x51, 0x5e, 0x95, 0xc7, 0x66, 0x10, 0xb3, 0x2e, 0xd4, 0x11, 0x2d,
	0xe6, 0xc6, 0xad, 0xdc, 0x07, 0xc7, 0x33, 0xbf, 0x9f, 0x1b, 0x76, 0x0e, 0x75, 0x2e, 0x47, 0x18,
	0x75, 0xab, 0x84, 0x67, 0x65, 0xbe, 0x58, 0x92, 0x57, 0xe3, 0x92, 0x4a, 0x7b, 0x0e, 0x2d, 0x1a,
	0x47, 0xef, 0x84, 0x1c, 0x71, 0xe9, 0xd6, 0xef, 0x2b, 0xb0, 0x89, 0xc8, 0x6f, 0x85, 0x1c, 0x48,
	0x76, 0x01, 0xa0, 0xd5, 0x64, 0x24, 0x69, 0x24, 0xb9, 0x8d, 0xf5, 0xd7, 0xa4, 0x13, 0xb1, 0xa9,
	0xd5, 0x24, 0x1d, 0x9d, 0x17, 0x00, 0x81, 0x36, 0x59, 0x4a, 0x73, 0x7b, 0x4a, 0xa0, 0x4d, 0x9a,
	0x72, 0x0a, 0x8e, 0xf4, 0x27, 0xd7, 0xa1, 0xd1, 0xa3, 0x89, 0x98, 0xc7, 0xc6, 0x75, 0x68, 0x04,
	0xec, 0xa7, 0x9b, 0x5f, 0xe3, 0x1e, 0x7b, 0x02, 0xad, 0xf1, 0x9d, 0x09, 0x33, 0xc8, 0x03, 0x82,
	0x00, 0x6d, 0x11, 0xa0, 0xf3, 0xb7, 0x05, 0xb5, 0xa4, 0x6a, 0x14, 0x1e, 0x69, 0x73, 0x99, 0x2a,
	0xb3, 0xed, 0x26, 0xb5, 0x9a, 0x0c, 0x24, 0xa2, 0x91, 0x31, 0x97, 0x6e, 0xed, 0x5e, 0x74, 0xa0,
	0xcd, 0x40, 0xb2, 0x4f, 0x0b, 0x6d, 0x07, 0x84, 0x7f, 0xb4, 0xa9, 0xb5, 0x0a, 0xcd, 0x78, 0x0e,
	0x0d, 0x12, 0x11, 0xc7, 0x48, 0x8b, 0x32, 0x0e, 0x56, 0xc7, 0x85, 0x57, 0x47, 0x01, 0x71, 0xb6,
	0x9c, 0x43, 0x83, 0xe4, 0x43, 0xf0, 0xfe, 0x36, 0x30, 0x4a, 0x27, 0x94, 0xe9, 0xfc, 0x68, 0x03,
	0xd0, 0x35, 0x87, 0x13, 0xa1, 0x02, 0xb4, 0xf8, 0xd2, 0xbf, 0x4e, 0x66, 0xd4, 0x83, 0xa5, 0x51,
	0x9d, 0xc4, 0x97, 0xed, 0x92, 0x2f, 0x9d, 0xcc, 0x80, 0x1f, 0x96, 0x0d, 0xe8, 0xe4, 0x4e, 0x6b,
	0x97, 0xe4, 0x74, 0x32, 0xdd, 0xda, 0x25, 0xdd, 0x9c, 0x4c, 0xa0, 0xc7, 0xeb, 0x56, 0x73, 0x8a,
	0x9e, 0x3a, 0x5e, 0xf3, 0x94, 0x53, 0xf4, 0xcf, 0xf1, 0x9a, 0x7f, 0x9c, 0xa2, 0x57, 0x8e, 0x56,
	0xe4, 0x2f, 0x4e, 0xbd, 0xc3, 0x15, 0xa1, 0x9d, 0xa5, 0xac, 0x87, 0x2b, 0xb2, 0x3a, 0xb9, 0x88,
	0xff, 0x93, 0xfb, 0x7e, 0xb2, 0xa1, 0x81, 0x57, 0x71, 0xc5, 0x67, 0x21, 0x7b, 0x01, 0xfb, 0x5c,
	0x8e, 0xfc, 0xc4, 0x38, 0xa1, 0x76, 0x2d, 0x9a, 0x69, 0x5b, 0x7c, 0xd5, 0xe2, 0xb2, 0x9f, 0x21,
	0x59, 0x17, 0xcb, 0xc3, 0x42, 0x43, 0x9d, 0x0e, 0xee, 0x4d, 0xbd, 0x93, 0x63, 0xd8, 0x29, 0x54,
	0x7c, 0x1d, 0xeb, 0xf4, 0x6b, 0xe9, 0xe1, 0xca, 0x4f, 0xb6, 0x47, 0x41, 0x04, 0x2d, 0x54, 0xa4,
	0xd3, 0xd1, 0xba, 0x36, 0xab, 0x28, 0xc8, 0x9e, 0x42, 0x15, 0xd5, 0xd1, 0x6e, 0x95, 0x50, 0xeb,
	0xae, 0x4b, 0xc2, 0xec, 0x12, 0x9a, 0x99, 0xe0, 0xda, 0xad, 0x11, 0x76, 0x73, 0x03, 0x2c, 0x61,
	0xec, 0x39, 0x00, 0xcf, 0xe6, 0x97, 0x76, 0xeb, 0x1b, 0xd4, 0xc8, 0xa7, 0x5b, 0x01, 0xc8, 0x5e,
	0xc2, 0x3e, 0x86, 0x47, 0x8a, 0xec, 0xad, 0xdd, 0x06, 0x25, 0x7e, 0x50, 0x4c, 0x5c, 0xba, 0xdf,
	0x6b, 0x45, 0xf9, 0x5a, 0x7f, 0x75, 0xf6, 0xc3, 0xd3, 0xff, 0xf6, 0xa9, 0x3b, 0xae, 0x11, 0xcd,
	0x67, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xa5, 0x38, 0x20, 0x1b, 0x0b, 0x00, 0x00,
}
