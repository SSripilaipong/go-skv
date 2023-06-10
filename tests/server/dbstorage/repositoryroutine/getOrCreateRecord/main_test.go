package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbstorage/repositoryroutine/repositoryroutinetest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_call_success_with_newly_created_record(t *testing.T) {
	storageChan := make(chan any)
	newlyCreatedRecord := &dbstoragetest.RecordMock{}
	factory := &repositoryroutinetest.RecordFactoryMock{New_Return: newlyCreatedRecord}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	var successRecord dbstorage.Record
	message := repositoryroutine.GetOrCreateRecordMessage{Key: "", Success: func(record dbstorage.Record) { successRecord = record }}
	goutil.SendWithTimeoutOrPanic[any](storageChan, message, defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, newlyCreatedRecord, successRecord)
}

func Test_should_not_create_same_record_twice(t *testing.T) {
	storageChan := make(chan any)
	success := func(record dbstorage.Record) {}
	factory := &repositoryroutinetest.RecordFactoryMock{}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	message := repositoryroutine.GetOrCreateRecordMessage{Key: "aaa", Success: success}
	goutil.SendWithTimeoutOrPanic[any](storageChan, message, defaultTimeout)
	factory.New_CaptureReset()

	goutil.SendWithTimeoutOrPanic[any](storageChan, message, defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.False(t, factory.New_IsCalled)
}

func Test_should_create_new_record_if_the_key_is_not_duplicate_to_existing_ones(t *testing.T) {
	storageChan := make(chan any)
	success := func(record dbstorage.Record) {}
	factory := &repositoryroutinetest.RecordFactoryMock{}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic[any](storageChan, repositoryroutine.GetOrCreateRecordMessage{Key: "aaa", Success: success}, defaultTimeout)
	factory.New_CaptureReset()

	goutil.SendWithTimeoutOrPanic[any](storageChan, repositoryroutine.GetOrCreateRecordMessage{Key: "bbb", Success: success}, defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.True(t, factory.New_IsCalled)
}

func Test_should_pass_context_that_would_be_cancelled_when_stops(t *testing.T) {
	storageChan := make(chan any)
	factory := &repositoryroutinetest.RecordFactoryMock{}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic[any](storageChan, repositoryroutine.GetOrCreateRecordMessage{Key: "", Success: func(dbstorage.Record) {}}, defaultTimeout)
	passedContext := factory.New_ctx

	goutil.PanicUnhandledError(storage.Stop())

	_, isCancelled := goutil.ReceiveNoBlock(passedContext.Done())
	assert.True(t, isCancelled)
}

func Test_should_call_success_with_the_same_record_if_key_is_the_same(t *testing.T) {
	storageChan := make(chan any)
	storage := repositoryroutinetest.NewStorageWithChannel(storageChan)
	goutil.PanicUnhandledError(storage.Start())

	var firstRecord dbstorage.Record
	goutil.SendWithTimeoutOrPanic[any](storageChan, repositoryroutine.GetOrCreateRecordMessage{Key: "aaa", Success: func(record dbstorage.Record) {
		firstRecord = record
	}}, defaultTimeout)

	var secondRecord dbstorage.Record
	goutil.SendWithTimeoutOrPanic[any](storageChan, repositoryroutine.GetOrCreateRecordMessage{Key: "aaa", Success: func(record dbstorage.Record) {
		secondRecord = record
	}}, defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, firstRecord, secondRecord)
}
