package recordreplicator

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/recordreplicator"
	"go-skv/tests"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"testing"
)

func Test_should_create_new_record_if_retrieved_record_channel_is_empty(t *testing.T) {
	storage := make(chan any)
	recordFactory := &storagerepositorytest.RecordFactoryMock{}
	factory := recordreplicator.NewFactory(storage, recordFactory)

	tests.ContextScope(func(ctx context.Context) {
		replicator, _ := factory.New(context.WithValue(ctx, "test", "the same ctx"), "", "")
		defer close(replicator)

		getRecord, _ := waitForMessageWithTimeout[dbstoragecontract.GetRecord](storage)

		recordFactory.NewActor_WaitUntilCalledOnce(defaultTimeout, func() {
			sendWithTimeout(getRecord.ReplyTo, dbstoragecontract.RecordChannel{Ch: nil})
		})

		assert.Equal(t, "the same ctx", recordFactory.NewActor_ctx.Value("test"))
	})
}

func Test_should_set_mode_to_replica_on_the_created_record(t *testing.T) {
	storage := make(chan any)
	createdRecord := make(chan any)
	recordFactory := &storagerepositorytest.RecordFactoryMock{}
	factory := recordreplicator.NewFactory(storage, recordFactory)

	tests.ContextScope(func(ctx context.Context) {
		replicator, _ := factory.New(ctx, "", "")
		defer close(replicator)

		getRecord, _ := waitForMessageWithTimeout[dbstoragecontract.GetRecord](storage)

		recordFactory.NewActor_Return = createdRecord
		recordFactory.NewActor_WaitUntilCalledOnce(defaultTimeout, func() {
			sendWithTimeout(getRecord.ReplyTo, dbstoragecontract.RecordChannel{Ch: nil})
		})

		setRecordMode, ok := waitForMessageWithTimeout[dbstoragecontract.SetRecordMode](createdRecord)

		assert.True(t, ok)
		assert.Equal(t, dbstoragecontract.ReplicaMode, setRecordMode.Mode)
		assert.Equal(t, "set record mode", setRecordMode.Memo)
	})
}

func Test_should_save_the_created_record_to_storage(t *testing.T) {
	storage := make(chan any)
	createdRecord := make(chan any)
	recordFactory := &storagerepositorytest.RecordFactoryMock{}
	factory := recordreplicator.NewFactory(storage, recordFactory)

	tests.ContextScope(func(ctx context.Context) {
		replicator, _ := factory.New(ctx, "fff", "")
		defer close(replicator)

		getRecord, _ := waitForMessageWithTimeout[dbstoragecontract.GetRecord](storage)

		recordFactory.NewActor_Return = createdRecord
		recordFactory.NewActor_WaitUntilCalledOnce(defaultTimeout, func() {
			sendWithTimeout(getRecord.ReplyTo, dbstoragecontract.RecordChannel{Ch: nil})
		})

		setRecordMode, _ := waitForMessageWithTimeout[dbstoragecontract.SetRecordMode](createdRecord)

		sendWithTimeout(setRecordMode.ReplyTo, commonmessage.Ok{Memo: "set record mode"})

		saveRecord, ok := waitForMessageWithTimeout[dbstoragecontract.SaveRecord](storage)

		assert.True(t, ok)
		assert.Equal(t, "fff", saveRecord.Key)
		assert.Equal(t, chan<- any(createdRecord), saveRecord.Ch)
		assert.Equal(t, "save record", saveRecord.Memo)
	})
}

func Test_should_not_save_the_created_record_to_storage_before_ok_from_set_record_mode(t *testing.T) {
	storage := make(chan any)
	createdRecord := make(chan any)
	recordFactory := &storagerepositorytest.RecordFactoryMock{}
	factory := recordreplicator.NewFactory(storage, recordFactory)

	tests.ContextScope(func(ctx context.Context) {
		replicator, _ := factory.New(ctx, "", "")
		defer close(replicator)

		getRecord, _ := waitForMessageWithTimeout[dbstoragecontract.GetRecord](storage)

		recordFactory.NewActor_Return = createdRecord
		recordFactory.NewActor_WaitUntilCalledOnce(defaultTimeout, func() {
			sendWithTimeout(getRecord.ReplyTo, dbstoragecontract.RecordChannel{Ch: nil})
		})

		setRecordMode, _ := waitForMessageWithTimeout[dbstoragecontract.SetRecordMode](createdRecord)

		sendWithTimeout(setRecordMode.ReplyTo, commonmessage.Ok{Memo: "from something else"})

		_, ok := waitForMessageWithTimeout[dbstoragecontract.SaveRecord](storage)

		assert.False(t, ok)
	})
}

func Test_should_retry_updating_after_record_is_saved(t *testing.T) {
	storage := make(chan any)
	createdRecord := make(chan any)
	recordFactory := &storagerepositorytest.RecordFactoryMock{}
	factory := recordreplicator.NewFactory(storage, recordFactory)

	tests.ContextScope(func(ctx context.Context) {
		replicator, _ := factory.New(ctx, "fff", "ggg")
		defer close(replicator)

		getRecord, _ := waitForMessageWithTimeout[dbstoragecontract.GetRecord](storage)

		recordFactory.NewActor_Return = createdRecord
		recordFactory.NewActor_WaitUntilCalledOnce(defaultTimeout, func() {
			sendWithTimeout(getRecord.ReplyTo, dbstoragecontract.RecordChannel{Ch: nil})
		})

		setRecordMode, _ := waitForMessageWithTimeout[dbstoragecontract.SetRecordMode](createdRecord)

		sendWithTimeout(setRecordMode.ReplyTo, commonmessage.Ok{Memo: "set record mode"})

		saveRecord, _ := waitForMessageWithTimeout[dbstoragecontract.SaveRecord](storage)

		sendWithTimeout(saveRecord.ReplyTo, commonmessage.Ok{Memo: "save record"})

		retryUpdating, ok := waitForMessageWithTimeout[dbstoragecontract.UpdateReplicaValue](createdRecord)
		assert.True(t, ok)
		assert.Equal(t, "ggg", retryUpdating.Value)
		assert.Equal(t, "update replica", retryUpdating.Memo)
	})
}
