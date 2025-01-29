// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package pot

import (
	"context"
	"fmt"
	"io"

	memclnt "github.com/ljkiraly/govpp/binapi/memclnt"
	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service pot.
type RPCService interface {
	PotProfileActivate(ctx context.Context, in *PotProfileActivate) (*PotProfileActivateReply, error)
	PotProfileAdd(ctx context.Context, in *PotProfileAdd) (*PotProfileAddReply, error)
	PotProfileDel(ctx context.Context, in *PotProfileDel) (*PotProfileDelReply, error)
	PotProfileShowConfigDump(ctx context.Context, in *PotProfileShowConfigDump) (RPCService_PotProfileShowConfigDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) PotProfileActivate(ctx context.Context, in *PotProfileActivate) (*PotProfileActivateReply, error) {
	out := new(PotProfileActivateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PotProfileAdd(ctx context.Context, in *PotProfileAdd) (*PotProfileAddReply, error) {
	out := new(PotProfileAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PotProfileDel(ctx context.Context, in *PotProfileDel) (*PotProfileDelReply, error) {
	out := new(PotProfileDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PotProfileShowConfigDump(ctx context.Context, in *PotProfileShowConfigDump) (RPCService_PotProfileShowConfigDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_PotProfileShowConfigDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_PotProfileShowConfigDumpClient interface {
	Recv() (*PotProfileShowConfigDetails, error)
	api.Stream
}

type serviceClient_PotProfileShowConfigDumpClient struct {
	api.Stream
}

func (c *serviceClient_PotProfileShowConfigDumpClient) Recv() (*PotProfileShowConfigDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *PotProfileShowConfigDetails:
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
