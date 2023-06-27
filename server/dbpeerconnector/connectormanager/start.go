package connectormanager

import (
	"errors"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/util/goutil"
)

func (m manager) Start() error {
	addr, peer := m.connectToExistingPeer()
	if peer != nil {
		goutil.PanicUnhandledError(m.peerRepo.Save(m.ctx, addr, peer))
	}
	return nil
}

func (m manager) connectToExistingPeer() (string, peerconnectorcontract.Peer) {
	var peer peerconnectorcontract.Peer
	var err error
	var addr string
	for _, addr = range m.existingPeerAddresses {
		peer, err = m.client.ConnectToPeer(m.subCtx, addr)
		if err == nil {
			break
		}
		if !errors.Is(err, peerclientcontract.ConnectionError{}) {
			goutil.PanicUnhandledError(err)
		}
	}
	return addr, peer
}
