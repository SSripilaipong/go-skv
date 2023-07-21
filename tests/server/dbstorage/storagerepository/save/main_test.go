package save

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/tests"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
	"testing"
)

func Test_should_be_able_to_retrieve_the_same_record_after_saved(t *testing.T) {
	storage := storagerepository.New(0, &storagerepositorytest.RecordFactoryMock{})
	recordToSave := &dbstoragetest.RecordMock{}

	var retrievedRecord dbstoragecontract.Record

	tests.ContextScope(func(ctx context.Context) {
		ctx, _ = context.WithTimeout(ctx, defaultTimeout)
		goutil.PanicUnhandledError(storage.Start(ctx))
		goutil.PanicUnhandledError(storage.Save(ctx, "kkk", recordToSave, func(error) {}))

		goutil.PanicUnhandledError(storage.GetRecord(context.Background(), "kkk", func(record dbstoragecontract.Record) {
			retrievedRecord = record
		}, func(error) {}))

	})

	goutil.PanicUnhandledError(storage.Join())
	assert.True(t, recordToSave == retrievedRecord)
}
