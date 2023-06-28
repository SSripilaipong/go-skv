package connectormanagertest

import (
	"go-skv/server/dbpeerconnector/connectormanager"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
	"go-skv/server/dbpeerconnector/peerserver/peerservercontract"
)

type Dependencies struct {
	addresses []string
	client    peerclientcontract.Client
	peerRepo  peerrepositorycontract.Repository
	server    peerservercontract.Server
}

func defaultDependencies() Dependencies {
	return Dependencies{
		addresses: []string{},
		client:    &PeerClientMock{},
		peerRepo:  &PeerRepositoryMock{},
		server:    &PeerServerMock{},
	}
}

func New(options ...func(deps *Dependencies)) peerconnectorcontract.Connector {
	deps := defaultDependencies()
	for _, option := range options {
		option(&deps)
	}

	return connectormanager.New(deps.addresses, deps.client, deps.peerRepo, deps.server)
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

func WithServer(server peerservercontract.Server) func(*Dependencies) {
	return func(deps *Dependencies) {
		deps.server = server
	}
}
