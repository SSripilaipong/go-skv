package peerclientmanager

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/util/goutil"
)

func (c client) ConnectToPeer(ctx context.Context, address string) (peerconnectorcontract.Peer, error) {
	_, err := c.peerFactory.New(ctx)
	goutil.PanicUnhandledError(err)
	return nil, nil
}
