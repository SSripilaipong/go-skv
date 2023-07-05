package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func (m *Manager) GetOrCreateRecord(ctx context.Context, key string, success GetOrCreateRecordSuccessCallback) error {
	return m.sendMessage(ctx, GetOrCreateRecordCommand{
		Key:     key,
		Success: success,
	})
}

type GetOrCreateRecordSuccessCallback func(storagerecord.Interface)

type GetOrCreateRecordCommand struct {
	Key     string
	Success GetOrCreateRecordSuccessCallback
}

func (c GetOrCreateRecordCommand) execute(s *state) {
	record, exists := s.records[c.Key]
	if !exists {
		record = s.recordFactory.New(s.ctx)
	}
	c.Success(record)
	s.records[c.Key] = record
}
