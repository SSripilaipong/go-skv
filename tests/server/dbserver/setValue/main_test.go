package setValue

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

func Test_should_call_set_value_usecase(t *testing.T) {
	usecase := &setValueUsecaseMock{}

	_ = dbserverTest.RunWithSetValueUsecase(usecase.New(), func(server dbmanager.DbServer) error {
		return dbserverTest.ConnectWithPort(server.Port(), func(client dbgrpc.DbServiceClient) error {
			return tests.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				_, _ = client.SetValue(ctx, &dbgrpc.SetValueRequest{Key: "Hello", Value: "World"})
				return nil
			})
		})
	})

	assert.Equal(t, &dbusecase.SetValueRequest{Key: "Hello", Value: "World"}, usecase.Request)
}
