package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_call_success_with_response(t *testing.T) {
	record := storagerecordtest.DoNewRecord(storagerecordtest.NewFactory())

	var successResponse dbstorage.SetValueResponse
	success := func(response storagerecord.SetValueResponse) { successResponse = response }
	goutil.PanicUnhandledError(record.SetValue("yyy", success))

	time.Sleep(time.Millisecond)
	goutil.PanicUnhandledError(record.Destroy())

	assert.Equal(t, storagerecord.SetValueResponse{Value: goutil.Pointer("yyy")}, successResponse)
}
