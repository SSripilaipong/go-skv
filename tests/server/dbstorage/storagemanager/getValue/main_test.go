package getValue

import (
	"github.com/stretchr/testify/assert"
	dbstoragetest2 "go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
	goutil2 "go-skv/util/goutil"
	"testing"
)

func Test_should_get_value_from_the_set_record(t *testing.T) {
	storageChan := make(chan any)
	record := &dbstoragetest2.RecordMock{}
	factory := &dbstoragetest2.RecordFactoryMock{New_Return: record}
	storage := dbstoragetest2.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil2.PanicUnhandledError(storage.Start())

	setMessage := &dbstoragetest2.SetValueMessage{KeyField: "kkk", ValueField: "vvv"}
	goutil2.SendWithTimeoutOrPanic(storageChan, any(setMessage), defaultTimeout)

	getMessage := &dbstoragetest2.GetValueMessage{KeyField: "kkk"}
	goutil2.SendWithTimeoutOrPanic(storageChan, any(getMessage), defaultTimeout)

	goutil2.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.GetValue_message, getMessage)
}
