package servertest

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"sync"
	"time"
)

type DbStorageMock struct {
	GetRecord_key                    string
	GetRecord_ctx                    context.Context
	GetRecord_execute_record         dbstoragecontract.Record
	GetRecord_failure_err            error
	GetRecord_wg                     *sync.WaitGroup
	GetOrCreateRecord_key            string
	GetOrCreateRecord_ctx            context.Context
	GetOrCreateRecord_execute_record dbstoragecontract.Record
	Start_ctx                        context.Context
	Join_IsCalled                    bool
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

	if s.GetRecord_failure_err != nil {
		failure(s.GetRecord_failure_err)
	} else {
		execute(goutil.Coalesce[dbstoragecontract.Record](s.GetRecord_execute_record, &dbstoragetest.RecordMock{}))
	}

	return nil
}

func (s *DbStorageMock) GetRecord_WaitUntilCalledOnce(timeout time.Duration, f func()) bool {
	defer func() {
		s.GetRecord_wg = nil
	}()
	s.GetRecord_wg = &sync.WaitGroup{}
	s.GetRecord_wg.Add(1)

	f()

	return goutil.WaitWithTimeout(s.GetRecord_wg, timeout)
}

func (s *DbStorageMock) GetOrCreateRecord(ctx context.Context, key string, execute func(dbstoragecontract.Record)) error {
	s.GetOrCreateRecord_ctx = ctx
	s.GetOrCreateRecord_key = key
	go execute(goutil.Coalesce[dbstoragecontract.Record](s.GetOrCreateRecord_execute_record, &dbstoragetest.RecordMock{}))
	return nil
}
