package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"testing"
)

func Test_should_call_execute_with_existing_record(t *testing.T) {
	storage := storagerepository.New(0, &storagerepositorytest.RecordFactoryMock{})

	var existingRecord, retrievedRecord dbstoragecontract.Record

	test.ContextScope(func(ctx context.Context) {
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

func Test_should_call_failure_when_record_not_exists(t *testing.T) {
	storage := storagerepository.New(0, &storagerepositorytest.RecordFactoryMock{})

	var failureErr error
	test.ContextScope(func(ctx context.Context) {
		ctx, _ = context.WithTimeout(ctx, defaultTimeout)
		goutil.PanicUnhandledError(storage.Start(ctx))

		goutil.PanicUnhandledError(storage.GetRecord(context.Background(), "xxx", func(dbstoragecontract.Record) {}, func(err error) {
			failureErr = err
		}))

	})

	goutil.PanicUnhandledError(storage.Join())
	assert.Equal(t, dbstoragecontract.RecordNotFoundError{}, failureErr)
}

func Test_should_not_call_execute_when_record_not_exists(t *testing.T) {
	storage := storagerepository.New(0, &storagerepositorytest.RecordFactoryMock{})

	var executeIsCalled bool
	test.ContextScope(func(ctx context.Context) {
		ctx, _ = context.WithTimeout(ctx, defaultTimeout)
		goutil.PanicUnhandledError(storage.Start(ctx))

		goutil.PanicUnhandledError(storage.GetRecord(context.Background(), "xxx", func(dbstoragecontract.Record) {
			executeIsCalled = true
		}, func(err error) {
		}))

	})

	goutil.PanicUnhandledError(storage.Join())
	assert.False(t, executeIsCalled)
}
