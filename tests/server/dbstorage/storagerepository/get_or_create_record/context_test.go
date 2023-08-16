package getValue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/common/commoncontract"
	"go-skv/common/test"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/dbstorage/storagerepository"
	"testing"
)

func Test_should_return_context_closed_error_when_context_is_closed(t *testing.T) {
	storage := storagerepository.New(0, nil)
	closedCtx := test.NewClosedContext()

	err := storage.GetOrCreateRecord(closedCtx, "", func(dbstoragecontract.Record) {})

	assert.Equal(t, commoncontract.ContextClosedError{}, err)
}
