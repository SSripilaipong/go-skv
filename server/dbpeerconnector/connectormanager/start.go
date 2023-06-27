package connectormanager

import (
	"errors"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/util/goutil"
)

func (c manager) Start() error {
	addr, peer := c.connectToExistingPeer()
	if peer != nil {
		goutil.PanicUnhandledError(c.peerRepo.Save(c.ctx, addr, peer))
	}
	return nil
}

func (c manager) connectToExistingPeer() (string, peerconnectorcontract.Peer) {
	var peer peerconnectorcontract.Peer
	var err error
	var addr string
	for _, addr = range c.existingPeerAddresses {
		peer, err = c.client.ConnectToPeer(c.subCtx, addr)
		if err == nil {
			break
		}
		if !errors.Is(err, peerclientcontract.ConnectionError{}) {
			goutil.PanicUnhandledError(err)
		}
	}
	return addr, peer
}
