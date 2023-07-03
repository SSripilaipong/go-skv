package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/commoncontract"
	goutil2 "go-skv/common/util/goutil"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository"
	"testing"
)

func Test_should_send_get_record_message(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storagerepository.NewInteractor(ch)

	_ = interactor.GetRecord(context.Background(), "", func(storagerecord.Interface) {})

	raw := goutil2.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	assert.True(t, goutil2.CanCast[storagerepository.GetRecordMessage](raw))
}

func Test_should_send_get_record_message_with_key_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storagerepository.NewInteractor(ch)

	_ = interactor.GetRecord(context.Background(), "aaa", func(storagerecord.Interface) {})

	raw := goutil2.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil2.CastOrPanic[storagerepository.GetRecordMessage](raw)
	assert.Equal(t, "aaa", message.Key)
}

func Test_should_send_get_record_message_with_success_callback_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storagerepository.NewInteractor(ch)

	var isTheSameFunction bool
	_ = interactor.GetRecord(context.Background(), "", func(storagerecord.Interface) { isTheSameFunction = true })

	raw := goutil2.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil2.CastOrPanic[storagerepository.GetRecordMessage](raw)

	isTheSameFunction = false
	message.Success(nil)
	assert.True(t, isTheSameFunction)
}

func Test_should_return_context_cancelled_error_when_context_is_cancelled(t *testing.T) {
	ch := make(chan any)
	interactor := storagerepository.NewInteractor(ch)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := interactor.GetRecord(ctx, "", func(storagerecord.Interface) {})

	assert.Equal(t, commoncontract.ContextClosedError{}, err)
}
