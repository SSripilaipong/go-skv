package update_replica_from_peer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeertest"
	"testing"
	"time"
)

func Test_should_create_inbound_replica_updater_if_not_exists(t *testing.T) {
	replicaUpdaterFactory := &clientsidepeertest.ReplicaUpdaterFactoryMock{
		NewInboundUpdater_Return: &clientsidepeertest.ReplicaInboundUpdaterMock{},
	}
	factory := clientsidepeertest.NewFactory(
		clientsidepeertest.WithReplicaUpdaterFactory(replicaUpdaterFactory),
	)
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
	factory := clientsidepeertest.NewFactory(
		clientsidepeertest.WithReplicaUpdaterFactory(replicaUpdaterFactory),
	)
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

func Test_should_pass_global_context_when_create_inbound_replica_updater(t *testing.T) {
	replicaUpdaterFactory := &clientsidepeertest.ReplicaUpdaterFactoryMock{
		NewInboundUpdater_Return: &clientsidepeertest.ReplicaInboundUpdaterMock{},
	}
	factory := clientsidepeertest.NewFactory(
		clientsidepeertest.WithReplicaUpdaterFactory(replicaUpdaterFactory),
	)
	var peer peerconnectorcontract.Peer
	tests.ContextScope(func(ctx context.Context) {
		var err error
		peer, err = factory.New(context.WithValue(ctx, "test", "abc555"))
		goutil.PanicUnhandledError(err)

		goutil.PanicUnhandledError(peer.UpdateReplicaFromPeer("", ""))

		time.Sleep(defaultTimeout)
	})
	goutil.PanicUnhandledError(peer.Join())
	assert.Equal(t, "abc555", replicaUpdaterFactory.NewInboundUpdater_ctx.Value("test"))
}

func Test_should_send_update_to_inbound_replica_updater_with_key_and_value(t *testing.T) {
	updater := &clientsidepeertest.ReplicaInboundUpdaterMock{}
	replicaUpdaterFactory := &clientsidepeertest.ReplicaUpdaterFactoryMock{NewInboundUpdater_Return: updater}
	factory := clientsidepeertest.NewFactory(
		clientsidepeertest.WithReplicaUpdaterFactory(replicaUpdaterFactory),
	)
	var peer peerconnectorcontract.Peer
	tests.ContextScope(func(ctx context.Context) {
		var err error
		peer, err = factory.New(ctx)
		goutil.PanicUnhandledError(err)

		goutil.PanicUnhandledError(peer.UpdateReplicaFromPeer("xxx", "yyy"))

		time.Sleep(defaultTimeout)
	})
	goutil.PanicUnhandledError(peer.Join())
	assert.Equal(t, "xxx", updater.Update_key)
	assert.Equal(t, "yyy", updater.Update_value)
}
