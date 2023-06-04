package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_create_new_record(t *testing.T) {
	storageChan := make(chan any)
	factory := &storagemanagertest.RecordFactoryMock{}
	storage := storagemanagertest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&storagemanagertest.SetValueMessage{}), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.True(t, factory.New_IsCalled)
}

func Test_should_not_create_same_record_twice(t *testing.T) {
	storageChan := make(chan any)
	factory := &storagemanagertest.RecordFactoryMock{}
	storage := storagemanagertest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&storagemanagertest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)
	factory.New_CaptureReset()

	goutil.SendWithTimeoutOrPanic(storageChan, any(&storagemanagertest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.False(t, factory.New_IsCalled)
}

func Test_should_pass_context_that_would_be_cancelled_when_stops(t *testing.T) {
	storageChan := make(chan any)
	factory := &storagemanagertest.RecordFactoryMock{}
	storage := storagemanagertest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&storagemanagertest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)
	passedContext := factory.New_ctx

	goutil.PanicUnhandledError(storage.Stop())

	_, isCancelled := goutil.ReceiveNoBlock(passedContext.Done())
	assert.True(t, isCancelled)
}

func Test_should_set_value_to_record(t *testing.T) {
	storageChan := make(chan any)
	record := &storagemanagertest.RecordMock{}
	factory := &storagemanagertest.RecordFactoryMock{New_Return: record}
	storage := storagemanagertest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	m := &storagemanagertest.SetValueMessage{ValueField: "vvv"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(m), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.SetValue_message, m)
}
