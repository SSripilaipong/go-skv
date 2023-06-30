package dbpeerconnectortest

import "go-skv/server/dbpeerconnector/peerconnectorcontract"

type PeerMock struct {
	UpdateReplica_key   string
	UpdateReplica_value string
}

func (p *PeerMock) UpdateReplica(key string, value string) error {
	p.UpdateReplica_key = key
	p.UpdateReplica_value = value
	return nil
}

var _ peerconnectorcontract.Peer = &PeerMock{}
