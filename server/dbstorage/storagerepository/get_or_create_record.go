package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func (m manager) GetOrCreateRecord(ctx context.Context, key string, success func(storagerecord.Interface)) error {
	return m.sendMessage(ctx, getOrCreateRecordCommand{
		Key:     key,
		Success: success,
	})
}

type getOrCreateRecordCommand struct {
	Key     string
	Success func(storagerecord.Interface)
}

func (c getOrCreateRecordCommand) execute(s *state) {
	record, exists := s.records[c.Key]
	if !exists {
		record = s.recordFactory.New(s.ctx)
	}
	c.Success(record)
	s.records[c.Key] = record
}
