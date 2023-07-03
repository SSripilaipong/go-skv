package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

type state struct {
	ctx           context.Context
	recordFactory storagerecord.Factory
	records       map[string]storagerecord.Interface
}

func mainLoop(ctx context.Context, ch chan any, stopped chan struct{}, recordFactory storagerecord.Factory) {
	s := state{
		ctx:           ctx,
		recordFactory: recordFactory,
		records:       make(map[string]storagerecord.Interface),
	}
	for {
		select {
		case raw := <-ch:
			if message, isCommand := raw.(command); isCommand {
				message.execute(&s)
			}
		case <-ctx.Done():
			goto stop
		}
	}
stop:
	stopped <- struct{}{}
}
