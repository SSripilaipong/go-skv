package storagerepository

import (
	"context"
	"errors"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/storagerecord"
)

func (m *Manager) GetRecord(ctx context.Context, key string, execute func(storagerecord.Interface)) error {
	return m.sendMessage(ctx, GetRecordCommand{
		Key:     key,
		Execute: execute,
	})
}

type GetRecordCommand struct {
	Key     string
	Execute func(storagerecord.Interface)
}

func (c GetRecordCommand) execute(s *state) {
	record, exists := s.records[c.Key]
	if !exists {
		goutil.Pointer(errors.New("record not found"))
	}
	c.Execute(record)
}
