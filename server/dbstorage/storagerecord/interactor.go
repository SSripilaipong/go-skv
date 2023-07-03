package storagerecord

import (
	"context"
	"go-skv/common/commoncontract"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type recordInteractor struct {
	ctx       context.Context
	ch        chan command
	ctxCancel context.CancelFunc

	stopped chan struct{}
}

func (r recordInteractor) sendCommand(ctx context.Context, cmd command) error {
	if r.isContextEnded() {
		return dbstoragecontract.RecordDestroyedError{}
	}
	select {
	case r.ch <- cmd:
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
	return nil
}

func (r recordInteractor) isContextEnded() bool {
	_, isEnded := goutil.ReceiveNoBlock(r.ctx.Done())
	return isEnded
}
