package update

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/actormodel"
	"go-skv/server/replicaupdater"
	"go-skv/tests"
	"go-skv/tests/server/replicaupdater/replicaupdatertest"
	"testing"
)

func Test_should_create_new_record_updater(t *testing.T) {
	storage := actormodel.NewTestActor()
	recordUpdater := actormodel.NewTestActor()
	recordUpdaterFactory := &replicaupdatertest.RecordUpdaterFactoryMock{New_Return: recordUpdater}
	factory := replicaupdater.NewFactory2(storage, recordUpdaterFactory)

	tests.ContextScope(func(ctx context.Context) {
		updater, _ := factory.NewInboundUpdater(context.WithValue(ctx, "test", "same context"))

		recordUpdaterFactory.New_WaitUntilCalledOnce(defaultTimeout, func() {
			_ = actormodel.NewTestActor().FakeTellBlocking(
				tests.ContextWithTimeout(defaultTimeout),
				updater,
				replicaupdater.InboundUpdate{Key: "kkk", Value: "vvv"},
			)
		})

	})

	assert.Equal(t, "same context", recordUpdaterFactory.New_ctx.Value("test"))
	assert.Equal(t, "kkk", recordUpdaterFactory.New_key)
	assert.Equal(t, "vvv", recordUpdaterFactory.New_value)
}
