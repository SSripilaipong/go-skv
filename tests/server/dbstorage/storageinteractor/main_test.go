package storageinteractor

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storageinteractor"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_send_get_record_message(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storageinteractor.New(ch)

	_ = interactor.GetRecord("", func(storagerecord.DbRecord) {})

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	assert.True(t, goutil.CanCast[storagerepository.GetRecordMessage](raw))
}

func Test_should_send_get_record_message_with_key_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storageinteractor.New(ch)

	_ = interactor.GetRecord("aaa", func(storagerecord.DbRecord) {})

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil.CastOrPanic[storagerepository.GetRecordMessage](raw)
	assert.Equal(t, "aaa", message.Key)
}

func Test_should_send_get_record_message_with_success_callback_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storageinteractor.New(ch)

	var isTheSameFunction bool
	_ = interactor.GetRecord("", func(storagerecord.DbRecord) { isTheSameFunction = true })

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil.CastOrPanic[storagerepository.GetRecordMessage](raw)

	isTheSameFunction = false
	message.Success(nil)
	assert.True(t, isTheSameFunction)
}
