package dbmanagertest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type PeerConnectorMock struct {
	Start_IsCalled bool
	Join_IsCalled  bool
}

func (p *PeerConnectorMock) Start(context.Context) error {
	p.Start_IsCalled = true
	return nil
}

func (p *PeerConnectorMock) Join() error {
	p.Join_IsCalled = true
	return nil
}

var _ peerconnectorcontract.Connector = &PeerConnectorMock{}
