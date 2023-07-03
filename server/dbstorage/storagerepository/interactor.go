package storagerepository

import (
	"context"
	"go-skv/common/commoncontract"
)

type interactor struct {
	ch chan<- any
}

func (i interactor) GetRecord(ctx context.Context, key string, success GetRecordSuccessCallback) error {
	return i.sendMessage(ctx, GetRecordMessage{
		Key:     key,
		Success: success,
	})
}

func (i interactor) GetOrCreateRecord(ctx context.Context, key string, success GetOrCreateRecordSuccessCallback) error {
	return i.sendMessage(ctx, GetOrCreateRecordMessage{
		Key:     key,
		Success: success,
	})
}

func (i interactor) sendMessage(ctx context.Context, message any) error {
	select {
	case i.ch <- message:
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
	return nil
}
