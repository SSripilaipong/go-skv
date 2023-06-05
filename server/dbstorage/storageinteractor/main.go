package storageinteractor

import (
	"go-skv/server/dbstorage/storagerepository"
)

func New(ch chan<- any) Interface {
	return interactor{ch: ch}
}

type interactor struct {
	ch chan<- any
}

func (i interactor) GetRecord(key string, success storagerepository.GetRecordSuccessCallback) error {
	i.ch <- storagerepository.GetRecordMessage{
		Key: key,
	}
	return nil
}
