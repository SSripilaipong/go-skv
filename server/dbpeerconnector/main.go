package dbpeerconnector

import "go-skv/server/dbpeerconnector/peerconnector"

func New() Interface {
	return peerconnector.New()
}
