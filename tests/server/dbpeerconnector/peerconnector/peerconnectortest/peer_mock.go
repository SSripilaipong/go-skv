package peerconnectortest

import "go-skv/server/dbpeerconnector/peerconnectorcontract"

type PeerMock struct {
	SubscribeUpdates_listener peerconnectorcontract.UpdateListener
}

func (p *PeerMock) SubscribeUpdates(listener peerconnectorcontract.UpdateListener) error {
	p.SubscribeUpdates_listener = listener
	return nil
}

var _ peerconnectorcontract.Peer = &PeerMock{}
