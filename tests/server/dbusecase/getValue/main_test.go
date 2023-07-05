package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/server/dbserver/dbusecase"
	"go-skv/server/dbstorage"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbusecase/dbusecasetest"
	"go-skv/tests/server/servertest"
	"testing"
)

func Test_should_get_record_from_repository_with_key(t *testing.T) {
	repoMock := &servertest.DbStorageMock{}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	_, _ = doExecuteWithRequest(usecase, dbusecase.GetValueRequest{Key: "abc"})

	assert.Equal(t, "abc", repoMock.GetRecord_key)
}

func Test_should_pass_context_to_repo(t *testing.T) {
	repoMock := &servertest.DbStorageMock{}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	ctx := context.WithValue(context.Background(), "Test", goutil.RandomString(8))
	_, _ = doExecuteWithContext(usecase, ctx)

	assert.Equal(t, ctx.Value("Test"), repoMock.GetRecord_ctx.Value("Test"))
}

func Test_should_return_value_from_record(t *testing.T) {
	record := &dbstoragetest.RecordMock{GetValue_success_response: dbstorage.GetValueResponse{Value: "Hello"}}
	repoMock := &servertest.DbStorageMock{GetRecord_success_record: record}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	response, _ := doExecute(usecase)

	assert.Equal(t, dbusecase.GetValueResponse{Value: "Hello"}, response)
}

func Test_should_return_error_when_context_cancelled(t *testing.T) {
	record := &dbstoragetest.RecordMock{GetValue_success_willFail: true}
	repoMock := &servertest.DbStorageMock{GetRecord_success_record: record}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	ctx, _ := contextWithDefaultTimeout()
	_, err := doExecuteWithContext(usecase, ctx)

	assert.Equal(t, dbusecase.ContextCancelledError{}, err)
}

func Test_should_pass_context_to_record(t *testing.T) {
	record := &dbstoragetest.RecordMock{}
	repoMock := &servertest.DbStorageMock{GetRecord_success_record: record}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	ctx := context.WithValue(context.Background(), "Test", goutil.RandomString(8))
	_, _ = doExecuteWithContext(usecase, ctx)

	assert.Equal(t, ctx.Value("Test"), record.GetValue_ctx.Value("Test"))
}
