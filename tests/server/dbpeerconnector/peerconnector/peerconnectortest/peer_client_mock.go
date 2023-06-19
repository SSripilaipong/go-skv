package peerconnectortest

import (
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type PeerClientMock struct {
	ConnectToPeer_address_array []string
}

func (c *PeerClientMock) ConnectToPeer(address string) (peerconnectorcontract.Peer, error) {
	c.ConnectToPeer_address_array = append(c.ConnectToPeer_address_array, address)
	return nil, nil
}

var _ peerclientcontract.Client = &PeerClientMock{}
