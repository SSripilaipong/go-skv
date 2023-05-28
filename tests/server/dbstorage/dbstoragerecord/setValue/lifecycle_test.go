package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage"
	"go-skv/tests/server/dbstorage/dbstoragerecord/dbstoragerecordtest"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"testing"
	"time"
)

func Test_should_return_error_when_context_is_cancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	factory := dbstoragerecordtest.NewFactoryWithContext(ctx)
	record := factory.New()

	cancel()
	time.Sleep(time.Millisecond)
	err := record.SetValue(&dbstoragetest.SetValueMessage{})

	assert.Equal(t, dbstorage.RecordDestroyedError{}, err)
}
