package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/repositoryroutine/repositoryroutinetest"
	"go-skv/tests/server/dbstorage/storagerecord/storagerecordtest"
	"testing"
	"time"
)

func Test_should_return_error_when_context_is_cancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	record := storagerecordtest.DoNewRecordWithContext(storagerecordtest.NewFactory(), ctx)

	cancel()
	time.Sleep(time.Millisecond)
	err := record.SetValue(&repositoryroutinetest.SetValueMessage{})

	assert.Equal(t, storagerecord.RecordDestroyedError{}, err)
}
