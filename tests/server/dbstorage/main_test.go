package dbstorage

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
	"testing"
)

func Test_should_receive_message_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := dbstorage.New(storageChan)
	goutil.PanicUnhandledError(storage.Start())
	defer goutil.PanicUnhandledError(storage.Stop())

	isReceived := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived)
}

func Test_should_receive_multiple_messages_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := dbstorage.New(storageChan)
	goutil.PanicUnhandledError(storage.Start())
	defer goutil.PanicUnhandledError(storage.Stop())

	goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)
	isReceived2 := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived2)
}
