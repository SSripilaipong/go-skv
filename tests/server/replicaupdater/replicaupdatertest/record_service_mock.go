package replicaupdatertest

import (
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater"
	"sync"
	"time"
)

type RecordServiceMock struct {
	UpdateReplicaValue_record      dbstoragecontract.Record
	UpdateReplicaValue_IsCalled    bool
	UpdateReplicaValue_value       string
	UpdateReplicaValue_wg          *sync.WaitGroup
	InitilizeReplicaRecord_record  dbstoragecontract.Record
	InitilizeReplicaRecord_value   string
	InitilizeReplicaRecord_wg      *sync.WaitGroup
	InitilizeReplicaRecord_execute func(record dbstoragecontract.Record)
}

func (s *RecordServiceMock) InitializeReplicaRecord(record dbstoragecontract.Record, value string, execute func(record dbstoragecontract.Record)) error {
	defer func() {
		if s.InitilizeReplicaRecord_wg != nil {
			s.InitilizeReplicaRecord_wg.Done()
		}
	}()
	s.InitilizeReplicaRecord_record = record
	s.InitilizeReplicaRecord_value = value
	s.InitilizeReplicaRecord_execute = execute
	return nil
}
func (s *RecordServiceMock) InitializeReplicaRecord_WaitUntilCalledOnce(timeout time.Duration, f func()) bool {
	defer func() {
		s.InitilizeReplicaRecord_wg = nil
	}()

	s.InitilizeReplicaRecord_wg = &sync.WaitGroup{}
	s.InitilizeReplicaRecord_wg.Add(1)

	f()

	return goutil.WaitWithTimeout(s.InitilizeReplicaRecord_wg, timeout)
}

func (s *RecordServiceMock) UpdateReplicaValue(record dbstoragecontract.Record, value string, onFailure func(err error)) {
	defer func() {
		if s.UpdateReplicaValue_wg != nil {
			s.UpdateReplicaValue_wg.Done()
		}
	}()

	s.UpdateReplicaValue_IsCalled = true
	s.UpdateReplicaValue_record = record
	s.UpdateReplicaValue_value = value
}
func (s *RecordServiceMock) UpdateReplicaValue_WaitUntilCalledOnce(timeout time.Duration, f func()) bool {
	defer func() {
		s.UpdateReplicaValue_wg = nil
	}()

	s.UpdateReplicaValue_wg = &sync.WaitGroup{}
	s.UpdateReplicaValue_wg.Add(1)

	f()

	return goutil.WaitWithTimeout(s.UpdateReplicaValue_wg, timeout)
}

var _ replicaupdater.RecordService = &RecordServiceMock{}
