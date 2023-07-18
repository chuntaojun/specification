// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config_file_response.proto

package config_manage // import "github.com/polarismesh/specification/source/go/api/v1/config_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type ConfigDiscoverResponse_ConfigDiscoverResponseType int32

const (
	ConfigDiscoverResponse_UNKNOWN           ConfigDiscoverResponse_ConfigDiscoverResponseType = 0
	ConfigDiscoverResponse_CONFIG_FILE       ConfigDiscoverResponse_ConfigDiscoverResponseType = 1
	ConfigDiscoverResponse_CONFIG_FILE_Names ConfigDiscoverResponse_ConfigDiscoverResponseType = 2
)

var ConfigDiscoverResponse_ConfigDiscoverResponseType_name = map[int32]string{
	0: "UNKNOWN",
	1: "CONFIG_FILE",
	2: "CONFIG_FILE_Names",
}
var ConfigDiscoverResponse_ConfigDiscoverResponseType_value = map[string]int32{
	"UNKNOWN":           0,
	"CONFIG_FILE":       1,
	"CONFIG_FILE_Names": 2,
}

func (x ConfigDiscoverResponse_ConfigDiscoverResponseType) String() string {
	return proto.EnumName(ConfigDiscoverResponse_ConfigDiscoverResponseType_name, int32(x))
}
func (ConfigDiscoverResponse_ConfigDiscoverResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{8, 0}
}

type ConfigSimpleResponse struct {
	Code                 *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigSimpleResponse) Reset()         { *m = ConfigSimpleResponse{} }
func (m *ConfigSimpleResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigSimpleResponse) ProtoMessage()    {}
func (*ConfigSimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{0}
}
func (m *ConfigSimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSimpleResponse.Unmarshal(m, b)
}
func (m *ConfigSimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSimpleResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigSimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSimpleResponse.Merge(dst, src)
}
func (m *ConfigSimpleResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigSimpleResponse.Size(m)
}
func (m *ConfigSimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSimpleResponse proto.InternalMessageInfo

func (m *ConfigSimpleResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigSimpleResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

type ConfigResponse struct {
	Code                     *wrapperspb.UInt32Value   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                     *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	ConfigFileGroup          *ConfigFileGroup          `protobuf:"bytes,3,opt,name=configFileGroup,proto3" json:"configFileGroup,omitempty"`
	ConfigFile               *ConfigFile               `protobuf:"bytes,4,opt,name=configFile,proto3" json:"configFile,omitempty"`
	ConfigFileRelease        *ConfigFileRelease        `protobuf:"bytes,5,opt,name=configFileRelease,proto3" json:"configFileRelease,omitempty"`
	ConfigFileReleaseHistory *ConfigFileReleaseHistory `protobuf:"bytes,6,opt,name=configFileReleaseHistory,proto3" json:"configFileReleaseHistory,omitempty"`
	ConfigFileTemplate       *ConfigFileTemplate       `protobuf:"bytes,7,opt,name=configFileTemplate,proto3" json:"configFileTemplate,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{1}
}
func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResponse.Unmarshal(m, b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(dst, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigResponse.Size(m)
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigResponse) GetConfigFileGroup() *ConfigFileGroup {
	if m != nil {
		return m.ConfigFileGroup
	}
	return nil
}

func (m *ConfigResponse) GetConfigFile() *ConfigFile {
	if m != nil {
		return m.ConfigFile
	}
	return nil
}

func (m *ConfigResponse) GetConfigFileRelease() *ConfigFileRelease {
	if m != nil {
		return m.ConfigFileRelease
	}
	return nil
}

func (m *ConfigResponse) GetConfigFileReleaseHistory() *ConfigFileReleaseHistory {
	if m != nil {
		return m.ConfigFileReleaseHistory
	}
	return nil
}

func (m *ConfigResponse) GetConfigFileTemplate() *ConfigFileTemplate {
	if m != nil {
		return m.ConfigFileTemplate
	}
	return nil
}

type ConfigBatchWriteResponse struct {
	Code                 *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Total                *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"`
	Responses            []*ConfigResponse       `protobuf:"bytes,4,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigBatchWriteResponse) Reset()         { *m = ConfigBatchWriteResponse{} }
func (m *ConfigBatchWriteResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigBatchWriteResponse) ProtoMessage()    {}
func (*ConfigBatchWriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{2}
}
func (m *ConfigBatchWriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigBatchWriteResponse.Unmarshal(m, b)
}
func (m *ConfigBatchWriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigBatchWriteResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigBatchWriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigBatchWriteResponse.Merge(dst, src)
}
func (m *ConfigBatchWriteResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigBatchWriteResponse.Size(m)
}
func (m *ConfigBatchWriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigBatchWriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigBatchWriteResponse proto.InternalMessageInfo

func (m *ConfigBatchWriteResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigBatchWriteResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigBatchWriteResponse) GetTotal() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *ConfigBatchWriteResponse) GetResponses() []*ConfigResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

type ConfigBatchQueryResponse struct {
	Code                       *wrapperspb.UInt32Value     `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                       *wrapperspb.StringValue     `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Total                      *wrapperspb.UInt32Value     `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"`
	ConfigFileGroups           []*ConfigFileGroup          `protobuf:"bytes,4,rep,name=configFileGroups,proto3" json:"configFileGroups,omitempty"`
	ConfigFiles                []*ConfigFile               `protobuf:"bytes,5,rep,name=configFiles,proto3" json:"configFiles,omitempty"`
	ConfigFileReleases         []*ConfigFileRelease        `protobuf:"bytes,6,rep,name=configFileReleases,proto3" json:"configFileReleases,omitempty"`
	ConfigFileReleaseHistories []*ConfigFileReleaseHistory `protobuf:"bytes,7,rep,name=configFileReleaseHistories,proto3" json:"configFileReleaseHistories,omitempty"`
	ConfigFileTemplates        []*ConfigFileTemplate       `protobuf:"bytes,8,rep,name=configFileTemplates,proto3" json:"configFileTemplates,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                    `json:"-"`
	XXX_unrecognized           []byte                      `json:"-"`
	XXX_sizecache              int32                       `json:"-"`
}

func (m *ConfigBatchQueryResponse) Reset()         { *m = ConfigBatchQueryResponse{} }
func (m *ConfigBatchQueryResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigBatchQueryResponse) ProtoMessage()    {}
func (*ConfigBatchQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{3}
}
func (m *ConfigBatchQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigBatchQueryResponse.Unmarshal(m, b)
}
func (m *ConfigBatchQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigBatchQueryResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigBatchQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigBatchQueryResponse.Merge(dst, src)
}
func (m *ConfigBatchQueryResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigBatchQueryResponse.Size(m)
}
func (m *ConfigBatchQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigBatchQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigBatchQueryResponse proto.InternalMessageInfo

func (m *ConfigBatchQueryResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetTotal() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFileGroups() []*ConfigFileGroup {
	if m != nil {
		return m.ConfigFileGroups
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFiles() []*ConfigFile {
	if m != nil {
		return m.ConfigFiles
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFileReleases() []*ConfigFileRelease {
	if m != nil {
		return m.ConfigFileReleases
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFileReleaseHistories() []*ConfigFileReleaseHistory {
	if m != nil {
		return m.ConfigFileReleaseHistories
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFileTemplates() []*ConfigFileTemplate {
	if m != nil {
		return m.ConfigFileTemplates
	}
	return nil
}

type ConfigClientResponse struct {
	Code                 *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	ConfigFile           *ClientConfigFileInfo   `protobuf:"bytes,3,opt,name=configFile,proto3" json:"configFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigClientResponse) Reset()         { *m = ConfigClientResponse{} }
func (m *ConfigClientResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigClientResponse) ProtoMessage()    {}
func (*ConfigClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{4}
}
func (m *ConfigClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigClientResponse.Unmarshal(m, b)
}
func (m *ConfigClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigClientResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigClientResponse.Merge(dst, src)
}
func (m *ConfigClientResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigClientResponse.Size(m)
}
func (m *ConfigClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigClientResponse proto.InternalMessageInfo

func (m *ConfigClientResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigClientResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigClientResponse) GetConfigFile() *ClientConfigFileInfo {
	if m != nil {
		return m.ConfigFile
	}
	return nil
}

type ConfigImportResponse struct {
	Code                 *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	CreateConfigFiles    []*ConfigFile           `protobuf:"bytes,3,rep,name=createConfigFiles,proto3" json:"createConfigFiles,omitempty"`
	SkipConfigFiles      []*ConfigFile           `protobuf:"bytes,4,rep,name=skipConfigFiles,proto3" json:"skipConfigFiles,omitempty"`
	OverwriteConfigFiles []*ConfigFile           `protobuf:"bytes,5,rep,name=overwriteConfigFiles,proto3" json:"overwriteConfigFiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigImportResponse) Reset()         { *m = ConfigImportResponse{} }
func (m *ConfigImportResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigImportResponse) ProtoMessage()    {}
func (*ConfigImportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{5}
}
func (m *ConfigImportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigImportResponse.Unmarshal(m, b)
}
func (m *ConfigImportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigImportResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigImportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigImportResponse.Merge(dst, src)
}
func (m *ConfigImportResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigImportResponse.Size(m)
}
func (m *ConfigImportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigImportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigImportResponse proto.InternalMessageInfo

func (m *ConfigImportResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigImportResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigImportResponse) GetCreateConfigFiles() []*ConfigFile {
	if m != nil {
		return m.CreateConfigFiles
	}
	return nil
}

func (m *ConfigImportResponse) GetSkipConfigFiles() []*ConfigFile {
	if m != nil {
		return m.SkipConfigFiles
	}
	return nil
}

func (m *ConfigImportResponse) GetOverwriteConfigFiles() []*ConfigFile {
	if m != nil {
		return m.OverwriteConfigFiles
	}
	return nil
}

type ConfigExportResponse struct {
	Code                 *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Data                 *wrapperspb.BytesValue  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigExportResponse) Reset()         { *m = ConfigExportResponse{} }
func (m *ConfigExportResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigExportResponse) ProtoMessage()    {}
func (*ConfigExportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{6}
}
func (m *ConfigExportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigExportResponse.Unmarshal(m, b)
}
func (m *ConfigExportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigExportResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigExportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigExportResponse.Merge(dst, src)
}
func (m *ConfigExportResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigExportResponse.Size(m)
}
func (m *ConfigExportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigExportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigExportResponse proto.InternalMessageInfo

func (m *ConfigExportResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigExportResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigExportResponse) GetData() *wrapperspb.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

type ConfigEncryptAlgorithmResponse struct {
	Code                 *wrapperspb.UInt32Value   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Algorithms           []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=algorithms,proto3" json:"algorithms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ConfigEncryptAlgorithmResponse) Reset()         { *m = ConfigEncryptAlgorithmResponse{} }
func (m *ConfigEncryptAlgorithmResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigEncryptAlgorithmResponse) ProtoMessage()    {}
func (*ConfigEncryptAlgorithmResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{7}
}
func (m *ConfigEncryptAlgorithmResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigEncryptAlgorithmResponse.Unmarshal(m, b)
}
func (m *ConfigEncryptAlgorithmResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigEncryptAlgorithmResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigEncryptAlgorithmResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigEncryptAlgorithmResponse.Merge(dst, src)
}
func (m *ConfigEncryptAlgorithmResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigEncryptAlgorithmResponse.Size(m)
}
func (m *ConfigEncryptAlgorithmResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigEncryptAlgorithmResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigEncryptAlgorithmResponse proto.InternalMessageInfo

func (m *ConfigEncryptAlgorithmResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigEncryptAlgorithmResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigEncryptAlgorithmResponse) GetAlgorithms() []*wrapperspb.StringValue {
	if m != nil {
		return m.Algorithms
	}
	return nil
}

type ConfigDiscoverResponse struct {
	Code                 uint32                                            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 string                                            `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Revision             string                                            `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	Type                 ConfigDiscoverResponse_ConfigDiscoverResponseType `protobuf:"varint,4,opt,name=type,proto3,enum=v1.ConfigDiscoverResponse_ConfigDiscoverResponseType" json:"type,omitempty"`
	Namespace            string                                            `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                string                                            `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
	FileName             string                                            `protobuf:"bytes,7,opt,name=file_name,proto3" json:"file_name,omitempty"`
	ConfigFile           *ClientConfigFileInfo                             `protobuf:"bytes,8,opt,name=config_file,proto3" json:"config_file,omitempty"`
	ConfigFileNames      []*ClientConfigFileInfo                           `protobuf:"bytes,9,rep,name=config_file_names,proto3" json:"config_file_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *ConfigDiscoverResponse) Reset()         { *m = ConfigDiscoverResponse{} }
func (m *ConfigDiscoverResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigDiscoverResponse) ProtoMessage()    {}
func (*ConfigDiscoverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_ecc7befa247015cc, []int{8}
}
func (m *ConfigDiscoverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigDiscoverResponse.Unmarshal(m, b)
}
func (m *ConfigDiscoverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigDiscoverResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigDiscoverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigDiscoverResponse.Merge(dst, src)
}
func (m *ConfigDiscoverResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigDiscoverResponse.Size(m)
}
func (m *ConfigDiscoverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigDiscoverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigDiscoverResponse proto.InternalMessageInfo

func (m *ConfigDiscoverResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ConfigDiscoverResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *ConfigDiscoverResponse) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ConfigDiscoverResponse) GetType() ConfigDiscoverResponse_ConfigDiscoverResponseType {
	if m != nil {
		return m.Type
	}
	return ConfigDiscoverResponse_UNKNOWN
}

func (m *ConfigDiscoverResponse) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ConfigDiscoverResponse) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *ConfigDiscoverResponse) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *ConfigDiscoverResponse) GetConfigFile() *ClientConfigFileInfo {
	if m != nil {
		return m.ConfigFile
	}
	return nil
}

func (m *ConfigDiscoverResponse) GetConfigFileNames() []*ClientConfigFileInfo {
	if m != nil {
		return m.ConfigFileNames
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigSimpleResponse)(nil), "v1.ConfigSimpleResponse")
	proto.RegisterType((*ConfigResponse)(nil), "v1.ConfigResponse")
	proto.RegisterType((*ConfigBatchWriteResponse)(nil), "v1.ConfigBatchWriteResponse")
	proto.RegisterType((*ConfigBatchQueryResponse)(nil), "v1.ConfigBatchQueryResponse")
	proto.RegisterType((*ConfigClientResponse)(nil), "v1.ConfigClientResponse")
	proto.RegisterType((*ConfigImportResponse)(nil), "v1.ConfigImportResponse")
	proto.RegisterType((*ConfigExportResponse)(nil), "v1.ConfigExportResponse")
	proto.RegisterType((*ConfigEncryptAlgorithmResponse)(nil), "v1.ConfigEncryptAlgorithmResponse")
	proto.RegisterType((*ConfigDiscoverResponse)(nil), "v1.ConfigDiscoverResponse")
	proto.RegisterEnum("v1.ConfigDiscoverResponse_ConfigDiscoverResponseType", ConfigDiscoverResponse_ConfigDiscoverResponseType_name, ConfigDiscoverResponse_ConfigDiscoverResponseType_value)
}

func init() {
	proto.RegisterFile("config_file_response.proto", fileDescriptor_config_file_response_ecc7befa247015cc)
}

var fileDescriptor_config_file_response_ecc7befa247015cc = []byte{
	// 845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x18, 0x25, 0x9b, 0xec, 0x4f, 0xbe, 0x88, 0xfd, 0x99, 0x6e, 0xcb, 0x28, 0x54, 0x55, 0xe5, 0xab,
	0x5e, 0xd9, 0xbb, 0xa9, 0x40, 0x15, 0x02, 0x21, 0x62, 0x36, 0x6d, 0x04, 0x4a, 0xa9, 0xdb, 0x52,
	0x84, 0x90, 0x56, 0xb3, 0xee, 0xc4, 0x3b, 0xc2, 0xf6, 0x8c, 0x66, 0x26, 0x29, 0xe1, 0x39, 0x10,
	0x2f, 0xc1, 0x2d, 0xe2, 0x01, 0xb8, 0xe4, 0x8a, 0x87, 0xe0, 0x41, 0x90, 0x67, 0xec, 0xb5, 0x13,
	0x3b, 0xd9, 0xbb, 0xa8, 0x97, 0xfe, 0xe6, 0x9c, 0x33, 0xe7, 0xf3, 0x1c, 0x7f, 0x1e, 0xe8, 0x87,
	0x3c, 0x9d, 0xb2, 0xe8, 0x72, 0xca, 0x62, 0x7a, 0x29, 0xa9, 0x12, 0x3c, 0x55, 0xd4, 0x15, 0x92,
	0x6b, 0x8e, 0x76, 0xe6, 0xe7, 0xfd, 0x07, 0x11, 0xe7, 0x51, 0x4c, 0x3d, 0x53, 0xb9, 0x9a, 0x4d,
	0xbd, 0x77, 0x92, 0x08, 0x41, 0xa5, 0xb2, 0x98, 0xfe, 0x49, 0x85, 0x6f, 0x4b, 0xce, 0xaf, 0x70,
	0xea, 0x9b, 0xe2, 0x4b, 0x96, 0x88, 0x98, 0x06, 0xb9, 0x28, 0x3a, 0x83, 0x4e, 0xc8, 0xdf, 0x52,
	0xdc, 0x7a, 0xd8, 0x7a, 0xd4, 0x1b, 0xdc, 0x77, 0xad, 0xb2, 0x5b, 0x28, 0xbb, 0xaf, 0xc7, 0xa9,
	0x7e, 0x3c, 0xf8, 0x9e, 0xc4, 0x33, 0x1a, 0x18, 0x64, 0xc6, 0x60, 0xe9, 0x94, 0xe3, 0x9d, 0x35,
	0x8c, 0x97, 0x5a, 0xb2, 0x34, 0xca, 0x19, 0x19, 0xd2, 0xf9, 0xa7, 0x0d, 0x87, 0x76, 0xf3, 0x6d,
	0x6e, 0x8b, 0xbe, 0x80, 0x23, 0xfb, 0x1e, 0x46, 0x2c, 0xa6, 0x4f, 0x25, 0x9f, 0x09, 0xdc, 0x36,
	0xe4, 0x3b, 0xee, 0xfc, 0xdc, 0xf5, 0x97, 0x97, 0x82, 0x55, 0x2c, 0x72, 0x01, 0xca, 0x12, 0xee,
	0x18, 0xe6, 0xe1, 0x32, 0x33, 0xa8, 0x20, 0x90, 0x0f, 0x27, 0xe5, 0x53, 0x40, 0x63, 0x4a, 0x14,
	0xc5, 0xbb, 0x86, 0x76, 0x77, 0x85, 0x66, 0x17, 0x83, 0x3a, 0x1e, 0xfd, 0x00, 0xb8, 0x56, 0x7c,
	0xc6, 0x94, 0xe6, 0x72, 0x81, 0xf7, 0xf2, 0xce, 0x9b, 0xb4, 0x72, 0x4c, 0xb0, 0x96, 0x8d, 0x46,
	0x80, 0xca, 0xb5, 0x57, 0x34, 0x11, 0x31, 0xd1, 0x14, 0xef, 0x1b, 0xcd, 0x7b, 0xcb, 0x9a, 0xc5,
	0x6a, 0xd0, 0xc0, 0x70, 0xfe, 0x6b, 0x01, 0xb6, 0xd0, 0x21, 0xd1, 0xe1, 0xf5, 0x1b, 0xc9, 0xf4,
	0x56, 0xd3, 0x84, 0x06, 0xb0, 0xab, 0xb9, 0x26, 0x71, 0x7e, 0x98, 0x9b, 0x37, 0xb1, 0x50, 0x74,
	0x06, 0xdd, 0xe2, 0x33, 0x52, 0xb8, 0xf3, 0xb0, 0xfd, 0xa8, 0x37, 0x40, 0x65, 0xcf, 0x85, 0xfd,
	0xa0, 0x04, 0x39, 0xbf, 0x75, 0x96, 0xda, 0x7c, 0x31, 0xa3, 0x72, 0xf1, 0xde, 0xb7, 0xf9, 0x25,
	0x1c, 0xaf, 0xa4, 0xb8, 0xe8, 0xb6, 0x31, 0xf2, 0x35, 0x30, 0x3a, 0x83, 0x5e, 0x59, 0x53, 0x78,
	0xd7, 0x70, 0x57, 0x43, 0x5f, 0x85, 0xa0, 0x8b, 0x6a, 0xac, 0xf2, 0xc8, 0x29, 0xbc, 0x67, 0x88,
	0x6b, 0x62, 0xdf, 0x40, 0x40, 0x3f, 0x15, 0x33, 0xaf, 0x96, 0x5c, 0x46, 0x15, 0xde, 0x37, 0x72,
	0x9b, 0x93, 0xbf, 0x81, 0x8f, 0x9e, 0xc1, 0x9d, 0x7a, 0x92, 0x15, 0x3e, 0x30, 0xb2, 0xeb, 0xc2,
	0xdf, 0x44, 0x71, 0xfe, 0x6c, 0x15, 0x73, 0xd4, 0x8f, 0x19, 0x4d, 0xf5, 0x56, 0x23, 0xf1, 0x64,
	0x69, 0x22, 0xd9, 0x5c, 0x60, 0xe3, 0xde, 0x78, 0x29, 0x7b, 0x18, 0xa7, 0x53, 0x5e, 0x9d, 0x4d,
	0xce, 0x5f, 0x3b, 0x85, 0xed, 0x71, 0x22, 0xb8, 0xdc, 0xae, 0xed, 0xcf, 0xe1, 0x24, 0x94, 0x94,
	0x68, 0xea, 0x57, 0xa2, 0xd5, 0x6e, 0x8c, 0x56, 0x1d, 0x88, 0x9e, 0xc0, 0x91, 0xfa, 0x99, 0x89,
	0x2a, 0xb7, 0xd3, 0xc8, 0x5d, 0x85, 0xa1, 0x21, 0x9c, 0xf2, 0x39, 0x95, 0xef, 0xb2, 0x09, 0xe5,
	0xdf, 0x9a, 0xea, 0x46, 0xac, 0xf3, 0xc7, 0xcd, 0x79, 0x5f, 0xfc, 0xb2, 0xf5, 0x17, 0xe7, 0x41,
	0xe7, 0x2d, 0xd1, 0x24, 0x3f, 0xe9, 0x8f, 0x6b, 0x8c, 0xe1, 0x42, 0x53, 0x95, 0x13, 0x32, 0xa0,
	0xf3, 0x77, 0x0b, 0x1e, 0xe4, 0x6e, 0xd3, 0x50, 0x2e, 0x84, 0xfe, 0x2a, 0x8e, 0xb8, 0x64, 0xfa,
	0x3a, 0xd9, 0xf2, 0x81, 0x03, 0x29, 0x36, 0x2e, 0x4e, 0x7a, 0x33, 0xaf, 0x82, 0x77, 0xfe, 0x6d,
	0xc3, 0x3d, 0xdb, 0xc4, 0xd7, 0x4c, 0x85, 0xd9, 0xa9, 0xdc, 0x98, 0x47, 0x15, 0xf3, 0x1f, 0xe6,
	0xf6, 0x50, 0xc5, 0x5e, 0x37, 0x37, 0xd0, 0x87, 0x03, 0x49, 0xe7, 0x4c, 0x31, 0x9e, 0x9a, 0x97,
	0xd7, 0x0d, 0x6e, 0x9e, 0xd1, 0x18, 0x3a, 0x7a, 0x21, 0xec, 0x0f, 0xfd, 0x70, 0xf0, 0x49, 0x99,
	0x82, 0xd5, 0xdd, 0xd6, 0x94, 0x5f, 0x2d, 0x04, 0x0d, 0x8c, 0x04, 0xba, 0x0f, 0xdd, 0x94, 0x24,
	0x54, 0x09, 0x12, 0xda, 0x3f, 0x7d, 0x37, 0x28, 0x0b, 0xe8, 0x14, 0x76, 0x23, 0x73, 0xe9, 0xd8,
	0x33, 0x2b, 0xf6, 0x21, 0xe3, 0x98, 0x5b, 0x5d, 0x86, 0x33, 0x7f, 0xdf, 0x6e, 0x50, 0x16, 0xd0,
	0x67, 0xc5, 0xfc, 0x35, 0x57, 0x37, 0x7c, 0x70, 0xcb, 0x27, 0x5e, 0x05, 0xa3, 0x11, 0x54, 0xaf,
	0x7d, 0x46, 0x4f, 0xe1, 0xae, 0x79, 0xf9, 0xeb, 0x15, 0xea, 0x14, 0xe7, 0x05, 0xf4, 0xd7, 0x77,
	0x8e, 0x7a, 0xb0, 0xff, 0x7a, 0xf2, 0xcd, 0xe4, 0xf9, 0x9b, 0xc9, 0xf1, 0x07, 0xe8, 0x08, 0x7a,
	0xfe, 0xf3, 0xc9, 0x68, 0xfc, 0xf4, 0x72, 0x34, 0xfe, 0xf6, 0xe2, 0xb8, 0x85, 0xee, 0xc2, 0x49,
	0xa5, 0x70, 0x39, 0xc9, 0x04, 0x8f, 0x77, 0x86, 0xbf, 0xb7, 0xe0, 0xd3, 0x90, 0x27, 0xae, 0xa6,
	0x69, 0x48, 0x53, 0xed, 0x0a, 0x1e, 0x13, 0xc9, 0x94, 0xab, 0x04, 0x0d, 0xd9, 0x94, 0x85, 0x44,
	0x33, 0x9e, 0xba, 0x44, 0xb0, 0xcc, 0xa7, 0x75, 0xe4, 0x26, 0x24, 0x25, 0x11, 0x1d, 0x7e, 0x54,
	0x1d, 0xf8, 0xd6, 0xc7, 0x77, 0x59, 0x82, 0x7e, 0xf4, 0x23, 0xa6, 0xaf, 0x67, 0x57, 0x6e, 0xc8,
	0x13, 0x2f, 0xd7, 0x4b, 0xa8, 0xba, 0xf6, 0x96, 0x34, 0x3d, 0xc5, 0x67, 0x32, 0xa4, 0x5e, 0xc4,
	0x3d, 0x22, 0x98, 0x37, 0x3f, 0xf7, 0xf2, 0x7e, 0xad, 0xfa, 0xd5, 0x9e, 0x49, 0xe3, 0xe3, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x86, 0x63, 0x4d, 0x71, 0x0b, 0x00, 0x00,
}
