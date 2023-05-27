package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	dbstorageTest "go-skv/tests/server/dbstorage"
	"testing"
)

func Test_should_create_new_record(t *testing.T) {
	storageChan := make(chan any)
	factory := &recordFactoryMock{}
	storage := dbstorageTest.NewStorageWithChannelAndRecordFactory(storageChan, factory)
	goutil.PanicUnhandledError(storage.Start())

	goutil.SendWithTimeoutOrPanic(storageChan, any(&message{}), defaultTimeout)
	goutil.PanicUnhandledError(storage.Stop())

	assert.True(t, factory.New_IsCalled)
}
