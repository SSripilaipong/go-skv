package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbserver"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/tests/server/dbserver/dbservertest"
	"testing"
)

func Test_should_pass_context_from_controller(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usecase := &dbservertest.UsecaseMock{}
	ctrl := dbserver.NewController(usecase)

	_, _ = ctrl.SetValue(ctx, &dbgrpc.SetValueRequest{})

	assert.Equal(t, ctx, usecase.SetValue_Context)
}
