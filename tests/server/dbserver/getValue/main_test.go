package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbmanager"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
	"go-skv/tests"
	grpcTest "go-skv/tests/grpc"
	"go-skv/tests/server/dbserver"
	"testing"
	"time"
)

func Test_should_call_get_value_usecase(t *testing.T) {
	port := grpcTest.GetAvailablePort()
	usecase := &getValueUsecaseMock{}

	_ = dbserverTest.RunWithPortAndGetValueUsecase(port, usecase.New(), func(server dbmanager.DbServer) error {
		return dbserverTest.ConnectWithPort(port, func(client dbgrpc.DbServiceClient) error {
			return tests.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				_, _ = client.GetValue(ctx, &dbgrpc.GetValueRequest{Key: "Hello"})
				return nil
			})
		})
	})

	assert.Equal(t, &dbusecase.GetValueRequest{Key: "Hello"}, usecase.Request)
}
