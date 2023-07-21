package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (m manager) Save(ctx context.Context, key string, record dbstoragecontract.Record, failure func(err error)) error {
	return m.sendMessage(ctx, saveCommand{
		Key:    key,
		Record: record,
	})
}

type saveCommand struct {
	Key    string
	Record dbstoragecontract.Record
}

func (c saveCommand) execute(s *state) {
	s.records[c.Key] = c.Record
}
