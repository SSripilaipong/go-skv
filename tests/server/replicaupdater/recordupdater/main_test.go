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
	storage := make(chan any, 1)
	factory := recordupdater.NewFactory(storage)

	var updater chan<- any
	tests.ContextScope(func(ctx context.Context) {
		updater = factory.New(ctx, "kkk", "")
	})

	msg, ok := waitForMessageWithTimeout(storage, dbstoragecontract.GetRecord{})
	assert.True(t, ok)
	assert.Equal(t, dbstoragecontract.GetRecord{Key: "kkk", ReplyTo: updater}, msg)
}
