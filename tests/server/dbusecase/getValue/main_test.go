package getValue

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/server/dbusecase"
	goutil2 "go-skv/util/goutil"
	"testing"
)

func Test_should_send_get_value_message_to_storage(t *testing.T) {
	storageChan := getStorageChannelAfterExecute(context.Background(), &dbusecase.GetValueRequest{Key: "Go"})

	message := goutil2.ReceiveWithTimeoutOrPanic(storageChan, defaultTimeout)

	assert.True(t, goutil2.CanCast[dbstorage.GetValueMessage](message))
}

func Test_should_send_get_value_message_with_key(t *testing.T) {
	storageChan := getStorageChannelAfterExecute(context.Background(), &dbusecase.GetValueRequest{Key: "Go"})

	message := goutil2.ReceiveWithTimeoutOrPanic(storageChan, defaultTimeout)
	assert.Equal(t, "Go", goutil2.CastOrPanic[dbstorage.GetValueMessage](message).Key())
}

func Test_should_return_value_when_get_value_completed(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		message := goutil2.ReceiveWithTimeoutOrPanic(storageChan, defaultTimeout)
		getValueMessage := goutil2.CastOrPanic[dbstorage.GetValueMessage](message)

		_ = getValueMessage.Completed(dbstorage.GetValueResponse{Value: goutil2.Pointer("Lang")})
	}()

	result, _ := execute(context.Background(), &dbusecase.GetValueRequest{Key: "Go"})

	assert.Equal(t, &dbusecase.GetValueResponse{Value: goutil2.Pointer("Lang")}, result)
}

func Test_should_return_error_when_context_is_closed(t *testing.T) {
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(make(chan any, 2)))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := execute(ctx, &dbusecase.GetValueRequest{Key: "Go"})

	assert.Equal(t, fmt.Errorf("context closed"), err)
}
