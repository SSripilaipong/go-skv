package getValue

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbusecase"
	"go-skv/server/storage"
	"testing"
	"time"
)

const defaultTimeout = 100 * time.Millisecond

func Test_should_send_get_value_message_to_storage(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		_, _ = execute(context.Background(), &dbusecase.GetValueRequest{Key: "Go"})
	}()

	message, ok := goutil.ReceiveWithTimeout(storageChan, defaultTimeout)
	if !ok {
		panic(fmt.Errorf("unexpected error"))
	}

	_, isGetValueMessage := message.(storage.GetValueMessage)
	assert.True(t, isGetValueMessage)
}

func Test_should_send_get_value_message_with_key(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		_, _ = execute(context.Background(), &dbusecase.GetValueRequest{Key: "Go"})
	}()

	message, ok := goutil.ReceiveWithTimeout(storageChan, defaultTimeout)
	if !ok {
		panic(fmt.Errorf("unexpected error"))
	}

	getValueMessage, isGetValueMessage := message.(storage.GetValueMessage)
	if !isGetValueMessage {
		panic(fmt.Errorf("unexpected error"))
	}

	assert.Equal(t, "Go", getValueMessage.Key())
}

func Test_should_return_value_when_get_value_completed(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		message, ok := goutil.ReceiveWithTimeout(storageChan, defaultTimeout)
		if !ok {
			panic(fmt.Errorf("unexpected error"))
		}

		getValueMessage, isGetValueMessage := message.(storage.GetValueMessage)
		if !isGetValueMessage {
			panic(fmt.Errorf("unexpected error"))
		}

		_ = getValueMessage.Completed(storage.GetValueResponse{Value: "Lang"})
	}()

	result, _ := execute(context.Background(), &dbusecase.GetValueRequest{Key: "Go"})

	assert.Equal(t, &dbusecase.GetValueResponse{Value: "Lang"}, result)
}

func Test_should_return_error_when_context_is_closed(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(storageChan))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := execute(ctx, &dbusecase.GetValueRequest{Key: "Go"})

	assert.Equal(t, fmt.Errorf("context closed"), err)
}
