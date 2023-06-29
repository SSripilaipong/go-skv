package peergrpccontroller

import (
	"go-skv/server/dbpeerconnector/peerserver/peerservercontract"
	"go-skv/server/dbpeerconnector/peerserver/peerserverusecase/peerserverusecase"
	"sync"
)

func New(port int, usecase peerserverusecase.Usecase) peerservercontract.Server {
	return &controller{
		port:    port,
		usecase: usecase,

		wg: new(sync.WaitGroup),
	}
}

type controller struct {
	port    int
	usecase peerserverusecase.Usecase

	wg *sync.WaitGroup
}

var _ peerservercontract.Server = &controller{}
