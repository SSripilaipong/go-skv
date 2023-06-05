package storagerepository

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_receive_message_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start())
	defer goutil.WillPanicUnhandledError(storage.Stop)()

	isReceived := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived)
}

func Test_should_receive_multiple_messages_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start())
	defer goutil.WillPanicUnhandledError(storage.Stop)()

	goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)
	isReceived2 := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived2)
}

func Test_should_not_receive_message_after_closed(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start())

	goutil.PanicUnhandledError(storage.Stop())
	isReceived := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.False(t, isReceived)
}
