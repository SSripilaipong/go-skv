package clientsidepeertest

import (
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer"
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
	"time"
)

type factoryDependency struct {
	bufferSize             int
	defaultSendingTimeout  time.Duration
	replicaUpdaterFactory  replicaupdatercontract.Factory
	replicaUpdaterFactory2 replicaupdatercontract.ActorFactory
}

func defaultFactoryDependency() factoryDependency {
	return factoryDependency{
		replicaUpdaterFactory:  &ReplicaUpdaterFactoryMock{},
		replicaUpdaterFactory2: &ReplicaUpdaterFactory2Mock{},
		defaultSendingTimeout:  100 * time.Millisecond,
	}
}

func NewFactory(options ...func(*factoryDependency)) clientsidepeercontract.Factory {
	deps := defaultFactoryDependency()
	for _, option := range options {
		option(&deps)
	}
	return clientsidepeer.NewFactory(deps.bufferSize, deps.defaultSendingTimeout, deps.replicaUpdaterFactory2)
}

func WithReplicaUpdaterFactory(factory replicaupdatercontract.ActorFactory) func(*factoryDependency) {
	return func(deps *factoryDependency) {
		deps.replicaUpdaterFactory2 = factory
	}
}

func WithBufferSize(bufferSize int) func(*factoryDependency) {
	return func(deps *factoryDependency) {
		deps.bufferSize = bufferSize
	}
}

func WithDefaultSendingTimeout(timeout time.Duration) func(*factoryDependency) {
	return func(deps *factoryDependency) {
		deps.defaultSendingTimeout = timeout
	}
}
