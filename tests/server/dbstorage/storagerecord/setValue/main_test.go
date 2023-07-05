package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"testing"
	"time"
)

func Test_should_call_success_with_response(t *testing.T) {
	record := storagerecordtest.DoNewRecord(storagerecordtest.NewFactory())

	var successResponse dbstoragecontract.RecordData
	success := func(response dbstoragecontract.RecordData) { successResponse = response }
	goutil.PanicUnhandledError(doSetValueWithValueAndSuccessFunc(record, "yyy", success))

	time.Sleep(defaultTimeout)
	goutil.PanicUnhandledError(record.Destroy())

	assert.Equal(t, dbstoragecontract.RecordData{Value: "yyy"}, successResponse)
}
