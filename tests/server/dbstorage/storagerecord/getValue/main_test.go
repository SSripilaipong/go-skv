package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_call_success_with_its_value(t *testing.T) {
	record := storagerecordtest.DoNewRecord(storagerecordtest.NewFactory())

	goutil.PanicUnhandledError(doSetValueWithValue(record, "bbb"))

	var successResponse storagerecord.GetValueResponse
	goutil.PanicUnhandledError(doGetValueWithSuccessFunc(record, func(response storagerecord.GetValueResponse) { successResponse = response }))

	time.Sleep(defaultTimeout)
	goutil.PanicUnhandledError(record.Destroy())

	assert.Equal(t, storagerecord.GetValueResponse{Value: goutil.Pointer("bbb")}, successResponse)
}
