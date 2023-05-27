package setValue

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbusecase"
	"go-skv/server/storage"
	"testing"
)

func Test_should_send_set_value_message_to_storage(t *testing.T) {
	getStorageChannelAfterExecute(context.Background(), &dbusecase.SetValueRequest{Key: "Go", Value: "Lang"}, func(storageChan chan any) {
		message, ok := goutil.ReceiveWithTimeout(storageChan, defaultTimeout)
		if !ok {
			panic(fmt.Errorf("unexpected error"))
		}

		_, isSetValueMessage := message.(storage.SetValueMessage)
		assert.True(t, isSetValueMessage)
	})
}

func Test_should_send_key_to_storage(t *testing.T) {
	getStorageChannelAfterExecute(context.Background(), &dbusecase.SetValueRequest{Key: "Go", Value: "Lang"}, func(storageChan chan any) {
		message, ok := goutil.ReceiveWithTimeout(storageChan, defaultTimeout)
		if !ok {
			panic(fmt.Errorf("unexpected error"))
		}

		parsedMessage, _ := message.(storage.SetValueMessage)
		assert.Equal(t, "Go", parsedMessage.Key())
	})
}

func Test_should_send_value_to_storage(t *testing.T) {
	getStorageChannelAfterExecute(context.Background(), &dbusecase.SetValueRequest{Key: "Go", Value: "Lang"}, func(storageChan chan any) {
		message, ok := goutil.ReceiveWithTimeout(storageChan, defaultTimeout)
		if !ok {
			panic(fmt.Errorf("unexpected error"))
		}

		parsedMessage, _ := message.(storage.SetValueMessage)
		assert.Equal(t, "Lang", parsedMessage.Value())
	})
}
