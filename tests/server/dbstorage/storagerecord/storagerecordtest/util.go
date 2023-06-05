package storagerecordtest

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerecordfactory"
)

func DoNewRecord(factory storagerecordfactory.Interface) storagerecord.DbRecord {
	return factory.New(context.Background())
}

func DoNewRecordWithContext(factory storagerecordfactory.Interface, ctx context.Context) storagerecord.DbRecord {
	return factory.New(ctx)
}
