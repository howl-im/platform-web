// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mem/mem.proto

package mem

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	basic "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/basic"
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

type VirtualMemoryStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Total                uint64               `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Available            uint64               `protobuf:"varint,3,opt,name=Available,proto3" json:"Available,omitempty"`
	Used                 uint64               `protobuf:"varint,4,opt,name=Used,proto3" json:"Used,omitempty"`
	UsedPercent          float64              `protobuf:"fixed64,5,opt,name=UsedPercent,proto3" json:"UsedPercent,omitempty"`
	Free                 uint64               `protobuf:"varint,6,opt,name=Free,proto3" json:"Free,omitempty"`
	Active               uint64               `protobuf:"varint,7,opt,name=Active,proto3" json:"Active,omitempty"`
	Inactive             uint64               `protobuf:"varint,8,opt,name=Inactive,proto3" json:"Inactive,omitempty"`
	Wired                uint64               `protobuf:"varint,9,opt,name=Wired,proto3" json:"Wired,omitempty"`
	Laundry              uint64               `protobuf:"varint,10,opt,name=Laundry,proto3" json:"Laundry,omitempty"`
	Buffers              uint64               `protobuf:"varint,11,opt,name=Buffers,proto3" json:"Buffers,omitempty"`
	Cached               uint64               `protobuf:"varint,12,opt,name=Cached,proto3" json:"Cached,omitempty"`
	Writeback            uint64               `protobuf:"varint,13,opt,name=Writeback,proto3" json:"Writeback,omitempty"`
	Dirty                uint64               `protobuf:"varint,14,opt,name=Dirty,proto3" json:"Dirty,omitempty"`
	WritebackTmp         uint64               `protobuf:"varint,15,opt,name=WritebackTmp,proto3" json:"WritebackTmp,omitempty"`
	Shared               uint64               `protobuf:"varint,16,opt,name=Shared,proto3" json:"Shared,omitempty"`
	Slab                 uint64               `protobuf:"varint,17,opt,name=Slab,proto3" json:"Slab,omitempty"`
	SReclaimable         uint64               `protobuf:"varint,18,opt,name=SReclaimable,proto3" json:"SReclaimable,omitempty"`
	PageTables           uint64               `protobuf:"varint,19,opt,name=PageTables,proto3" json:"PageTables,omitempty"`
	SwapCached           uint64               `protobuf:"varint,20,opt,name=SwapCached,proto3" json:"SwapCached,omitempty"`
	CommitLimit          uint64               `protobuf:"varint,21,opt,name=CommitLimit,proto3" json:"CommitLimit,omitempty"`
	CommittedAS          uint64               `protobuf:"varint,22,opt,name=CommittedAS,proto3" json:"CommittedAS,omitempty"`
	HighTotal            uint64               `protobuf:"varint,23,opt,name=HighTotal,proto3" json:"HighTotal,omitempty"`
	HighFree             uint64               `protobuf:"varint,24,opt,name=HighFree,proto3" json:"HighFree,omitempty"`
	LowTotal             uint64               `protobuf:"varint,25,opt,name=LowTotal,proto3" json:"LowTotal,omitempty"`
	LowFree              uint64               `protobuf:"varint,26,opt,name=LowFree,proto3" json:"LowFree,omitempty"`
	SwapTotal            uint64               `protobuf:"varint,27,opt,name=SwapTotal,proto3" json:"SwapTotal,omitempty"`
	SwapFree             uint64               `protobuf:"varint,28,opt,name=SwapFree,proto3" json:"SwapFree,omitempty"`
	Mapped               uint64               `protobuf:"varint,29,opt,name=Mapped,proto3" json:"Mapped,omitempty"`
	VMallocTotal         uint64               `protobuf:"varint,30,opt,name=VMallocTotal,proto3" json:"VMallocTotal,omitempty"`
	VMallocUsed          uint64               `protobuf:"varint,31,opt,name=VMallocUsed,proto3" json:"VMallocUsed,omitempty"`
	VMallocChunk         uint64               `protobuf:"varint,32,opt,name=VMallocChunk,proto3" json:"VMallocChunk,omitempty"`
	HugePagesTotal       uint64               `protobuf:"varint,33,opt,name=HugePagesTotal,proto3" json:"HugePagesTotal,omitempty"`
	HugePagesFree        uint64               `protobuf:"varint,34,opt,name=HugePagesFree,proto3" json:"HugePagesFree,omitempty"`
	HugePageSize         uint64               `protobuf:"varint,35,opt,name=HugePageSize,proto3" json:"HugePageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VirtualMemoryStat) Reset()         { *m = VirtualMemoryStat{} }
func (m *VirtualMemoryStat) String() string { return proto.CompactTextString(m) }
func (*VirtualMemoryStat) ProtoMessage()    {}
func (*VirtualMemoryStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2907961eca7e66, []int{0}
}

func (m *VirtualMemoryStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMemoryStat.Unmarshal(m, b)
}
func (m *VirtualMemoryStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMemoryStat.Marshal(b, m, deterministic)
}
func (m *VirtualMemoryStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMemoryStat.Merge(m, src)
}
func (m *VirtualMemoryStat) XXX_Size() int {
	return xxx_messageInfo_VirtualMemoryStat.Size(m)
}
func (m *VirtualMemoryStat) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMemoryStat.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMemoryStat proto.InternalMessageInfo

func (m *VirtualMemoryStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *VirtualMemoryStat) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *VirtualMemoryStat) GetAvailable() uint64 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *VirtualMemoryStat) GetUsed() uint64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *VirtualMemoryStat) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

func (m *VirtualMemoryStat) GetFree() uint64 {
	if m != nil {
		return m.Free
	}
	return 0
}

func (m *VirtualMemoryStat) GetActive() uint64 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *VirtualMemoryStat) GetInactive() uint64 {
	if m != nil {
		return m.Inactive
	}
	return 0
}

func (m *VirtualMemoryStat) GetWired() uint64 {
	if m != nil {
		return m.Wired
	}
	return 0
}

func (m *VirtualMemoryStat) GetLaundry() uint64 {
	if m != nil {
		return m.Laundry
	}
	return 0
}

func (m *VirtualMemoryStat) GetBuffers() uint64 {
	if m != nil {
		return m.Buffers
	}
	return 0
}

func (m *VirtualMemoryStat) GetCached() uint64 {
	if m != nil {
		return m.Cached
	}
	return 0
}

func (m *VirtualMemoryStat) GetWriteback() uint64 {
	if m != nil {
		return m.Writeback
	}
	return 0
}

func (m *VirtualMemoryStat) GetDirty() uint64 {
	if m != nil {
		return m.Dirty
	}
	return 0
}

func (m *VirtualMemoryStat) GetWritebackTmp() uint64 {
	if m != nil {
		return m.WritebackTmp
	}
	return 0
}

func (m *VirtualMemoryStat) GetShared() uint64 {
	if m != nil {
		return m.Shared
	}
	return 0
}

func (m *VirtualMemoryStat) GetSlab() uint64 {
	if m != nil {
		return m.Slab
	}
	return 0
}

func (m *VirtualMemoryStat) GetSReclaimable() uint64 {
	if m != nil {
		return m.SReclaimable
	}
	return 0
}

func (m *VirtualMemoryStat) GetPageTables() uint64 {
	if m != nil {
		return m.PageTables
	}
	return 0
}

func (m *VirtualMemoryStat) GetSwapCached() uint64 {
	if m != nil {
		return m.SwapCached
	}
	return 0
}

func (m *VirtualMemoryStat) GetCommitLimit() uint64 {
	if m != nil {
		return m.CommitLimit
	}
	return 0
}

func (m *VirtualMemoryStat) GetCommittedAS() uint64 {
	if m != nil {
		return m.CommittedAS
	}
	return 0
}

func (m *VirtualMemoryStat) GetHighTotal() uint64 {
	if m != nil {
		return m.HighTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetHighFree() uint64 {
	if m != nil {
		return m.HighFree
	}
	return 0
}

func (m *VirtualMemoryStat) GetLowTotal() uint64 {
	if m != nil {
		return m.LowTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetLowFree() uint64 {
	if m != nil {
		return m.LowFree
	}
	return 0
}

func (m *VirtualMemoryStat) GetSwapTotal() uint64 {
	if m != nil {
		return m.SwapTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetSwapFree() uint64 {
	if m != nil {
		return m.SwapFree
	}
	return 0
}

func (m *VirtualMemoryStat) GetMapped() uint64 {
	if m != nil {
		return m.Mapped
	}
	return 0
}

func (m *VirtualMemoryStat) GetVMallocTotal() uint64 {
	if m != nil {
		return m.VMallocTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetVMallocUsed() uint64 {
	if m != nil {
		return m.VMallocUsed
	}
	return 0
}

func (m *VirtualMemoryStat) GetVMallocChunk() uint64 {
	if m != nil {
		return m.VMallocChunk
	}
	return 0
}

func (m *VirtualMemoryStat) GetHugePagesTotal() uint64 {
	if m != nil {
		return m.HugePagesTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetHugePagesFree() uint64 {
	if m != nil {
		return m.HugePagesFree
	}
	return 0
}

func (m *VirtualMemoryStat) GetHugePageSize() uint64 {
	if m != nil {
		return m.HugePageSize
	}
	return 0
}

type SwapMemoryStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Total                uint64               `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Used                 uint64               `protobuf:"varint,3,opt,name=Used,proto3" json:"Used,omitempty"`
	Free                 uint64               `protobuf:"varint,4,opt,name=Free,proto3" json:"Free,omitempty"`
	UsedPercent          float64              `protobuf:"fixed64,5,opt,name=UsedPercent,proto3" json:"UsedPercent,omitempty"`
	Sin                  uint64               `protobuf:"varint,6,opt,name=Sin,proto3" json:"Sin,omitempty"`
	Sout                 uint64               `protobuf:"varint,7,opt,name=Sout,proto3" json:"Sout,omitempty"`
	PgIn                 uint64               `protobuf:"varint,8,opt,name=PgIn,proto3" json:"PgIn,omitempty"`
	PgOut                uint64               `protobuf:"varint,9,opt,name=PgOut,proto3" json:"PgOut,omitempty"`
	PgFault              uint64               `protobuf:"varint,10,opt,name=PgFault,proto3" json:"PgFault,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SwapMemoryStat) Reset()         { *m = SwapMemoryStat{} }
func (m *SwapMemoryStat) String() string { return proto.CompactTextString(m) }
func (*SwapMemoryStat) ProtoMessage()    {}
func (*SwapMemoryStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2907961eca7e66, []int{1}
}

func (m *SwapMemoryStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwapMemoryStat.Unmarshal(m, b)
}
func (m *SwapMemoryStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwapMemoryStat.Marshal(b, m, deterministic)
}
func (m *SwapMemoryStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapMemoryStat.Merge(m, src)
}
func (m *SwapMemoryStat) XXX_Size() int {
	return xxx_messageInfo_SwapMemoryStat.Size(m)
}
func (m *SwapMemoryStat) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapMemoryStat.DiscardUnknown(m)
}

var xxx_messageInfo_SwapMemoryStat proto.InternalMessageInfo

func (m *SwapMemoryStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *SwapMemoryStat) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *SwapMemoryStat) GetUsed() uint64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *SwapMemoryStat) GetFree() uint64 {
	if m != nil {
		return m.Free
	}
	return 0
}

func (m *SwapMemoryStat) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

func (m *SwapMemoryStat) GetSin() uint64 {
	if m != nil {
		return m.Sin
	}
	return 0
}

func (m *SwapMemoryStat) GetSout() uint64 {
	if m != nil {
		return m.Sout
	}
	return 0
}

func (m *SwapMemoryStat) GetPgIn() uint64 {
	if m != nil {
		return m.PgIn
	}
	return 0
}

func (m *SwapMemoryStat) GetPgOut() uint64 {
	if m != nil {
		return m.PgOut
	}
	return 0
}

func (m *SwapMemoryStat) GetPgFault() uint64 {
	if m != nil {
		return m.PgFault
	}
	return 0
}

type MemRequest struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IP                   string               `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	NodeName             string               `protobuf:"bytes,3,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	VirtualMemoryStat    []*VirtualMemoryStat `protobuf:"bytes,4,rep,name=virtualMemoryStat,proto3" json:"virtualMemoryStat,omitempty"`
	SwapMemoryStat       []*SwapMemoryStat    `protobuf:"bytes,5,rep,name=swapMemoryStat,proto3" json:"swapMemoryStat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MemRequest) Reset()         { *m = MemRequest{} }
func (m *MemRequest) String() string { return proto.CompactTextString(m) }
func (*MemRequest) ProtoMessage()    {}
func (*MemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2907961eca7e66, []int{2}
}

func (m *MemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemRequest.Unmarshal(m, b)
}
func (m *MemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemRequest.Marshal(b, m, deterministic)
}
func (m *MemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemRequest.Merge(m, src)
}
func (m *MemRequest) XXX_Size() int {
	return xxx_messageInfo_MemRequest.Size(m)
}
func (m *MemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemRequest proto.InternalMessageInfo

func (m *MemRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *MemRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *MemRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MemRequest) GetVirtualMemoryStat() []*VirtualMemoryStat {
	if m != nil {
		return m.VirtualMemoryStat
	}
	return nil
}

func (m *MemRequest) GetSwapMemoryStat() []*SwapMemoryStat {
	if m != nil {
		return m.SwapMemoryStat
	}
	return nil
}

type MemResponse struct {
	Error                *basic.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MemResponse) Reset()         { *m = MemResponse{} }
func (m *MemResponse) String() string { return proto.CompactTextString(m) }
func (*MemResponse) ProtoMessage()    {}
func (*MemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2907961eca7e66, []int{3}
}

func (m *MemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemResponse.Unmarshal(m, b)
}
func (m *MemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemResponse.Marshal(b, m, deterministic)
}
func (m *MemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemResponse.Merge(m, src)
}
func (m *MemResponse) XXX_Size() int {
	return xxx_messageInfo_MemResponse.Size(m)
}
func (m *MemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemResponse proto.InternalMessageInfo

func (m *MemResponse) GetError() *basic.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualMemoryStat)(nil), "protobuf.pb.mem.VirtualMemoryStat")
	proto.RegisterType((*SwapMemoryStat)(nil), "protobuf.pb.mem.SwapMemoryStat")
	proto.RegisterType((*MemRequest)(nil), "protobuf.pb.mem.MemRequest")
	proto.RegisterType((*MemResponse)(nil), "protobuf.pb.mem.MemResponse")
}

func init() { proto.RegisterFile("mem/mem.proto", fileDescriptor_6d2907961eca7e66) }

var fileDescriptor_6d2907961eca7e66 = []byte{
	// 853 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0x13, 0x27, 0x8d, 0xe9, 0xc6, 0x6d, 0xb8, 0xa6, 0xe3, 0xdc, 0xac, 0xf1, 0xbc, 0x61,
	0x08, 0x30, 0xc4, 0x06, 0xb2, 0x9b, 0x61, 0x77, 0x69, 0xd6, 0x9f, 0x00, 0xc9, 0xa6, 0x49, 0x5e,
	0x0b, 0xec, 0x8e, 0x92, 0x19, 0x99, 0xa8, 0x28, 0x6a, 0x12, 0x95, 0x20, 0x7b, 0x87, 0xdd, 0xed,
	0x2d, 0x76, 0xbd, 0xf7, 0x1b, 0xce, 0x39, 0x92, 0x2c, 0xdb, 0x03, 0x76, 0x51, 0xf4, 0x4a, 0xfc,
	0xbe, 0xf3, 0x47, 0x7e, 0x87, 0x3a, 0x64, 0xfb, 0x46, 0x99, 0xa9, 0x51, 0x66, 0x92, 0xe5, 0xd6,
	0x59, 0xfe, 0x08, 0x3f, 0x61, 0x79, 0x33, 0xc9, 0xc2, 0x89, 0x51, 0x66, 0x78, 0x10, 0xca, 0x42,
	0x47, 0x53, 0x95, 0xe7, 0x36, 0x27, 0x9f, 0xe1, 0x71, 0x6c, 0x6d, 0x9c, 0xa8, 0x69, 0xed, 0x3a,
	0x75, 0xda, 0xa8, 0xc2, 0x49, 0x93, 0x91, 0xc3, 0xf8, 0xef, 0x3d, 0x76, 0xf0, 0x56, 0xe7, 0xae,
	0x94, 0xc9, 0xb5, 0x32, 0x36, 0xbf, 0x0f, 0x9c, 0x74, 0xfc, 0x7b, 0xd6, 0x6b, 0x1c, 0x45, 0x67,
	0xd4, 0x39, 0xe9, 0x9f, 0x0d, 0x27, 0x94, 0x6a, 0xd2, 0x54, 0x9d, 0xd5, 0x1e, 0xfe, 0xd2, 0x99,
	0x3f, 0x61, 0x3b, 0x33, 0xeb, 0x64, 0x22, 0xb6, 0x46, 0x9d, 0x93, 0xae, 0x4f, 0x80, 0x1f, 0xb1,
	0xde, 0xf9, 0xad, 0xd4, 0x89, 0x0c, 0x13, 0x25, 0xb6, 0xd1, 0xb2, 0x24, 0x38, 0x67, 0xdd, 0x5f,
	0x0b, 0x35, 0x17, 0x5d, 0x34, 0xe0, 0x9a, 0x8f, 0x58, 0x1f, 0xbe, 0x9e, 0xca, 0x23, 0x95, 0x3a,
	0xb1, 0x33, 0xea, 0x9c, 0x74, 0xfc, 0x36, 0x05, 0x51, 0xaf, 0x72, 0xa5, 0xc4, 0x2e, 0x45, 0xc1,
	0x9a, 0x3f, 0x65, 0xbb, 0xe7, 0x91, 0xd3, 0xb7, 0x4a, 0x3c, 0x40, 0xb6, 0x42, 0x7c, 0xc8, 0xf6,
	0x2e, 0x53, 0x49, 0x96, 0x3d, 0xb4, 0x34, 0x18, 0x76, 0xfc, 0x4e, 0xe7, 0x6a, 0x2e, 0x7a, 0xb4,
	0x63, 0x04, 0x5c, 0xb0, 0x07, 0x57, 0xb2, 0x4c, 0xe7, 0xf9, 0xbd, 0x60, 0xc8, 0xd7, 0x10, 0x2c,
	0x2f, 0xca, 0x9b, 0x1b, 0x95, 0x17, 0xa2, 0x4f, 0x96, 0x0a, 0x42, 0xf5, 0x0b, 0x19, 0x2d, 0xd4,
	0x5c, 0x3c, 0xa4, 0xea, 0x84, 0xe0, 0xf4, 0xef, 0x72, 0xed, 0x54, 0x28, 0xa3, 0xf7, 0x62, 0x9f,
	0x4e, 0xdf, 0x10, 0x50, 0xff, 0x47, 0x9d, 0xbb, 0x7b, 0x31, 0xa0, 0xfa, 0x08, 0xf8, 0x98, 0x3d,
	0x6c, 0x5c, 0x66, 0x26, 0x13, 0x8f, 0xd0, 0xb8, 0xc2, 0x41, 0xbd, 0x60, 0x21, 0x61, 0xeb, 0x8f,
	0xa9, 0x1e, 0x21, 0x50, 0x26, 0x48, 0x64, 0x28, 0x0e, 0x48, 0x19, 0x58, 0x43, 0xbe, 0xc0, 0x57,
	0x51, 0x22, 0xb5, 0xc1, 0x26, 0x70, 0xca, 0xd7, 0xe6, 0xf8, 0x73, 0xc6, 0x3c, 0x19, 0xab, 0x19,
	0x80, 0x42, 0x7c, 0x8a, 0x1e, 0x2d, 0x06, 0xec, 0xc1, 0x9d, 0xcc, 0xaa, 0x33, 0x3e, 0x21, 0xfb,
	0x92, 0x81, 0x9e, 0x5d, 0x58, 0x63, 0xb4, 0xbb, 0xd2, 0x46, 0x3b, 0x71, 0x88, 0x0e, 0x6d, 0x6a,
	0xe9, 0xe1, 0xd4, 0xfc, 0x3c, 0x10, 0x4f, 0xdb, 0x1e, 0x48, 0x81, 0x56, 0x6f, 0x74, 0xbc, 0xa0,
	0x3b, 0xf4, 0x19, 0x69, 0xd5, 0x10, 0xd0, 0x47, 0x00, 0xd8, 0x77, 0x41, 0x7d, 0xac, 0x31, 0xd8,
	0xae, 0xec, 0x1d, 0x05, 0x7e, 0x4e, 0xb6, 0x1a, 0x63, 0x37, 0xed, 0x1d, 0x86, 0x0d, 0xab, 0x6e,
	0x12, 0x84, 0x7a, 0x70, 0x02, 0x0a, 0x7b, 0x46, 0xf5, 0x1a, 0x02, 0x72, 0x02, 0xc0, 0xc0, 0x23,
	0xca, 0x59, 0x63, 0x50, 0xff, 0x5a, 0x66, 0x99, 0x9a, 0x8b, 0x2f, 0x48, 0x7d, 0x42, 0xa0, 0xf4,
	0xdb, 0x6b, 0x99, 0x24, 0x36, 0xa2, 0xa4, 0xcf, 0x49, 0xe9, 0x36, 0x07, 0x3a, 0x54, 0x18, 0x2f,
	0xfe, 0x31, 0xe9, 0xd0, 0xa2, 0x5a, 0x59, 0x2e, 0x16, 0x65, 0xfa, 0x5e, 0x8c, 0x56, 0xb2, 0x20,
	0xc7, 0xbf, 0x61, 0x83, 0x37, 0x65, 0xac, 0xa0, 0x43, 0x05, 0xd5, 0xfa, 0x12, 0xbd, 0xd6, 0x58,
	0xfe, 0x35, 0xdb, 0x6f, 0x18, 0x3c, 0xca, 0x18, 0xdd, 0x56, 0x49, 0xa8, 0x58, 0x13, 0x81, 0xfe,
	0x43, 0x89, 0xaf, 0xa8, 0x62, 0x9b, 0x1b, 0xff, 0xb5, 0xc5, 0x06, 0x20, 0xc0, 0x47, 0x1c, 0x15,
	0xf5, 0x30, 0xd8, 0x6e, 0x0d, 0x83, 0xfa, 0x57, 0xef, 0xb6, 0x7e, 0xf5, 0xff, 0x1f, 0x10, 0x8f,
	0xd9, 0x76, 0xa0, 0xd3, 0x6a, 0x3e, 0xc0, 0x12, 0x7f, 0x0c, 0x5b, 0xba, 0x6a, 0x38, 0xe0, 0x1a,
	0x38, 0x2f, 0xbe, 0x4c, 0xab, 0xb1, 0x80, 0x6b, 0xd8, 0x99, 0x17, 0xff, 0x5c, 0xba, 0x7a, 0x24,
	0x20, 0x80, 0x4b, 0xe4, 0xc5, 0xaf, 0x64, 0x99, 0xb8, 0x7a, 0x24, 0x54, 0x70, 0xfc, 0xe7, 0x16,
	0x63, 0xd7, 0xca, 0xf8, 0xea, 0xf7, 0x52, 0x15, 0x1f, 0x22, 0xc9, 0x80, 0x6d, 0x5d, 0x7a, 0xa8,
	0x47, 0xcf, 0xdf, 0xba, 0xf4, 0xe0, 0xfe, 0xa5, 0x76, 0xae, 0x7e, 0x92, 0x86, 0xc6, 0x66, 0xcf,
	0x6f, 0x30, 0xf7, 0xd8, 0xc1, 0xed, 0xfa, 0xe0, 0x16, 0xdd, 0xd1, 0xf6, 0x49, 0xff, 0x6c, 0x3c,
	0x59, 0x7b, 0x1a, 0x26, 0x1b, 0x23, 0xde, 0xdf, 0x0c, 0xe6, 0xaf, 0xd9, 0xa0, 0x58, 0x69, 0xae,
	0xd8, 0xc1, 0x74, 0xc7, 0x1b, 0xe9, 0x56, 0xef, 0x80, 0xbf, 0x16, 0x36, 0xfe, 0x81, 0xf5, 0x51,
	0x8e, 0x22, 0xb3, 0x69, 0xa1, 0xf8, 0xb7, 0x6c, 0x07, 0xdf, 0xa4, 0x4a, 0x8b, 0xc3, 0x65, 0x3a,
	0x7c, 0xb0, 0x26, 0x2f, 0xc1, 0xe8, 0x93, 0xcf, 0xd9, 0x3f, 0x1d, 0xd4, 0x32, 0x50, 0xf9, 0xad,
	0x8e, 0x14, 0x9f, 0xb1, 0x43, 0xaf, 0x2c, 0x16, 0x9b, 0x4f, 0xd4, 0xb3, 0x8d, 0x4d, 0x2d, 0x3b,
	0x30, 0x3c, 0xfa, 0x6f, 0x23, 0xed, 0x67, 0xfc, 0x09, 0xff, 0x85, 0x71, 0xc8, 0xba, 0x76, 0x95,
	0x3f, 0x24, 0xe5, 0x8b, 0xd7, 0xbf, 0xbd, 0x8c, 0xb5, 0x5b, 0x94, 0xe1, 0x24, 0xb2, 0x66, 0x6a,
	0x74, 0x94, 0xdb, 0x53, 0x9d, 0x9e, 0x46, 0xe9, 0x34, 0x4b, 0xa4, 0xbb, 0xb1, 0xb9, 0x39, 0xbd,
	0x53, 0xe1, 0x54, 0x16, 0x85, 0x32, 0x61, 0x72, 0x7f, 0x9a, 0xe8, 0xb4, 0xf5, 0x2c, 0xc7, 0x76,
	0x5a, 0x3d, 0xee, 0xe1, 0x2e, 0x92, 0xdf, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x3b, 0xaa,
	0x4a, 0xee, 0x07, 0x00, 0x00,
}
