package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_call_completed_with_its_value(t *testing.T) {
	record := storagerecordtest.DoNewRecord(storagerecordtest.NewFactory())

	goutil.PanicUnhandledError(record.SetValue(&storagemanagertest.SetValueMessage{KeyField: "aaa", ValueField: "bbb"}))

	getValueMessage := &storagemanagertest.GetValueMessage{KeyField: "aaa"}
	goutil.PanicUnhandledError(record.GetValue(getValueMessage))

	time.Sleep(time.Millisecond)
	goutil.PanicUnhandledError(record.Destroy())

	assert.Equal(t, storagerecord.GetValueResponse{Value: goutil.Pointer("bbb")}, getValueMessage.Completed_Response)
}
