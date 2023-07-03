package getValue

import (
	"github.com/stretchr/testify/assert"
	goutil2 "go-skv/common/util/goutil"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"testing"
)

func Test_should_call_success_with_newly_created_record(t *testing.T) {
	storageChan := make(chan any)
	newlyCreatedRecord := &dbstoragetest.RecordMock{}
	factory := &storagerepositorytest.RecordFactoryMock{New_Return: newlyCreatedRecord}
	storage := storagerepositorytest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil2.PanicUnhandledError(storage.Start())

	var successRecord dbstorage.Record
	message := storagerepository.GetOrCreateRecordMessage{Key: "", Success: func(record dbstorage.Record) { successRecord = record }}
	goutil2.SendWithTimeoutOrPanic[any](storageChan, message, defaultTimeout)

	goutil2.PanicUnhandledError(storage.Stop())

	assert.Equal(t, newlyCreatedRecord, successRecord)
}

func Test_should_not_create_same_record_twice(t *testing.T) {
	storageChan := make(chan any)
	success := func(record dbstorage.Record) {}
	factory := &storagerepositorytest.RecordFactoryMock{}
	storage := storagerepositorytest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil2.PanicUnhandledError(storage.Start())

	message := storagerepository.GetOrCreateRecordMessage{Key: "aaa", Success: success}
	goutil2.SendWithTimeoutOrPanic[any](storageChan, message, defaultTimeout)
	factory.New_CaptureReset()

	goutil2.SendWithTimeoutOrPanic[any](storageChan, message, defaultTimeout)

	goutil2.PanicUnhandledError(storage.Stop())

	assert.False(t, factory.New_IsCalled)
}

func Test_should_create_new_record_if_the_key_is_not_duplicate_to_existing_ones(t *testing.T) {
	storageChan := make(chan any)
	success := func(record dbstorage.Record) {}
	factory := &storagerepositorytest.RecordFactoryMock{}
	storage := storagerepositorytest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil2.PanicUnhandledError(storage.Start())

	goutil2.SendWithTimeoutOrPanic[any](storageChan, storagerepository.GetOrCreateRecordMessage{Key: "aaa", Success: success}, defaultTimeout)
	factory.New_CaptureReset()

	goutil2.SendWithTimeoutOrPanic[any](storageChan, storagerepository.GetOrCreateRecordMessage{Key: "bbb", Success: success}, defaultTimeout)

	goutil2.PanicUnhandledError(storage.Stop())

	assert.True(t, factory.New_IsCalled)
}

func Test_should_pass_context_that_would_be_cancelled_when_stops(t *testing.T) {
	storageChan := make(chan any)
	factory := &storagerepositorytest.RecordFactoryMock{}
	storage := storagerepositorytest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil2.PanicUnhandledError(storage.Start())

	goutil2.SendWithTimeoutOrPanic[any](storageChan, storagerepository.GetOrCreateRecordMessage{Key: "", Success: func(dbstorage.Record) {}}, defaultTimeout)
	passedContext := factory.New_ctx

	goutil2.PanicUnhandledError(storage.Stop())

	_, isCancelled := goutil2.ReceiveNoBlock(passedContext.Done())
	assert.True(t, isCancelled)
}

func Test_should_call_success_with_the_same_record_if_key_is_the_same(t *testing.T) {
	storageChan := make(chan any)
	storage := storagerepositorytest.NewStorageWithChannel(storageChan)
	goutil2.PanicUnhandledError(storage.Start())

	var firstRecord dbstorage.Record
	goutil2.SendWithTimeoutOrPanic[any](storageChan, storagerepository.GetOrCreateRecordMessage{Key: "aaa", Success: func(record dbstorage.Record) {
		firstRecord = record
	}}, defaultTimeout)

	var secondRecord dbstorage.Record
	goutil2.SendWithTimeoutOrPanic[any](storageChan, storagerepository.GetOrCreateRecordMessage{Key: "aaa", Success: func(record dbstorage.Record) {
		secondRecord = record
	}}, defaultTimeout)

	goutil2.PanicUnhandledError(storage.Stop())

	assert.Equal(t, firstRecord, secondRecord)
}
