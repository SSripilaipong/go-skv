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

func Test_should_have_timeout_when_sending_command(t *testing.T) {
	var isTimeoutApplied bool
	blockingMainLoop, unblockMainLoop := context.WithCancel(context.Background())
	replicaUpdaterFactory := &clientsidepeertest.ReplicaUpdaterFactoryMock{
		NewInboundUpdater_Return: &clientsidepeertest.ReplicaInboundUpdaterMock{
			Update_Do: func(_, _ string) { <-blockingMainLoop.Done() },
		},
	}
	factory := clientsidepeertest.NewFactory(
		clientsidepeertest.WithBufferSize(1),
		clientsidepeertest.WithDefaultSendingTimeout(defaultTimeout),
		clientsidepeertest.WithReplicaUpdaterFactory(replicaUpdaterFactory),
	)
	var peer peerconnectorcontract.Peer
	tests.ContextScope(func(ctx context.Context) {
		var err error
		peer, err = factory.New(ctx)
		goutil.PanicUnhandledError(err)
		goutil.PanicUnhandledError(peer.UpdateReplicaFromPeer("", "")) // blocks main loop
		goutil.PanicUnhandledError(peer.UpdateReplicaFromPeer("", "")) // go to channel buffer (1)

		done := make(chan struct{})
		go func() {
			goutil.PanicUnhandledError(peer.UpdateReplicaFromPeer("", "")) // buffer full, cannot send
			done <- struct{}{}
		}()
		_, isTimeoutApplied = goutil.ReceiveWithTimeout(done, 2*defaultTimeout)

		unblockMainLoop()
		time.Sleep(defaultTimeout)
	})
	goutil.PanicUnhandledError(peer.Join())
	assert.True(t, isTimeoutApplied)
}
