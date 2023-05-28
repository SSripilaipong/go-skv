package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
	"go-skv/tests/server/dbstorage/dbstoragerecord/dbstoragerecordtest"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"testing"
	"time"
)

func Test_should_call_completed_with_its_value(t *testing.T) {
	factory := dbstoragerecordtest.NewFactory()
	record := factory.New()

	goutil.PanicUnhandledError(record.SetValue(&dbstoragetest.SetValueMessage{KeyField: "aaa", ValueField: "bbb"}))

	getValueMessage := &dbstoragetest.GetValueMessage{KeyField: "aaa"}
	goutil.PanicUnhandledError(record.GetValue(getValueMessage))

	time.Sleep(time.Millisecond)
	goutil.PanicUnhandledError(record.Destroy())

	assert.Equal(t, dbstorage.GetValueResponse{Value: goutil.Pointer("bbb")}, getValueMessage.Completed_Response)
}
