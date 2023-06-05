package getRecord

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storageinteractor"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_send_get_or_create_record_message(t *testing.T) {
	ch := make(chan any, 1)
	interactor := storageinteractor.New(ch)

	_ = interactor.GetOrCreateRecord("", func(storagerecord.DbRecord) {}, 0)

	raw := goutil.ReceiveWithTimeoutOrPanic(ch, defaultTimeout)
	assert.True(t, goutil.CanCast[storagerepository.GetOrCreateRecordMessage](raw))
}
