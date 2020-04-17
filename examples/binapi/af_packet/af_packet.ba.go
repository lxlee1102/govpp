// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/af_packet.api.json

/*
Package af_packet is a generated VPP binary API for 'af_packet' module.

It consists of:
	  6 enums
	  2 aliases
	  8 messages
	  4 services
*/
package af_packet

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"

	ethernet_types "git.fd.io/govpp.git/examples/binapi/ethernet_types"
	interface_types "git.fd.io/govpp.git/examples/binapi/interface_types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "af_packet"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xba745e20
)

type IfStatusFlags = interface_types.IfStatusFlags

type IfType = interface_types.IfType

type LinkDuplex = interface_types.LinkDuplex

type MtuProto = interface_types.MtuProto

type RxMode = interface_types.RxMode

type SubIfFlags = interface_types.SubIfFlags

type InterfaceIndex = interface_types.InterfaceIndex

type MacAddress = ethernet_types.MacAddress

// AfPacketCreate represents VPP binary API message 'af_packet_create'.
type AfPacketCreate struct {
	HwAddr          MacAddress
	UseRandomHwAddr bool
	HostIfName      string `struc:"[64]byte"`
}

func (m *AfPacketCreate) Reset()                        { *m = AfPacketCreate{} }
func (*AfPacketCreate) GetMessageName() string          { return "af_packet_create" }
func (*AfPacketCreate) GetCrcString() string            { return "a190415f" }
func (*AfPacketCreate) GetMessageType() api.MessageType { return api.RequestMessage }

// AfPacketCreateReply represents VPP binary API message 'af_packet_create_reply'.
type AfPacketCreateReply struct {
	Retval    int32
	SwIfIndex InterfaceIndex
}

func (m *AfPacketCreateReply) Reset()                        { *m = AfPacketCreateReply{} }
func (*AfPacketCreateReply) GetMessageName() string          { return "af_packet_create_reply" }
func (*AfPacketCreateReply) GetCrcString() string            { return "5383d31f" }
func (*AfPacketCreateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// AfPacketDelete represents VPP binary API message 'af_packet_delete'.
type AfPacketDelete struct {
	HostIfName string `struc:"[64]byte"`
}

func (m *AfPacketDelete) Reset()                        { *m = AfPacketDelete{} }
func (*AfPacketDelete) GetMessageName() string          { return "af_packet_delete" }
func (*AfPacketDelete) GetCrcString() string            { return "863fa648" }
func (*AfPacketDelete) GetMessageType() api.MessageType { return api.RequestMessage }

// AfPacketDeleteReply represents VPP binary API message 'af_packet_delete_reply'.
type AfPacketDeleteReply struct {
	Retval int32
}

func (m *AfPacketDeleteReply) Reset()                        { *m = AfPacketDeleteReply{} }
func (*AfPacketDeleteReply) GetMessageName() string          { return "af_packet_delete_reply" }
func (*AfPacketDeleteReply) GetCrcString() string            { return "e8d4e804" }
func (*AfPacketDeleteReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// AfPacketDetails represents VPP binary API message 'af_packet_details'.
type AfPacketDetails struct {
	SwIfIndex  InterfaceIndex
	HostIfName string `struc:"[64]byte"`
}

func (m *AfPacketDetails) Reset()                        { *m = AfPacketDetails{} }
func (*AfPacketDetails) GetMessageName() string          { return "af_packet_details" }
func (*AfPacketDetails) GetCrcString() string            { return "58c7c042" }
func (*AfPacketDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// AfPacketDump represents VPP binary API message 'af_packet_dump'.
type AfPacketDump struct{}

func (m *AfPacketDump) Reset()                        { *m = AfPacketDump{} }
func (*AfPacketDump) GetMessageName() string          { return "af_packet_dump" }
func (*AfPacketDump) GetCrcString() string            { return "51077d14" }
func (*AfPacketDump) GetMessageType() api.MessageType { return api.RequestMessage }

// AfPacketSetL4CksumOffload represents VPP binary API message 'af_packet_set_l4_cksum_offload'.
type AfPacketSetL4CksumOffload struct {
	SwIfIndex InterfaceIndex
	Set       bool
}

func (m *AfPacketSetL4CksumOffload) Reset()                        { *m = AfPacketSetL4CksumOffload{} }
func (*AfPacketSetL4CksumOffload) GetMessageName() string          { return "af_packet_set_l4_cksum_offload" }
func (*AfPacketSetL4CksumOffload) GetCrcString() string            { return "319cd5c8" }
func (*AfPacketSetL4CksumOffload) GetMessageType() api.MessageType { return api.RequestMessage }

// AfPacketSetL4CksumOffloadReply represents VPP binary API message 'af_packet_set_l4_cksum_offload_reply'.
type AfPacketSetL4CksumOffloadReply struct {
	Retval int32
}

func (m *AfPacketSetL4CksumOffloadReply) Reset() { *m = AfPacketSetL4CksumOffloadReply{} }
func (*AfPacketSetL4CksumOffloadReply) GetMessageName() string {
	return "af_packet_set_l4_cksum_offload_reply"
}
func (*AfPacketSetL4CksumOffloadReply) GetCrcString() string            { return "e8d4e804" }
func (*AfPacketSetL4CksumOffloadReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*AfPacketCreate)(nil), "af_packet.AfPacketCreate")
	api.RegisterMessage((*AfPacketCreateReply)(nil), "af_packet.AfPacketCreateReply")
	api.RegisterMessage((*AfPacketDelete)(nil), "af_packet.AfPacketDelete")
	api.RegisterMessage((*AfPacketDeleteReply)(nil), "af_packet.AfPacketDeleteReply")
	api.RegisterMessage((*AfPacketDetails)(nil), "af_packet.AfPacketDetails")
	api.RegisterMessage((*AfPacketDump)(nil), "af_packet.AfPacketDump")
	api.RegisterMessage((*AfPacketSetL4CksumOffload)(nil), "af_packet.AfPacketSetL4CksumOffload")
	api.RegisterMessage((*AfPacketSetL4CksumOffloadReply)(nil), "af_packet.AfPacketSetL4CksumOffloadReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AfPacketCreate)(nil),
		(*AfPacketCreateReply)(nil),
		(*AfPacketDelete)(nil),
		(*AfPacketDeleteReply)(nil),
		(*AfPacketDetails)(nil),
		(*AfPacketDump)(nil),
		(*AfPacketSetL4CksumOffload)(nil),
		(*AfPacketSetL4CksumOffloadReply)(nil),
	}
}

// RPCService represents RPC service API for af_packet module.
type RPCService interface {
	DumpAfPacket(ctx context.Context, in *AfPacketDump) (RPCService_DumpAfPacketClient, error)
	AfPacketCreate(ctx context.Context, in *AfPacketCreate) (*AfPacketCreateReply, error)
	AfPacketDelete(ctx context.Context, in *AfPacketDelete) (*AfPacketDeleteReply, error)
	AfPacketSetL4CksumOffload(ctx context.Context, in *AfPacketSetL4CksumOffload) (*AfPacketSetL4CksumOffloadReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpAfPacket(ctx context.Context, in *AfPacketDump) (RPCService_DumpAfPacketClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpAfPacketClient{stream}
	return x, nil
}

type RPCService_DumpAfPacketClient interface {
	Recv() (*AfPacketDetails, error)
}

type serviceClient_DumpAfPacketClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpAfPacketClient) Recv() (*AfPacketDetails, error) {
	m := new(AfPacketDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) AfPacketCreate(ctx context.Context, in *AfPacketCreate) (*AfPacketCreateReply, error) {
	out := new(AfPacketCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AfPacketDelete(ctx context.Context, in *AfPacketDelete) (*AfPacketDeleteReply, error) {
	out := new(AfPacketDeleteReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AfPacketSetL4CksumOffload(ctx context.Context, in *AfPacketSetL4CksumOffload) (*AfPacketSetL4CksumOffloadReply, error) {
	out := new(AfPacketSetL4CksumOffloadReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
