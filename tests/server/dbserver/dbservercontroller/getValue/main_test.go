package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	"go-skv/common/util/goutil"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbserver/dbservercontroller"
	"go-skv/server/dbserver/dbusecase"
	"go-skv/tests/server/dbserver/dbservercontroller/dbservercontrollertest"
	"testing"
	"time"
)

func Test_should_call_get_value_usecase(t *testing.T) {
	usecase := &dbservercontrollertest.UsecaseMock{}

	goutil.PanicUnhandledError(dbservercontrollertest.RunWithGetValueUsecase(usecase, func(controller dbservercontroller.Interface) error {
		return dbservercontrollertest.ConnectWithPort(controller.Port(), func(client dbgrpc.DbServiceClient) error {
			return test.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				_, err := client.GetValue(ctx, &dbgrpc.GetValueRequest{Key: "Hello"})
				return err
			})
		})
	}))

	assert.Equal(t, dbusecase.GetValueRequest{Key: "Hello"}, usecase.GetValue_Request)
}

func Test_should_return_value_from_usecase(t *testing.T) {
	usecase := &dbservercontrollertest.UsecaseMock{GetValue_Return: dbusecase.GetValueResponse{Value: "World"}}

	var result *dbgrpc.GetValueResponse
	goutil.PanicUnhandledError(dbservercontrollertest.RunWithGetValueUsecase(usecase, func(controller dbservercontroller.Interface) error {
		return dbservercontrollertest.ConnectWithPort(controller.Port(), func(client dbgrpc.DbServiceClient) error {
			return test.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				var err error
				result, err = client.GetValue(ctx, &dbgrpc.GetValueRequest{Key: "Hello"})
				return err
			})
		})
	}))

	assert.Equal(t, goutil.Pointer("World"), result.Value)
}

func Test_should_pass_context(t *testing.T) {
	usecase := &dbservercontrollertest.UsecaseMock{}

	goutil.PanicUnhandledError(dbservercontrollertest.RunWithGetValueUsecase(usecase, func(controller dbservercontroller.Interface) error {
		return dbservercontrollertest.ConnectWithPort(controller.Port(), func(client dbgrpc.DbServiceClient) error {
			return test.ExecuteWithTimeout(100*time.Millisecond, func(ctx context.Context) error {
				_, err := client.GetValue(ctx, &dbgrpc.GetValueRequest{})
				return err
			})
		})
	}))

	_, isContextDone := goutil.ReceiveNoBlock(usecase.GetValue_Context.Done())
	assert.True(t, isContextDone)
}
