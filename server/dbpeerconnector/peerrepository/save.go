package peerrepository

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func (t interactor) Save(ctx context.Context, name string, peer peerconnectorcontract.Peer) error {
	t.sendCommand(ctx, saveCommand{
		name: name,
		peer: peer,
	})
	return nil
}

type saveCommand struct {
	name string
	peer peerconnectorcontract.Peer
}

func (c saveCommand) execute(s *state) {
	s.peers[c.name] = c.peer
}
