package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/server/dbserver/dbusecase"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbusecase/dbusecasetest"
	"go-skv/tests/server/servertest"
	"testing"
)

func Test_should_get_record_from_repository_with_key(t *testing.T) {
	repoMock := &servertest.DbStorageMock{}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	_, _ = doExecuteWithRequest(usecase, dbusecase.SetValueRequest{Key: "abc"})

	assert.Equal(t, "abc", repoMock.GetOrCreateRecord_key)
}

func Test_should_pass_context_to_repo(t *testing.T) {
	repoMock := &servertest.DbStorageMock{}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	ctx := context.WithValue(context.Background(), "Test", goutil.RandomString(8))
	_, _ = doExecuteWithContext(usecase, ctx)

	assert.Equal(t, ctx.Value("Test"), repoMock.GetOrCreateRecord_ctx.Value("Test"))
}

func Test_should_set_value_to_record(t *testing.T) {
	record := &dbstoragetest.RecordMock{}
	repoMock := &servertest.DbStorageMock{GetOrCreateRecord_execute_record: record}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	_, _ = doExecuteWithRequest(usecase, dbusecase.SetValueRequest{Value: "xxx"})

	assert.Equal(t, "xxx", record.SetValue_value)
}

func Test_should_return_error_when_context_cancelled(t *testing.T) {
	record := &dbstoragetest.RecordMock{SetValue_success_willFail: true}
	repoMock := &servertest.DbStorageMock{GetOrCreateRecord_execute_record: record}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	ctx, _ := contextWithDefaultTimeout()
	_, err := doExecuteWithContext(usecase, ctx)

	assert.Equal(t, dbusecase.ContextCancelledError{}, err)
}

func Test_should_pass_context_to_record(t *testing.T) {
	record := &dbstoragetest.RecordMock{}
	repoMock := &servertest.DbStorageMock{GetOrCreateRecord_execute_record: record}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	ctx := context.WithValue(context.Background(), "Test", goutil.RandomString(8))
	_, _ = doExecuteWithContext(usecase, ctx)

	assert.Equal(t, ctx.Value("Test"), record.SetValue_ctx.Value("Test"))
}
