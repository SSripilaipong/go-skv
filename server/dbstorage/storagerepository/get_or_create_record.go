package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (m manager) GetOrCreateRecord(ctx context.Context, key string, success func(dbstoragecontract.Record)) error {
	return m.sendMessage(ctx, getOrCreateRecordCommand{
		Key:     key,
		Success: success,
	})
}

type getOrCreateRecordCommand struct {
	Key     string
	Success func(dbstoragecontract.Record)
}

func (c getOrCreateRecordCommand) execute(s *state) {
	record, exists := s.records[c.Key]
	if !exists {
		record = s.recordFactory.New(s.ctx)
	}
	c.Success(record)
	s.records[c.Key] = record
}
