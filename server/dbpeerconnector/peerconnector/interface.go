package peerconnector

type Interface interface {
	Start() error
	Stop() error
}
