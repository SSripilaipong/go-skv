package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbmanager"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
	"go-skv/tests"
	"go-skv/tests/server/dbserver"
	"testing"
	"time"
)

func Test_should_call_get_value_usecase(t *testing.T) {
	usecase := &getValueUsecaseMock{}

	_ = dbserverTest.RunWithGetValueUsecase(usecase.New(), func(server dbmanager.DbServer) error {
		return dbserverTest.ConnectWithPort(server.Port(), func(client dbgrpc.DbServiceClient) error {
			return tests.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				_, _ = client.GetValue(ctx, &dbgrpc.GetValueRequest{Key: "Hello"})
				return nil
			})
		})
	})

	assert.Equal(t, &dbusecase.GetValueRequest{Key: "Hello"}, usecase.Request)
}
