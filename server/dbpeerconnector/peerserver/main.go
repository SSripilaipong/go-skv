package peerserver

import (
	"go-skv/server/dbpeerconnector/peerserver/peergrpccontroller"
	"go-skv/server/dbpeerconnector/peerserver/peerservercontract"
	"go-skv/server/dbpeerconnector/peerserver/peerserverusecase"
)

func New(port int) peerservercontract.Server {
	return peergrpccontroller.New(port, peerserverusecase.New())
}
