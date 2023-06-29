package peergrpctest

import (
	"go-skv/server/dbpeerconnector/peerserver/peergrpccontroller"
	"go-skv/server/dbpeerconnector/peerserver/peerservercontract"
	"go-skv/server/dbpeerconnector/peerserver/peerserverusecase/peerserverusecase"
)

type controllerDependency struct {
	port    int
	usecase peerserverusecase.Usecase
}

func defaultControllerDependency() controllerDependency {
	return controllerDependency{
		port:    0,
		usecase: &ServerUsecaseMock{},
	}
}

func NewController(options ...func(*controllerDependency)) peerservercontract.Server {
	deps := defaultControllerDependency()
	for _, option := range options {
		option(&deps)
	}
	return peergrpccontroller.New(deps.port, deps.usecase)
}

func WithServerUsecase(usecase peerserverusecase.Usecase) func(*controllerDependency) {
	return func(deps *controllerDependency) {
		deps.usecase = usecase
	}
}
