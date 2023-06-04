package storagemanagertest

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
)

func NewStorageWithChannel(ch chan any) storagemanager.Interface {
	return storagemanager.New(context.Background(), ch, nil)
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory storagemanager.RecordFactory) storagemanager.Interface {
	return storagemanager.New(context.Background(), ch, factory)
}

func NewStorageWithChannelAndContext(ch chan any, ctx context.Context) storagemanager.Interface {
	return storagemanager.New(ctx, ch, nil)
}
