package peerconnectortest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
)

type PeerRepositoryMock struct {
	Save_name     string
	Save_peer     peerconnectorcontract.Peer
	Save_ctx      context.Context
	Save_IsCalled bool
}

func (r *PeerRepositoryMock) Save(ctx context.Context, name string, peer peerconnectorcontract.Peer) error {
	r.Save_IsCalled = true
	r.Save_ctx = ctx
	r.Save_name = name
	r.Save_peer = peer
	return nil
}

var _ peerrepositorycontract.Repository = &PeerRepositoryMock{}
