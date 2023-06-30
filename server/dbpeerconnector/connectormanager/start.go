package connectormanager

import (
	"context"
	"errors"
	"fmt"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/util/goutil"
)

func (m manager) Start(ctx context.Context) error {
	addr, peer := m.connectToExistingPeer(ctx)
	if peer != nil {
		fmt.Printf("PeerRepo: %#v\n", m.peerRepo)
		goutil.PanicUnhandledError(m.peerRepo.Save(ctx, addr, peer))
	}
	goutil.PanicUnhandledError(m.server.Start(ctx))
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
