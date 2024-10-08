// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package service_manage // import "github.com/polarismesh/specification/source/go/api/v1/service_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import model "github.com/polarismesh/specification/source/go/api/v1/model"
import wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Client_ClientType int32

const (
	Client_UNKNOWN Client_ClientType = 0
	Client_SDK     Client_ClientType = 1
	Client_AGENT   Client_ClientType = 2
)

var Client_ClientType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SDK",
	2: "AGENT",
}
var Client_ClientType_value = map[string]int32{
	"UNKNOWN": 0,
	"SDK":     1,
	"AGENT":   2,
}

func (x Client_ClientType) String() string {
	return proto.EnumName(Client_ClientType_name, int32(x))
}
func (Client_ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_8c42d25563430f24, []int{0, 0}
}

type Client struct {
	Host                 *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Type                 Client_ClientType       `protobuf:"varint,2,opt,name=type,proto3,enum=v1.Client_ClientType" json:"type,omitempty"`
	Version              *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Location             *model.Location         `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Id                   *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	Stat                 []*StatInfo             `protobuf:"bytes,6,rep,name=stat,proto3" json:"stat,omitempty"`
	Ctime                *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_8c42d25563430f24, []int{0}
}
func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (dst *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(dst, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetHost() *wrapperspb.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Client) GetType() Client_ClientType {
	if m != nil {
		return m.Type
	}
	return Client_UNKNOWN
}

func (m *Client) GetVersion() *wrapperspb.StringValue {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Client) GetLocation() *model.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Client) GetId() *wrapperspb.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Client) GetStat() []*StatInfo {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *Client) GetCtime() *wrapperspb.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Client) GetMtime() *wrapperspb.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type StatInfo struct {
	Target               *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Port                 *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Path                 *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Protocol             *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StatInfo) Reset()         { *m = StatInfo{} }
func (m *StatInfo) String() string { return proto.CompactTextString(m) }
func (*StatInfo) ProtoMessage()    {}
func (*StatInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_8c42d25563430f24, []int{1}
}
func (m *StatInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatInfo.Unmarshal(m, b)
}
func (m *StatInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatInfo.Marshal(b, m, deterministic)
}
func (dst *StatInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatInfo.Merge(dst, src)
}
func (m *StatInfo) XXX_Size() int {
	return xxx_messageInfo_StatInfo.Size(m)
}
func (m *StatInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StatInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StatInfo proto.InternalMessageInfo

func (m *StatInfo) GetTarget() *wrapperspb.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *StatInfo) GetPort() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *StatInfo) GetPath() *wrapperspb.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *StatInfo) GetProtocol() *wrapperspb.StringValue {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func init() {
	proto.RegisterType((*Client)(nil), "v1.Client")
	proto.RegisterType((*StatInfo)(nil), "v1.StatInfo")
	proto.RegisterEnum("v1.Client_ClientType", Client_ClientType_name, Client_ClientType_value)
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_client_8c42d25563430f24) }

var fileDescriptor_client_8c42d25563430f24 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0xfa, 0xd7, 0xd3, 0x45, 0xca, 0x80, 0x10, 0x16, 0x91, 0xd2, 0xab, 0x0a, 0x3a,
	0x63, 0xbb, 0xa2, 0xde, 0xba, 0xae, 0xc8, 0xb2, 0x52, 0xa5, 0xdd, 0x55, 0xf0, 0x46, 0xa6, 0xd3,
	0xd3, 0x74, 0x20, 0xc9, 0x0c, 0x33, 0xa7, 0x91, 0x7d, 0x07, 0x9f, 0xc5, 0x47, 0xf2, 0x59, 0x24,
	0x93, 0x74, 0x45, 0x10, 0xc9, 0x55, 0xc8, 0x9c, 0xdf, 0xf7, 0xcb, 0x70, 0xbe, 0xc0, 0x89, 0xca,
	0x34, 0x16, 0xc4, 0xad, 0x33, 0x64, 0x58, 0x5c, 0xce, 0x4f, 0x1f, 0xa7, 0xc6, 0xa4, 0x19, 0x8a,
	0x70, 0xb2, 0x39, 0xec, 0xc4, 0x77, 0x27, 0xad, 0x45, 0xe7, 0x6b, 0xe6, 0x74, 0x94, 0x9b, 0x2d,
	0x66, 0xf5, 0xcb, 0xf4, 0x67, 0x07, 0xfa, 0x6f, 0x83, 0x81, 0x3d, 0x87, 0xee, 0xde, 0x78, 0x4a,
	0xa2, 0x49, 0x34, 0x1b, 0x2d, 0x1e, 0xf1, 0x5a, 0xc3, 0x8f, 0x1a, 0xbe, 0x26, 0xa7, 0x8b, 0xf4,
	0xb3, 0xcc, 0x0e, 0xb8, 0x0a, 0x24, 0x7b, 0x02, 0x5d, 0xba, 0xb5, 0x98, 0xc4, 0x93, 0x68, 0xf6,
	0x60, 0xf1, 0x90, 0x97, 0x73, 0x5e, 0xbb, 0x9a, 0xc7, 0xf5, 0xad, 0xc5, 0x55, 0x40, 0xd8, 0x4b,
	0x18, 0x94, 0xe8, 0xbc, 0x36, 0x45, 0xd2, 0x69, 0xe1, 0x3f, 0xc2, 0x6c, 0x06, 0xc3, 0xcc, 0x28,
	0x49, 0x55, 0xb0, 0x1b, 0x82, 0x27, 0xd5, 0x67, 0x3e, 0x34, 0x67, 0xab, 0xbb, 0x29, 0x7b, 0x0a,
	0xb1, 0xde, 0x26, 0xbd, 0x16, 0xf2, 0x58, 0x6f, 0xd9, 0x04, 0xba, 0x9e, 0x24, 0x25, 0xfd, 0x49,
	0xe7, 0xe8, 0x5c, 0x93, 0xa4, 0xcb, 0x62, 0x67, 0x56, 0x61, 0xc2, 0x16, 0xd0, 0x53, 0xa4, 0x73,
	0x4c, 0x06, 0x2d, 0x94, 0x35, 0x5a, 0x65, 0xf2, 0x90, 0x19, 0xb6, 0xc9, 0x04, 0x74, 0xfa, 0x0c,
	0xe0, 0xcf, 0xb6, 0xd8, 0x08, 0x06, 0x37, 0xcb, 0xab, 0xe5, 0xc7, 0x2f, 0xcb, 0xf1, 0x3d, 0x36,
	0x80, 0xce, 0xfa, 0xe2, 0x6a, 0x1c, 0xb1, 0xfb, 0xd0, 0x7b, 0xf3, 0xfe, 0xdd, 0xf2, 0x7a, 0x1c,
	0x4f, 0x7f, 0x45, 0x30, 0x3c, 0xde, 0x94, 0xbd, 0x80, 0x3e, 0x49, 0x97, 0x62, 0xbb, 0xd2, 0x1a,
	0xb6, 0x2a, 0xda, 0x1a, 0x47, 0xa1, 0xb6, 0x7f, 0x65, 0x6e, 0x2e, 0x0b, 0x3a, 0x5b, 0x34, 0x45,
	0x57, 0x64, 0x48, 0x48, 0xda, 0xb7, 0xaa, 0x2e, 0x90, 0xec, 0x35, 0x0c, 0xc3, 0x54, 0x99, 0xac,
	0xe9, 0xed, 0xff, 0xa9, 0x3b, 0xfa, 0xfc, 0x47, 0x04, 0xaf, 0x94, 0xc9, 0x39, 0x61, 0xa1, 0xc2,
	0x8f, 0x6d, 0x32, 0xe9, 0xb4, 0xe7, 0xde, 0xa2, 0xd2, 0x3b, 0x5d, 0xb7, 0xcd, 0xa5, 0xd5, 0x55,
	0x67, 0x1e, 0x5d, 0xa9, 0x15, 0xf2, 0x5c, 0x16, 0x32, 0xc5, 0xf3, 0x51, 0xbd, 0xc9, 0x4f, 0x95,
	0xeb, 0xeb, 0x45, 0xaa, 0x69, 0x7f, 0xd8, 0x70, 0x65, 0x72, 0xd1, 0x48, 0x72, 0xf4, 0x7b, 0xf1,
	0x97, 0x48, 0x78, 0x73, 0x70, 0x0a, 0x45, 0x6a, 0x84, 0xb4, 0x5a, 0x94, 0x73, 0xd1, 0x28, 0xbf,
	0xd5, 0xca, 0x4d, 0x3f, 0x5c, 0xec, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0xf4, 0x49,
	0x5d, 0x68, 0x03, 0x00, 0x00,
}
