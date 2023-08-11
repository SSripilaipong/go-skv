package clientsidepeer

import (
	"context"
)

func mainLoop(ctx context.Context, ch <-chan command, onStopped func()) {
	defer onStopped()

	s := state{
		ctx: ctx,
	}

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
	ctx            context.Context
	inboundUpdater chan<- any
}

type command interface {
	execute(s *state)
}
