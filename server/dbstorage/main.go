package dbstorage

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/server/dbstorage/storagerecord"
)

func New(storageBufferSize int, recordBufferSize int) (storagemanager.Interface, chan<- any) {
	ctx := context.Background()
	ch := make(chan any, storageBufferSize)
	s := storagemanager.New(ctx, ch, storagerecord.NewFactory(ctx, recordBufferSize))
	return s, ch
}
