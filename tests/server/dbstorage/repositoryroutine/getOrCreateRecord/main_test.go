package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/tests/server/dbstorage/repositoryroutine/repositoryroutinetest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_call_success_with_newly_created_record(t *testing.T) {
	storageChan := make(chan any)
	newlyCreatedRecord := &repositoryroutinetest.RecordMock{}
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
