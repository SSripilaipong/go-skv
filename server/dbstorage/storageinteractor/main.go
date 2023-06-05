package storageinteractor

import (
	"go-skv/server/dbstorage/storagerepository"
	"time"
)

func New(ch chan<- any) Interface {
	return interactor{ch: ch}
}

type interactor struct {
	ch chan<- any
}

func (i interactor) GetRecord(key string, success storagerepository.GetRecordSuccessCallback, timeout time.Duration) error {
	return i.sendMessage(storagerepository.GetRecordMessage{
		Key:     key,
		Success: success,
	}, timeout)
}

func (i interactor) GetOrCreateRecord(key string, success storagerepository.GetOrCreateRecordSuccessCallback, timeout time.Duration) error {
	return i.sendMessage(storagerepository.GetOrCreateRecordMessage{
		Key:     key,
		Success: success,
	}, timeout)
}

func (i interactor) sendMessage(message any, timeout time.Duration) error {
	select {
	case i.ch <- message:
	case <-time.After(timeout):
		return TimeoutError{}
	}
	return nil
}
