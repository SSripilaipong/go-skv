package update

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater"
	"go-skv/server/replicaupdater/replicaupdatercontract"
	"go-skv/tests"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"go-skv/tests/server/replicaupdater/replicaupdatertest"
	"go-skv/tests/server/servertest"
	"testing"
)

func Test_should_create_new_replica_record_when_record_doesnt_exist(t *testing.T) {
	storage := &servertest.DbStorageMock{}
	recordService := &replicaupdatertest.RecordServiceMock{}
	recordFactory := &storagerepositorytest.RecordFactoryMock{}
	factory := replicaupdater.NewFactory(storage, recordService, recordFactory)

	var updater replicaupdatercontract.InboundUpdater
	tests.ContextScope(func(ctx context.Context) {
		updater, _ = factory.NewInboundUpdater(context.WithValue(ctx, "test", "global context"))

		doUpdate := func() {
			_ = updater.Update("", "")
		}
		storage.GetRecord_WaitUntilCalledOnce(defaultTimeout, doUpdate)

		storageFails := func() {
			storage.GetRecord_failure(dbstoragecontract.RecordNotFoundError{})
		}
		recordFactory.New_WaitUntilCalledOnce(defaultTimeout, storageFails)
	})
	updater.Join()

	assert.Equal(t, "global context", recordFactory.New_ctx.Value("test"))
}

func Test_should_set_initialize_replica_record_using_record_service_when_record_doesnt_exist(t *testing.T) {
	newlyCreatedRecord := &dbstoragetest.RecordMock{}
	storage := &servertest.DbStorageMock{}
	recordService := &replicaupdatertest.RecordServiceMock{}
	recordFactory := &storagerepositorytest.RecordFactoryMock{New_Return: newlyCreatedRecord}
	factory := replicaupdater.NewFactory(storage, recordService, recordFactory)

	var updater replicaupdatercontract.InboundUpdater
	tests.ContextScope(func(ctx context.Context) {
		updater, _ = factory.NewInboundUpdater(ctx)

		doUpdate := func() {
			_ = updater.Update("", "xxx")
		}
		storage.GetRecord_WaitUntilCalledOnce(defaultTimeout, doUpdate)

		storageFails := func() {
			storage.GetRecord_failure(dbstoragecontract.RecordNotFoundError{})
		}
		recordService.InitializeReplicaRecord_WaitUntilCalledOnce(defaultTimeout, storageFails)
	})
	updater.Join()

	assert.True(t, recordService.InitilizeReplicaRecord_record == newlyCreatedRecord)
	assert.Equal(t, recordService.InitilizeReplicaRecord_value, "xxx")
}
