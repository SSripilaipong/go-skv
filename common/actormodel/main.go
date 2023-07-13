package actormodel

import (
	"context"
)

type Command[S any] interface {
	Execute(*S)
}

func GoRunActor[S any](ctx context.Context, state S, cmdCh <-chan Command[S], onStopped func(state S, panicValue any)) context.CancelFunc {
	defer func() {
		onStopped(state, recover())
	}()

	subCtx, cancelSubCtx := context.WithCancel(ctx)
	go func() {
		for {
			select {
			case cmd := <-cmdCh:
				cmd.Execute(&state)
			case <-subCtx.Done():
				return
			}
		}
	}()

	return cancelSubCtx
}
