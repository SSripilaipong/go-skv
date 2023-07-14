package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/tests"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"testing"
)

func Test_should_call_success_with_existing_record(t *testing.T) {
	storage := storagerepository.New(0, &storagerepositorytest.RecordFactoryMock{})

	var existingRecord, retrievedRecord dbstoragecontract.Record

	tests.ContextScope(func(ctx context.Context) {
		ctx, _ = context.WithTimeout(ctx, defaultTimeout)
		goutil.PanicUnhandledError(storage.Start(ctx))
		goutil.PanicUnhandledError(storage.GetOrCreateRecord(context.Background(), "aaa", func(record dbstoragecontract.Record) {
			existingRecord = record
		}))

		goutil.PanicUnhandledError(storage.GetRecord(context.Background(), "aaa", func(record dbstoragecontract.Record) {
			retrievedRecord = record
		}, func(error) {}))

	})

	goutil.PanicUnhandledError(storage.Join())
	assert.Equal(t, existingRecord, retrievedRecord)
}
