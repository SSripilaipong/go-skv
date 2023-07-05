package storagerepository

import (
	"context"
	"go-skv/common/commoncontract"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type state struct {
	ctx           context.Context
	recordFactory dbstoragecontract.Factory
	records       map[string]dbstoragecontract.Record
}

func mainLoop(ch chan command, stopped chan struct{}, recordFactory dbstoragecontract.Factory) {
	s := state{
		recordFactory: recordFactory,
		records:       make(map[string]dbstoragecontract.Record),
	}
	waitUntilStart(ch, &s)
	for {
		select {
		case cmd := <-ch:
			cmd.execute(&s)
		case <-s.ctx.Done():
			goto stop
		}
	}
stop:
	stopped <- struct{}{}
}

type command interface {
	execute(s *state)
}

func (m manager) sendMessage(ctx context.Context, cmd command) error {
	select {
	case m.ch <- cmd:
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
	return nil
}

func waitUntilStart(ch chan command, s *state) {
	startCmd := (<-ch).(command)
	startCmd.execute(s)
}
