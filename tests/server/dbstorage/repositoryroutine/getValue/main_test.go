package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/server/dbstorage/repositoryroutine/repositoryroutinetest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_get_value_from_the_set_record(t *testing.T) {
	storageChan := make(chan any)
	record := &repositoryroutinetest.RecordMock{}
	factory := &repositoryroutinetest.RecordFactoryMock{New_Return: record}
	storage := repositoryroutinetest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	setMessage := &repositoryroutinetest.SetValueMessage{KeyField: "kkk", ValueField: "vvv"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(setMessage), defaultTimeout)

	getMessage := &repositoryroutinetest.GetValueMessage{KeyField: "kkk"}
	goutil.SendWithTimeoutOrPanic(storageChan, any(getMessage), defaultTimeout)

	goutil.PanicUnhandledError(storage.Stop())

	assert.Equal(t, record.GetValue_message, getMessage)
}
