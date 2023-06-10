package getRecord

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/repositoryinteractor"
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_send_get_record_message(t *testing.T) {
	ch := make(chan any, 1)
	interactor := repositoryinteractor.New(ch)

	_ = interactor.GetRecord(context.Background(), "", func(storagerecord.Interface) {})

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	assert.True(t, goutil.CanCast[repositoryroutine.GetRecordMessage](raw))
}

func Test_should_send_get_record_message_with_key_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := repositoryinteractor.New(ch)

	_ = interactor.GetRecord(context.Background(), "aaa", func(storagerecord.Interface) {})

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil.CastOrPanic[repositoryroutine.GetRecordMessage](raw)
	assert.Equal(t, "aaa", message.Key)
}

func Test_should_send_get_record_message_with_success_callback_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := repositoryinteractor.New(ch)

	var isTheSameFunction bool
	_ = interactor.GetRecord(context.Background(), "", func(storagerecord.Interface) { isTheSameFunction = true })

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil.CastOrPanic[repositoryroutine.GetRecordMessage](raw)

	isTheSameFunction = false
	message.Success(nil)
	assert.True(t, isTheSameFunction)
}

func Test_should_return_context_cancelled_error_when_context_is_cancelled(t *testing.T) {
	ch := make(chan any)
	interactor := repositoryinteractor.New(ch)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := interactor.GetRecord(ctx, "", func(storagerecord.Interface) {})

	assert.Equal(t, repositoryinteractor.ContextCancelledError{}, err)
}
