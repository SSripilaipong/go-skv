package storagerepository

import (
	"context"
	"go-skv/common/commoncontract"
)

type command interface {
	execute(s *state)
}

func (m *Manager) sendMessage(ctx context.Context, message any) error {
	select {
	case m.ch <- message:
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
	return nil
}
