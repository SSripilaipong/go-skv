package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"testing"
	"time"
)

func Test_should_call_success_with_its_value(t *testing.T) {
	record := storagerecordtest.DoNewRecord(storagerecordtest.NewFactory())

	goutil.PanicUnhandledError(doSetValueWithValue(record, "bbb"))

	var successResponse dbstoragecontract.RecordData
	goutil.PanicUnhandledError(doGetValueWithSuccessFunc(record, func(response dbstoragecontract.RecordData) { successResponse = response }))

	time.Sleep(defaultTimeout)
	goutil.PanicUnhandledError(record.Destroy())

	assert.Equal(t, dbstoragecontract.RecordData{Value: "bbb"}, successResponse)
}
