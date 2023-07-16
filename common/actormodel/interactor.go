package actormodel

import (
	"context"
	"go-skv/common/commoncontract"
	"sync"
)

type Interactor[S any] struct {
	cmdCh chan<- Command[S]
	wg    *sync.WaitGroup
}

func NewInteractor[S any](cmdCh chan<- Command[S], wg *sync.WaitGroup) Interactor[S] {
	return Interactor[S]{
		cmdCh: cmdCh,
		wg:    wg,
	}
}

func (t Interactor[S]) SendCommand(ctx context.Context, cmd Command[S]) error {
	select {
	case t.cmdCh <- cmd:
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
	return nil
}

func (t Interactor[S]) SendCommandOrPanic(cmd Command[S]) {
	t.cmdCh <- cmd
}

func (t Interactor[S]) Join() {
	t.wg.Wait()
}
