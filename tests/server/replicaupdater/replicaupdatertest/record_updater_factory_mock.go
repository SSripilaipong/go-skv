package replicaupdatertest

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/tests"
	"sync"
	"time"
)

type RecordUpdaterFactoryMock struct {
	New_Return actormodel.ActorRef
	New_ctx    context.Context
	New_key    string
	New_value  string
	New_wg     *sync.WaitGroup
}

func (t *RecordUpdaterFactoryMock) New(ctx context.Context, key string, value string) actormodel.ActorRef {
	defer func() {
		if t.New_wg != nil {
			t.New_wg.Done()
		}
	}()

	t.New_ctx = ctx
	t.New_key = key
	t.New_value = value
	return t.New_Return
}

func (t *RecordUpdaterFactoryMock) New_WaitUntilCalledOnce(timeout time.Duration, f func()) {
	tests.MockWaitUntilCalledNthTimes(&t.New_wg, 1, timeout, f)
}
