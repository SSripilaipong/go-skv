package connectormanagertest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/util/goutil"
)

type PeerClientMock struct {
	ConnectToPeer_ctx_array     []context.Context
	ConnectToPeer_address_array []string
	ConnectToPeer_Return_array  []peerconnectorcontract.Peer
	ConnectToPeer_Error_array   []error
	ConnectToPeer_Call_index    int
}

func (c *PeerClientMock) ConnectToPeer(ctx context.Context, address string) (peerconnectorcontract.Peer, error) {
	c.ConnectToPeer_ctx_array = append(c.ConnectToPeer_ctx_array, ctx)
	c.ConnectToPeer_address_array = append(c.ConnectToPeer_address_array, address)

	r, _ := goutil.ElementAt(c.ConnectToPeer_Return_array, c.ConnectToPeer_Call_index)
	e, _ := goutil.ElementAt(c.ConnectToPeer_Error_array, c.ConnectToPeer_Call_index)

	c.ConnectToPeer_Call_index += 1
	return r, e
}

var _ peerclientcontract.Client = &PeerClientMock{}
