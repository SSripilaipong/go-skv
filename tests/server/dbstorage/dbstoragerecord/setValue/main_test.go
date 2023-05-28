package setValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
	"go-skv/tests/server/dbstorage/dbstoragerecord/dbstoragerecordtest"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"testing"
)

func Test_should_call_completed(t *testing.T) {
	factory := dbstoragerecordtest.NewFactory()
	record := factory.New()
	defer goutil.WillPanicUnhandledError(record.Destroy)

	message := &dbstoragetest.SetValueMessage{}
	goutil.PanicUnhandledError(record.SetValue(message))

	assert.Equal(t, dbstorage.SetValueResponse{}, message.Completed_Response)
}
