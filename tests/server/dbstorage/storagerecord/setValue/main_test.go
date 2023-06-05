package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_call_completed(t *testing.T) {
	record := storagerecordtest.DoNewRecord(storagerecordtest.NewFactory())

	message := &storagemanagertest.SetValueMessage{KeyField: "xxx"}
	goutil.PanicUnhandledError(record.SetValue(message))

	time.Sleep(time.Millisecond)
	goutil.PanicUnhandledError(record.Destroy())

	assert.Equal(t, &storagerecord.SetValueResponse{}, message.Completed_Response)
}
