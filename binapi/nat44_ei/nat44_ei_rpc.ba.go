// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package nat44_ei

import (
	"context"
	"fmt"
	"io"

	memclnt "github.com/ljkiraly/govpp/binapi/memclnt"
	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service nat44_ei.
type RPCService interface {
	Nat44EiAddDelAddressRange(ctx context.Context, in *Nat44EiAddDelAddressRange) (*Nat44EiAddDelAddressRangeReply, error)
	Nat44EiAddDelIdentityMapping(ctx context.Context, in *Nat44EiAddDelIdentityMapping) (*Nat44EiAddDelIdentityMappingReply, error)
	Nat44EiAddDelInterfaceAddr(ctx context.Context, in *Nat44EiAddDelInterfaceAddr) (*Nat44EiAddDelInterfaceAddrReply, error)
	Nat44EiAddDelOutputInterface(ctx context.Context, in *Nat44EiAddDelOutputInterface) (*Nat44EiAddDelOutputInterfaceReply, error)
	Nat44EiAddDelStaticMapping(ctx context.Context, in *Nat44EiAddDelStaticMapping) (*Nat44EiAddDelStaticMappingReply, error)
	Nat44EiAddressDump(ctx context.Context, in *Nat44EiAddressDump) (RPCService_Nat44EiAddressDumpClient, error)
	Nat44EiDelSession(ctx context.Context, in *Nat44EiDelSession) (*Nat44EiDelSessionReply, error)
	Nat44EiDelUser(ctx context.Context, in *Nat44EiDelUser) (*Nat44EiDelUserReply, error)
	Nat44EiForwardingEnableDisable(ctx context.Context, in *Nat44EiForwardingEnableDisable) (*Nat44EiForwardingEnableDisableReply, error)
	Nat44EiGetAddrAndPortAllocAlg(ctx context.Context, in *Nat44EiGetAddrAndPortAllocAlg) (*Nat44EiGetAddrAndPortAllocAlgReply, error)
	Nat44EiGetMssClamping(ctx context.Context, in *Nat44EiGetMssClamping) (*Nat44EiGetMssClampingReply, error)
	Nat44EiHaFlush(ctx context.Context, in *Nat44EiHaFlush) (*Nat44EiHaFlushReply, error)
	Nat44EiHaGetFailover(ctx context.Context, in *Nat44EiHaGetFailover) (*Nat44EiHaGetFailoverReply, error)
	Nat44EiHaGetListener(ctx context.Context, in *Nat44EiHaGetListener) (*Nat44EiHaGetListenerReply, error)
	Nat44EiHaResync(ctx context.Context, in *Nat44EiHaResync) (*Nat44EiHaResyncReply, error)
	Nat44EiHaSetFailover(ctx context.Context, in *Nat44EiHaSetFailover) (*Nat44EiHaSetFailoverReply, error)
	Nat44EiHaSetListener(ctx context.Context, in *Nat44EiHaSetListener) (*Nat44EiHaSetListenerReply, error)
	Nat44EiIdentityMappingDump(ctx context.Context, in *Nat44EiIdentityMappingDump) (RPCService_Nat44EiIdentityMappingDumpClient, error)
	Nat44EiInterfaceAddDelFeature(ctx context.Context, in *Nat44EiInterfaceAddDelFeature) (*Nat44EiInterfaceAddDelFeatureReply, error)
	Nat44EiInterfaceAddDelOutputFeature(ctx context.Context, in *Nat44EiInterfaceAddDelOutputFeature) (*Nat44EiInterfaceAddDelOutputFeatureReply, error)
	Nat44EiInterfaceAddrDump(ctx context.Context, in *Nat44EiInterfaceAddrDump) (RPCService_Nat44EiInterfaceAddrDumpClient, error)
	Nat44EiInterfaceDump(ctx context.Context, in *Nat44EiInterfaceDump) (RPCService_Nat44EiInterfaceDumpClient, error)
	Nat44EiInterfaceOutputFeatureDump(ctx context.Context, in *Nat44EiInterfaceOutputFeatureDump) (RPCService_Nat44EiInterfaceOutputFeatureDumpClient, error)
	Nat44EiIpfixEnableDisable(ctx context.Context, in *Nat44EiIpfixEnableDisable) (*Nat44EiIpfixEnableDisableReply, error)
	Nat44EiOutputInterfaceGet(ctx context.Context, in *Nat44EiOutputInterfaceGet) (RPCService_Nat44EiOutputInterfaceGetClient, error)
	Nat44EiPluginEnableDisable(ctx context.Context, in *Nat44EiPluginEnableDisable) (*Nat44EiPluginEnableDisableReply, error)
	Nat44EiSetAddrAndPortAllocAlg(ctx context.Context, in *Nat44EiSetAddrAndPortAllocAlg) (*Nat44EiSetAddrAndPortAllocAlgReply, error)
	Nat44EiSetFqOptions(ctx context.Context, in *Nat44EiSetFqOptions) (*Nat44EiSetFqOptionsReply, error)
	Nat44EiSetLogLevel(ctx context.Context, in *Nat44EiSetLogLevel) (*Nat44EiSetLogLevelReply, error)
	Nat44EiSetMssClamping(ctx context.Context, in *Nat44EiSetMssClamping) (*Nat44EiSetMssClampingReply, error)
	Nat44EiSetTimeouts(ctx context.Context, in *Nat44EiSetTimeouts) (*Nat44EiSetTimeoutsReply, error)
	Nat44EiSetWorkers(ctx context.Context, in *Nat44EiSetWorkers) (*Nat44EiSetWorkersReply, error)
	Nat44EiShowFqOptions(ctx context.Context, in *Nat44EiShowFqOptions) (*Nat44EiShowFqOptionsReply, error)
	Nat44EiShowRunningConfig(ctx context.Context, in *Nat44EiShowRunningConfig) (*Nat44EiShowRunningConfigReply, error)
	Nat44EiStaticMappingDump(ctx context.Context, in *Nat44EiStaticMappingDump) (RPCService_Nat44EiStaticMappingDumpClient, error)
	Nat44EiUserDump(ctx context.Context, in *Nat44EiUserDump) (RPCService_Nat44EiUserDumpClient, error)
	Nat44EiUserSessionDump(ctx context.Context, in *Nat44EiUserSessionDump) (RPCService_Nat44EiUserSessionDumpClient, error)
	Nat44EiUserSessionV2Dump(ctx context.Context, in *Nat44EiUserSessionV2Dump) (RPCService_Nat44EiUserSessionV2DumpClient, error)
	Nat44EiWorkerDump(ctx context.Context, in *Nat44EiWorkerDump) (RPCService_Nat44EiWorkerDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) Nat44EiAddDelAddressRange(ctx context.Context, in *Nat44EiAddDelAddressRange) (*Nat44EiAddDelAddressRangeReply, error) {
	out := new(Nat44EiAddDelAddressRangeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiAddDelIdentityMapping(ctx context.Context, in *Nat44EiAddDelIdentityMapping) (*Nat44EiAddDelIdentityMappingReply, error) {
	out := new(Nat44EiAddDelIdentityMappingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiAddDelInterfaceAddr(ctx context.Context, in *Nat44EiAddDelInterfaceAddr) (*Nat44EiAddDelInterfaceAddrReply, error) {
	out := new(Nat44EiAddDelInterfaceAddrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiAddDelOutputInterface(ctx context.Context, in *Nat44EiAddDelOutputInterface) (*Nat44EiAddDelOutputInterfaceReply, error) {
	out := new(Nat44EiAddDelOutputInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiAddDelStaticMapping(ctx context.Context, in *Nat44EiAddDelStaticMapping) (*Nat44EiAddDelStaticMappingReply, error) {
	out := new(Nat44EiAddDelStaticMappingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiAddressDump(ctx context.Context, in *Nat44EiAddressDump) (RPCService_Nat44EiAddressDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiAddressDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiAddressDumpClient interface {
	Recv() (*Nat44EiAddressDetails, error)
	api.Stream
}

type serviceClient_Nat44EiAddressDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiAddressDumpClient) Recv() (*Nat44EiAddressDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiAddressDetails:
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

func (c *serviceClient) Nat44EiDelSession(ctx context.Context, in *Nat44EiDelSession) (*Nat44EiDelSessionReply, error) {
	out := new(Nat44EiDelSessionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiDelUser(ctx context.Context, in *Nat44EiDelUser) (*Nat44EiDelUserReply, error) {
	out := new(Nat44EiDelUserReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiForwardingEnableDisable(ctx context.Context, in *Nat44EiForwardingEnableDisable) (*Nat44EiForwardingEnableDisableReply, error) {
	out := new(Nat44EiForwardingEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiGetAddrAndPortAllocAlg(ctx context.Context, in *Nat44EiGetAddrAndPortAllocAlg) (*Nat44EiGetAddrAndPortAllocAlgReply, error) {
	out := new(Nat44EiGetAddrAndPortAllocAlgReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiGetMssClamping(ctx context.Context, in *Nat44EiGetMssClamping) (*Nat44EiGetMssClampingReply, error) {
	out := new(Nat44EiGetMssClampingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiHaFlush(ctx context.Context, in *Nat44EiHaFlush) (*Nat44EiHaFlushReply, error) {
	out := new(Nat44EiHaFlushReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiHaGetFailover(ctx context.Context, in *Nat44EiHaGetFailover) (*Nat44EiHaGetFailoverReply, error) {
	out := new(Nat44EiHaGetFailoverReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiHaGetListener(ctx context.Context, in *Nat44EiHaGetListener) (*Nat44EiHaGetListenerReply, error) {
	out := new(Nat44EiHaGetListenerReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiHaResync(ctx context.Context, in *Nat44EiHaResync) (*Nat44EiHaResyncReply, error) {
	out := new(Nat44EiHaResyncReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiHaSetFailover(ctx context.Context, in *Nat44EiHaSetFailover) (*Nat44EiHaSetFailoverReply, error) {
	out := new(Nat44EiHaSetFailoverReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiHaSetListener(ctx context.Context, in *Nat44EiHaSetListener) (*Nat44EiHaSetListenerReply, error) {
	out := new(Nat44EiHaSetListenerReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiIdentityMappingDump(ctx context.Context, in *Nat44EiIdentityMappingDump) (RPCService_Nat44EiIdentityMappingDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiIdentityMappingDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiIdentityMappingDumpClient interface {
	Recv() (*Nat44EiIdentityMappingDetails, error)
	api.Stream
}

type serviceClient_Nat44EiIdentityMappingDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiIdentityMappingDumpClient) Recv() (*Nat44EiIdentityMappingDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiIdentityMappingDetails:
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

func (c *serviceClient) Nat44EiInterfaceAddDelFeature(ctx context.Context, in *Nat44EiInterfaceAddDelFeature) (*Nat44EiInterfaceAddDelFeatureReply, error) {
	out := new(Nat44EiInterfaceAddDelFeatureReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiInterfaceAddDelOutputFeature(ctx context.Context, in *Nat44EiInterfaceAddDelOutputFeature) (*Nat44EiInterfaceAddDelOutputFeatureReply, error) {
	out := new(Nat44EiInterfaceAddDelOutputFeatureReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiInterfaceAddrDump(ctx context.Context, in *Nat44EiInterfaceAddrDump) (RPCService_Nat44EiInterfaceAddrDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiInterfaceAddrDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiInterfaceAddrDumpClient interface {
	Recv() (*Nat44EiInterfaceAddrDetails, error)
	api.Stream
}

type serviceClient_Nat44EiInterfaceAddrDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiInterfaceAddrDumpClient) Recv() (*Nat44EiInterfaceAddrDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiInterfaceAddrDetails:
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

func (c *serviceClient) Nat44EiInterfaceDump(ctx context.Context, in *Nat44EiInterfaceDump) (RPCService_Nat44EiInterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiInterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiInterfaceDumpClient interface {
	Recv() (*Nat44EiInterfaceDetails, error)
	api.Stream
}

type serviceClient_Nat44EiInterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiInterfaceDumpClient) Recv() (*Nat44EiInterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiInterfaceDetails:
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

func (c *serviceClient) Nat44EiInterfaceOutputFeatureDump(ctx context.Context, in *Nat44EiInterfaceOutputFeatureDump) (RPCService_Nat44EiInterfaceOutputFeatureDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiInterfaceOutputFeatureDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiInterfaceOutputFeatureDumpClient interface {
	Recv() (*Nat44EiInterfaceOutputFeatureDetails, error)
	api.Stream
}

type serviceClient_Nat44EiInterfaceOutputFeatureDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiInterfaceOutputFeatureDumpClient) Recv() (*Nat44EiInterfaceOutputFeatureDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiInterfaceOutputFeatureDetails:
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

func (c *serviceClient) Nat44EiIpfixEnableDisable(ctx context.Context, in *Nat44EiIpfixEnableDisable) (*Nat44EiIpfixEnableDisableReply, error) {
	out := new(Nat44EiIpfixEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiOutputInterfaceGet(ctx context.Context, in *Nat44EiOutputInterfaceGet) (RPCService_Nat44EiOutputInterfaceGetClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiOutputInterfaceGetClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiOutputInterfaceGetClient interface {
	Recv() (*Nat44EiOutputInterfaceDetails, *Nat44EiOutputInterfaceGetReply, error)
	api.Stream
}

type serviceClient_Nat44EiOutputInterfaceGetClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiOutputInterfaceGetClient) Recv() (*Nat44EiOutputInterfaceDetails, *Nat44EiOutputInterfaceGetReply, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiOutputInterfaceDetails:
		return m, nil, nil
	case *Nat44EiOutputInterfaceGetReply:
		if err := api.RetvalToVPPApiError(m.Retval); err != nil {
			c.Stream.Close()
			return nil, m, err
		}
		err = c.Stream.Close()
		if err != nil {
			return nil, m, err
		}
		return nil, m, io.EOF
	default:
		return nil, nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat44EiPluginEnableDisable(ctx context.Context, in *Nat44EiPluginEnableDisable) (*Nat44EiPluginEnableDisableReply, error) {
	out := new(Nat44EiPluginEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiSetAddrAndPortAllocAlg(ctx context.Context, in *Nat44EiSetAddrAndPortAllocAlg) (*Nat44EiSetAddrAndPortAllocAlgReply, error) {
	out := new(Nat44EiSetAddrAndPortAllocAlgReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiSetFqOptions(ctx context.Context, in *Nat44EiSetFqOptions) (*Nat44EiSetFqOptionsReply, error) {
	out := new(Nat44EiSetFqOptionsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiSetLogLevel(ctx context.Context, in *Nat44EiSetLogLevel) (*Nat44EiSetLogLevelReply, error) {
	out := new(Nat44EiSetLogLevelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiSetMssClamping(ctx context.Context, in *Nat44EiSetMssClamping) (*Nat44EiSetMssClampingReply, error) {
	out := new(Nat44EiSetMssClampingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiSetTimeouts(ctx context.Context, in *Nat44EiSetTimeouts) (*Nat44EiSetTimeoutsReply, error) {
	out := new(Nat44EiSetTimeoutsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiSetWorkers(ctx context.Context, in *Nat44EiSetWorkers) (*Nat44EiSetWorkersReply, error) {
	out := new(Nat44EiSetWorkersReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiShowFqOptions(ctx context.Context, in *Nat44EiShowFqOptions) (*Nat44EiShowFqOptionsReply, error) {
	out := new(Nat44EiShowFqOptionsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiShowRunningConfig(ctx context.Context, in *Nat44EiShowRunningConfig) (*Nat44EiShowRunningConfigReply, error) {
	out := new(Nat44EiShowRunningConfigReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EiStaticMappingDump(ctx context.Context, in *Nat44EiStaticMappingDump) (RPCService_Nat44EiStaticMappingDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiStaticMappingDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiStaticMappingDumpClient interface {
	Recv() (*Nat44EiStaticMappingDetails, error)
	api.Stream
}

type serviceClient_Nat44EiStaticMappingDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiStaticMappingDumpClient) Recv() (*Nat44EiStaticMappingDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiStaticMappingDetails:
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

func (c *serviceClient) Nat44EiUserDump(ctx context.Context, in *Nat44EiUserDump) (RPCService_Nat44EiUserDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiUserDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiUserDumpClient interface {
	Recv() (*Nat44EiUserDetails, error)
	api.Stream
}

type serviceClient_Nat44EiUserDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiUserDumpClient) Recv() (*Nat44EiUserDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiUserDetails:
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

func (c *serviceClient) Nat44EiUserSessionDump(ctx context.Context, in *Nat44EiUserSessionDump) (RPCService_Nat44EiUserSessionDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiUserSessionDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiUserSessionDumpClient interface {
	Recv() (*Nat44EiUserSessionDetails, error)
	api.Stream
}

type serviceClient_Nat44EiUserSessionDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiUserSessionDumpClient) Recv() (*Nat44EiUserSessionDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiUserSessionDetails:
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

func (c *serviceClient) Nat44EiUserSessionV2Dump(ctx context.Context, in *Nat44EiUserSessionV2Dump) (RPCService_Nat44EiUserSessionV2DumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiUserSessionV2DumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiUserSessionV2DumpClient interface {
	Recv() (*Nat44EiUserSessionV2Details, error)
	api.Stream
}

type serviceClient_Nat44EiUserSessionV2DumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiUserSessionV2DumpClient) Recv() (*Nat44EiUserSessionV2Details, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiUserSessionV2Details:
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

func (c *serviceClient) Nat44EiWorkerDump(ctx context.Context, in *Nat44EiWorkerDump) (RPCService_Nat44EiWorkerDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44EiWorkerDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44EiWorkerDumpClient interface {
	Recv() (*Nat44EiWorkerDetails, error)
	api.Stream
}

type serviceClient_Nat44EiWorkerDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44EiWorkerDumpClient) Recv() (*Nat44EiWorkerDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44EiWorkerDetails:
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
