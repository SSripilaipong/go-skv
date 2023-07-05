package storagerepository

import (
	"context"
	"go-skv/common/commoncontract"
	"go-skv/server/dbstorage/storagerecord"
)

type state struct {
	ctx           context.Context
	recordFactory storagerecord.Factory
	records       map[string]storagerecord.Interface
}

func mainLoop(ch chan any, stopped chan struct{}, recordFactory storagerecord.Factory) {
	s := state{
		recordFactory: recordFactory,
		records:       make(map[string]storagerecord.Interface),
	}
	waitUntilStart(ch, &s)
	for {
		select {
		case raw := <-ch:
			if message, isCommand := raw.(command); isCommand {
				message.execute(&s)
			}
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

func (m manager) sendMessage(ctx context.Context, message any) error {
	select {
	case m.ch <- message:
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
	return nil
}

func waitUntilStart(ch chan any, s *state) {
	startCmd := (<-ch).(command)
	startCmd.execute(s)
}
