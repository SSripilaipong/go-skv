package dbmanager

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbmanager"
	"testing"
)

func Test_should_start_peer_server(t *testing.T) {
	peerServer := &PeerServerMock{}
	mgr := dbmanager.New(peerServer)

	_ = mgr.Start()

	assert.True(t, peerServer.Start_IsCalled)
}
