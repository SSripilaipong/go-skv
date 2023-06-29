package connectormanagertest

import "go-skv/server/dbpeerconnector/peerconnectorcontract"

type PeerMock struct{}

func (p *PeerMock) UpdateReplica(key string, value string) error {
	return nil
}

var _ peerconnectorcontract.Peer = &PeerMock{}
