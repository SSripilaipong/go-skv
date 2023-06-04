package storagerecordtest

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
)

func DoNewRecord(factory storagemanager.RecordFactory) storagemanager.DbRecord {
	return factory.New(context.Background())
}

func DoNewRecordWithContext(factory storagemanager.RecordFactory, ctx context.Context) storagemanager.DbRecord {
	return factory.New(ctx)
}
