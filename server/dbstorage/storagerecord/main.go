package storagerecord

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
)

func NewFactory(ctx context.Context, channelBufferSize int) storagemanager.RecordFactory {
	return &recordFactory{
		ctx:          ctx,
		chBufferSize: channelBufferSize,
	}
}
