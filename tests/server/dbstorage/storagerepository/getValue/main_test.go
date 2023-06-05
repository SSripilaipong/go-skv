package getValue

import (
	"github.com/stretchr/testify/assert"
	storagemanagertest2 "go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	goutil2 "go-skv/util/goutil"
	"testing"
)

func Test_should_get_value_from_the_set_record(t *testing.T) {
	storageChan := make(chan any)
	record := &storagemanagertest2.RecordMock{}
	factory := &storagemanagertest2.RecordFactoryMock{New_Return: record}
	storage := storagemanagertest2.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil2.PanicUnhandledError(storage.Start())

	setMessage := &storagemanagertest2.SetValueMessage{KeyField: "kkk", ValueField: "vvv"}
	goutil2.SendWithTimeoutOrPanic(storageChan, any(setMessage), defaultTimeout)

	getMessage := &storagemanagertest2.GetValueMessage{KeyField: "kkk"}
	goutil2.SendWithTimeoutOrPanic(storageChan, any(getMessage), defaultTimeout)

	goutil2.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.GetValue_message, getMessage)
}
