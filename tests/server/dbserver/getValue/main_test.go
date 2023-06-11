package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbserver"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
	"go-skv/tests"
	"go-skv/tests/server/dbserver"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_call_get_value_usecase(t *testing.T) {
	usecase := &getValueUsecaseMock{}

	_ = dbserverTest.RunWithGetValueUsecase(usecase.New(), func(server dbserver.Interface) error {
		return dbserverTest.ConnectWithPort(server.Port(), func(client dbgrpc.DbServiceClient) error {
			return tests.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				_, _ = client.GetValue(ctx, &dbgrpc.GetValueRequest{Key: "Hello"})
				return nil
			})
		})
	})

	assert.Equal(t, dbusecase.GetValueRequest{Key: "Hello"}, usecase.Request)
}

func Test_should_return_value_from_usecase(t *testing.T) {
	usecase := &getValueUsecaseMock{Return: dbusecase.GetValueResponse{Value: "World"}}

	var result *dbgrpc.GetValueResponse
	_ = dbserverTest.RunWithGetValueUsecase(usecase.New(), func(server dbserver.Interface) error {
		return dbserverTest.ConnectWithPort(server.Port(), func(client dbgrpc.DbServiceClient) error {
			return tests.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				result, _ = client.GetValue(ctx, &dbgrpc.GetValueRequest{Key: "Hello"})
				return nil
			})
		})
	})

	assert.Equal(t, goutil.Pointer("World"), result.Value)
}
