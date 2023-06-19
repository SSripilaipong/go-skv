package peerconnector

import "go-skv/util/goutil"

func (p connector) Start() error {
	_, err := p.client.ConnectToPeer(p.existingPeerAddresses[0])
	goutil.PanicUnhandledError(err)
	return nil
}
