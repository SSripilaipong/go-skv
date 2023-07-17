package update

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/actormodel"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater"
	"go-skv/tests"
	"testing"
)

func Test_should_get_record_from_repo_in_the_first_time(t *testing.T) {
	storage := actormodel.NewTestActor()
	factory := replicaupdater.NewFactory2(storage)

	var updater actormodel.ActorRef
	var getRecordMsg any
	var doGetRecord bool
	tests.ContextScope(func(ctx context.Context) {
		updater, _ = factory.NewInboundUpdater(ctx)

		_ = actormodel.NewTestActor().FakeTellBlocking(
			tests.ContextWithTimeout(defaultTimeout),
			updater,
			replicaupdater.InboundUpdate{Key: "aaa"},
		)

		getRecordMsg, doGetRecord = storage.SeekMessage(
			tests.ContextWithTimeout(defaultTimeout),
			dbstoragecontract.GetRecord{},
		)
	})

	assert.True(t, doGetRecord)
	assert.Equal(t, "aaa", getRecordMsg.(dbstoragecontract.GetRecord).Key)
}
