package storagerepository

import (
	"github.com/stretchr/testify/assert"
	goutil2 "go-skv/common/util/goutil"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"testing"
)

func Test_should_receive_message_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil2.PanicUnhandledError(storage.Start())
	defer goutil2.WillPanicUnhandledError(storage.Stop)()

	isReceived := goutil2.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived)
}

func Test_should_receive_multiple_messages_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil2.PanicUnhandledError(storage.Start())
	defer goutil2.WillPanicUnhandledError(storage.Stop)()

	goutil2.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)
	isReceived2 := goutil2.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived2)
}

func Test_should_not_receive_message_after_closed(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil2.PanicUnhandledError(storage.Start())

	goutil2.PanicUnhandledError(storage.Stop())
	isReceived := goutil2.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.False(t, isReceived)
}
