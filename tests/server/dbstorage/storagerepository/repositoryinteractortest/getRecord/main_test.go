package getRecord

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagerecord"
	repositoryinteractor2 "go-skv/server/dbstorage/storagerepository/repositoryinteractor"
	"go-skv/server/dbstorage/storagerepository/repositoryroutine"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_send_get_record_message(t *testing.T) {
	ch := make(chan any, 1)
	interactor := repositoryinteractor2.New(ch)

	_ = interactor.GetRecord("", func(storagerecord.DbRecord) {}, defaultTimeout)

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	assert.True(t, goutil.CanCast[repositoryroutine.GetRecordMessage](raw))
}

func Test_should_send_get_record_message_with_key_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := repositoryinteractor2.New(ch)

	_ = interactor.GetRecord("aaa", func(storagerecord.DbRecord) {}, defaultTimeout)

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil.CastOrPanic[repositoryroutine.GetRecordMessage](raw)
	assert.Equal(t, "aaa", message.Key)
}

func Test_should_send_get_record_message_with_success_callback_to_repository(t *testing.T) {
	ch := make(chan any, 1)
	interactor := repositoryinteractor2.New(ch)

	var isTheSameFunction bool
	_ = interactor.GetRecord("", func(storagerecord.DbRecord) { isTheSameFunction = true }, defaultTimeout)

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	message := goutil.CastOrPanic[repositoryroutine.GetRecordMessage](raw)

	isTheSameFunction = false
	message.Success(nil)
	assert.True(t, isTheSameFunction)
}

func Test_should_return_timeout_error_when_cannot_send_message_within_timeout(t *testing.T) {
	ch := make(chan any)
	interactor := repositoryinteractor2.New(ch)

	err := interactor.GetRecord("", func(storagerecord.DbRecord) {}, defaultTimeout)

	assert.Equal(t, repositoryinteractor2.TimeoutError{}, err)
}

func Test_should_not_return_timeout_error_when_message_is_sent_within_timeout(t *testing.T) {
	ch := make(chan any, 1)
	interactor := repositoryinteractor2.New(ch)

	err := interactor.GetRecord("", func(storagerecord.DbRecord) {}, defaultTimeout)

	assert.Nil(t, err)
}
