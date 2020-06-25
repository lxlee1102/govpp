// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package mactime

import (
	"context"
	"io"

	api "git.fd.io/govpp.git/api"
)

// RPCService represents RPC service API for mactime module.
type RPCService interface {
	DumpMactime(ctx context.Context, in *MactimeDump) (RPCService_DumpMactimeClient, error)
	MactimeAddDelRange(ctx context.Context, in *MactimeAddDelRange) (*MactimeAddDelRangeReply, error)
	MactimeEnableDisable(ctx context.Context, in *MactimeEnableDisable) (*MactimeEnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpMactime(ctx context.Context, in *MactimeDump) (RPCService_DumpMactimeClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMactimeClient{stream}
	return x, nil
}

type RPCService_DumpMactimeClient interface {
	Recv() (*MactimeDetails, error)
}

type serviceClient_DumpMactimeClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMactimeClient) Recv() (*MactimeDetails, error) {
	m := new(MactimeDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) MactimeAddDelRange(ctx context.Context, in *MactimeAddDelRange) (*MactimeAddDelRangeReply, error) {
	out := new(MactimeAddDelRangeReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MactimeEnableDisable(ctx context.Context, in *MactimeEnableDisable) (*MactimeEnableDisableReply, error) {
	out := new(MactimeEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = context.Background
var _ = io.Copy
