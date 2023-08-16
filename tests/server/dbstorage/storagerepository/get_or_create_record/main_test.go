package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"testing"
)

func Test_should_call_success_with_newly_created_record(t *testing.T) {
	newlyCreatedRecord := &dbstoragetest.RecordMock{}
	factory := &storagerepositorytest.RecordFactoryMock{New_Return: newlyCreatedRecord}
	storage := storagerepository.New(0, factory)

	var successRecord dbstoragecontract.Record
	test.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(storage.Start(ctx))

		goutil.PanicUnhandledError(storage.GetOrCreateRecord(ctx, "", func(record dbstoragecontract.Record) { successRecord = record }))
	})

	goutil.PanicUnhandledError(storage.Join())
	assert.Equal(t, newlyCreatedRecord, successRecord)
}

func Test_should_not_create_same_record_twice(t *testing.T) {
	done := make(chan struct{})
	signalDone := func(record dbstoragecontract.Record) { done <- struct{}{} }
	factory := &storagerepositorytest.RecordFactoryMock{}
	storage := storagerepository.New(0, factory)

	test.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(storage.Start(ctx))
		goutil.PanicUnhandledError(storage.GetOrCreateRecord(ctx, "aaa", signalDone))
		goutil.ReceiveWithTimeoutOrPanic(done, defaultTimeout)
		factory.New_CaptureReset()

		goutil.PanicUnhandledError(storage.GetOrCreateRecord(ctx, "aaa", signalDone))
		goutil.ReceiveWithTimeoutOrPanic(done, defaultTimeout)

		assert.False(t, factory.New_IsCalled)
	})
	goutil.PanicUnhandledError(storage.Join())
}

func Test_should_create_new_record_if_the_key_is_not_duplicate_to_existing_ones(t *testing.T) {
	done := make(chan struct{})
	signalDone := func(record dbstoragecontract.Record) { done <- struct{}{} }
	factory := &storagerepositorytest.RecordFactoryMock{}
	storage := storagerepository.New(0, factory)

	test.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(storage.Start(ctx))
		goutil.PanicUnhandledError(storage.GetOrCreateRecord(ctx, "aaa", signalDone))
		goutil.ReceiveWithTimeoutOrPanic(done, defaultTimeout)
		factory.New_CaptureReset()

		goutil.PanicUnhandledError(storage.GetOrCreateRecord(ctx, "bbb", signalDone))
		goutil.ReceiveWithTimeoutOrPanic(done, defaultTimeout)

		assert.True(t, factory.New_IsCalled)
	})
	goutil.PanicUnhandledError(storage.Join())
}

func Test_should_pass_global_context_when_creating_new_record(t *testing.T) {
	doNothing := func(record dbstoragecontract.Record) {}
	factory := &storagerepositorytest.RecordFactoryMock{}
	storage := storagerepository.New(0, factory)

	test.ContextScope(func(ctx context.Context) {
		globalCtx := context.WithValue(ctx, "test", "abc123")
		goutil.PanicUnhandledError(storage.Start(globalCtx))

		goutil.PanicUnhandledError(storage.GetOrCreateRecord(ctx, "", doNothing))
	})

	goutil.PanicUnhandledError(storage.Join())
	assert.Equal(t, "abc123", factory.New_ctx.Value("test"))
}

func Test_should_call_success_with_the_same_record_if_key_is_the_same(t *testing.T) {
	factory := &storagerepositorytest.RecordFactoryMock{}
	storage := storagerepository.New(0, factory)

	var firstRecord, secondRecord dbstoragecontract.Record
	test.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(storage.Start(ctx))

		goutil.PanicUnhandledError(storage.GetOrCreateRecord(ctx, "aaa", func(record dbstoragecontract.Record) {
			firstRecord = record
		}))
		goutil.PanicUnhandledError(storage.GetOrCreateRecord(ctx, "aaa", func(record dbstoragecontract.Record) {
			secondRecord = record
		}))
	})

	goutil.PanicUnhandledError(storage.Join())
	assert.Equal(t, firstRecord, secondRecord)
}
