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

func (i interactor) GetRecord(key string, callback storagerepository.GetRecordSuccessCallback, timeout time.Duration) error {
	i.ch <- storagerepository.GetRecordMessage{
		Key:     key,
		Success: callback,
	}
	return TimeoutError{}
}
