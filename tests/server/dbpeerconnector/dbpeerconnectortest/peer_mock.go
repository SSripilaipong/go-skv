package dbpeerconnectortest

import "go-skv/server/dbpeerconnector/peerconnectorcontract"

type PeerMock struct {
	UpdateReplica_key   string
	UpdateReplica_value string
}

func (p *PeerMock) UpdateReplicaFromPeer(key string, value string) error {
	p.UpdateReplica_key = key
	p.UpdateReplica_value = value
	return nil
}

func (p *PeerMock) Join() error {
	return nil
}

var _ peerconnectorcontract.Peer = &PeerMock{}
