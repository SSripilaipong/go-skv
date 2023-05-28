package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"testing"
)

func Test_should_get_value_from_the_set_record(t *testing.T) {
	storageChan := make(chan any)
	record := &dbstoragetest.RecordMock{}
	factory := &dbstoragetest.RecordFactoryMock{New_Return: record}
	storage := dbstoragetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	setMessage := &dbstoragetest.SetValueMessage{KeyField: "kkk", ValueField: "vvv"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(setMessage), defaultTimeout)

	getMessage := &dbstoragetest.GetValueMessage{KeyField: "kkk"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(getMessage), defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.GetValue_message, getMessage)
}
