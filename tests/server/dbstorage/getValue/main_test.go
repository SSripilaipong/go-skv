package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	dbstorageTest "go-skv/tests/server/dbstorage"
	"testing"
)

func Test_should_get_value_from_the_set_record(t *testing.T) {
	storageChan := make(chan any)
	record := &dbstorageTest.RecordMock{}
	factory := &dbstorageTest.RecordFactoryMock{New_Return: record}
	storage := dbstorageTest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	setMessage := &dbstorageTest.SetValueMessage{KeyField: "kkk", ValueField: "vvv"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(setMessage), defaultTimeout)

	getMessage := &dbstorageTest.GetValueMessage{KeyField: "kkk"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(getMessage), defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.GetValue_message, getMessage)
}
