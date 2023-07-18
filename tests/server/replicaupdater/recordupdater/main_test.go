package recordupdater

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/recordupdater"
	"go-skv/tests"
	"testing"
)

func Test_should_get_record_from_storage(t *testing.T) {
	storage := make(chan any)
	factory := recordupdater.NewFactory(storage)

	tests.ContextScope(func(ctx context.Context) {
		updater := factory.New(ctx, "kkk", "")

		msg, ok := waitForMessageWithTimeout[dbstoragecontract.GetRecord](storage)
		assert.True(t, ok)
		assert.Equal(t, dbstoragecontract.GetRecord{Key: "kkk", ReplyTo: updater}, msg)
	})
}

func Test_should_update_replica_value_on_the_retrieved(t *testing.T) {
	storage := make(chan any)
	factory := recordupdater.NewFactory(storage)

	tests.ContextScope(func(ctx context.Context) {
		updater := factory.New(ctx, "", "vvv")
		tests.ClearMessages(storage)

		recordChan := make(chan any)
		sendWithTimeout(updater, dbstoragecontract.RecordChannel{Ch: recordChan})

		msg, ok := waitForMessageWithTimeout[dbstoragecontract.UpdateReplicaValue](recordChan)
		assert.True(t, ok)
		assert.Equal(t, dbstoragecontract.UpdateReplicaValue{Value: "vvv"}, msg)
	})
}
