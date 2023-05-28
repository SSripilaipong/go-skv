package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/tests/server/dbstorage/dbstoragerecord/dbstoragerecordtest"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	goutil2 "go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_call_completed_with_its_value(t *testing.T) {
	factory := dbstoragerecordtest.NewFactory()
	record := factory.New()

	goutil2.PanicUnhandledError(record.SetValue(&dbstoragetest.SetValueMessage{KeyField: "aaa", ValueField: "bbb"}))

	getValueMessage := &dbstoragetest.GetValueMessage{KeyField: "aaa"}
	goutil2.PanicUnhandledError(record.GetValue(getValueMessage))

	time.Sleep(time.Millisecond)
	goutil2.PanicUnhandledError(record.Destroy())

	assert.Equal(t, dbstorage.GetValueResponse{Value: goutil2.Pointer("bbb")}, getValueMessage.Completed_Response)
}
