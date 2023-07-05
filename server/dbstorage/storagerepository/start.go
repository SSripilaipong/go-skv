package storagerepository

import "context"

func (m manager) Start(ctx context.Context) error {
	return m.sendMessage(ctx, startCommand{ctx: ctx})
}

type startCommand struct {
	ctx context.Context
}

func (c startCommand) execute(s *state) {
	s.ctx = c.ctx
}
