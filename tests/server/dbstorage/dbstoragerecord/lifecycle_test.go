package dbstoragerecord

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/tests/server/dbstorage/dbstoragerecord/dbstoragerecordtest"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_return_error_when_destroyed(t *testing.T) {
	factory := dbstoragerecordtest.NewFactory()
	record := factory.New()

	goutil.PanicUnhandledError(record.Destroy())
	time.Sleep(time.Millisecond)
	err := dbstoragerecordtest.SendAnyMessage(record)

	assert.Equal(t, dbstorage.RecordDestroyedError{}, err)
}
