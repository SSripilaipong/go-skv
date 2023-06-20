package peerconnector

import (
	"errors"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/util/goutil"
)

func (c connector) Start() error {
	_ = c.connectToExistingPeer()
	return nil
}

func (c connector) connectToExistingPeer() peerconnectorcontract.Peer {
	var peer peerconnectorcontract.Peer
	var err error
	for _, addr := range c.existingPeerAddresses {
		peer, err = c.client.ConnectToPeer(addr)
		if err == nil {
			break
		}
		if !errors.Is(err, peerclientcontract.ConnectionError{}) {
			goutil.PanicUnhandledError(err)
		}
	}
	return peer
}
