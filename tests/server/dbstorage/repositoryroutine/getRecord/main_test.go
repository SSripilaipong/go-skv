package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/tests/server/dbstorage/repositoryroutine/repositoryroutinetest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_call_success_with_existing_record(t *testing.T) {
	storageChan := make(chan any)
	newlyCreatedRecord := &repositoryroutinetest.RecordMock{}
	factory := &repositoryroutinetest.RecordFactoryMock{New_Return: newlyCreatedRecord}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	var existingRecord dbstorage.Record
	goutil.SendWithTimeoutOrPanic[any](storageChan, repositoryroutine.GetOrCreateRecordMessage{Key: "aaa", Success: func(record dbstorage.Record) {
		existingRecord = record
	}}, defaultTimeout)

	var retrievedRecord dbstorage.Record
	goutil.SendWithTimeoutOrPanic[any](storageChan, repositoryroutine.GetRecordMessage{Key: "aaa", Success: func(record dbstorage.Record) {
		retrievedRecord = record
	}}, defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, existingRecord, retrievedRecord)
}
