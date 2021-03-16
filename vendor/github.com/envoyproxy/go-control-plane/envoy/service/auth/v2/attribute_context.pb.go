// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v2/attribute_context.proto

package envoy_service_auth_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AttributeContext struct {
	Source               *AttributeContext_Peer    `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination          *AttributeContext_Peer    `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Request              *AttributeContext_Request `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	ContextExtensions    map[string]string         `protobuf:"bytes,10,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MetadataContext      *core.Metadata            `protobuf:"bytes,11,opt,name=metadata_context,json=metadataContext,proto3" json:"metadata_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AttributeContext) Reset()         { *m = AttributeContext{} }
func (m *AttributeContext) String() string { return proto.CompactTextString(m) }
func (*AttributeContext) ProtoMessage()    {}
func (*AttributeContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6030c9468e3591b, []int{0}
}

func (m *AttributeContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext.Unmarshal(m, b)
}
func (m *AttributeContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext.Marshal(b, m, deterministic)
}
func (m *AttributeContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext.Merge(m, src)
}
func (m *AttributeContext) XXX_Size() int {
	return xxx_messageInfo_AttributeContext.Size(m)
}
func (m *AttributeContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext proto.InternalMessageInfo

func (m *AttributeContext) GetSource() *AttributeContext_Peer {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AttributeContext) GetDestination() *AttributeContext_Peer {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *AttributeContext) GetRequest() *AttributeContext_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AttributeContext) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

func (m *AttributeContext) GetMetadataContext() *core.Metadata {
	if m != nil {
		return m.MetadataContext
	}
	return nil
}

type AttributeContext_Peer struct {
	Address              *core.Address     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Service              string            `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Principal            string            `protobuf:"bytes,4,opt,name=principal,proto3" json:"principal,omitempty"`
	Certificate          string            `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AttributeContext_Peer) Reset()         { *m = AttributeContext_Peer{} }
func (m *AttributeContext_Peer) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_Peer) ProtoMessage()    {}
func (*AttributeContext_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6030c9468e3591b, []int{0, 0}
}

func (m *AttributeContext_Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_Peer.Unmarshal(m, b)
}
func (m *AttributeContext_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_Peer.Marshal(b, m, deterministic)
}
func (m *AttributeContext_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_Peer.Merge(m, src)
}
func (m *AttributeContext_Peer) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_Peer.Size(m)
}
func (m *AttributeContext_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_Peer proto.InternalMessageInfo

func (m *AttributeContext_Peer) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AttributeContext_Peer) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *AttributeContext_Peer) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *AttributeContext_Peer) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func (m *AttributeContext_Peer) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

type AttributeContext_Request struct {
	Time                 *timestamp.Timestamp          `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Http                 *AttributeContext_HttpRequest `protobuf:"bytes,2,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AttributeContext_Request) Reset()         { *m = AttributeContext_Request{} }
func (m *AttributeContext_Request) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_Request) ProtoMessage()    {}
func (*AttributeContext_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6030c9468e3591b, []int{0, 1}
}

func (m *AttributeContext_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_Request.Unmarshal(m, b)
}
func (m *AttributeContext_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_Request.Marshal(b, m, deterministic)
}
func (m *AttributeContext_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_Request.Merge(m, src)
}
func (m *AttributeContext_Request) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_Request.Size(m)
}
func (m *AttributeContext_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_Request.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_Request proto.InternalMessageInfo

func (m *AttributeContext_Request) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *AttributeContext_Request) GetHttp() *AttributeContext_HttpRequest {
	if m != nil {
		return m.Http
	}
	return nil
}

type AttributeContext_HttpRequest struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Method               string            `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Headers              map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Path                 string            `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Host                 string            `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Scheme               string            `protobuf:"bytes,6,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Query                string            `protobuf:"bytes,7,opt,name=query,proto3" json:"query,omitempty"`
	Fragment             string            `protobuf:"bytes,8,opt,name=fragment,proto3" json:"fragment,omitempty"`
	Size                 int64             `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Protocol             string            `protobuf:"bytes,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Body                 string            `protobuf:"bytes,11,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AttributeContext_HttpRequest) Reset()         { *m = AttributeContext_HttpRequest{} }
func (m *AttributeContext_HttpRequest) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_HttpRequest) ProtoMessage()    {}
func (*AttributeContext_HttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6030c9468e3591b, []int{0, 2}
}

func (m *AttributeContext_HttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_HttpRequest.Unmarshal(m, b)
}
func (m *AttributeContext_HttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_HttpRequest.Marshal(b, m, deterministic)
}
func (m *AttributeContext_HttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_HttpRequest.Merge(m, src)
}
func (m *AttributeContext_HttpRequest) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_HttpRequest.Size(m)
}
func (m *AttributeContext_HttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_HttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_HttpRequest proto.InternalMessageInfo

func (m *AttributeContext_HttpRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *AttributeContext_HttpRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetFragment() string {
	if m != nil {
		return m.Fragment
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *AttributeContext_HttpRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*AttributeContext)(nil), "envoy.service.auth.v2.AttributeContext")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v2.AttributeContext.ContextExtensionsEntry")
	proto.RegisterType((*AttributeContext_Peer)(nil), "envoy.service.auth.v2.AttributeContext.Peer")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v2.AttributeContext.Peer.LabelsEntry")
	proto.RegisterType((*AttributeContext_Request)(nil), "envoy.service.auth.v2.AttributeContext.Request")
	proto.RegisterType((*AttributeContext_HttpRequest)(nil), "envoy.service.auth.v2.AttributeContext.HttpRequest")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v2.AttributeContext.HttpRequest.HeadersEntry")
}

func init() {
	proto.RegisterFile("envoy/service/auth/v2/attribute_context.proto", fileDescriptor_a6030c9468e3591b)
}

var fileDescriptor_a6030c9468e3591b = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x7e, 0x9a, 0x34, 0x13, 0x04, 0x61, 0x45, 0x2b, 0xcb, 0x14, 0x35, 0x82, 0x4b, 0x0f,
	0x60, 0x4b, 0x29, 0x87, 0xd2, 0x03, 0xa2, 0xa5, 0x2d, 0x45, 0x02, 0x14, 0x59, 0x9c, 0xb8, 0x54,
	0x1b, 0x7b, 0xda, 0xac, 0x48, 0xbc, 0xee, 0xee, 0x38, 0x6a, 0x38, 0x21, 0x78, 0x12, 0xc4, 0x5b,
	0xf0, 0x04, 0x5c, 0x79, 0x23, 0xb4, 0x3f, 0x2e, 0x51, 0xc9, 0xa1, 0xed, 0xc9, 0x33, 0xe3, 0x6f,
	0xbe, 0x9d, 0x99, 0x6f, 0x06, 0x9e, 0x61, 0x3e, 0x93, 0xf3, 0x58, 0xa3, 0x9a, 0x89, 0x14, 0x63,
	0x5e, 0xd2, 0x38, 0x9e, 0x0d, 0x62, 0x4e, 0xa4, 0xc4, 0xa8, 0x24, 0x3c, 0x49, 0x65, 0x4e, 0x78,
	0x41, 0x51, 0xa1, 0x24, 0x49, 0xb6, 0x66, 0xe1, 0x91, 0x87, 0x47, 0x06, 0x1e, 0xcd, 0x06, 0xe1,
	0xa6, 0x63, 0xe1, 0x85, 0x30, 0xc9, 0xa9, 0x54, 0x18, 0xf3, 0x2c, 0x53, 0xa8, 0xb5, 0xcb, 0x0b,
	0x37, 0xfe, 0x07, 0x8c, 0xb8, 0x46, 0xff, 0x77, 0xf3, 0x4c, 0xca, 0xb3, 0x09, 0xc6, 0xd6, 0x1b,
	0x95, 0xa7, 0x31, 0x89, 0x29, 0x6a, 0xe2, 0xd3, 0xc2, 0x03, 0x1e, 0x95, 0x59, 0xc1, 0x63, 0x9e,
	0xe7, 0x92, 0x38, 0x09, 0x99, 0xeb, 0x58, 0x13, 0xa7, 0xd2, 0xb3, 0x3f, 0xfe, 0x01, 0xd0, 0xdb,
	0xab, 0x2a, 0x7e, 0xed, 0x0a, 0x66, 0x07, 0xd0, 0xd2, 0xb2, 0x54, 0x29, 0x06, 0xb5, 0x7e, 0x6d,
	0xab, 0x3b, 0x78, 0x1a, 0x2d, 0xad, 0x3d, 0xba, 0x9a, 0x18, 0x0d, 0x11, 0x55, 0xe2, 0x73, 0xd9,
	0x07, 0xe8, 0x66, 0xa8, 0x49, 0xe4, 0xf6, 0xdd, 0xa0, 0x7e, 0x0b, 0xaa, 0x45, 0x02, 0xf6, 0x16,
	0xda, 0x0a, 0xcf, 0x4b, 0xd4, 0x14, 0x34, 0x2d, 0x57, 0x7c, 0x5d, 0xae, 0xc4, 0xa5, 0x25, 0x55,
	0x3e, 0x9b, 0x02, 0xf3, 0xe2, 0x9c, 0xe0, 0x05, 0x61, 0xae, 0xcd, 0x60, 0x02, 0xe8, 0x37, 0xb6,
	0xba, 0x83, 0x97, 0xd7, 0x65, 0xf5, 0xdf, 0xc3, 0x4b, 0x82, 0xc3, 0x9c, 0xd4, 0x3c, 0xb9, 0x9f,
	0x5e, 0x8d, 0xb3, 0x23, 0xe8, 0x4d, 0x91, 0x78, 0xc6, 0x89, 0x57, 0x4b, 0x11, 0x74, 0x6d, 0x0b,
	0x0f, 0xfd, 0x63, 0xbc, 0x10, 0xe6, 0x0d, 0xa3, 0x6e, 0xf4, 0xde, 0x43, 0x93, 0x7b, 0x55, 0x92,
	0x7f, 0x29, 0xfc, 0x59, 0x87, 0xa6, 0x99, 0x0b, 0x7b, 0x0e, 0x6d, 0xbf, 0x24, 0x5e, 0xa1, 0x70,
	0x09, 0xcf, 0x9e, 0x43, 0x24, 0x15, 0x94, 0x05, 0xd0, 0xf6, 0x4d, 0x59, 0x31, 0x3a, 0x49, 0xe5,
	0xb2, 0x21, 0xb4, 0x26, 0x7c, 0x84, 0x13, 0x1d, 0x34, 0xec, 0x0c, 0x76, 0x6e, 0xa2, 0x52, 0xf4,
	0xce, 0xa6, 0xba, 0xee, 0x3d, 0x0f, 0xdb, 0x80, 0x4e, 0xa1, 0x44, 0x9e, 0x8a, 0x82, 0x4f, 0xac,
	0x5c, 0x9d, 0xe4, 0x5f, 0x80, 0xf5, 0xa1, 0x9b, 0xa2, 0x22, 0x71, 0x2a, 0x52, 0x4e, 0x18, 0xac,
	0xd8, 0xff, 0x8b, 0xa1, 0xf0, 0x05, 0x74, 0x17, 0x68, 0x59, 0x0f, 0x1a, 0x9f, 0x71, 0x6e, 0x9b,
	0xed, 0x24, 0xc6, 0x64, 0x0f, 0x60, 0x65, 0xc6, 0x27, 0x65, 0xd5, 0x8a, 0x73, 0x76, 0xeb, 0x3b,
	0xb5, 0xf0, 0x5b, 0x0d, 0xda, 0x5e, 0x71, 0x16, 0x41, 0xd3, 0x1c, 0xc4, 0xe5, 0x94, 0xdc, 0xb5,
	0x44, 0xd5, 0xb5, 0x44, 0x1f, 0xab, 0x6b, 0x49, 0x2c, 0x8e, 0xbd, 0x81, 0xe6, 0x98, 0xa8, 0xf0,
	0xcb, 0xba, 0x7d, 0xdd, 0x31, 0x1c, 0x13, 0x15, 0xd5, 0x92, 0x59, 0x82, 0xf0, 0x7b, 0x03, 0xba,
	0x0b, 0x51, 0x76, 0x17, 0xea, 0x22, 0xf3, 0xf5, 0xd7, 0x45, 0xc6, 0xd6, 0xa1, 0x35, 0x45, 0x1a,
	0xcb, 0xcc, 0xd7, 0xef, 0x3d, 0xf6, 0x09, 0xda, 0x63, 0xe4, 0x19, 0xaa, 0x4a, 0x8a, 0x57, 0xb7,
	0xa8, 0x21, 0x3a, 0x76, 0x14, 0x4e, 0x92, 0x8a, 0x90, 0x31, 0x68, 0x16, 0x9c, 0xc6, 0x5e, 0x0e,
	0x6b, 0x9b, 0xd8, 0x58, 0x6a, 0xf2, 0x12, 0x58, 0xdb, 0xd4, 0xa6, 0xd3, 0x31, 0x4e, 0x31, 0x68,
	0xb9, 0xda, 0x9c, 0x67, 0x46, 0x7e, 0x5e, 0xa2, 0x9a, 0x07, 0x6d, 0x37, 0x72, 0xeb, 0xb0, 0x10,
	0x56, 0x4f, 0x15, 0x3f, 0x9b, 0x62, 0x4e, 0xc1, 0xaa, 0xfd, 0x71, 0xe9, 0x1b, 0x76, 0x2d, 0xbe,
	0x60, 0xd0, 0xe9, 0xd7, 0xb6, 0x1a, 0x89, 0xb5, 0x0d, 0xde, 0x8e, 0x3f, 0x95, 0x93, 0x00, 0x1c,
	0xbe, 0xf2, 0x0d, 0x7e, 0x24, 0xb3, 0xb9, 0x3d, 0x8e, 0x4e, 0x62, 0xed, 0x70, 0x17, 0xee, 0x2c,
	0xb6, 0x73, 0xa3, 0x55, 0x38, 0x80, 0xf5, 0xe5, 0x57, 0x7a, 0x13, 0x96, 0xfd, 0xa3, 0x5f, 0x5f,
	0x7f, 0xff, 0x69, 0xd5, 0x7b, 0x35, 0x78, 0x22, 0xa4, 0x93, 0xa3, 0x50, 0xf2, 0x62, 0xbe, 0x5c,
	0x99, 0xfd, 0xb5, 0xab, 0xd2, 0x0c, 0x4d, 0x7b, 0xc3, 0xda, 0xa8, 0x65, 0xfb, 0xdc, 0xfe, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xd1, 0xda, 0x79, 0x9c, 0x39, 0x06, 0x00, 0x00,
}