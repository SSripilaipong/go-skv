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
	GetRecord_success_record         dbstoragecontract.Record
	GetRecord_wg                     *sync.WaitGroup
	GetOrCreateRecord_key            string
	GetOrCreateRecord_ctx            context.Context
	GetOrCreateRecord_success_record dbstoragecontract.Record
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

func (s *DbStorageMock) GetRecord(ctx context.Context, key string, success func(dbstoragecontract.Record)) error {
	defer func() {
		if s.GetRecord_wg != nil {
			s.GetRecord_wg.Done()
		}
	}()
	s.GetRecord_key = key
	s.GetRecord_ctx = ctx
	go success(goutil.Coalesce[dbstoragecontract.Record](s.GetRecord_success_record, &dbstoragetest.RecordMock{}))
	return nil
}

func (s *DbStorageMock) GetRecord_WaitUntilCalledOnce(timeout time.Duration, f func()) bool {
	s.GetRecord_wg = &sync.WaitGroup{}
	s.GetRecord_wg.Add(1)

	f()

	called := make(chan struct{})
	go func() {
		s.GetRecord_wg.Wait()
		called <- struct{}{}
	}()
	select {
	case <-called:
		return true
	case <-time.After(timeout):
		return false
	}
}

func (s *DbStorageMock) GetOrCreateRecord(ctx context.Context, key string, success func(dbstoragecontract.Record)) error {
	s.GetOrCreateRecord_ctx = ctx
	s.GetOrCreateRecord_key = key
	go success(goutil.Coalesce[dbstoragecontract.Record](s.GetOrCreateRecord_success_record, &dbstoragetest.RecordMock{}))
	return nil
}
