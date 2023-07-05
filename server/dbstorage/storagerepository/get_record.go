package storagerepository

import (
	"context"
	"errors"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/storagerecord"
)

func (m manager) GetRecord(ctx context.Context, key string, execute func(storagerecord.Interface)) error {
	return m.sendMessage(ctx, getRecordCommand{
		Key:     key,
		Execute: execute,
	})
}

type getRecordCommand struct {
	Key     string
	Execute func(storagerecord.Interface)
}

func (c getRecordCommand) execute(s *state) {
	record, exists := s.records[c.Key]
	if !exists {
		goutil.PanicUnhandledError(errors.New("record not found"))
	}
	c.Execute(record)
}
