package storagerecord

import "context"

type state struct {
	value string
}

type command interface {
	execute(s *state)
}

func runRecordMainLoop(ctx context.Context, ch chan command, stopped chan struct{}) {
	var s state

	for {
		select {
		case cmd := <-ch:
			cmd.execute(&s)
		case <-ctx.Done():
			goto stop
		}
	}
stop:
	stopped <- struct{}{}
}
