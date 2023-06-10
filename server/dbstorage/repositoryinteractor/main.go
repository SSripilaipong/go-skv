package repositoryinteractor

import (
	"context"
	"go-skv/server/dbstorage/repositoryroutine"
)

func New(ch chan<- any) Interface {
	return interactor{ch: ch}
}

type interactor struct {
	ch chan<- any
}

func (i interactor) GetRecord(ctx context.Context, key string, success repositoryroutine.GetRecordSuccessCallback) error {
	return i.sendMessage(ctx, repositoryroutine.GetRecordMessage{
		Key:     key,
		Success: success,
	})
}

func (i interactor) GetOrCreateRecord(ctx context.Context, key string, success repositoryroutine.GetOrCreateRecordSuccessCallback) error {
	return i.sendMessage(ctx, repositoryroutine.GetOrCreateRecordMessage{
		Key:     key,
		Success: success,
	})
}

func (i interactor) sendMessage(ctx context.Context, message any) error {
	select {
	case i.ch <- message:
	case <-ctx.Done():
		return ContextCancelledError{}
	}
	return nil
}
