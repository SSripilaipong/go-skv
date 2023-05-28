package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"testing"
)

func Test_should_create_new_record(t *testing.T) {
	storageChan := make(chan any)
	factory := &dbstoragetest.RecordFactoryMock{}
	storage := dbstoragetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&dbstoragetest.SetValueMessage{}), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.True(t, factory.New_IsCalled)
}

func Test_should_not_create_same_record_twice(t *testing.T) {
	storageChan := make(chan any)
	factory := &dbstoragetest.RecordFactoryMock{}
	storage := dbstoragetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&dbstoragetest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)
	factory.New_CaptureReset()

	goutil.SendWithTimeoutOrPanic(storageChan, any(&dbstoragetest.SetValueMessage{KeyField: "aaa"}), defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.False(t, factory.New_IsCalled)
}

func Test_should_set_value_to_record(t *testing.T) {
	storageChan := make(chan any)
	record := &dbstoragetest.RecordMock{}
	factory := &dbstoragetest.RecordFactoryMock{New_Return: record}
	storage := dbstoragetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	m := &dbstoragetest.SetValueMessage{ValueField: "vvv"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(m), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.SetValue_message, m)
}
