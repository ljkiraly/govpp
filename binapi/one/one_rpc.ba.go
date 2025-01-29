// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package one

import (
	"context"
	"fmt"
	"io"

	memclnt "github.com/ljkiraly/govpp/binapi/memclnt"
	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service one.
type RPCService interface {
	OneAddDelAdjacency(ctx context.Context, in *OneAddDelAdjacency) (*OneAddDelAdjacencyReply, error)
	OneAddDelL2ArpEntry(ctx context.Context, in *OneAddDelL2ArpEntry) (*OneAddDelL2ArpEntryReply, error)
	OneAddDelLocalEid(ctx context.Context, in *OneAddDelLocalEid) (*OneAddDelLocalEidReply, error)
	OneAddDelLocator(ctx context.Context, in *OneAddDelLocator) (*OneAddDelLocatorReply, error)
	OneAddDelLocatorSet(ctx context.Context, in *OneAddDelLocatorSet) (*OneAddDelLocatorSetReply, error)
	OneAddDelMapRequestItrRlocs(ctx context.Context, in *OneAddDelMapRequestItrRlocs) (*OneAddDelMapRequestItrRlocsReply, error)
	OneAddDelMapResolver(ctx context.Context, in *OneAddDelMapResolver) (*OneAddDelMapResolverReply, error)
	OneAddDelMapServer(ctx context.Context, in *OneAddDelMapServer) (*OneAddDelMapServerReply, error)
	OneAddDelNdpEntry(ctx context.Context, in *OneAddDelNdpEntry) (*OneAddDelNdpEntryReply, error)
	OneAddDelRemoteMapping(ctx context.Context, in *OneAddDelRemoteMapping) (*OneAddDelRemoteMappingReply, error)
	OneAdjacenciesGet(ctx context.Context, in *OneAdjacenciesGet) (*OneAdjacenciesGetReply, error)
	OneEidTableAddDelMap(ctx context.Context, in *OneEidTableAddDelMap) (*OneEidTableAddDelMapReply, error)
	OneEidTableDump(ctx context.Context, in *OneEidTableDump) (RPCService_OneEidTableDumpClient, error)
	OneEidTableMapDump(ctx context.Context, in *OneEidTableMapDump) (RPCService_OneEidTableMapDumpClient, error)
	OneEidTableVniDump(ctx context.Context, in *OneEidTableVniDump) (RPCService_OneEidTableVniDumpClient, error)
	OneEnableDisable(ctx context.Context, in *OneEnableDisable) (*OneEnableDisableReply, error)
	OneEnableDisablePetrMode(ctx context.Context, in *OneEnableDisablePetrMode) (*OneEnableDisablePetrModeReply, error)
	OneEnableDisablePitrMode(ctx context.Context, in *OneEnableDisablePitrMode) (*OneEnableDisablePitrModeReply, error)
	OneEnableDisableXtrMode(ctx context.Context, in *OneEnableDisableXtrMode) (*OneEnableDisableXtrModeReply, error)
	OneGetMapRequestItrRlocs(ctx context.Context, in *OneGetMapRequestItrRlocs) (*OneGetMapRequestItrRlocsReply, error)
	OneGetTransportProtocol(ctx context.Context, in *OneGetTransportProtocol) (*OneGetTransportProtocolReply, error)
	OneL2ArpBdGet(ctx context.Context, in *OneL2ArpBdGet) (*OneL2ArpBdGetReply, error)
	OneL2ArpEntriesGet(ctx context.Context, in *OneL2ArpEntriesGet) (*OneL2ArpEntriesGetReply, error)
	OneLocatorDump(ctx context.Context, in *OneLocatorDump) (RPCService_OneLocatorDumpClient, error)
	OneLocatorSetDump(ctx context.Context, in *OneLocatorSetDump) (RPCService_OneLocatorSetDumpClient, error)
	OneMapRegisterEnableDisable(ctx context.Context, in *OneMapRegisterEnableDisable) (*OneMapRegisterEnableDisableReply, error)
	OneMapRegisterFallbackThreshold(ctx context.Context, in *OneMapRegisterFallbackThreshold) (*OneMapRegisterFallbackThresholdReply, error)
	OneMapRegisterSetTTL(ctx context.Context, in *OneMapRegisterSetTTL) (*OneMapRegisterSetTTLReply, error)
	OneMapRequestMode(ctx context.Context, in *OneMapRequestMode) (*OneMapRequestModeReply, error)
	OneMapResolverDump(ctx context.Context, in *OneMapResolverDump) (RPCService_OneMapResolverDumpClient, error)
	OneMapServerDump(ctx context.Context, in *OneMapServerDump) (RPCService_OneMapServerDumpClient, error)
	OneNdpBdGet(ctx context.Context, in *OneNdpBdGet) (*OneNdpBdGetReply, error)
	OneNdpEntriesGet(ctx context.Context, in *OneNdpEntriesGet) (*OneNdpEntriesGetReply, error)
	OneNshSetLocatorSet(ctx context.Context, in *OneNshSetLocatorSet) (*OneNshSetLocatorSetReply, error)
	OnePitrSetLocatorSet(ctx context.Context, in *OnePitrSetLocatorSet) (*OnePitrSetLocatorSetReply, error)
	OneRlocProbeEnableDisable(ctx context.Context, in *OneRlocProbeEnableDisable) (*OneRlocProbeEnableDisableReply, error)
	OneSetTransportProtocol(ctx context.Context, in *OneSetTransportProtocol) (*OneSetTransportProtocolReply, error)
	OneShowPetrMode(ctx context.Context, in *OneShowPetrMode) (*OneShowPetrModeReply, error)
	OneShowPitrMode(ctx context.Context, in *OneShowPitrMode) (*OneShowPitrModeReply, error)
	OneShowXtrMode(ctx context.Context, in *OneShowXtrMode) (*OneShowXtrModeReply, error)
	OneStatsDump(ctx context.Context, in *OneStatsDump) (RPCService_OneStatsDumpClient, error)
	OneStatsEnableDisable(ctx context.Context, in *OneStatsEnableDisable) (*OneStatsEnableDisableReply, error)
	OneStatsFlush(ctx context.Context, in *OneStatsFlush) (*OneStatsFlushReply, error)
	OneUsePetr(ctx context.Context, in *OneUsePetr) (*OneUsePetrReply, error)
	ShowOneMapRegisterFallbackThreshold(ctx context.Context, in *ShowOneMapRegisterFallbackThreshold) (*ShowOneMapRegisterFallbackThresholdReply, error)
	ShowOneMapRegisterState(ctx context.Context, in *ShowOneMapRegisterState) (*ShowOneMapRegisterStateReply, error)
	ShowOneMapRegisterTTL(ctx context.Context, in *ShowOneMapRegisterTTL) (*ShowOneMapRegisterTTLReply, error)
	ShowOneMapRequestMode(ctx context.Context, in *ShowOneMapRequestMode) (*ShowOneMapRequestModeReply, error)
	ShowOneNshMapping(ctx context.Context, in *ShowOneNshMapping) (*ShowOneNshMappingReply, error)
	ShowOnePitr(ctx context.Context, in *ShowOnePitr) (*ShowOnePitrReply, error)
	ShowOneRlocProbeState(ctx context.Context, in *ShowOneRlocProbeState) (*ShowOneRlocProbeStateReply, error)
	ShowOneStatsEnableDisable(ctx context.Context, in *ShowOneStatsEnableDisable) (*ShowOneStatsEnableDisableReply, error)
	ShowOneStatus(ctx context.Context, in *ShowOneStatus) (*ShowOneStatusReply, error)
	ShowOneUsePetr(ctx context.Context, in *ShowOneUsePetr) (*ShowOneUsePetrReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) OneAddDelAdjacency(ctx context.Context, in *OneAddDelAdjacency) (*OneAddDelAdjacencyReply, error) {
	out := new(OneAddDelAdjacencyReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAddDelL2ArpEntry(ctx context.Context, in *OneAddDelL2ArpEntry) (*OneAddDelL2ArpEntryReply, error) {
	out := new(OneAddDelL2ArpEntryReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAddDelLocalEid(ctx context.Context, in *OneAddDelLocalEid) (*OneAddDelLocalEidReply, error) {
	out := new(OneAddDelLocalEidReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAddDelLocator(ctx context.Context, in *OneAddDelLocator) (*OneAddDelLocatorReply, error) {
	out := new(OneAddDelLocatorReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAddDelLocatorSet(ctx context.Context, in *OneAddDelLocatorSet) (*OneAddDelLocatorSetReply, error) {
	out := new(OneAddDelLocatorSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAddDelMapRequestItrRlocs(ctx context.Context, in *OneAddDelMapRequestItrRlocs) (*OneAddDelMapRequestItrRlocsReply, error) {
	out := new(OneAddDelMapRequestItrRlocsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAddDelMapResolver(ctx context.Context, in *OneAddDelMapResolver) (*OneAddDelMapResolverReply, error) {
	out := new(OneAddDelMapResolverReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAddDelMapServer(ctx context.Context, in *OneAddDelMapServer) (*OneAddDelMapServerReply, error) {
	out := new(OneAddDelMapServerReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAddDelNdpEntry(ctx context.Context, in *OneAddDelNdpEntry) (*OneAddDelNdpEntryReply, error) {
	out := new(OneAddDelNdpEntryReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAddDelRemoteMapping(ctx context.Context, in *OneAddDelRemoteMapping) (*OneAddDelRemoteMappingReply, error) {
	out := new(OneAddDelRemoteMappingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneAdjacenciesGet(ctx context.Context, in *OneAdjacenciesGet) (*OneAdjacenciesGetReply, error) {
	out := new(OneAdjacenciesGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneEidTableAddDelMap(ctx context.Context, in *OneEidTableAddDelMap) (*OneEidTableAddDelMapReply, error) {
	out := new(OneEidTableAddDelMapReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneEidTableDump(ctx context.Context, in *OneEidTableDump) (RPCService_OneEidTableDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_OneEidTableDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_OneEidTableDumpClient interface {
	Recv() (*OneEidTableDetails, error)
	api.Stream
}

type serviceClient_OneEidTableDumpClient struct {
	api.Stream
}

func (c *serviceClient_OneEidTableDumpClient) Recv() (*OneEidTableDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *OneEidTableDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) OneEidTableMapDump(ctx context.Context, in *OneEidTableMapDump) (RPCService_OneEidTableMapDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_OneEidTableMapDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_OneEidTableMapDumpClient interface {
	Recv() (*OneEidTableMapDetails, error)
	api.Stream
}

type serviceClient_OneEidTableMapDumpClient struct {
	api.Stream
}

func (c *serviceClient_OneEidTableMapDumpClient) Recv() (*OneEidTableMapDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *OneEidTableMapDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) OneEidTableVniDump(ctx context.Context, in *OneEidTableVniDump) (RPCService_OneEidTableVniDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_OneEidTableVniDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_OneEidTableVniDumpClient interface {
	Recv() (*OneEidTableVniDetails, error)
	api.Stream
}

type serviceClient_OneEidTableVniDumpClient struct {
	api.Stream
}

func (c *serviceClient_OneEidTableVniDumpClient) Recv() (*OneEidTableVniDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *OneEidTableVniDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) OneEnableDisable(ctx context.Context, in *OneEnableDisable) (*OneEnableDisableReply, error) {
	out := new(OneEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneEnableDisablePetrMode(ctx context.Context, in *OneEnableDisablePetrMode) (*OneEnableDisablePetrModeReply, error) {
	out := new(OneEnableDisablePetrModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneEnableDisablePitrMode(ctx context.Context, in *OneEnableDisablePitrMode) (*OneEnableDisablePitrModeReply, error) {
	out := new(OneEnableDisablePitrModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneEnableDisableXtrMode(ctx context.Context, in *OneEnableDisableXtrMode) (*OneEnableDisableXtrModeReply, error) {
	out := new(OneEnableDisableXtrModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneGetMapRequestItrRlocs(ctx context.Context, in *OneGetMapRequestItrRlocs) (*OneGetMapRequestItrRlocsReply, error) {
	out := new(OneGetMapRequestItrRlocsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneGetTransportProtocol(ctx context.Context, in *OneGetTransportProtocol) (*OneGetTransportProtocolReply, error) {
	out := new(OneGetTransportProtocolReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneL2ArpBdGet(ctx context.Context, in *OneL2ArpBdGet) (*OneL2ArpBdGetReply, error) {
	out := new(OneL2ArpBdGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneL2ArpEntriesGet(ctx context.Context, in *OneL2ArpEntriesGet) (*OneL2ArpEntriesGetReply, error) {
	out := new(OneL2ArpEntriesGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneLocatorDump(ctx context.Context, in *OneLocatorDump) (RPCService_OneLocatorDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_OneLocatorDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_OneLocatorDumpClient interface {
	Recv() (*OneLocatorDetails, error)
	api.Stream
}

type serviceClient_OneLocatorDumpClient struct {
	api.Stream
}

func (c *serviceClient_OneLocatorDumpClient) Recv() (*OneLocatorDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *OneLocatorDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) OneLocatorSetDump(ctx context.Context, in *OneLocatorSetDump) (RPCService_OneLocatorSetDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_OneLocatorSetDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_OneLocatorSetDumpClient interface {
	Recv() (*OneLocatorSetDetails, error)
	api.Stream
}

type serviceClient_OneLocatorSetDumpClient struct {
	api.Stream
}

func (c *serviceClient_OneLocatorSetDumpClient) Recv() (*OneLocatorSetDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *OneLocatorSetDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) OneMapRegisterEnableDisable(ctx context.Context, in *OneMapRegisterEnableDisable) (*OneMapRegisterEnableDisableReply, error) {
	out := new(OneMapRegisterEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneMapRegisterFallbackThreshold(ctx context.Context, in *OneMapRegisterFallbackThreshold) (*OneMapRegisterFallbackThresholdReply, error) {
	out := new(OneMapRegisterFallbackThresholdReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneMapRegisterSetTTL(ctx context.Context, in *OneMapRegisterSetTTL) (*OneMapRegisterSetTTLReply, error) {
	out := new(OneMapRegisterSetTTLReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneMapRequestMode(ctx context.Context, in *OneMapRequestMode) (*OneMapRequestModeReply, error) {
	out := new(OneMapRequestModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneMapResolverDump(ctx context.Context, in *OneMapResolverDump) (RPCService_OneMapResolverDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_OneMapResolverDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_OneMapResolverDumpClient interface {
	Recv() (*OneMapResolverDetails, error)
	api.Stream
}

type serviceClient_OneMapResolverDumpClient struct {
	api.Stream
}

func (c *serviceClient_OneMapResolverDumpClient) Recv() (*OneMapResolverDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *OneMapResolverDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) OneMapServerDump(ctx context.Context, in *OneMapServerDump) (RPCService_OneMapServerDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_OneMapServerDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_OneMapServerDumpClient interface {
	Recv() (*OneMapServerDetails, error)
	api.Stream
}

type serviceClient_OneMapServerDumpClient struct {
	api.Stream
}

func (c *serviceClient_OneMapServerDumpClient) Recv() (*OneMapServerDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *OneMapServerDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) OneNdpBdGet(ctx context.Context, in *OneNdpBdGet) (*OneNdpBdGetReply, error) {
	out := new(OneNdpBdGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneNdpEntriesGet(ctx context.Context, in *OneNdpEntriesGet) (*OneNdpEntriesGetReply, error) {
	out := new(OneNdpEntriesGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneNshSetLocatorSet(ctx context.Context, in *OneNshSetLocatorSet) (*OneNshSetLocatorSetReply, error) {
	out := new(OneNshSetLocatorSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OnePitrSetLocatorSet(ctx context.Context, in *OnePitrSetLocatorSet) (*OnePitrSetLocatorSetReply, error) {
	out := new(OnePitrSetLocatorSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneRlocProbeEnableDisable(ctx context.Context, in *OneRlocProbeEnableDisable) (*OneRlocProbeEnableDisableReply, error) {
	out := new(OneRlocProbeEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneSetTransportProtocol(ctx context.Context, in *OneSetTransportProtocol) (*OneSetTransportProtocolReply, error) {
	out := new(OneSetTransportProtocolReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneShowPetrMode(ctx context.Context, in *OneShowPetrMode) (*OneShowPetrModeReply, error) {
	out := new(OneShowPetrModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneShowPitrMode(ctx context.Context, in *OneShowPitrMode) (*OneShowPitrModeReply, error) {
	out := new(OneShowPitrModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneShowXtrMode(ctx context.Context, in *OneShowXtrMode) (*OneShowXtrModeReply, error) {
	out := new(OneShowXtrModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneStatsDump(ctx context.Context, in *OneStatsDump) (RPCService_OneStatsDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_OneStatsDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_OneStatsDumpClient interface {
	Recv() (*OneStatsDetails, error)
	api.Stream
}

type serviceClient_OneStatsDumpClient struct {
	api.Stream
}

func (c *serviceClient_OneStatsDumpClient) Recv() (*OneStatsDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *OneStatsDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) OneStatsEnableDisable(ctx context.Context, in *OneStatsEnableDisable) (*OneStatsEnableDisableReply, error) {
	out := new(OneStatsEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneStatsFlush(ctx context.Context, in *OneStatsFlush) (*OneStatsFlushReply, error) {
	out := new(OneStatsFlushReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OneUsePetr(ctx context.Context, in *OneUsePetr) (*OneUsePetrReply, error) {
	out := new(OneUsePetrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOneMapRegisterFallbackThreshold(ctx context.Context, in *ShowOneMapRegisterFallbackThreshold) (*ShowOneMapRegisterFallbackThresholdReply, error) {
	out := new(ShowOneMapRegisterFallbackThresholdReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOneMapRegisterState(ctx context.Context, in *ShowOneMapRegisterState) (*ShowOneMapRegisterStateReply, error) {
	out := new(ShowOneMapRegisterStateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOneMapRegisterTTL(ctx context.Context, in *ShowOneMapRegisterTTL) (*ShowOneMapRegisterTTLReply, error) {
	out := new(ShowOneMapRegisterTTLReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOneMapRequestMode(ctx context.Context, in *ShowOneMapRequestMode) (*ShowOneMapRequestModeReply, error) {
	out := new(ShowOneMapRequestModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOneNshMapping(ctx context.Context, in *ShowOneNshMapping) (*ShowOneNshMappingReply, error) {
	out := new(ShowOneNshMappingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOnePitr(ctx context.Context, in *ShowOnePitr) (*ShowOnePitrReply, error) {
	out := new(ShowOnePitrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOneRlocProbeState(ctx context.Context, in *ShowOneRlocProbeState) (*ShowOneRlocProbeStateReply, error) {
	out := new(ShowOneRlocProbeStateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOneStatsEnableDisable(ctx context.Context, in *ShowOneStatsEnableDisable) (*ShowOneStatsEnableDisableReply, error) {
	out := new(ShowOneStatsEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOneStatus(ctx context.Context, in *ShowOneStatus) (*ShowOneStatusReply, error) {
	out := new(ShowOneStatusReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowOneUsePetr(ctx context.Context, in *ShowOneUsePetr) (*ShowOneUsePetrReply, error) {
	out := new(ShowOneUsePetrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
