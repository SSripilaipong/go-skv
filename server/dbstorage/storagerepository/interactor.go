package storagerepository

import (
	"context"
	"go-skv/common/commoncontract"
)

type interactor struct {
	ch chan<- any
}

type command interface {
	execute(s *state)
}

func (i interactor) sendMessage(ctx context.Context, message any) error {
	select {
	case i.ch <- message:
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
	return nil
}
