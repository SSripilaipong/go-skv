package storagerecordtest

import (
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/dbstorage/storagerecord"
)

func NewFactory() dbstoragecontract.Factory {
	return storagerecord.NewFactory(1)
}

func NewFactoryWIthChannelBufferSize(channelBufferSize int) dbstoragecontract.Factory {
	return storagerecord.NewFactory(channelBufferSize)
}
