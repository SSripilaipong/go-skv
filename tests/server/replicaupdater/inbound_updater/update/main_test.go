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

func Test_should_get_record_from_repo_in_the_first_time(t *testing.T) {
	storage := &servertest.DbStorageMock{}
	factory := replicaupdater.NewFactory(storage, &replicaupdatertest.RecordServiceMock{}, nil)

	var updater replicaupdatercontract.InboundUpdater
	tests.ContextScope(func(ctx context.Context) {
		updater, _ = factory.NewInboundUpdater(ctx)

		storage.GetRecord_WaitUntilCalledOnce(defaultTimeout, func() {
			_ = updater.Update("aaa", "")
		})
	})
	updater.Join()

	assert.Equal(t, "aaa", storage.GetRecord_key)
}

func Test_should_update_replica_record_value_using_record_service(t *testing.T) {
	record := &dbstoragetest.RecordMock{}
	storage := &servertest.DbStorageMock{GetRecord_execute_record: record}
	recordService := &replicaupdatertest.RecordServiceMock{}
	factory := replicaupdater.NewFactory(storage, recordService, nil)

	var updater replicaupdatercontract.InboundUpdater
	tests.ContextScope(func(ctx context.Context) {
		updater, _ = factory.NewInboundUpdater(ctx)

		recordService.UpdateReplicaValue_WaitUntilCalledOnce(defaultTimeout, func() {
			_ = updater.Update("", "xxx")
		})
	})
	updater.Join()

	assert.True(t, recordService.UpdateReplicaValue_record == record)
	assert.Equal(t, "xxx", recordService.UpdateReplicaValue_value)
}

func Test_should_create_new_replica_record_when_record_doesnt_exist(t *testing.T) {
	storage := &servertest.DbStorageMock{GetRecord_failure_err: dbstoragecontract.RecordNotFoundError{}}
	recordService := &replicaupdatertest.RecordServiceMock{}
	recordFactory := &storagerepositorytest.RecordFactoryMock{}
	factory := replicaupdater.NewFactory(storage, recordService, recordFactory)

	var updater replicaupdatercontract.InboundUpdater
	tests.ContextScope(func(ctx context.Context) {
		updater, _ = factory.NewInboundUpdater(context.WithValue(ctx, "test", "global context"))

		recordFactory.New_WaitUntilCalledOnce(defaultTimeout, func() {
			_ = updater.Update("", "")
		})
	})
	updater.Join()

	assert.Equal(t, "global context", recordFactory.New_ctx.Value("test"))
}
