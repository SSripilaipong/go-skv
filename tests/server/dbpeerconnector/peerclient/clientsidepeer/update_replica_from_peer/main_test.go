package update_replica_from_peer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeertest"
	"testing"
	"time"
)

func Test_should_create_inbound_replica_updater_if_not_exists(t *testing.T) {
	replicaUpdaterFactory := &clientsidepeertest.ReplicaUpdaterFactoryMock{}
	factory := clientsidepeer.NewFactory(replicaUpdaterFactory)
	var peer peerconnectorcontract.Peer
	tests.ContextScope(func(ctx context.Context) {
		var err error
		peer, err = factory.New(ctx)
		goutil.PanicUnhandledError(err)

		goutil.PanicUnhandledError(peer.UpdateReplicaFromPeer("", ""))

		time.Sleep(defaultTimeout)
	})
	goutil.PanicUnhandledError(peer.Join())
	assert.True(t, replicaUpdaterFactory.NewInboundUpdater_IsCalled)
}

func Test_should_not_create_inbound_replica_updater_if_already_exists(t *testing.T) {
	replicaUpdaterFactory := &clientsidepeertest.ReplicaUpdaterFactoryMock{
		NewInboundUpdater_Return: &clientsidepeertest.ReplicaInboundUpdaterMock{},
	}
	factory := clientsidepeer.NewFactory(replicaUpdaterFactory)
	var peer peerconnectorcontract.Peer
	tests.ContextScope(func(ctx context.Context) {
		var err error
		peer, err = factory.New(ctx)
		goutil.PanicUnhandledError(err)
		goutil.PanicUnhandledError(peer.UpdateReplicaFromPeer("", "")) // 1st time

		goutil.PanicUnhandledError(peer.UpdateReplicaFromPeer("", "")) // 2nd time

		time.Sleep(defaultTimeout)
	})
	goutil.PanicUnhandledError(peer.Join())
	assert.Equal(t, 1, replicaUpdaterFactory.NewInboundUpdater_CallCount)
}
