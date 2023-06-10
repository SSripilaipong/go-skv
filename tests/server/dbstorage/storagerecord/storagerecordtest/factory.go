package storagerecordtest

import (
	"go-skv/server/dbstorage/storagerecord"
)

func NewFactory() storagerecord.Factory {
	return storagerecord.NewFactory(1)
}

func NewFactoryWIthChannelBufferSize(channelBufferSize int) storagerecord.Factory {
	return storagerecord.NewFactory(channelBufferSize)
}
