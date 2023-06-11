package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_call_success_with_existing_record(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start())

	var existingRecord dbstorage.Record
	goutil.SendWithTimeoutOrPanic[any](storageChan, storagerepository.GetOrCreateRecordMessage{Key: "aaa", Success: func(record dbstorage.Record) {
		existingRecord = record
	}}, defaultTimeout)

	var retrievedRecord dbstorage.Record
	goutil.SendWithTimeoutOrPanic[any](storageChan, storagerepository.GetRecordMessage{Key: "aaa", Success: func(record dbstorage.Record) {
		retrievedRecord = record
	}}, defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, existingRecord, retrievedRecord)
}
