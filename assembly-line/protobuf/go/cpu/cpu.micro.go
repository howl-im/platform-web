// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cpu/cpu.proto

package cpu

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/basic"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CPUService service

type CPUService interface {
	PushCPUTimesStat(ctx context.Context, in *CPURequest, opts ...client.CallOption) (*CPUResponse, error)
	PushCPUInfoStat(ctx context.Context, in *CPURequest, opts ...client.CallOption) (*CPUResponse, error)
	PushCPUPercent(ctx context.Context, in *CPURequest, opts ...client.CallOption) (*CPUResponse, error)
	PushCPUCounts(ctx context.Context, in *CPURequest, opts ...client.CallOption) (*CPUResponse, error)
}

type cPUService struct {
	c    client.Client
	name string
}

func NewCPUService(name string, c client.Client) CPUService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "protobuf.pb.cpu.cpu"
	}
	return &cPUService{
		c:    c,
		name: name,
	}
}

func (c *cPUService) PushCPUTimesStat(ctx context.Context, in *CPURequest, opts ...client.CallOption) (*CPUResponse, error) {
	req := c.c.NewRequest(c.name, "CPUService.PushCPUTimesStat", in)
	out := new(CPUResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPUService) PushCPUInfoStat(ctx context.Context, in *CPURequest, opts ...client.CallOption) (*CPUResponse, error) {
	req := c.c.NewRequest(c.name, "CPUService.PushCPUInfoStat", in)
	out := new(CPUResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPUService) PushCPUPercent(ctx context.Context, in *CPURequest, opts ...client.CallOption) (*CPUResponse, error) {
	req := c.c.NewRequest(c.name, "CPUService.PushCPUPercent", in)
	out := new(CPUResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPUService) PushCPUCounts(ctx context.Context, in *CPURequest, opts ...client.CallOption) (*CPUResponse, error) {
	req := c.c.NewRequest(c.name, "CPUService.PushCPUCounts", in)
	out := new(CPUResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CPUService service

type CPUServiceHandler interface {
	PushCPUTimesStat(context.Context, *CPURequest, *CPUResponse) error
	PushCPUInfoStat(context.Context, *CPURequest, *CPUResponse) error
	PushCPUPercent(context.Context, *CPURequest, *CPUResponse) error
	PushCPUCounts(context.Context, *CPURequest, *CPUResponse) error
}

func RegisterCPUServiceHandler(s server.Server, hdlr CPUServiceHandler, opts ...server.HandlerOption) error {
	type cPUService interface {
		PushCPUTimesStat(ctx context.Context, in *CPURequest, out *CPUResponse) error
		PushCPUInfoStat(ctx context.Context, in *CPURequest, out *CPUResponse) error
		PushCPUPercent(ctx context.Context, in *CPURequest, out *CPUResponse) error
		PushCPUCounts(ctx context.Context, in *CPURequest, out *CPUResponse) error
	}
	type CPUService struct {
		cPUService
	}
	h := &cPUServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CPUService{h}, opts...))
}

type cPUServiceHandler struct {
	CPUServiceHandler
}

func (h *cPUServiceHandler) PushCPUTimesStat(ctx context.Context, in *CPURequest, out *CPUResponse) error {
	return h.CPUServiceHandler.PushCPUTimesStat(ctx, in, out)
}

func (h *cPUServiceHandler) PushCPUInfoStat(ctx context.Context, in *CPURequest, out *CPUResponse) error {
	return h.CPUServiceHandler.PushCPUInfoStat(ctx, in, out)
}

func (h *cPUServiceHandler) PushCPUPercent(ctx context.Context, in *CPURequest, out *CPUResponse) error {
	return h.CPUServiceHandler.PushCPUPercent(ctx, in, out)
}

func (h *cPUServiceHandler) PushCPUCounts(ctx context.Context, in *CPURequest, out *CPUResponse) error {
	return h.CPUServiceHandler.PushCPUCounts(ctx, in, out)
}
