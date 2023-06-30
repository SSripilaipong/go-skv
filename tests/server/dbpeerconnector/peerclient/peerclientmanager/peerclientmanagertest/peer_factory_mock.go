package peerclientmanagertest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type PeerFactoryMock struct {
	New_ctx context.Context
}

func (p *PeerFactoryMock) New(ctx context.Context) (peerconnectorcontract.Peer, error) {
	p.New_ctx = ctx
	return nil, nil
}

var _ clientsidepeercontract.Factory = &PeerFactoryMock{}
