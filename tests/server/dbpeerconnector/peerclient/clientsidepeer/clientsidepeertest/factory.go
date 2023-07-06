package clientsidepeertest

import (
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer"
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

type factoryDependency struct {
	bufferSize            int
	replicaUpdaterFactory replicaupdatercontract.Factory
}

func defaultFactoryDependency() factoryDependency {
	return factoryDependency{
		bufferSize:            0,
		replicaUpdaterFactory: &ReplicaUpdaterFactoryMock{},
	}
}

func NewFactory(options ...func(*factoryDependency)) clientsidepeercontract.Factory {
	deps := defaultFactoryDependency()
	for _, option := range options {
		option(&deps)
	}
	return clientsidepeer.NewFactory(deps.bufferSize, deps.replicaUpdaterFactory)
}

func WithReplicaUpdaterFactory(factory replicaupdatercontract.Factory) func(*factoryDependency) {
	return func(deps *factoryDependency) {
		deps.replicaUpdaterFactory = factory
	}
}

func WithBufferSize(bufferSize int) func(*factoryDependency) {
	return func(deps *factoryDependency) {
		deps.bufferSize = bufferSize
	}
}
