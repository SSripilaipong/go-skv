package peerconnectortest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnector"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
)

type Dependencies struct {
	addresses []string
	client    peerclientcontract.Client
	peerRepo  peerrepositorycontract.Repository
	ctx       context.Context
}

func defaultDependencies() Dependencies {
	return Dependencies{
		addresses: []string{},
		client:    &PeerClientMock{},
		peerRepo:  &PeerRepositoryMock{},
	}
}

func New(options ...func(deps *Dependencies)) peerconnectorcontract.Connector {
	deps := defaultDependencies()
	for _, option := range options {
		option(&deps)
	}

	return peerconnector.New(deps.ctx, deps.addresses, deps.client, deps.peerRepo)
}

func WithContext(ctx context.Context) func(*Dependencies) {
	return func(deps *Dependencies) {
		deps.ctx = ctx
	}
}

func WithNonEmptyAddresses() func(*Dependencies) {
	return func(deps *Dependencies) {
		deps.addresses = []string{"1.2.3.4:5678"}
	}
}

func WithAddresses(addresses []string) func(*Dependencies) {
	return func(deps *Dependencies) {
		deps.addresses = addresses
	}
}

func WithClient(client peerclientcontract.Client) func(*Dependencies) {
	return func(deps *Dependencies) {
		deps.client = client
	}
}

func WithPeerRepo(peerRepo peerrepositorycontract.Repository) func(*Dependencies) {
	return func(deps *Dependencies) {
		deps.peerRepo = peerRepo
	}
}
