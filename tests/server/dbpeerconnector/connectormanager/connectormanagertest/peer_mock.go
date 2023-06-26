package connectormanagertest

import "go-skv/server/dbpeerconnector/peerconnectorcontract"

type PeerMock struct{}

var _ peerconnectorcontract.Peer = &PeerMock{}
