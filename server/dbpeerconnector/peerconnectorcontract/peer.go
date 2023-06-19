package peerconnectorcontract

type Peer interface {
	SubscribeUpdates(listener UpdateListener) error
}
