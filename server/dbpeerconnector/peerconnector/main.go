package peerconnector

func New() Interface {
	return &connector{}
}

type connector struct{}

var _ Interface = &connector{}
