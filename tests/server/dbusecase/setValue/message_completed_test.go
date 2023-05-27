package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbusecase"
	"go-skv/server/storage"
	"testing"
)

func Test_should_return_value_when_set_value_completed(t *testing.T) {
	storageChan := make(chan any, 2)
	execute := dbusecase.SetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		message := goutil.ReceiveWithTimeoutOrPanic(storageChan, defaultTimeout)
		setValueMessage := goutil.CastOrPanic[storage.SetValueMessage](message)

		_ = setValueMessage.Completed(storage.SetValueResponse{})
	}()

	result, _ := execute(context.Background(), &dbusecase.SetValueRequest{})

	assert.Equal(t, &dbusecase.SetValueResponse{}, result)
}
