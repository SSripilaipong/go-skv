package peerconnectortest

import (
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnector"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
)

func NewWithAddressesAndClient(addresses []string, client peerclientcontract.Client) peerconnectorcontract.Connector {
	return peerconnector.New(addresses, client, &PeerRepositoryMock{})
}

func NewWithAddressesAndClientAndPeerRepo(addresses []string, client peerclientcontract.Client, peerRepo peerrepositorycontract.Repository) peerconnectorcontract.Connector {
	return peerconnector.New(addresses, client, peerRepo)
}

func New() peerconnectorcontract.Connector {
	return peerconnector.New(nil, &PeerClientMock{}, &PeerRepositoryMock{})
}
