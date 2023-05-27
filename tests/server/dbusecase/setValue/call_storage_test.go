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
	storageChan := make(chan any, 2)
	execute := dbusecase.SetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		_, _ = execute(context.Background(), &dbusecase.SetValueRequest{Key: "Go", Value: "Lang"})
	}()

	message, ok := goutil.ReceiveWithTimeout(storageChan, defaultTimeout)
	if !ok {
		panic(fmt.Errorf("unexpected error"))
	}

	_, isSetValueMessage := message.(storage.SetValueMessage)
	assert.True(t, isSetValueMessage)
}

func Test_should_send_key_to_storage(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.SetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		_, _ = execute(context.Background(), &dbusecase.SetValueRequest{Key: "Go", Value: "Lang"})
	}()

	message, ok := goutil.ReceiveWithTimeout(storageChan, defaultTimeout)
	if !ok {
		panic(fmt.Errorf("unexpected error"))
	}

	parsedMessage, _ := message.(storage.SetValueMessage)
	assert.Equal(t, "Go", parsedMessage.Key())
}
