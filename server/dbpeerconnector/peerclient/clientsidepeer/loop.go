package clientsidepeer

import "context"

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

type state struct{}

type command interface {
	execute(s *state)
}
