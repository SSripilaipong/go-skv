package update

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/replicaupdater"
	"go-skv/server/replicaupdater/replicaupdatercontract"
	"go-skv/tests"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/replicaupdater/replicaupdatertest"
	"go-skv/tests/server/servertest"
	"testing"
)

func Test_should_get_record_from_repo_in_the_first_time(t *testing.T) {
	storage := &servertest.DbStorageMock{}
	factory := replicaupdater.NewFactory(storage, &replicaupdatertest.RecordServiceMock{}, nil)

	var updater replicaupdatercontract.InboundUpdater
	tests.ContextScope(func(ctx context.Context) {
		updater, _ = factory.NewInboundUpdater(ctx)

		doUpdate := func() {
			_ = updater.Update("aaa", "")
		}
		storage.GetRecord_WaitUntilCalledOnce(defaultTimeout, doUpdate)
	})
	updater.Join()

	assert.Equal(t, "aaa", storage.GetRecord_key)
}

func Test_should_update_replica_record_value_using_record_service(t *testing.T) {
	record := &dbstoragetest.RecordMock{}
	storage := &servertest.DbStorageMock{}
	recordService := &replicaupdatertest.RecordServiceMock{}
	factory := replicaupdater.NewFactory(storage, recordService, nil)

	var updater replicaupdatercontract.InboundUpdater
	tests.ContextScope(func(ctx context.Context) {
		updater, _ = factory.NewInboundUpdater(ctx)

		doUpdate := func() {
			_ = updater.Update("", "xxx")
		}

		storage.GetRecord_WaitUntilCalledOnce(defaultTimeout, doUpdate)

		recordService.UpdateReplicaValue_WaitUntilCalledOnce(defaultTimeout, func() {
			storage.GetRecord_execute(record)
		})
	})
	updater.Join()

	assert.True(t, recordService.UpdateReplicaValue_record == record)
	assert.Equal(t, "xxx", recordService.UpdateReplicaValue_value)
}
