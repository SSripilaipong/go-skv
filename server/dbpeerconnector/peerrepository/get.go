package peerrepository

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func (t interactor) Get(ctx context.Context, name string, execute func(peer peerconnectorcontract.Peer)) error {
	t.sendCommand(ctx, getCommand{
		name: name,
		exec: execute,
	})
	return nil
}

type getCommand struct {
	name string
	exec func(peer peerconnectorcontract.Peer)
}

func (c getCommand) execute(s *state) {
	c.exec(s.temp)
}
