package setValue

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbusecase"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_return_value_when_set_value_completed(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := newUsecaseWithStorageChan(storageChan)

	go func() {
		message := goutil.ReceiveWithTimeoutOrPanic(storageChan, defaultTimeout)
		setValueMessage := goutil.CastOrPanic[storagerecord.SetValueMessage](message)

		_ = setValueMessage.Completed(storagerecord.SetValueResponse{})
	}()

	result, _ := execute(context.Background(), dbusecase.SetValueRequest{})

	assert.Equal(t, dbusecase.SetValueResponse{}, result)
}

func Test_should_return_error_when_context_is_closed(t *testing.T) {
	execute := newUsecase()
	ctx := newClosedContext()

	_, err := execute(ctx, dbusecase.SetValueRequest{Key: "Go"})

	assert.Equal(t, fmt.Errorf("context closed"), err)
}
