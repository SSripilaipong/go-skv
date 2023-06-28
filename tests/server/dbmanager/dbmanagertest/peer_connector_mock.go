package dbmanagertest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type PeerConnectorMock struct {
	Start_ctx     context.Context
	Join_IsCalled bool
}

func (p *PeerConnectorMock) Start(ctx context.Context) error {
	p.Start_ctx = ctx
	return nil
}

func (p *PeerConnectorMock) Join() error {
	p.Join_IsCalled = true
	return nil
}

var _ peerconnectorcontract.Connector = &PeerConnectorMock{}
