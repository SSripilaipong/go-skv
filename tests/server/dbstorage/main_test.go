package dbstorageTest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	goutil2 "go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_receive_message_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := dbstoragetest.NewStorageWithChannel(storageChan)
	goutil2.PanicUnhandledError(storage.Start())
	defer goutil2.WillPanicUnhandledError(storage.Stop)()

	isReceived := goutil2.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived)
}

func Test_should_receive_multiple_messages_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := dbstoragetest.NewStorageWithChannel(storageChan)
	goutil2.PanicUnhandledError(storage.Start())
	defer goutil2.WillPanicUnhandledError(storage.Stop)()

	goutil2.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)
	isReceived2 := goutil2.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived2)
}

func Test_should_not_receive_message_after_closed(t *testing.T) {
	storageChan := make(chan any)
	storage := dbstoragetest.NewStorageWithChannel(storageChan)
	goutil2.PanicUnhandledError(storage.Start())

	goutil2.PanicUnhandledError(storage.Stop())
	isReceived := goutil2.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.False(t, isReceived)
}

func Test_should_not_receive_message_after_context_completed(t *testing.T) {
	storageChan := make(chan any)
	ctx, cancel := context.WithCancel(context.Background())
	storage := dbstoragetest.NewStorageWithChannelAndContext(storageChan, ctx)
	goutil2.PanicUnhandledError(storage.Start())

	cancel()
	time.Sleep(time.Millisecond)
	isReceived := goutil2.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.False(t, isReceived)
}
