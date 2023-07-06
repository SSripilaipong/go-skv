package clientsidepeertest

import (
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer"
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
	"time"
)

type factoryDependency struct {
	bufferSize            int
	defaultSendingTimeout time.Duration
	replicaUpdaterFactory replicaupdatercontract.Factory
}

func defaultFactoryDependency() factoryDependency {
	return factoryDependency{
		replicaUpdaterFactory: &ReplicaUpdaterFactoryMock{},
		defaultSendingTimeout: 100 * time.Millisecond,
	}
}

func NewFactory(options ...func(*factoryDependency)) clientsidepeercontract.Factory {
	deps := defaultFactoryDependency()
	for _, option := range options {
		option(&deps)
	}
	return clientsidepeer.NewFactory(deps.bufferSize, deps.defaultSendingTimeout, deps.replicaUpdaterFactory)
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

func WithDefaultSendingTimeout(timeout time.Duration) func(*factoryDependency) {
	return func(deps *factoryDependency) {
		deps.defaultSendingTimeout = timeout
	}
}
