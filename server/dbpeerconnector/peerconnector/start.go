package peerconnector

import (
	"errors"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/util/goutil"
)

func (p connector) Start() error {
	for _, addr := range p.existingPeerAddresses {
		_, err := p.client.ConnectToPeer(addr)
		if !errors.Is(err, peerclientcontract.ConnectionError{}) {
			goutil.PanicUnhandledError(err)
		}
	}
	return nil
}
