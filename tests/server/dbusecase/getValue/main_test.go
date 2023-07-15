package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/server/dbserver/dbusecase"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbusecase/dbusecasetest"
	"go-skv/tests/server/servertest"
	"testing"
)

func Test_should_get_record_from_repository_with_key(t *testing.T) {
	repoMock := &servertest.DbStorageMock{}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	ctx, _ := contextWithDefaultTimeout()
	_, _ = doExecuteWithContextAndRequest(usecase, ctx, dbusecase.GetValueRequest{Key: "abc"})

	assert.Equal(t, "abc", repoMock.GetRecord_key)
}

func Test_should_pass_context_to_repo(t *testing.T) {
	repoMock := &servertest.DbStorageMock{}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)
	ctxId := goutil.RandomString(8)
	ctx, _ := contextWithDefaultTimeout()

	ctx = context.WithValue(ctx, "Test", ctxId)
	_, _ = doExecuteWithContext(usecase, ctx)

	assert.Equal(t, ctxId, repoMock.GetRecord_ctx.Value("Test"))
}

func Test_should_return_value_from_record(t *testing.T) {
	repoMock := &servertest.DbStorageMock{}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	var response dbusecase.GetValueResponse

	usecaseDone := make(chan struct{})
	executeInBackground := func() {
		go func() {
			response, _ = doExecute(usecase)
			usecaseDone <- struct{}{}
		}()
	}

	repoMock.GetRecord_WaitUntilCalledOnce(defaultTimeout, executeInBackground)
	repoMock.GetRecord_execute(&dbstoragetest.RecordMock{
		GetValue_success_response: dbstoragecontract.RecordData{Value: "Hello"},
	})

	goutil.ReceiveWithTimeout(usecaseDone, defaultTimeout)
	assert.Equal(t, dbusecase.GetValueResponse{Value: "Hello"}, response)
}

func Test_should_return_error_when_context_cancelled(t *testing.T) {
	repoMock := &servertest.DbStorageMock{}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	ctx, _ := contextWithDefaultTimeout()
	_, err := doExecuteWithContext(usecase, ctx)

	assert.Equal(t, dbusecase.ContextCancelledError{}, err)
}

func Test_should_pass_context_to_record(t *testing.T) {
	record := &dbstoragetest.RecordMock{}
	repoMock := &servertest.DbStorageMock{}
	usecase := dbusecasetest.NewUsecaseWithRepo(repoMock)

	ctxId := goutil.RandomString(8)
	ctx := context.WithValue(context.Background(), "Test", ctxId)

	usecaseDone := make(chan struct{})
	executeInBackground := func() {
		go func() {
			_, _ = doExecuteWithContext(usecase, ctx)
			usecaseDone <- struct{}{}
		}()
	}

	repoMock.GetRecord_WaitUntilCalledOnce(defaultTimeout, executeInBackground)
	repoMock.GetRecord_execute(record)

	assert.Equal(t, ctxId, record.GetValue_ctx.Value("Test"))
}
