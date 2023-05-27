package dbstorageTest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"testing"
)

func Test_should_receive_message_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start())
	defer goutil.WillPanicUnhandledError(storage.Stop)()

	isReceived := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived)
}

func Test_should_receive_multiple_messages_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start())
	defer goutil.WillPanicUnhandledError(storage.Stop)()

	goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)
	isReceived2 := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived2)
}

func Test_should_not_receive_message_after_closed(t *testing.T) {
	storageChan := make(chan any)
	storage := NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start())

	goutil.PanicUnhandledError(storage.Stop())
	isReceived := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.False(t, isReceived)
}

func Test_should_not_receive_message_after_context_completed(t *testing.T) {
	storageChan := make(chan any)
	ctx, cancel := context.WithCancel(context.Background())
	storage := NewStorageWithChannelAndContext(storageChan, ctx)
	goutil.PanicUnhandledError(storage.Start())

	cancel()
	isReceived := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.False(t, isReceived)
}
