package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"testing"
	"time"
)

func Test_should_return_error_when_record_context_is_cancelled(t *testing.T) {
	ctx, cancel := contextWithDefaultTimeout()
	record := storagerecordtest.DoNewRecordWithContext(storagerecordtest.NewFactory(), ctx)

	cancel()
	time.Sleep(defaultTimeout)
	err := record.GetValue(context.Background(), func(storagerecord.GetValueResponse) {})

	assert.Equal(t, storagerecord.RecordDestroyedError{}, err)
}

func Test_should_return_error_when_request_context_is_cancelled(t *testing.T) {
	recordCtx, recordCancel := contextWithDefaultTimeout()
	defer recordCancel()
	record := storagerecordtest.DoNewRecordWithContext(storagerecordtest.NewFactory(), recordCtx)

	requestCtx, requestCancel := context.WithCancel(context.Background())
	requestCancel()
	err := doGetValueWithContext(record, requestCtx)

	assert.Equal(t, storagerecord.ContextCancelledError{}, err)
}
