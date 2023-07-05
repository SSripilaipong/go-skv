package storagerepository

import (
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"testing"
)

func Test_should_receive_message_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start(nil))
	defer goutil.WillPanicUnhandledError(storage.Join)()

	isReceived := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived)
}

func Test_should_receive_multiple_messages_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start(nil))
	defer goutil.WillPanicUnhandledError(storage.Join)()

	goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)
	isReceived2 := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived2)
}

func Test_should_not_receive_message_after_closed(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start(nil))

	goutil.PanicUnhandledError(storage.Join())
	isReceived := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.False(t, isReceived)
}
