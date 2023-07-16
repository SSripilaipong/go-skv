package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (m manager) Add(ctx context.Context, key string, record dbstoragecontract.Record, failure func(err error)) error {
	//TODO implement me
	panic("implement me")
}
