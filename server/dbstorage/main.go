package dbstorage

import (
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/server/dbstorage/storagerecord"
)

func New(storageBufferSize int, recordBufferSize int) (storagemanager.Interface, chan<- any) {
	ch := make(chan any, storageBufferSize)
	s := storagemanager.New(ch, storagerecord.NewFactory(recordBufferSize))
	return s, ch
}
