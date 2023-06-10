package setValue

import (
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
	err := doSetValue(record)

	assert.Equal(t, storagerecord.RecordDestroyedError{}, err)
}

func Test_should_return_error_when_request_context_is_cancelled(t *testing.T) {
	recordCtx, recordCancel := contextWithDefaultTimeout()
	defer recordCancel()
	record := storagerecordtest.DoNewRecordWithContext(storagerecordtest.NewFactoryWIthChannelBufferSize(0), recordCtx)

	requestCtx, requestCancel := contextWithDefaultTimeout()
	requestCancel()
	err := doSetValueWithContext(record, requestCtx)

	assert.Equal(t, storagerecord.ContextCancelledError{}, err)
}
