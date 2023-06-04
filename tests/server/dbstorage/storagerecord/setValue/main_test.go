package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_call_completed(t *testing.T) {
	factory := storagerecordtest.NewFactory()
	record := factory.New()

	message := &storagemanagertest.SetValueMessage{KeyField: "xxx"}
	goutil.PanicUnhandledError(record.SetValue(message))

	time.Sleep(time.Millisecond)
	goutil.PanicUnhandledError(record.Destroy())

	assert.Equal(t, &storagemanager.SetValueResponse{}, message.Completed_Response)
}
