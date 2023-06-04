package setValue

import (
	"github.com/stretchr/testify/assert"
	dbstoragetest2 "go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
	goutil2 "go-skv/util/goutil"
	"testing"
)

func Test_should_create_new_record(t *testing.T) {
	storageChan := make(chan any)
	factory := &dbstoragetest2.RecordFactoryMock{}
	storage := dbstoragetest2.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil2.PanicUnhandledError(storage.Start())

	goutil2.SendWithTimeoutOrPanic(storageChan, any(&dbstoragetest2.SetValueMessage{}), defaultTimeout)
	goutil2.PanicUnhandledError(storage.Stop())

	assert.True(t, factory.New_IsCalled)
}

func Test_should_not_create_same_record_twice(t *testing.T) {
	storageChan := make(chan any)
	factory := &dbstoragetest2.RecordFactoryMock{}
	storage := dbstoragetest2.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil2.PanicUnhandledError(storage.Start())

	goutil2.SendWithTimeoutOrPanic(storageChan, any(&dbstoragetest2.SetValueMessage{KeyField: "aaa"}), defaultTimeout)
	factory.New_CaptureReset()

	goutil2.SendWithTimeoutOrPanic(storageChan, any(&dbstoragetest2.SetValueMessage{KeyField: "aaa"}), defaultTimeout)

	goutil2.PanicUnhandledError(storage.Stop())

	assert.False(t, factory.New_IsCalled)
}

func Test_should_set_value_to_record(t *testing.T) {
	storageChan := make(chan any)
	record := &dbstoragetest2.RecordMock{}
	factory := &dbstoragetest2.RecordFactoryMock{New_Return: record}
	storage := dbstoragetest2.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil2.PanicUnhandledError(storage.Start())

	m := &dbstoragetest2.SetValueMessage{ValueField: "vvv"}
	goutil2.SendWithTimeoutOrPanic(storageChan, any(m), defaultTimeout)
	goutil2.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.SetValue_message, m)
}
