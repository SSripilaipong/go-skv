package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
	"go-skv/tests/server/dbstorage/dbstoragerecord/dbstoragerecordtest"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"testing"
	"time"
)

func Test_should_call_completed(t *testing.T) {
	factory := dbstoragerecordtest.NewFactory()
	record := factory.New()

	message := &dbstoragetest.SetValueMessage{KeyField: "xxx"}
	goutil.PanicUnhandledError(record.SetValue(message))

	time.Sleep(time.Millisecond)
	goutil.PanicUnhandledError(record.Destroy())

	assert.Equal(t, &dbstorage.SetValueResponse{}, message.Completed_Response)
}
