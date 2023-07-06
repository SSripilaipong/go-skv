package peerconnectorcontract

type Peer interface {
	UpdateReplicaFromPeer(key string, value string) error
	Join() error
}
