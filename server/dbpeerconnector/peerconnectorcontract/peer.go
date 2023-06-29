package peerconnectorcontract

type Peer interface {
	UpdateReplica(key string, value string) error
}
