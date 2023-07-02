package peerrepository

import "context"

func (t interactor) Start(ctx context.Context) error {
	t.sendCommand(ctx, startCommand{ctx: ctx})
	return nil
}

type startCommand struct {
	ctx context.Context
}

func (startCommand) execute(*state) {}

func waitUntilStart(ch <-chan command) context.Context {
	for {
		if start, ok := (<-ch).(startCommand); ok {
			return start.ctx
		}
	}
}
