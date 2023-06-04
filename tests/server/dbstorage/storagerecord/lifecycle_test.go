package storagerecord

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_return_error_when_destroyed(t *testing.T) {
	record := storagerecordtest.DoNewRecord(storagerecordtest.NewFactory())

	goutil.PanicUnhandledError(record.Destroy())
	time.Sleep(time.Millisecond)
	err := storagerecordtest.SendAnyMessage(record)

	assert.Equal(t, storagemanager.RecordDestroyedError{}, err)
}
