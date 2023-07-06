package clientsidepeer

import (
	"context"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func mainLoop(ctx context.Context, ch <-chan command, onStopped func()) {
	defer onStopped()

	s := state{}

	for {
		select {
		case cmd := <-ch:
			cmd.execute(&s)
		case <-ctx.Done():
			return
		}
	}
}

type state struct {
	inboundUpdater replicaupdatercontract.InboundUpdater
}

type command interface {
	execute(s *state)
}
