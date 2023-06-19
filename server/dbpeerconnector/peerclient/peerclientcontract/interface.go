package peerclientcontract

import "go-skv/server/dbpeerconnector/peerconnectorcontract"

type Client interface {
	ConnectToPeer(address string) (peerconnectorcontract.Peer, error)
}
