package connectormanagertest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerserver/peerservercontract"
)

type PeerServerMock struct {
	Start_ctx      context.Context
	Start_IsCalled bool
}

func (s *PeerServerMock) Start(ctx context.Context) error {
	s.Start_IsCalled = true
	s.Start_ctx = ctx
	return nil
}

var _ peerservercontract.Server = &PeerServerMock{}
