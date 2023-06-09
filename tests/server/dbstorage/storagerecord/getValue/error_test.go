package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/common/commoncontract"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"testing"
	"time"
)

func Test_should_return_error_when_record_context_is_cancelled(t *testing.T) {
	ctx, cancel := contextWithDefaultTimeout()
	record := storagerecordtest.DoNewRecordWithContext(storagerecordtest.NewFactory(), ctx)

	cancel()
	time.Sleep(defaultTimeout)
	err := doGetValue(record)

	assert.Equal(t, dbstoragecontract.RecordDestroyedError{}, err)
}

func Test_should_return_error_when_request_context_is_cancelled(t *testing.T) {
	recordCtx, recordCancel := contextWithDefaultTimeout()
	defer recordCancel()
	record := storagerecordtest.DoNewRecordWithContext(storagerecordtest.NewFactoryWIthChannelBufferSize(0), recordCtx)

	requestCtx, requestCancel := contextWithDefaultTimeout()
	requestCancel()
	err := doGetValueWithContext(record, requestCtx)

	assert.Equal(t, commoncontract.ContextClosedError{}, err)
}
