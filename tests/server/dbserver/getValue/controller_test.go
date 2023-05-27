package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbserver"
	"go-skv/server/dbserver/dbgrpc"
	"testing"
)

func Test_should_pass_context_from_controller(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usecase := &getValueUsecaseMock{}
	ctrl := dbserver.NewController(dbserver.Dependency{GetValueUsecase: usecase.New()})

	_, _ = ctrl.GetValue(ctx, &dbgrpc.GetValueRequest{})

	assert.Equal(t, ctx, usecase.Context)
}
