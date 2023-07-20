package recordreplicator

import (
	"context"
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, dbstoragecontract.SetRecordMode{Mode: dbstoragecontract.ReplicaMode}, setRecordMode)
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

		_, ok := waitForMessageWithTimeout[dbstoragecontract.SetRecordMode](createdRecord)
		assert.True(t, ok)

		saveRecord, ok := waitForMessageWithTimeout[dbstoragecontract.SaveRecord](storage)
		assert.True(t, ok)
		assert.Equal(t, dbstoragecontract.SaveRecord{Key: "fff", Ch: createdRecord}, saveRecord)
	})
}
