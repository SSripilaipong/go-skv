package peerconnector

import (
	"errors"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/util/goutil"
)

func (p connector) Start() error {
	_ = p.connectToExistingPeer()
	return nil
}

func (p connector) connectToExistingPeer() peerconnectorcontract.Peer {
	var peer peerconnectorcontract.Peer
	var err error
	for _, addr := range p.existingPeerAddresses {
		peer, err = p.client.ConnectToPeer(addr)
		if err == nil {
			break
		}
		if !errors.Is(err, peerclientcontract.ConnectionError{}) {
			goutil.PanicUnhandledError(err)
		}
	}
	return peer
}
