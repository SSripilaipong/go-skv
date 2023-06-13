package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbserver"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
	"go-skv/tests"
	"go-skv/tests/server/dbserver/dbservertest"
	"go-skv/util/goutil"
	"testing"
	"time"
)

func Test_should_call_set_value_usecase(t *testing.T) {
	usecase := &dbservertest.UsecaseMock{}

	err := dbservertest.RunWithSetValueUsecase(usecase, func(server dbserver.Interface) error {
		return dbservertest.ConnectWithPort(server.Port(), func(client dbgrpc.DbServiceClient) error {
			return tests.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				_, err := client.SetValue(ctx, &dbgrpc.SetValueRequest{Key: "Hello", Value: "World"})
				return err
			})
		})
	})
	goutil.PanicUnhandledError(err)

	assert.Equal(t, dbusecase.SetValueRequest{Key: "Hello", Value: "World"}, usecase.SetValue_Request)
}

func Test_should_return_nonempty_response(t *testing.T) {
	var result *dbgrpc.SetValueResponse
	err := dbservertest.RunWithSetValueUsecase(&dbservertest.UsecaseMock{}, func(server dbserver.Interface) error {
		return dbservertest.ConnectWithPort(server.Port(), func(client dbgrpc.DbServiceClient) error {
			return tests.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				var err error
				result, err = client.SetValue(ctx, &dbgrpc.SetValueRequest{Key: "Hello"})
				return err
			})
		})
	})
	goutil.PanicUnhandledError(err)

	assert.NotNil(t, result)
}
