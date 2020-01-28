// Code generated by protoc-gen-go. DO NOT EDIT.
// source: osprey.proto

package pb

import (
	fmt "fmt"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_osprey_968553557d684e87, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

type LoginResponse struct {
	Cluster              *Cluster      `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Provider             *AuthProvider `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	User                 *User         `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_osprey_968553557d684e87, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *LoginResponse) GetProvider() *AuthProvider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ClusterInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterInfoRequest) Reset()         { *m = ClusterInfoRequest{} }
func (m *ClusterInfoRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterInfoRequest) ProtoMessage()    {}
func (*ClusterInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_osprey_968553557d684e87, []int{2}
}
func (m *ClusterInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterInfoRequest.Unmarshal(m, b)
}
func (m *ClusterInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterInfoRequest.Marshal(b, m, deterministic)
}
func (dst *ClusterInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterInfoRequest.Merge(dst, src)
}
func (m *ClusterInfoRequest) XXX_Size() int {
	return xxx_messageInfo_ClusterInfoRequest.Size(m)
}
func (m *ClusterInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterInfoRequest proto.InternalMessageInfo

type ClusterInfoResponse struct {
	Cluster              *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterInfoResponse) Reset()         { *m = ClusterInfoResponse{} }
func (m *ClusterInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterInfoResponse) ProtoMessage()    {}
func (*ClusterInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_osprey_968553557d684e87, []int{3}
}
func (m *ClusterInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterInfoResponse.Unmarshal(m, b)
}
func (m *ClusterInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterInfoResponse.Marshal(b, m, deterministic)
}
func (dst *ClusterInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterInfoResponse.Merge(dst, src)
}
func (m *ClusterInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterInfoResponse.Size(m)
}
func (m *ClusterInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterInfoResponse proto.InternalMessageInfo

func (m *ClusterInfoResponse) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type Cluster struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ApiServerURL         string   `protobuf:"bytes,2,opt,name=apiServerURL" json:"apiServerURL,omitempty"`
	ApiServerCA          string   `protobuf:"bytes,3,opt,name=apiServerCA" json:"apiServerCA,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_osprey_968553557d684e87, []int{4}
}
func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (dst *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(dst, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetApiServerURL() string {
	if m != nil {
		return m.ApiServerURL
	}
	return ""
}

func (m *Cluster) GetApiServerCA() string {
	if m != nil {
		return m.ApiServerCA
	}
	return ""
}

type AuthProvider struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=clientSecret" json:"clientSecret,omitempty"`
	IssuerURL            string   `protobuf:"bytes,3,opt,name=issuerURL" json:"issuerURL,omitempty"`
	IssuerCA             string   `protobuf:"bytes,4,opt,name=issuerCA" json:"issuerCA,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthProvider) Reset()         { *m = AuthProvider{} }
func (m *AuthProvider) String() string { return proto.CompactTextString(m) }
func (*AuthProvider) ProtoMessage()    {}
func (*AuthProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_osprey_968553557d684e87, []int{5}
}
func (m *AuthProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthProvider.Unmarshal(m, b)
}
func (m *AuthProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthProvider.Marshal(b, m, deterministic)
}
func (dst *AuthProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthProvider.Merge(dst, src)
}
func (m *AuthProvider) XXX_Size() int {
	return xxx_messageInfo_AuthProvider.Size(m)
}
func (m *AuthProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthProvider.DiscardUnknown(m)
}

var xxx_messageInfo_AuthProvider proto.InternalMessageInfo

func (m *AuthProvider) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *AuthProvider) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *AuthProvider) GetIssuerURL() string {
	if m != nil {
		return m.IssuerURL
	}
	return ""
}

func (m *AuthProvider) GetIssuerCA() string {
	if m != nil {
		return m.IssuerCA
	}
	return ""
}

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_osprey_968553557d684e87, []int{6}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
	proto.RegisterType((*ClusterInfoRequest)(nil), "pb.ClusterInfoRequest")
	proto.RegisterType((*ClusterInfoResponse)(nil), "pb.ClusterInfoResponse")
	proto.RegisterType((*Cluster)(nil), "pb.Cluster")
	proto.RegisterType((*AuthProvider)(nil), "pb.AuthProvider")
	proto.RegisterType((*User)(nil), "pb.User")
}

func init() { proto.RegisterFile("osprey.proto", fileDescriptor_osprey_968553557d684e87) }

var fileDescriptor_osprey_968553557d684e87 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x95, 0x12, 0x68, 0x73, 0x0d, 0x08, 0x1d, 0x1d, 0x22, 0xd4, 0xa1, 0xb2, 0x84, 0xc4,
	0x80, 0x3a, 0xc0, 0xc2, 0xc0, 0x52, 0x95, 0xa5, 0x52, 0x07, 0xe4, 0xaa, 0x1f, 0xa0, 0x0d, 0x07,
	0x44, 0x14, 0xdb, 0xd8, 0x4e, 0x25, 0x56, 0x26, 0x3e, 0x36, 0xf2, 0x9f, 0x84, 0x74, 0x64, 0xf3,
	0xfb, 0xe5, 0xe9, 0xde, 0xbb, 0x53, 0x20, 0x97, 0x46, 0x69, 0xfa, 0x9a, 0x2a, 0x2d, 0xad, 0xc4,
	0x9e, 0xda, 0xb2, 0x33, 0xc8, 0x97, 0xf2, 0xb5, 0x12, 0x9c, 0x3e, 0x6b, 0x32, 0x96, 0x7d, 0x27,
	0x70, 0x1a, 0x81, 0x51, 0x52, 0x18, 0xc2, 0x2b, 0xe8, 0x97, 0xbb, 0xda, 0x58, 0xd2, 0x45, 0x32,
	0x49, 0xae, 0x87, 0xb7, 0xc3, 0xa9, 0xda, 0x4e, 0xe7, 0x01, 0xf1, 0xe6, 0x1b, 0xde, 0xc0, 0x40,
	0x69, 0xb9, 0xaf, 0x9e, 0x49, 0x17, 0x3d, 0xef, 0x3b, 0x77, 0xbe, 0x59, 0x6d, 0xdf, 0x9e, 0x22,
	0xe7, 0xad, 0x03, 0xc7, 0x90, 0xd6, 0x86, 0x74, 0x71, 0xe4, 0x9d, 0x03, 0xe7, 0x5c, 0x1b, 0xd2,
	0xdc, 0x53, 0x36, 0x02, 0x8c, 0xf3, 0x17, 0xe2, 0x45, 0x36, 0xd5, 0x1e, 0xe0, 0xe2, 0x80, 0xfe,
	0xab, 0x1f, 0x2b, 0xa1, 0x1f, 0x19, 0x22, 0xa4, 0x62, 0xf3, 0x41, 0xde, 0x9e, 0x71, 0xff, 0x46,
	0x06, 0xf9, 0x46, 0x55, 0x2b, 0xd2, 0x7b, 0xd2, 0x6b, 0xbe, 0xf4, 0x2b, 0x64, 0xfc, 0x80, 0xe1,
	0x04, 0x86, 0xad, 0x9e, 0xcf, 0x7c, 0xf7, 0x8c, 0x77, 0x11, 0xfb, 0x49, 0x20, 0xef, 0x6e, 0x8c,
	0x97, 0x30, 0x28, 0x77, 0x15, 0x09, 0xbb, 0x78, 0x8c, 0x71, 0xad, 0x76, 0x91, 0xe1, 0xbd, 0xa2,
	0x52, 0x93, 0x6d, 0x22, 0xbb, 0x0c, 0xc7, 0x90, 0x55, 0xc6, 0xd4, 0xa1, 0x53, 0x08, 0xfc, 0x03,
	0x6e, 0x7a, 0x10, 0xf3, 0x59, 0x91, 0x86, 0xe9, 0x8d, 0x66, 0xf7, 0x90, 0xba, 0x8b, 0x3a, 0x8f,
	0xbb, 0x69, 0x67, 0xe1, 0x56, 0xe3, 0x08, 0x8e, 0xad, 0x7c, 0x27, 0x11, 0xa3, 0x83, 0xd8, 0x9e,
	0xf8, 0xbf, 0xe3, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x33, 0xa5, 0xcf, 0x06, 0x2d, 0x02, 0x00,
	0x00,
}
