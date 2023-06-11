package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbusecase"
	"go-skv/tests/server/dbusecase/dbusecasetest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_get_record_from_repository_with_key(t *testing.T) {
	repoMock := &dbusecasetest.RepoMock{}
	usecase := newUsecaseWithRepo(repoMock)

	_, _ = doExecuteWithRequest(usecase, dbusecase.SetValueRequest{Key: "abc"})

	assert.Equal(t, "abc", repoMock.GetOrCreateRecord_key)
}

func Test_should_pass_context_to_repo(t *testing.T) {
	repoMock := &dbusecasetest.RepoMock{}
	usecase := newUsecaseWithRepo(repoMock)

	ctx := context.WithValue(context.Background(), "Test", goutil.RandomString(8))
	_, _ = doExecuteWithContext(usecase, ctx)

	assert.Equal(t, ctx.Value("Test"), repoMock.GetOrCreateRecord_ctx.Value("Test"))
}
