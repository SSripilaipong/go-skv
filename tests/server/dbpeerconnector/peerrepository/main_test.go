package peerrepository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/dbpeerconnectortest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_save_then_get_the_same_peer_using_the_same_name(t *testing.T) {
	savedPeer := &dbpeerconnectortest.PeerMock{}
	var receivedPeer peerconnectorcontract.Peer
	repo := peerrepository.New()
	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(repo.Start(ctx))

		goutil.PanicUnhandledError(repo.Save(ctx, "xxx", savedPeer))

		done := make(chan struct{})
		goutil.PanicUnhandledError(repo.Get(ctx, "xxx", func(peer peerconnectorcontract.Peer) {
			receivedPeer = peer
			done <- struct{}{}
		}))
		goutil.ReceiveWithTimeout(done, defaultTimeout)
	})

	assert.Equal(t, savedPeer, receivedPeer)
}

func Test_should_get_peer_1_from_name(t *testing.T) {
	peer1 := new(dbpeerconnectortest.PeerMock)
	peer2 := new(dbpeerconnectortest.PeerMock)
	var receivedPeer peerconnectorcontract.Peer
	repo := peerrepository.New()
	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(repo.Start(ctx))

		goutil.PanicUnhandledError(repo.Save(ctx, "xxx", peer1))
		goutil.PanicUnhandledError(repo.Save(ctx, "yyy", peer2))

		done := make(chan struct{})
		goutil.PanicUnhandledError(repo.Get(ctx, "xxx", func(peer peerconnectorcontract.Peer) {
			receivedPeer = peer
			done <- struct{}{}
		}))
		goutil.ReceiveWithTimeout(done, defaultTimeout)
	})

	assert.True(t, peer1 == receivedPeer)
}
