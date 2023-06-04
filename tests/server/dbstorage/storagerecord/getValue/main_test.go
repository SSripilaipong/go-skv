package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagemanager"
	dbstoragetest2 "go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	goutil2 "go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_call_completed_with_its_value(t *testing.T) {
	factory := storagerecordtest.NewFactory()
	record := factory.New()

	goutil2.PanicUnhandledError(record.SetValue(&dbstoragetest2.SetValueMessage{KeyField: "aaa", ValueField: "bbb"}))

	getValueMessage := &dbstoragetest2.GetValueMessage{KeyField: "aaa"}
	goutil2.PanicUnhandledError(record.GetValue(getValueMessage))

	time.Sleep(time.Millisecond)
	goutil2.PanicUnhandledError(record.Destroy())

	assert.Equal(t, storagemanager.GetValueResponse{Value: goutil2.Pointer("bbb")}, getValueMessage.Completed_Response)
}
