package peerconnectorcontract

type Connector interface {
	Start() error
	Stop() error
}
