package update

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/actormodel"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater"
	"go-skv/tests"
	"go-skv/tests/server/replicaupdater/replicaupdatertest"
	"testing"
)

func Test_should_get_record_from_repo_in_the_first_time(t *testing.T) {
	storage := actormodel.NewTestActor()
	recordUpdaterFactory := &replicaupdatertest.RecordUpdaterFactoryMock{}
	factory := replicaupdater.NewFactory2(storage, recordUpdaterFactory)

	var getRecordMsg any
	var doGetRecord bool
	tests.ContextScope(func(ctx context.Context) {
		updater, _ := factory.NewInboundUpdater(ctx)

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

func Test_should_let_storage_reply_to_newly_spawned_record_updater(t *testing.T) {
	storage := actormodel.NewTestActor()
	recordUpdater := actormodel.NewTestActor()
	recordUpdaterFactory := &replicaupdatertest.RecordUpdaterFactoryMock{New_Return: recordUpdater}
	factory := replicaupdater.NewFactory2(storage, recordUpdaterFactory)

	var getRecordMsg any
	tests.ContextScope(func(ctx context.Context) {
		updater, _ := factory.NewInboundUpdater(ctx)

		_ = actormodel.NewTestActor().FakeTellBlocking(
			tests.ContextWithTimeout(defaultTimeout),
			updater,
			replicaupdater.InboundUpdate{},
		)

		getRecordMsg, _ = storage.SeekMessage(
			tests.ContextWithTimeout(defaultTimeout),
			dbstoragecontract.GetRecord{},
		)
	})

	assert.True(t, recordUpdater == getRecordMsg.(dbstoragecontract.GetRecord).ReplyTo)
}
