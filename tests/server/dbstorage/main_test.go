package dbstorage

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
	"testing"
)

func Test_should_read_message_from_channel(t *testing.T) {
	storageChan := make(chan any)
	storage := dbstorage.New(storageChan)
	goutil.PanicUnhandledError(storage.Start())
	defer goutil.PanicUnhandledError(storage.Stop())

	isReceived := goutil.SendWithTimeout(storageChan, any(struct{}{}), defaultTimeout)

	assert.True(t, isReceived)
}
