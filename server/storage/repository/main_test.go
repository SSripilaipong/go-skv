package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	"go-skv/common/util/goutil"
	"go-skv/server/storage/message"
	"testing"
	"time"
)

var testDefaultTimeout = 100 * time.Millisecond

func Test_should_notify_termination(t *testing.T) {
	var isNotified bool
	done := make(chan struct{})

	test.ContextScope(func(ctx context.Context) {
		repo := newRepository(ctx, 1)

		goutil.SendWithTimeout[any](repo, message.Terminate{Notify: done}, testDefaultTimeout)
		_, isNotified = goutil.ReceiveWithTimeout(done, testDefaultTimeout)
	})

	assert.True(t, isNotified)
}
