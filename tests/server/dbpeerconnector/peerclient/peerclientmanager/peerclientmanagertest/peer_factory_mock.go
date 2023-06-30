package peerclientmanagertest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type PeerFactoryMock struct {
	New_ctx    context.Context
	New_Return peerconnectorcontract.Peer
}

func (p *PeerFactoryMock) New(ctx context.Context) (peerconnectorcontract.Peer, error) {
	p.New_ctx = ctx
	return p.New_Return, nil
}

var _ clientsidepeercontract.Factory = &PeerFactoryMock{}
