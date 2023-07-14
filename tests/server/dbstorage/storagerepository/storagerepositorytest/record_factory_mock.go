package storagerepositorytest

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"sync"
	"time"
)

type RecordFactoryMock struct {
	New_Return   dbstoragecontract.Record
	New_IsCalled bool
	New_ctx      context.Context
	New_wg       *sync.WaitGroup
}

func (t *RecordFactoryMock) New(ctx context.Context) dbstoragecontract.Record {
	defer func() {
		if t.New_wg != nil {
			t.New_wg.Done()
		}
	}()

	t.New_IsCalled = true
	t.New_ctx = ctx
	return goutil.Coalesce[dbstoragecontract.Record](t.New_Return, &dbstoragetest.RecordMock{})
}

func (t *RecordFactoryMock) New_WaitUntilCalledOnce(timeout time.Duration, f func()) bool {
	defer func() {
		t.New_wg = nil
	}()

	t.New_wg = &sync.WaitGroup{}
	t.New_wg.Add(1)
	f()
	return goutil.WaitWithTimeout(t.New_wg, timeout)
}

func (t *RecordFactoryMock) New_CaptureReset() {
	t.New_IsCalled = false
}
