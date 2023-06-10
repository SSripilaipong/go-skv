package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbusecase"
	"go-skv/tests/server/dbusecase/dbusecasetest"
	"testing"
)

func Test_should_get_record_from_repository_with_key(t *testing.T) {
	repoMock := &dbusecasetest.RepoMock{}
	usecase := dbusecase.GetValueUsecaseV2(dbusecase.NewDependencyV2(nil, repoMock))

	_, _ = usecase(context.Background(), dbusecase.GetValueRequest{Key: "abc"})

	assert.Equal(t, "abc", repoMock.GetRecord_key)
}
