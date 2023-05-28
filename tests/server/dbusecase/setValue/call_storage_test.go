package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/server/dbusecase"
	goutil2 "go-skv/util/goutil"
	"testing"
)

func Test_should_send_set_value_message_to_storage(t *testing.T) {
	storageChan := getStorageChannelAfterExecute(context.Background(), &dbusecase.SetValueRequest{Key: "Go", Value: "Lang"})
	message := goutil2.ReceiveWithTimeoutOrPanic(storageChan, defaultTimeout)

	_, isSetValueMessage := message.(dbstorage.SetValueMessage)
	assert.True(t, isSetValueMessage)
}

func Test_should_send_key_to_storage(t *testing.T) {
	storageChan := getStorageChannelAfterExecute(context.Background(), &dbusecase.SetValueRequest{Key: "Go", Value: "Lang"})
	message := goutil2.ReceiveWithTimeoutOrPanic(storageChan, defaultTimeout)

	parsedMessage := goutil2.CastOrPanic[dbstorage.SetValueMessage](message)
	assert.Equal(t, "Go", parsedMessage.Key())
}

func Test_should_send_value_to_storage(t *testing.T) {
	storageChan := getStorageChannelAfterExecute(context.Background(), &dbusecase.SetValueRequest{Key: "Go", Value: "Lang"})
	message := goutil2.ReceiveWithTimeoutOrPanic(storageChan, defaultTimeout)

	parsedMessage := goutil2.CastOrPanic[dbstorage.SetValueMessage](message)
	assert.Equal(t, "Lang", parsedMessage.Value())
}
