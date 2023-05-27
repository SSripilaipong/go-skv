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

	goutil.SendWithTimeoutOrPanic(storageChan, any(&message{}), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.True(t, factory.New_IsCalled)
}

func Test_should_set_value_to_record(t *testing.T) {
	storageChan := make(chan any)
	record := &dbstorageTest.RecordMock{}
	factory := &dbstorageTest.RecordFactoryMock{New_Return: record}
	storage := dbstorageTest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&message{value: "vvv"}), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.SetValue_value, "vvv")
}
