package setValue

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

func Test_should_call_set_value_usecase(t *testing.T) {
	usecase := &dbservercontrollertest.UsecaseMock{}

	goutil.PanicUnhandledError(dbservercontrollertest.RunWithSetValueUsecase(usecase, func(controller dbservercontroller.Interface) error {
		return dbservercontrollertest.ConnectWithPort(controller.Port(), func(client dbgrpc.DbServiceClient) error {
			return test.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				_, err := client.SetValue(ctx, &dbgrpc.SetValueRequest{Key: "Hello", Value: "World"})
				return err
			})
		})
	}))

	assert.Equal(t, dbusecase.SetValueRequest{Key: "Hello", Value: "World"}, usecase.SetValue_Request)
}

func Test_should_return_nonempty_response(t *testing.T) {
	var result *dbgrpc.SetValueResponse
	goutil.PanicUnhandledError(dbservercontrollertest.RunWithSetValueUsecase(&dbservercontrollertest.UsecaseMock{}, func(controller dbservercontroller.Interface) error {
		return dbservercontrollertest.ConnectWithPort(controller.Port(), func(client dbgrpc.DbServiceClient) error {
			return test.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				var err error
				result, err = client.SetValue(ctx, &dbgrpc.SetValueRequest{Key: "Hello"})
				return err
			})
		})
	}))

	assert.NotNil(t, result)
}

func Test_should_pass_context(t *testing.T) {
	usecase := &dbservercontrollertest.UsecaseMock{}

	goutil.PanicUnhandledError(dbservercontrollertest.RunWithSetValueUsecase(usecase, func(controller dbservercontroller.Interface) error {
		return dbservercontrollertest.ConnectWithPort(controller.Port(), func(client dbgrpc.DbServiceClient) error {
			return test.ExecuteWithTimeout(time.Second, func(ctx context.Context) error {
				_, err := client.SetValue(ctx, &dbgrpc.SetValueRequest{})
				return err
			})
		})
	}))

	_, isContextDone := goutil.ReceiveNoBlock(usecase.SetValue_Context.Done())
	assert.True(t, isContextDone)
}
