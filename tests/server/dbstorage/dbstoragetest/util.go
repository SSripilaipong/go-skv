package dbstoragetest

import (
	"context"
	"go-skv/server/dbstorage"
)

func NewStorageWithChannel(ch chan any) dbstorage.Interface {
	return dbstorage.New(context.Background(), ch, nil)
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory dbstorage.RecordFactory) dbstorage.Interface {
	return dbstorage.New(context.Background(), ch, factory)
}

func NewStorageWithChannelAndContext(ch chan any, ctx context.Context) dbstorage.Interface {
	return dbstorage.New(ctx, ch, nil)
}
