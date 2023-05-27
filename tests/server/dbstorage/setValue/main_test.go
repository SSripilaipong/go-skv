package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	dbstorageTest "go-skv/tests/server/dbstorage"
	"testing"
)

func Test_should_create_new_record(t *testing.T) {
	storageChan := make(chan any)
	factory := &dbstorageTest.RecordFactoryMock{}
	storage := dbstorageTest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&dbstorageTest.SetValueMessage{}), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.True(t, factory.New_IsCalled)
}

func Test_should_not_create_same_record_twice(t *testing.T) {
	storageChan := make(chan any)
	factory := &dbstorageTest.RecordFactoryMock{}
	storage := dbstorageTest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&dbstorageTest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)
	factory.New_CaptureReset()

	goutil.SendWithTimeoutOrPanic(storageChan, any(&dbstorageTest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.False(t, factory.New_IsCalled)
}

func Test_should_set_value_to_record(t *testing.T) {
	storageChan := make(chan any)
	record := &dbstorageTest.RecordMock{}
	factory := &dbstorageTest.RecordFactoryMock{New_Return: record}
	storage := dbstorageTest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	m := &dbstorageTest.SetValueMessage{ValueField: "vvv"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(m), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.SetValue_message, m)
}
