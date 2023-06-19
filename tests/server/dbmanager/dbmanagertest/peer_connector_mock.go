package dbmanagertest

import "go-skv/server/dbpeerconnector/peerconnectorcontract"

type PeerConnectorMock struct {
	Start_IsCalled bool
	Stop_IsCalled  bool
}

func (p *PeerConnectorMock) Start() error {
	p.Start_IsCalled = true
	return nil
}

func (p *PeerConnectorMock) Stop() error {
	p.Stop_IsCalled = true
	return nil
}

var _ peerconnectorcontract.Connector = &PeerConnectorMock{}
