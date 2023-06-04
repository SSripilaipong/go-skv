package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"testing"
	"time"
)

func Test_should_return_error_when_context_is_cancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	record := storagerecordtest.DoNewRecordWithContext(storagerecordtest.NewFactory(), ctx)

	cancel()
	time.Sleep(time.Millisecond)
	err := record.SetValue(&storagemanagertest.SetValueMessage{})

	assert.Equal(t, storagemanager.RecordDestroyedError{}, err)
}
