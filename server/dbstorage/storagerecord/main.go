package storagerecord

import (
	"go-skv/server/dbstorage/storagemanager"
)

func NewFactory(channelBufferSize int) storagemanager.RecordFactory {
	return &recordFactory{
		chBufferSize: channelBufferSize,
	}
}
