package storagerepository

import (
	"context"
	"errors"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/storagerecord"
)

func (i interactor) GetRecord(ctx context.Context, key string, success GetRecordSuccessCallback) error {
	return i.sendMessage(ctx, GetRecordCommand{
		Key:     key,
		Success: success,
	})
}

type GetRecordSuccessCallback func(storagerecord.Interface)

type GetRecordCommand struct {
	Key     string
	Success GetRecordSuccessCallback
}

func (c GetRecordCommand) execute(s *state) {
	record, exists := s.records[c.Key]
	if !exists {
		goutil.Pointer(errors.New("record not found"))
	}
	c.Success(record)
}
