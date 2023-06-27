package connectormanager

import (
	"context"
	"errors"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/util/goutil"
)

func (m manager) Start(ctx context.Context) error {
	addr, peer := m.connectToExistingPeer(ctx)
	if peer != nil {
		goutil.PanicUnhandledError(m.peerRepo.Save(ctx, addr, peer))
	}
	return nil
}

func (m manager) connectToExistingPeer(ctx context.Context) (string, peerconnectorcontract.Peer) {
	var peer peerconnectorcontract.Peer
	var err error
	var addr string
	for _, addr = range m.existingPeerAddresses {
		peer, err = m.client.ConnectToPeer(ctx, addr)
		if err == nil {
			break
		}
		if !errors.Is(err, peerclientcontract.ConnectionError{}) {
			goutil.PanicUnhandledError(err)
		}
	}
	return addr, peer
}