package setValue

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
	"go-skv/server/dbusecase"
	"testing"
)

func Test_should_return_value_when_set_value_completed(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.SetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		message := goutil.ReceiveWithTimeoutOrPanic(storageChan, defaultTimeout)
		setValueMessage := goutil.CastOrPanic[dbstorage.SetValueMessage](message)

		_ = setValueMessage.Completed(dbstorage.SetValueResponse{})
	}()

	result, _ := execute(context.Background(), &dbusecase.SetValueRequest{})

	assert.Equal(t, &dbusecase.SetValueResponse{}, result)
}

func Test_should_return_error_when_context_is_closed(t *testing.T) {
	execute := dbusecase.SetValueUsecase(dbusecase.NewDependency(make(chan any, 2)))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := execute(ctx, &dbusecase.SetValueRequest{Key: "Go"})

	assert.Equal(t, fmt.Errorf("context closed"), err)
}