// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frontgate.proto

/*
Package openpitrix_frontgate is a generated protocol buffer package.

It is generated from these files:
	frontgate.proto

It has these top-level messages:
	Info
	SubTaskResult
	GetConfdInfoRequest
	ConfdInfo
	StartConfdRequest
	StopConfdRequest
	RegisterMetadataRequest
	DeregisterMetadataRequest
	RegisterCmdRequest
	DeregisterCmdRequest
	GetSubTaskResultRequest
	Empty
*/
package openpitrix_frontgate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import openpitrix_drone "openpitrix.io/openpitrix/pkg/pb/drone"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Info struct {
	FrontgateId     string `protobuf:"bytes,1,opt,name=frontgate_id,json=frontgateId" json:"frontgate_id,omitempty"`
	FrontgateNodeId string `protobuf:"bytes,2,opt,name=frontgate_node_id,json=frontgateNodeId" json:"frontgate_node_id,omitempty"`
	FrontgateNodeIp string `protobuf:"bytes,3,opt,name=frontgate_node_ip,json=frontgateNodeIp" json:"frontgate_node_ip,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Info) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *Info) GetFrontgateNodeId() string {
	if m != nil {
		return m.FrontgateNodeId
	}
	return ""
}

func (m *Info) GetFrontgateNodeIp() string {
	if m != nil {
		return m.FrontgateNodeIp
	}
	return ""
}

type SubTaskResult struct {
	DroneIp   string `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	SubtaskId string `protobuf:"bytes,2,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *SubTaskResult) Reset()                    { *m = SubTaskResult{} }
func (m *SubTaskResult) String() string            { return proto.CompactTextString(m) }
func (*SubTaskResult) ProtoMessage()               {}
func (*SubTaskResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubTaskResult) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTaskResult) GetSubtaskId() string {
	if m != nil {
		return m.SubtaskId
	}
	return ""
}

func (m *SubTaskResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetConfdInfoRequest struct {
	DroneIp string `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
}

func (m *GetConfdInfoRequest) Reset()                    { *m = GetConfdInfoRequest{} }
func (m *GetConfdInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetConfdInfoRequest) ProtoMessage()               {}
func (*GetConfdInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetConfdInfoRequest) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

type ConfdInfo struct {
	DroneIp string `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *ConfdInfo) Reset()                    { *m = ConfdInfo{} }
func (m *ConfdInfo) String() string            { return proto.CompactTextString(m) }
func (*ConfdInfo) ProtoMessage()               {}
func (*ConfdInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ConfdInfo) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *ConfdInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type StartConfdRequest struct {
	DroneIp            string                               `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	TimeoutSecond      int32                                `protobuf:"varint,2,opt,name=timeout_second,json=timeoutSecond" json:"timeout_second,omitempty"`
	ConfdConfig        *openpitrix_drone.ConfdConfig        `protobuf:"bytes,3,opt,name=confd_config,json=confdConfig" json:"confd_config,omitempty"`
	ConfdBackendConfig *openpitrix_drone.ConfdBackendConfig `protobuf:"bytes,4,opt,name=confd_backend_config,json=confdBackendConfig" json:"confd_backend_config,omitempty"`
}

func (m *StartConfdRequest) Reset()                    { *m = StartConfdRequest{} }
func (m *StartConfdRequest) String() string            { return proto.CompactTextString(m) }
func (*StartConfdRequest) ProtoMessage()               {}
func (*StartConfdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StartConfdRequest) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *StartConfdRequest) GetTimeoutSecond() int32 {
	if m != nil {
		return m.TimeoutSecond
	}
	return 0
}

func (m *StartConfdRequest) GetConfdConfig() *openpitrix_drone.ConfdConfig {
	if m != nil {
		return m.ConfdConfig
	}
	return nil
}

func (m *StartConfdRequest) GetConfdBackendConfig() *openpitrix_drone.ConfdBackendConfig {
	if m != nil {
		return m.ConfdBackendConfig
	}
	return nil
}

type StopConfdRequest struct {
	DroneIp       string `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	TimeoutSecond int32  `protobuf:"varint,2,opt,name=timeout_second,json=timeoutSecond" json:"timeout_second,omitempty"`
}

func (m *StopConfdRequest) Reset()                    { *m = StopConfdRequest{} }
func (m *StopConfdRequest) String() string            { return proto.CompactTextString(m) }
func (*StopConfdRequest) ProtoMessage()               {}
func (*StopConfdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StopConfdRequest) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *StopConfdRequest) GetTimeoutSecond() int32 {
	if m != nil {
		return m.TimeoutSecond
	}
	return 0
}

type RegisterMetadataRequest struct {
	DroneIp       string `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	FrontgateId   string `protobuf:"bytes,2,opt,name=frontgate_id,json=frontgateId" json:"frontgate_id,omitempty"`
	SubtaskId     string `protobuf:"bytes,3,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
	TimeoutSecond int32  `protobuf:"varint,4,opt,name=timeout_second,json=timeoutSecond" json:"timeout_second,omitempty"`
	Retry         int32  `protobuf:"varint,5,opt,name=retry" json:"retry,omitempty"`
}

func (m *RegisterMetadataRequest) Reset()                    { *m = RegisterMetadataRequest{} }
func (m *RegisterMetadataRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterMetadataRequest) ProtoMessage()               {}
func (*RegisterMetadataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RegisterMetadataRequest) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *RegisterMetadataRequest) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *RegisterMetadataRequest) GetSubtaskId() string {
	if m != nil {
		return m.SubtaskId
	}
	return ""
}

func (m *RegisterMetadataRequest) GetTimeoutSecond() int32 {
	if m != nil {
		return m.TimeoutSecond
	}
	return 0
}

func (m *RegisterMetadataRequest) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

type DeregisterMetadataRequest struct {
	DroneIp       string `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	FrontgateId   string `protobuf:"bytes,2,opt,name=frontgate_id,json=frontgateId" json:"frontgate_id,omitempty"`
	SubtaskId     string `protobuf:"bytes,3,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
	TimeoutSecond int32  `protobuf:"varint,4,opt,name=timeout_second,json=timeoutSecond" json:"timeout_second,omitempty"`
	Retry         int32  `protobuf:"varint,5,opt,name=retry" json:"retry,omitempty"`
}

func (m *DeregisterMetadataRequest) Reset()                    { *m = DeregisterMetadataRequest{} }
func (m *DeregisterMetadataRequest) String() string            { return proto.CompactTextString(m) }
func (*DeregisterMetadataRequest) ProtoMessage()               {}
func (*DeregisterMetadataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeregisterMetadataRequest) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *DeregisterMetadataRequest) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *DeregisterMetadataRequest) GetSubtaskId() string {
	if m != nil {
		return m.SubtaskId
	}
	return ""
}

func (m *DeregisterMetadataRequest) GetTimeoutSecond() int32 {
	if m != nil {
		return m.TimeoutSecond
	}
	return 0
}

func (m *DeregisterMetadataRequest) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

type RegisterCmdRequest struct {
	DroneIp       string `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	FrontgateId   string `protobuf:"bytes,2,opt,name=frontgate_id,json=frontgateId" json:"frontgate_id,omitempty"`
	SubtaskId     string `protobuf:"bytes,3,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
	TimeoutSecond int32  `protobuf:"varint,4,opt,name=timeout_second,json=timeoutSecond" json:"timeout_second,omitempty"`
	Retry         int32  `protobuf:"varint,5,opt,name=retry" json:"retry,omitempty"`
}

func (m *RegisterCmdRequest) Reset()                    { *m = RegisterCmdRequest{} }
func (m *RegisterCmdRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterCmdRequest) ProtoMessage()               {}
func (*RegisterCmdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RegisterCmdRequest) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *RegisterCmdRequest) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *RegisterCmdRequest) GetSubtaskId() string {
	if m != nil {
		return m.SubtaskId
	}
	return ""
}

func (m *RegisterCmdRequest) GetTimeoutSecond() int32 {
	if m != nil {
		return m.TimeoutSecond
	}
	return 0
}

func (m *RegisterCmdRequest) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

type DeregisterCmdRequest struct {
	DroneIp       string `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	FrontgateId   string `protobuf:"bytes,2,opt,name=frontgate_id,json=frontgateId" json:"frontgate_id,omitempty"`
	SubtaskId     string `protobuf:"bytes,3,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
	TimeoutSecond int32  `protobuf:"varint,4,opt,name=timeout_second,json=timeoutSecond" json:"timeout_second,omitempty"`
	Retry         int32  `protobuf:"varint,5,opt,name=retry" json:"retry,omitempty"`
}

func (m *DeregisterCmdRequest) Reset()                    { *m = DeregisterCmdRequest{} }
func (m *DeregisterCmdRequest) String() string            { return proto.CompactTextString(m) }
func (*DeregisterCmdRequest) ProtoMessage()               {}
func (*DeregisterCmdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeregisterCmdRequest) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *DeregisterCmdRequest) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *DeregisterCmdRequest) GetSubtaskId() string {
	if m != nil {
		return m.SubtaskId
	}
	return ""
}

func (m *DeregisterCmdRequest) GetTimeoutSecond() int32 {
	if m != nil {
		return m.TimeoutSecond
	}
	return 0
}

func (m *DeregisterCmdRequest) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

type GetSubTaskResultRequest struct {
	DroneIp   string `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	SubtaskId string `protobuf:"bytes,2,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
}

func (m *GetSubTaskResultRequest) Reset()                    { *m = GetSubTaskResultRequest{} }
func (m *GetSubTaskResultRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSubTaskResultRequest) ProtoMessage()               {}
func (*GetSubTaskResultRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetSubTaskResultRequest) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *GetSubTaskResultRequest) GetSubtaskId() string {
	if m != nil {
		return m.SubtaskId
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*Info)(nil), "openpitrix.frontgate.Info")
	proto.RegisterType((*SubTaskResult)(nil), "openpitrix.frontgate.SubTaskResult")
	proto.RegisterType((*GetConfdInfoRequest)(nil), "openpitrix.frontgate.GetConfdInfoRequest")
	proto.RegisterType((*ConfdInfo)(nil), "openpitrix.frontgate.ConfdInfo")
	proto.RegisterType((*StartConfdRequest)(nil), "openpitrix.frontgate.StartConfdRequest")
	proto.RegisterType((*StopConfdRequest)(nil), "openpitrix.frontgate.StopConfdRequest")
	proto.RegisterType((*RegisterMetadataRequest)(nil), "openpitrix.frontgate.RegisterMetadataRequest")
	proto.RegisterType((*DeregisterMetadataRequest)(nil), "openpitrix.frontgate.DeregisterMetadataRequest")
	proto.RegisterType((*RegisterCmdRequest)(nil), "openpitrix.frontgate.RegisterCmdRequest")
	proto.RegisterType((*DeregisterCmdRequest)(nil), "openpitrix.frontgate.DeregisterCmdRequest")
	proto.RegisterType((*GetSubTaskResultRequest)(nil), "openpitrix.frontgate.GetSubTaskResultRequest")
	proto.RegisterType((*Empty)(nil), "openpitrix.frontgate.Empty")
}

func init() { proto.RegisterFile("frontgate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xb6, 0x75, 0xa5, 0xa7, 0x1d, 0x74, 0x5e, 0xc5, 0xb6, 0xa2, 0x31, 0x08, 0x7f, 0x63,
	0x12, 0x2d, 0x1a, 0xf7, 0x08, 0xad, 0xc0, 0xe8, 0x05, 0x08, 0xa5, 0x13, 0x20, 0x2e, 0x98, 0xdc,
	0xe6, 0xb4, 0x44, 0x6b, 0x63, 0x63, 0x3b, 0xc0, 0x2e, 0x79, 0x21, 0x84, 0x84, 0x78, 0x16, 0x1e,
	0x83, 0x57, 0x40, 0x75, 0xb2, 0x64, 0x49, 0x13, 0xaf, 0x17, 0x5c, 0xb0, 0x9b, 0x48, 0x3e, 0xfe,
	0xce, 0xf7, 0x9d, 0x63, 0x1f, 0x7f, 0x2d, 0x5c, 0x19, 0x0a, 0xe6, 0xab, 0x11, 0x55, 0xd8, 0xe2,
	0x82, 0x29, 0x46, 0x1a, 0x8c, 0xa3, 0xcf, 0x3d, 0x25, 0xbc, 0xaf, 0xad, 0x78, 0xaf, 0xb9, 0x3d,
	0x62, 0x6c, 0x34, 0xc6, 0xb6, 0xc6, 0xf4, 0x83, 0x61, 0x5b, 0x79, 0x13, 0x94, 0x8a, 0x4e, 0x78,
	0x98, 0xd6, 0xbc, 0x9e, 0x05, 0x7c, 0x11, 0x94, 0x73, 0x14, 0x32, 0xda, 0x5f, 0x75, 0x05, 0xf3,
	0xb1, 0xad, 0xbf, 0x61, 0xc8, 0xfe, 0x66, 0xc1, 0x52, 0xd7, 0x1f, 0x32, 0x72, 0x13, 0x6a, 0xb1,
	0xd2, 0x91, 0xe7, 0x6e, 0x58, 0x37, 0xac, 0x9d, 0x8a, 0x53, 0x8d, 0x63, 0x5d, 0x97, 0xec, 0xc2,
	0x6a, 0x02, 0xf1, 0x99, 0xab, 0x71, 0x0b, 0x1a, 0x97, 0x74, 0xf0, 0x8a, 0xb9, 0x05, 0x58, 0xbe,
	0xb1, 0x98, 0x87, 0xe5, 0x36, 0x85, 0x95, 0x5e, 0xd0, 0x3f, 0xa4, 0xf2, 0xd8, 0x41, 0x19, 0x8c,
	0x15, 0xd9, 0x84, 0x4b, 0xba, 0xc6, 0x69, 0x4e, 0x58, 0x47, 0x59, 0xaf, 0xbb, 0x9c, 0x6c, 0x01,
	0xc8, 0xa0, 0xaf, 0xa8, 0x3c, 0x4e, 0xc4, 0x2b, 0x51, 0xa4, 0xeb, 0x92, 0xab, 0xb0, 0x2c, 0x15,
	0x55, 0x81, 0x8c, 0xb4, 0xa2, 0x95, 0xfd, 0x10, 0xd6, 0x0e, 0x50, 0x75, 0x98, 0x3f, 0x74, 0xa7,
	0xdd, 0x3a, 0xf8, 0x29, 0x40, 0x69, 0x12, 0xb2, 0x1f, 0x43, 0x25, 0x86, 0x9b, 0x0a, 0x4a, 0x14,
	0x17, 0x52, 0x8a, 0x7f, 0x2c, 0x58, 0xed, 0x29, 0x2a, 0x42, 0xd1, 0xf3, 0x05, 0xc9, 0x1d, 0xb8,
	0x3c, 0xbd, 0x4f, 0x16, 0xa8, 0x23, 0x89, 0x03, 0xe6, 0x87, 0xdd, 0x95, 0x9c, 0x95, 0x28, 0xda,
	0xd3, 0x41, 0xf2, 0x04, 0x6a, 0x83, 0x29, 0xe3, 0xd1, 0xf4, 0xeb, 0x8d, 0x74, 0x9f, 0xd5, 0xbd,
	0xad, 0xd6, 0x99, 0x89, 0x09, 0xef, 0x57, 0xeb, 0x76, 0x34, 0xc8, 0xa9, 0x0e, 0x92, 0x05, 0x79,
	0x03, 0x8d, 0x90, 0xa1, 0x4f, 0x07, 0xc7, 0xe8, 0xc7, 0x4c, 0x4b, 0x9a, 0xe9, 0x76, 0x01, 0xd3,
	0x7e, 0x08, 0x8e, 0x08, 0xc9, 0x60, 0x26, 0x66, 0x1f, 0x42, 0xbd, 0xa7, 0x18, 0xff, 0xb7, 0xfd,
	0xda, 0x3f, 0x2d, 0x58, 0x77, 0x70, 0xe4, 0x49, 0x85, 0xe2, 0x25, 0x2a, 0xea, 0x52, 0x45, 0xe7,
	0x60, 0xcf, 0x8e, 0xf3, 0xc2, 0xec, 0x38, 0xa7, 0x47, 0x69, 0x31, 0x3b, 0x4a, 0xb3, 0xf5, 0x2d,
	0xe5, 0xdd, 0x47, 0x03, 0x4a, 0x02, 0x95, 0x38, 0xd9, 0x28, 0xe9, 0xdd, 0x70, 0x61, 0xff, 0xb2,
	0x60, 0xf3, 0x29, 0x8a, 0x0b, 0x57, 0xf7, 0x77, 0x0b, 0xc8, 0xe9, 0x69, 0x77, 0x26, 0xee, 0xff,
	0x5f, 0xf0, 0x0f, 0x0b, 0x1a, 0xc9, 0x41, 0x5f, 0x88, 0x92, 0x7b, 0xb0, 0x7e, 0x80, 0x2a, 0xe5,
	0x78, 0x73, 0x14, 0x6d, 0x36, 0x3e, 0xbb, 0x0c, 0xa5, 0x67, 0x13, 0xae, 0x4e, 0xf6, 0x7e, 0x97,
	0xa1, 0xfe, 0xfc, 0xb4, 0x93, 0x1e, 0x8a, 0xcf, 0xde, 0x00, 0xc9, 0x0b, 0xa8, 0x75, 0xc6, 0x4c,
	0x62, 0xe7, 0x23, 0xf5, 0x7d, 0x1c, 0x93, 0x6b, 0xad, 0xbc, 0x1f, 0x98, 0x96, 0x66, 0x68, 0x9a,
	0x36, 0xc9, 0x3e, 0x94, 0x0f, 0x50, 0x69, 0x53, 0x34, 0x92, 0x34, 0xf3, 0x37, 0x75, 0xe2, 0x7b,
	0xa8, 0x9d, 0x35, 0x63, 0x72, 0x3f, 0x1f, 0x9b, 0x63, 0xd8, 0xcd, 0xed, 0x7c, 0x68, 0xc2, 0xe5,
	0x00, 0x24, 0xae, 0x4b, 0xee, 0xe5, 0xc3, 0x67, 0x7c, 0xd9, 0xdc, 0xf3, 0x6b, 0xa8, 0xc4, 0xc6,
	0x46, 0xee, 0x16, 0x51, 0xa6, 0x9d, 0xcf, 0xcc, 0xf8, 0x01, 0xea, 0x59, 0x4f, 0x23, 0x0f, 0xf2,
	0x13, 0x0a, 0xbc, 0xcf, 0xcc, 0xdf, 0x07, 0x32, 0xeb, 0x3e, 0xa4, 0x9d, 0x9f, 0x52, 0xe8, 0x53,
	0x66, 0x8d, 0x43, 0xa8, 0x9e, 0x71, 0x0a, 0xb2, 0x63, 0x2e, 0x3f, 0x79, 0x99, 0x66, 0xd6, 0x77,
	0xb0, 0x92, 0x7a, 0xce, 0x64, 0xf7, 0xbc, 0xa2, 0xe7, 0x65, 0x1e, 0x42, 0x3d, 0xfb, 0xec, 0x8a,
	0xce, 0xbc, 0xe0, 0x79, 0x36, 0x6f, 0x15, 0xdc, 0x7d, 0x8a, 0xf3, 0x2d, 0xac, 0x39, 0xc8, 0x99,
	0xc8, 0x48, 0xcd, 0x93, 0x6b, 0x6c, 0xa0, 0xbf, 0xac, 0xff, 0xb1, 0x3d, 0xfa, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x6d, 0xfc, 0x2b, 0xc1, 0x2e, 0x0a, 0x00, 0x00,
}
