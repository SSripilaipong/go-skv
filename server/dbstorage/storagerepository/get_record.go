package storagerepository

import (
	"context"
	"errors"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (m manager) GetRecord(ctx context.Context, key string, execute func(dbstoragecontract.Record)) error {
	return m.sendMessage(ctx, getRecordCommand{
		Key:     key,
		Execute: execute,
	})
}

type getRecordCommand struct {
	Key     string
	Execute func(dbstoragecontract.Record)
}

func (c getRecordCommand) execute(s *state) {
	record, exists := s.records[c.Key]
	if !exists {
		goutil.PanicUnhandledError(errors.New("record not found"))
	}
	c.Execute(record)
}
