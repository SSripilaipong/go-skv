package update

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/replicaupdater"
	"go-skv/server/replicaupdater/replicaupdatercontract"
	"go-skv/tests"
	"go-skv/tests/server/servertest"
	"testing"
)

func Test_should_get_record_from_repo_in_the_first_time(t *testing.T) {
	storage := &servertest.DbStorageMock{}
	factory := replicaupdater.NewFactory(storage)

	var updater replicaupdatercontract.InboundUpdater
	tests.ContextScope(func(ctx context.Context) {
		updater, _ = factory.NewInboundUpdater(ctx)

		storage.GetRecord_WaitUntilCalledOnce(defaultTimeout, func() {
			_ = updater.Update("aaa", "")
		})
	})
	updater.Join()

	assert.Equal(t, "aaa", storage.GetRecord_key)
}
