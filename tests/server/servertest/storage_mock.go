package servertest

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/tests"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"sync"
	"time"
)

type DbStorageMock struct {
	GetRecord_key                    string
	GetRecord_ctx                    context.Context
	GetRecord_execute                func(dbstoragecontract.Record)
	GetRecord_failure                func(error)
	GetRecord_wg                     *sync.WaitGroup
	GetOrCreateRecord_key            string
	GetOrCreateRecord_ctx            context.Context
	GetOrCreateRecord_execute_record dbstoragecontract.Record
	Start_ctx                        context.Context
	Join_IsCalled                    bool
	Save_ctx                         context.Context
	Save_key                         string
	Save_record                      dbstoragecontract.Record
	Save_wg                          *sync.WaitGroup
}

var _ dbstoragecontract.Storage = &DbStorageMock{}

func (s *DbStorageMock) Start(ctx context.Context) error {
	s.Start_ctx = ctx
	return nil
}

func (s *DbStorageMock) Join() error {
	s.Join_IsCalled = true
	return nil
}

func (s *DbStorageMock) GetRecord(ctx context.Context, key string, execute func(dbstoragecontract.Record), failure func(err error)) error {
	defer func() {
		if s.GetRecord_wg != nil {
			s.GetRecord_wg.Done()
		}
	}()
	s.GetRecord_key = key
	s.GetRecord_ctx = ctx
	s.GetRecord_execute = execute
	s.GetRecord_failure = failure

	return nil
}

func (s *DbStorageMock) GetRecord_WaitUntilCalledOnce(timeout time.Duration, f func()) bool {
	return tests.MockWaitUntilCalledNthTimes(&s.GetRecord_wg, 1, timeout, f)
}

func (s *DbStorageMock) GetOrCreateRecord(ctx context.Context, key string, execute func(dbstoragecontract.Record)) error {
	s.GetOrCreateRecord_ctx = ctx
	s.GetOrCreateRecord_key = key
	go execute(goutil.Coalesce[dbstoragecontract.Record](s.GetOrCreateRecord_execute_record, &dbstoragetest.RecordMock{}))
	return nil
}

func (s *DbStorageMock) Save(ctx context.Context, key string, record dbstoragecontract.Record) error {
	defer func() {
		if s.Save_wg != nil {
			s.Save_wg.Done()
		}
	}()
	s.Save_ctx = ctx
	s.Save_key = key
	s.Save_record = record
	return nil
}
func (s *DbStorageMock) Save_WaitUntillCalledOnce(timeout time.Duration, f func()) bool {
	return tests.MockWaitUntilCalledNthTimes(&s.Save_wg, 1, timeout, f)
}
