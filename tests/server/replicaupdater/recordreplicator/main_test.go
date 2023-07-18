package recordreplicator

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/recordreplicator"
	"go-skv/tests"
	"testing"
)

func Test_should_get_record_from_storage(t *testing.T) {
	storage := make(chan any)
	factory := recordreplicator.NewFactory(storage)

	tests.ContextScope(func(ctx context.Context) {
		replicator := factory.New(ctx, "kkk", "")

		msg, ok := waitForMessageWithTimeout[dbstoragecontract.GetRecord](storage)
		assert.True(t, ok)
		assert.Equal(t, dbstoragecontract.GetRecord{Key: "kkk", ReplyTo: replicator}, msg)
	})
}

func Test_should_update_replica_value_on_the_retrieved(t *testing.T) {
	storage := make(chan any)
	factory := recordreplicator.NewFactory(storage)

	tests.ContextScope(func(ctx context.Context) {
		replicator := factory.New(ctx, "", "vvv")
		tests.ClearMessages(storage)

		recordChan := make(chan any)
		sendWithTimeout(replicator, dbstoragecontract.RecordChannel{Ch: recordChan})

		msg, ok := waitForMessageWithTimeout[dbstoragecontract.UpdateReplicaValue](recordChan)
		assert.True(t, ok)
		assert.Equal(t, dbstoragecontract.UpdateReplicaValue{Value: "vvv", ReplyTo: replicator}, msg)
	})
}
