package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (m manager) GetRecord(ctx context.Context, key string, execute func(dbstoragecontract.Record), failure func(err error)) error {
	return m.sendMessage(ctx, getRecordCommand{
		Key:     key,
		Execute: execute,
		Failure: failure,
	})
}

type getRecordCommand struct {
	Key     string
	Execute func(dbstoragecontract.Record)
	Failure func(error)
}

func (c getRecordCommand) execute(s *state) {
	record, exists := s.records[c.Key]
	if !exists {
		c.Failure(dbstoragecontract.RecordNotFoundError{})
		return
	}
	c.Execute(record)
}
