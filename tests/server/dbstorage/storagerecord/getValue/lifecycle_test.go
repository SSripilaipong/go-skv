package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"testing"
	"time"
)

func Test_should_return_error_when_context_is_cancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	record := storagerecordtest.DoNewRecordWithContext(storagerecordtest.NewFactory(), ctx)

	cancel()
	time.Sleep(time.Millisecond)
	err := record.GetValue(&storagerepositorytest.GetValueMessage{})

	assert.Equal(t, storagerecord.RecordDestroyedError{}, err)
}
