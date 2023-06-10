package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbstorage/repositoryroutine/repositoryroutinetest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_create_new_record(t *testing.T) {
	storageChan := make(chan any)
	factory := &repositoryroutinetest.RecordFactoryMock{}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&repositoryroutinetest.SetValueMessage{}), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.True(t, factory.New_IsCalled)
}

func Test_should_not_create_same_record_twice(t *testing.T) {
	storageChan := make(chan any)
	factory := &repositoryroutinetest.RecordFactoryMock{}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&repositoryroutinetest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)
	factory.New_CaptureReset()

	goutil.SendWithTimeoutOrPanic(storageChan, any(&repositoryroutinetest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.False(t, factory.New_IsCalled)
}

func Test_should_pass_context_that_would_be_cancelled_when_stops(t *testing.T) {
	storageChan := make(chan any)
	factory := &repositoryroutinetest.RecordFactoryMock{}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&repositoryroutinetest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	passedContext := factory.New_ctx
	_, isCancelled := goutil.ReceiveNoBlock(passedContext.Done())
	assert.True(t, isCancelled)
}

func Test_should_set_value_to_record(t *testing.T) {
	storageChan := make(chan any)
	record := &dbstoragetest.RecordMock{}
	factory := &repositoryroutinetest.RecordFactoryMock{New_Return: record}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	m := &repositoryroutinetest.SetValueMessage{ValueField: "vvv"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(m), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.SetValue_value, "vvv")
}
