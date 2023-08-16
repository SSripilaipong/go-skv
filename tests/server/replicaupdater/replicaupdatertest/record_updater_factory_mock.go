package replicaupdatertest

import (
	"context"
	"go-skv/common/test"
	"sync"
	"time"
)

type RecordUpdaterFactoryMock struct {
	New_Return chan<- any
	New_ctx    context.Context
	New_key    string
	New_value  string
	New_wg     *sync.WaitGroup
}

func (t *RecordUpdaterFactoryMock) New(ctx context.Context, key string, value string) (chan<- any, func()) {
	defer func() {
		if t.New_wg != nil {
			t.New_wg.Done()
		}
	}()

	t.New_ctx = ctx
	t.New_key = key
	t.New_value = value
	return t.New_Return, nil
}

func (t *RecordUpdaterFactoryMock) New_WaitUntilCalledOnce(timeout time.Duration, f func()) {
	test.MockWaitUntilCalledNthTimes(&t.New_wg, 1, timeout, f)
}
