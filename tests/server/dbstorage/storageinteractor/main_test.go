package storageinteractor

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storageinteractor"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_send_get_record_message(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storageinteractor.New(ch)

	_ = interactor.GetRecord("", func(storagerecord.DbRecord) {}, 0)

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	assert.True(t, goutil.CanCast[storagerepository.GetRecordMessage](raw))
}

func Test_should_send_get_record_message_with_key_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storageinteractor.New(ch)

	_ = interactor.GetRecord("aaa", func(storagerecord.DbRecord) {}, 0)

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil.CastOrPanic[storagerepository.GetRecordMessage](raw)
	assert.Equal(t, "aaa", message.Key)
}

func Test_should_send_get_record_message_with_success_callback_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storageinteractor.New(ch)

	var isTheSameFunction bool
	_ = interactor.GetRecord("", func(storagerecord.DbRecord) { isTheSameFunction = true }, 0)

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil.CastOrPanic[storagerepository.GetRecordMessage](raw)

	isTheSameFunction = false
	message.Success(nil)
	assert.True(t, isTheSameFunction)
}

func Test_should_return_timeout_error_when_cannot_send_message_within_timeout(t *testing.T) {
	ch := make(chan any)
	interactor := storageinteractor.New(ch)

	err := interactor.GetRecord("", func(storagerecord.DbRecord) {}, time.Second)

	assert.Equal(t, storageinteractor.TimeoutError{}, err)
}

func Test_should_not_return_timeout_error_when_message_is_sent_within_timeout(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storageinteractor.New(ch)

	err := interactor.GetRecord("", func(storagerecord.DbRecord) {}, time.Second)

	assert.Nil(t, err)
}
