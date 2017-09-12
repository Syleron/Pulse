// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/pulse.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/pulse.proto

It has these top-level messages:
	HealthCheckRequest
	HealthCheckResponse
	PulseJoin
	PulseLeave
	PulseCreate
	PulseGroupNew
	PulseGroupDelete
	PulseGroupAdd
	PulseGroupRemove
	PulseGroupAssign
	PulseGroupUnassign
	PulseGroupList
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheckRequest_ServingRequest int32

const (
	HealthCheckRequest_SETUP  HealthCheckRequest_ServingRequest = 0
	HealthCheckRequest_STATUS HealthCheckRequest_ServingRequest = 1
)

var HealthCheckRequest_ServingRequest_name = map[int32]string{
	0: "SETUP",
	1: "STATUS",
}
var HealthCheckRequest_ServingRequest_value = map[string]int32{
	"SETUP":  0,
	"STATUS": 1,
}

func (x HealthCheckRequest_ServingRequest) String() string {
	return proto1.EnumName(HealthCheckRequest_ServingRequest_name, int32(x))
}
func (HealthCheckRequest_ServingRequest) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN      HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_UNCONFIGURED HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_CONFIGURED   HealthCheckResponse_ServingStatus = 2
	HealthCheckResponse_FAILVER      HealthCheckResponse_ServingStatus = 3
	HealthCheckResponse_HEALTHY      HealthCheckResponse_ServingStatus = 4
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "UNCONFIGURED",
	2: "CONFIGURED",
	3: "FAILVER",
	4: "HEALTHY",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":      0,
	"UNCONFIGURED": 1,
	"CONFIGURED":   2,
	"FAILVER":      3,
	"HEALTHY":      4,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto1.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

// Types
type HealthCheckRequest struct {
	Request HealthCheckRequest_ServingRequest `protobuf:"varint,1,opt,name=request,enum=proto.HealthCheckRequest_ServingRequest" json:"request,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto1.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HealthCheckRequest) GetRequest() HealthCheckRequest_ServingRequest {
	if m != nil {
		return m.Request
	}
	return HealthCheckRequest_SETUP
}

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=proto.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto1.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

// Pulse Cluster Messages
type PulseJoin struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *PulseJoin) Reset()                    { *m = PulseJoin{} }
func (m *PulseJoin) String() string            { return proto1.CompactTextString(m) }
func (*PulseJoin) ProtoMessage()               {}
func (*PulseJoin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PulseJoin) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseJoin) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type PulseLeave struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PulseLeave) Reset()                    { *m = PulseLeave{} }
func (m *PulseLeave) String() string            { return proto1.CompactTextString(m) }
func (*PulseLeave) ProtoMessage()               {}
func (*PulseLeave) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PulseLeave) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseLeave) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PulseCreate struct {
	Success  bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	BindIp   string `protobuf:"bytes,3,opt,name=bind_ip,json=bindIp" json:"bind_ip,omitempty"`
	BindPort string `protobuf:"bytes,4,opt,name=bind_port,json=bindPort" json:"bind_port,omitempty"`
}

func (m *PulseCreate) Reset()                    { *m = PulseCreate{} }
func (m *PulseCreate) String() string            { return proto1.CompactTextString(m) }
func (*PulseCreate) ProtoMessage()               {}
func (*PulseCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PulseCreate) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseCreate) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PulseCreate) GetBindIp() string {
	if m != nil {
		return m.BindIp
	}
	return ""
}

func (m *PulseCreate) GetBindPort() string {
	if m != nil {
		return m.BindPort
	}
	return ""
}

// Pulse Group Messages
type PulseGroupNew struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PulseGroupNew) Reset()                    { *m = PulseGroupNew{} }
func (m *PulseGroupNew) String() string            { return proto1.CompactTextString(m) }
func (*PulseGroupNew) ProtoMessage()               {}
func (*PulseGroupNew) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PulseGroupNew) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseGroupNew) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PulseGroupDelete struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *PulseGroupDelete) Reset()                    { *m = PulseGroupDelete{} }
func (m *PulseGroupDelete) String() string            { return proto1.CompactTextString(m) }
func (*PulseGroupDelete) ProtoMessage()               {}
func (*PulseGroupDelete) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PulseGroupDelete) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseGroupDelete) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PulseGroupDelete) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PulseGroupAdd struct {
	Success bool     `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Name    string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Ips     []string `protobuf:"bytes,4,rep,name=ips" json:"ips,omitempty"`
}

func (m *PulseGroupAdd) Reset()                    { *m = PulseGroupAdd{} }
func (m *PulseGroupAdd) String() string            { return proto1.CompactTextString(m) }
func (*PulseGroupAdd) ProtoMessage()               {}
func (*PulseGroupAdd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PulseGroupAdd) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseGroupAdd) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PulseGroupAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PulseGroupAdd) GetIps() []string {
	if m != nil {
		return m.Ips
	}
	return nil
}

type PulseGroupRemove struct {
	Success bool     `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Name    string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Ips     []string `protobuf:"bytes,4,rep,name=ips" json:"ips,omitempty"`
}

func (m *PulseGroupRemove) Reset()                    { *m = PulseGroupRemove{} }
func (m *PulseGroupRemove) String() string            { return proto1.CompactTextString(m) }
func (*PulseGroupRemove) ProtoMessage()               {}
func (*PulseGroupRemove) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PulseGroupRemove) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseGroupRemove) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PulseGroupRemove) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PulseGroupRemove) GetIps() []string {
	if m != nil {
		return m.Ips
	}
	return nil
}

type PulseGroupAssign struct {
	Success   bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Group     string `protobuf:"bytes,3,opt,name=group" json:"group,omitempty"`
	Interface string `protobuf:"bytes,4,opt,name=interface" json:"interface,omitempty"`
	Node      string `protobuf:"bytes,5,opt,name=node" json:"node,omitempty"`
}

func (m *PulseGroupAssign) Reset()                    { *m = PulseGroupAssign{} }
func (m *PulseGroupAssign) String() string            { return proto1.CompactTextString(m) }
func (*PulseGroupAssign) ProtoMessage()               {}
func (*PulseGroupAssign) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PulseGroupAssign) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseGroupAssign) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PulseGroupAssign) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *PulseGroupAssign) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *PulseGroupAssign) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type PulseGroupUnassign struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *PulseGroupUnassign) Reset()                    { *m = PulseGroupUnassign{} }
func (m *PulseGroupUnassign) String() string            { return proto1.CompactTextString(m) }
func (*PulseGroupUnassign) ProtoMessage()               {}
func (*PulseGroupUnassign) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PulseGroupUnassign) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseGroupUnassign) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PulseGroupUnassign) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PulseGroupList struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PulseGroupList) Reset()                    { *m = PulseGroupList{} }
func (m *PulseGroupList) String() string            { return proto1.CompactTextString(m) }
func (*PulseGroupList) ProtoMessage()               {}
func (*PulseGroupList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PulseGroupList) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PulseGroupList) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*HealthCheckRequest)(nil), "proto.HealthCheckRequest")
	proto1.RegisterType((*HealthCheckResponse)(nil), "proto.HealthCheckResponse")
	proto1.RegisterType((*PulseJoin)(nil), "proto.PulseJoin")
	proto1.RegisterType((*PulseLeave)(nil), "proto.PulseLeave")
	proto1.RegisterType((*PulseCreate)(nil), "proto.PulseCreate")
	proto1.RegisterType((*PulseGroupNew)(nil), "proto.PulseGroupNew")
	proto1.RegisterType((*PulseGroupDelete)(nil), "proto.PulseGroupDelete")
	proto1.RegisterType((*PulseGroupAdd)(nil), "proto.PulseGroupAdd")
	proto1.RegisterType((*PulseGroupRemove)(nil), "proto.PulseGroupRemove")
	proto1.RegisterType((*PulseGroupAssign)(nil), "proto.PulseGroupAssign")
	proto1.RegisterType((*PulseGroupUnassign)(nil), "proto.PulseGroupUnassign")
	proto1.RegisterType((*PulseGroupList)(nil), "proto.PulseGroupList")
	proto1.RegisterEnum("proto.HealthCheckRequest_ServingRequest", HealthCheckRequest_ServingRequest_name, HealthCheckRequest_ServingRequest_value)
	proto1.RegisterEnum("proto.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Requester service

type RequesterClient interface {
	// Perform GRPC Health Check
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	// Join Cluster
	Join(ctx context.Context, in *PulseJoin, opts ...grpc.CallOption) (*PulseJoin, error)
	// Leave Cluster
	Leave(ctx context.Context, in *PulseLeave, opts ...grpc.CallOption) (*PulseLeave, error)
	// Create Cluster
	Create(ctx context.Context, in *PulseCreate, opts ...grpc.CallOption) (*PulseCreate, error)
	// Create floating ip group
	NewGroup(ctx context.Context, in *PulseGroupNew, opts ...grpc.CallOption) (*PulseGroupNew, error)
	// Delete floating ip group
	DeleteGroup(ctx context.Context, in *PulseGroupDelete, opts ...grpc.CallOption) (*PulseGroupDelete, error)
	// Add floating IP
	GroupIPAdd(ctx context.Context, in *PulseGroupAdd, opts ...grpc.CallOption) (*PulseGroupAdd, error)
	// Remove floating IP
	GroupIPRemove(ctx context.Context, in *PulseGroupRemove, opts ...grpc.CallOption) (*PulseGroupRemove, error)
	// Assign a group
	GroupAssign(ctx context.Context, in *PulseGroupAssign, opts ...grpc.CallOption) (*PulseGroupAssign, error)
	// Unassign a group
	GroupUnassign(ctx context.Context, in *PulseGroupUnassign, opts ...grpc.CallOption) (*PulseGroupUnassign, error)
	// Get group list
	GroupList(ctx context.Context, in *PulseGroupList, opts ...grpc.CallOption) (*PulseGroupList, error)
}

type requesterClient struct {
	cc *grpc.ClientConn
}

func NewRequesterClient(cc *grpc.ClientConn) RequesterClient {
	return &requesterClient{cc}
}

func (c *requesterClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/proto.Requester/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) Join(ctx context.Context, in *PulseJoin, opts ...grpc.CallOption) (*PulseJoin, error) {
	out := new(PulseJoin)
	err := grpc.Invoke(ctx, "/proto.Requester/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) Leave(ctx context.Context, in *PulseLeave, opts ...grpc.CallOption) (*PulseLeave, error) {
	out := new(PulseLeave)
	err := grpc.Invoke(ctx, "/proto.Requester/Leave", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) Create(ctx context.Context, in *PulseCreate, opts ...grpc.CallOption) (*PulseCreate, error) {
	out := new(PulseCreate)
	err := grpc.Invoke(ctx, "/proto.Requester/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) NewGroup(ctx context.Context, in *PulseGroupNew, opts ...grpc.CallOption) (*PulseGroupNew, error) {
	out := new(PulseGroupNew)
	err := grpc.Invoke(ctx, "/proto.Requester/NewGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) DeleteGroup(ctx context.Context, in *PulseGroupDelete, opts ...grpc.CallOption) (*PulseGroupDelete, error) {
	out := new(PulseGroupDelete)
	err := grpc.Invoke(ctx, "/proto.Requester/DeleteGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) GroupIPAdd(ctx context.Context, in *PulseGroupAdd, opts ...grpc.CallOption) (*PulseGroupAdd, error) {
	out := new(PulseGroupAdd)
	err := grpc.Invoke(ctx, "/proto.Requester/GroupIPAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) GroupIPRemove(ctx context.Context, in *PulseGroupRemove, opts ...grpc.CallOption) (*PulseGroupRemove, error) {
	out := new(PulseGroupRemove)
	err := grpc.Invoke(ctx, "/proto.Requester/GroupIPRemove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) GroupAssign(ctx context.Context, in *PulseGroupAssign, opts ...grpc.CallOption) (*PulseGroupAssign, error) {
	out := new(PulseGroupAssign)
	err := grpc.Invoke(ctx, "/proto.Requester/GroupAssign", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) GroupUnassign(ctx context.Context, in *PulseGroupUnassign, opts ...grpc.CallOption) (*PulseGroupUnassign, error) {
	out := new(PulseGroupUnassign)
	err := grpc.Invoke(ctx, "/proto.Requester/GroupUnassign", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requesterClient) GroupList(ctx context.Context, in *PulseGroupList, opts ...grpc.CallOption) (*PulseGroupList, error) {
	out := new(PulseGroupList)
	err := grpc.Invoke(ctx, "/proto.Requester/GroupList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Requester service

type RequesterServer interface {
	// Perform GRPC Health Check
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	// Join Cluster
	Join(context.Context, *PulseJoin) (*PulseJoin, error)
	// Leave Cluster
	Leave(context.Context, *PulseLeave) (*PulseLeave, error)
	// Create Cluster
	Create(context.Context, *PulseCreate) (*PulseCreate, error)
	// Create floating ip group
	NewGroup(context.Context, *PulseGroupNew) (*PulseGroupNew, error)
	// Delete floating ip group
	DeleteGroup(context.Context, *PulseGroupDelete) (*PulseGroupDelete, error)
	// Add floating IP
	GroupIPAdd(context.Context, *PulseGroupAdd) (*PulseGroupAdd, error)
	// Remove floating IP
	GroupIPRemove(context.Context, *PulseGroupRemove) (*PulseGroupRemove, error)
	// Assign a group
	GroupAssign(context.Context, *PulseGroupAssign) (*PulseGroupAssign, error)
	// Unassign a group
	GroupUnassign(context.Context, *PulseGroupUnassign) (*PulseGroupUnassign, error)
	// Get group list
	GroupList(context.Context, *PulseGroupList) (*PulseGroupList, error)
}

func RegisterRequesterServer(s *grpc.Server, srv RequesterServer) {
	s.RegisterService(&_Requester_serviceDesc, srv)
}

func _Requester_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseJoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).Join(ctx, req.(*PulseJoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseLeave)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).Leave(ctx, req.(*PulseLeave))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).Create(ctx, req.(*PulseCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_NewGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseGroupNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).NewGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/NewGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).NewGroup(ctx, req.(*PulseGroupNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseGroupDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).DeleteGroup(ctx, req.(*PulseGroupDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_GroupIPAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseGroupAdd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).GroupIPAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/GroupIPAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).GroupIPAdd(ctx, req.(*PulseGroupAdd))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_GroupIPRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseGroupRemove)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).GroupIPRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/GroupIPRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).GroupIPRemove(ctx, req.(*PulseGroupRemove))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_GroupAssign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseGroupAssign)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).GroupAssign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/GroupAssign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).GroupAssign(ctx, req.(*PulseGroupAssign))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_GroupUnassign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseGroupUnassign)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).GroupUnassign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/GroupUnassign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).GroupUnassign(ctx, req.(*PulseGroupUnassign))
	}
	return interceptor(ctx, in, info, handler)
}

func _Requester_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PulseGroupList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequesterServer).GroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Requester/GroupList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequesterServer).GroupList(ctx, req.(*PulseGroupList))
	}
	return interceptor(ctx, in, info, handler)
}

var _Requester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Requester",
	HandlerType: (*RequesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Requester_Check_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Requester_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Requester_Leave_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Requester_Create_Handler,
		},
		{
			MethodName: "NewGroup",
			Handler:    _Requester_NewGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _Requester_DeleteGroup_Handler,
		},
		{
			MethodName: "GroupIPAdd",
			Handler:    _Requester_GroupIPAdd_Handler,
		},
		{
			MethodName: "GroupIPRemove",
			Handler:    _Requester_GroupIPRemove_Handler,
		},
		{
			MethodName: "GroupAssign",
			Handler:    _Requester_GroupAssign_Handler,
		},
		{
			MethodName: "GroupUnassign",
			Handler:    _Requester_GroupUnassign_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _Requester_GroupList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pulse.proto",
}

func init() { proto1.RegisterFile("proto/pulse.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xad, 0x9b, 0x38, 0xa9, 0xa7, 0xbf, 0x46, 0xee, 0xfc, 0x8a, 0x9a, 0x1a, 0x0e, 0x95, 0x2f,
	0x44, 0x48, 0x14, 0x54, 0x24, 0x04, 0x17, 0x5a, 0x93, 0xa6, 0x6d, 0x20, 0x72, 0x23, 0x27, 0xe1,
	0x9f, 0x90, 0x90, 0x1b, 0x0f, 0xa9, 0x45, 0x6a, 0x1b, 0xaf, 0xd3, 0x9e, 0xf9, 0x02, 0x48, 0x7c,
	0x14, 0xbe, 0x21, 0xf2, 0xee, 0xba, 0x71, 0xe5, 0xa4, 0x07, 0x0b, 0x4e, 0xde, 0x79, 0x6f, 0xe6,
	0xcd, 0xb3, 0xbd, 0x33, 0xb0, 0x19, 0xc5, 0x61, 0x12, 0x3e, 0x89, 0x66, 0x53, 0x46, 0x7b, 0xfc,
	0x8c, 0x2a, 0x7f, 0x98, 0x3f, 0x14, 0xc0, 0x53, 0x72, 0xa7, 0xc9, 0x45, 0xfb, 0x82, 0xc6, 0xdf,
	0x1c, 0xfa, 0x3e, 0x23, 0x96, 0xe0, 0x6b, 0xa8, 0xc7, 0xe2, 0xd8, 0x54, 0x76, 0x95, 0x56, 0x63,
	0xbf, 0x25, 0xca, 0xf6, 0x8a, 0xb9, 0x7b, 0x03, 0x8a, 0xaf, 0xfc, 0x60, 0x22, 0x43, 0x27, 0x2b,
	0x34, 0x1f, 0x42, 0xe3, 0x36, 0x85, 0x1a, 0xa8, 0x83, 0xce, 0x70, 0xd4, 0xd7, 0x57, 0x10, 0xa0,
	0x36, 0x18, 0x5a, 0xc3, 0xd1, 0x40, 0x57, 0xcc, 0xdf, 0x0a, 0xfc, 0x7f, 0x4b, 0x97, 0x45, 0x61,
	0xc0, 0x08, 0x0f, 0xa1, 0xc6, 0x12, 0x37, 0x99, 0xb1, 0xbb, 0x3c, 0x88, 0xdc, 0xcc, 0xc4, 0x80,
	0xe7, 0x3b, 0xb2, 0xce, 0xfc, 0x00, 0x1b, 0xb7, 0x08, 0x5c, 0x87, 0xfa, 0xc8, 0x7e, 0x6b, 0x9f,
	0xbd, 0xb7, 0xf5, 0x15, 0xd4, 0xe1, 0xbf, 0x91, 0xdd, 0x3e, 0xb3, 0x8f, 0xbb, 0x27, 0x23, 0xa7,
	0x73, 0xa4, 0x2b, 0xd8, 0x00, 0xc8, 0xc5, 0xab, 0x69, 0xfa, 0xb1, 0xd5, 0xed, 0xbd, 0xeb, 0x38,
	0x7a, 0x25, 0x0d, 0x4e, 0x3b, 0x56, 0x6f, 0x78, 0xfa, 0x51, 0xaf, 0x9a, 0x07, 0xa0, 0xf5, 0xd3,
	0xaf, 0xf9, 0x26, 0xf4, 0x03, 0x6c, 0x42, 0x9d, 0xcd, 0xc6, 0x63, 0x62, 0xc2, 0xe9, 0x9a, 0x93,
	0x85, 0x29, 0xe3, 0x7a, 0x5e, 0x9c, 0x32, 0xab, 0xbb, 0x4a, 0x4b, 0x73, 0xb2, 0xd0, 0x3c, 0x04,
	0xe0, 0x02, 0x3d, 0x72, 0xaf, 0xe8, 0x6e, 0x85, 0x4b, 0x62, 0xcc, 0x9d, 0x50, 0xa6, 0x20, 0x43,
	0xf3, 0x1a, 0xd6, 0xb9, 0x42, 0x3b, 0x26, 0x37, 0x29, 0x25, 0x81, 0xdb, 0x50, 0x3f, 0xf7, 0x03,
	0xef, 0x8b, 0x1f, 0x35, 0x2b, 0x9c, 0xa9, 0xa5, 0x61, 0x37, 0xc2, 0xfb, 0xa0, 0x71, 0x22, 0x0a,
	0xe3, 0xa4, 0x59, 0xe5, 0xd4, 0x5a, 0x0a, 0xf4, 0xc3, 0x38, 0x31, 0xdb, 0xb0, 0xc1, 0x1b, 0x9f,
	0xc4, 0xe1, 0x2c, 0xb2, 0xe9, 0xba, 0x94, 0xfb, 0x4f, 0xa0, 0xcf, 0x45, 0x8e, 0x68, 0x4a, 0x25,
	0x5f, 0x01, 0xa1, 0x1a, 0xb8, 0x97, 0x24, 0xfd, 0xf3, 0xb3, 0xe9, 0xe7, 0x0d, 0x5a, 0x9e, 0xf7,
	0xb7, 0x84, 0x51, 0x87, 0x8a, 0x1f, 0xb1, 0x66, 0x75, 0xb7, 0xd2, 0xd2, 0x9c, 0xf4, 0x68, 0x4e,
	0xf3, 0xaf, 0xe1, 0xd0, 0x65, 0x78, 0x45, 0xff, 0xb0, 0xdb, 0x4f, 0x25, 0xdf, 0xce, 0x62, 0xcc,
	0x9f, 0x04, 0xa5, 0xda, 0x6d, 0x81, 0x3a, 0x49, 0x25, 0x64, 0x3f, 0x11, 0xe0, 0x03, 0xd0, 0xfc,
	0x20, 0xa1, 0xf8, 0xab, 0x3b, 0x26, 0xf9, 0xd7, 0xe7, 0x00, 0xb7, 0x18, 0x7a, 0xd4, 0x54, 0xa5,
	0xc5, 0xd0, 0x23, 0xf3, 0x33, 0xe0, 0xdc, 0xcf, 0x28, 0x70, 0xcb, 0x3b, 0x5a, 0xf4, 0x1f, 0x8f,
	0xa0, 0x31, 0x57, 0xef, 0xf9, 0x2c, 0x29, 0xa3, 0xbc, 0xff, 0x4b, 0x05, 0x4d, 0x6e, 0x20, 0x8a,
	0xf1, 0x15, 0xa8, 0x7c, 0x73, 0xe0, 0xce, 0xd2, 0x8d, 0x66, 0x18, 0xcb, 0x17, 0x0d, 0x3e, 0x82,
	0x2a, 0x9f, 0x79, 0x5d, 0xe6, 0xdc, 0x6c, 0x01, 0xa3, 0x80, 0xe0, 0x63, 0x50, 0xc5, 0x78, 0x6f,
	0xe6, 0x29, 0x0e, 0x19, 0x45, 0x08, 0x9f, 0x42, 0x4d, 0xce, 0x32, 0xe6, 0x49, 0x81, 0x19, 0x0b,
	0x30, 0x7c, 0x0e, 0x6b, 0x36, 0x5d, 0xf3, 0xcf, 0x83, 0x5b, 0x79, 0x3e, 0x1b, 0x4d, 0x63, 0x21,
	0x8a, 0x07, 0xb0, 0x2e, 0x46, 0x4e, 0x94, 0x6e, 0x17, 0x92, 0x04, 0x6b, 0x2c, 0x23, 0xf0, 0x05,
	0x00, 0x0f, 0xbb, 0xfd, 0x74, 0xbc, 0x8a, 0x4d, 0x2c, 0xcf, 0x33, 0x16, 0xa2, 0x68, 0xc1, 0x86,
	0xac, 0x94, 0xd3, 0x52, 0xec, 0x21, 0x08, 0x63, 0x19, 0x91, 0xba, 0xcf, 0xdf, 0xff, 0x62, 0x9e,
	0x20, 0x8c, 0x65, 0x04, 0x76, 0xa4, 0x87, 0x9b, 0x0b, 0xbb, 0x53, 0xc8, 0xcc, 0x28, 0x63, 0x39,
	0x85, 0x2f, 0x41, 0x9b, 0xdf, 0xcc, 0x7b, 0x85, 0xbc, 0x14, 0x36, 0x16, 0xc3, 0xe7, 0x35, 0x8e,
	0x3e, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x08, 0x7a, 0x8e, 0x39, 0x99, 0x07, 0x00, 0x00,
}