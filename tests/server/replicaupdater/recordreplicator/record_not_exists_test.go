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
