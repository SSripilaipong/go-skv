package getValue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbusecase"
	"go-skv/server/storage"
	"testing"
)

func Test_should_send_get_value_message_to_storage(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(storageChan))

	_, _ = execute(&dbusecase.GetValueRequest{Key: "Go"})

	message, ok := goutil.ReceiveNoBlock(storageChan)
	if !ok {
		panic(fmt.Errorf("unexpected error"))
	}

	_, isGetValueMessage := message.(storage.GetValueMessage)
	assert.True(t, isGetValueMessage)
}

func Test_should_send_get_value_message_with_key(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(storageChan))

	_, _ = execute(&dbusecase.GetValueRequest{Key: "Go"})

	message, ok := goutil.ReceiveNoBlock(storageChan)
	if !ok {
		panic(fmt.Errorf("unexpected error"))
	}

	getValueMessage, isGetValueMessage := message.(storage.GetValueMessage)
	if !isGetValueMessage {
		panic(fmt.Errorf("unexpected error"))
	}

	assert.Equal(t, "Go", getValueMessage.Key())
}
