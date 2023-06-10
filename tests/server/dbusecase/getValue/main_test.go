package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/server/dbusecase"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/tests/server/dbusecase/dbusecasetest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_get_record_from_repository_with_key(t *testing.T) {
	repoMock := &dbusecasetest.RepoMock{}
	usecase := dbusecase.GetValueUsecase(dbusecase.NewDependencyV2(nil, repoMock))

	_, _ = usecase(context.Background(), dbusecase.GetValueRequest{Key: "abc"})

	assert.Equal(t, "abc", repoMock.GetRecord_key)
}

func Test_should_pass_context_to_repo(t *testing.T) {
	repoMock := &dbusecasetest.RepoMock{}
	usecase := dbusecase.GetValueUsecase(dbusecase.NewDependencyV2(nil, repoMock))

	ctx := context.WithValue(context.Background(), "Test", goutil.RandomString(8))
	_, _ = usecase(ctx, dbusecase.GetValueRequest{Key: "abc"})

	assert.Equal(t, ctx.Value("Test"), repoMock.GetRecord_ctx.Value("Test"))
}

func Test_should_return_value_from_record(t *testing.T) {
	record := &dbstoragetest.RecordMock{GetValue_success_response: dbstorage.GetValueResponse{Value: goutil.Pointer("Hello")}}
	repoMock := &dbusecasetest.RepoMock{GetRecord_success_record: record}
	usecase := dbusecase.GetValueUsecase(dbusecase.NewDependencyV2(nil, repoMock))

	response, _ := usecase(context.Background(), dbusecase.GetValueRequest{})

	assert.Equal(t, dbusecase.GetValueResponse{Value: goutil.Pointer("Hello")}, response)
}

func Test_should_return_error_when_context_cancelled(t *testing.T) {
	record := &dbstoragetest.RecordMock{GetValue_success_willFail: true}
	repoMock := &dbusecasetest.RepoMock{GetRecord_success_record: record}
	usecase := dbusecase.GetValueUsecase(dbusecase.NewDependencyV2(nil, repoMock))

	ctx, cancel := contextWithDefaultTimeout()
	cancel()
	_, err := usecase(ctx, dbusecase.GetValueRequest{})

	assert.Equal(t, dbusecase.ContextCancelledError{}, err)
}

func Test_should_pass_context_to_record(t *testing.T) {
	record := &dbstoragetest.RecordMock{}
	repoMock := &dbusecasetest.RepoMock{GetRecord_success_record: record}
	usecase := dbusecase.GetValueUsecase(dbusecase.NewDependencyV2(nil, repoMock))

	ctx := context.WithValue(context.Background(), "Test", goutil.RandomString(8))
	_, _ = usecase(ctx, dbusecase.GetValueRequest{})

	assert.Equal(t, ctx.Value("Test"), record.GetValue_ctx.Value("Test"))
}
