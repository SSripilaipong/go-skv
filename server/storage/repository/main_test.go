package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	"go-skv/common/util/goutil"
	storageMessage "go-skv/server/storage/message"
	"testing"
	"time"
)

func Test_should_notify_termination(t *testing.T) {
	var isNotified bool
	done := make(chan struct{})

	test.ContextScope(func(ctx context.Context) {
		repo := newRepository(ctx, 1)

		send(repo, storageMessage.Terminate{Notify: done})

		_, isNotified = receive(done)
	})

	assert.True(t, isNotified)
}

func Test_should_acknowledge_save_with_memo(t *testing.T) {
	test.ContextScope(func(ctx context.Context) {
		repo := newRepository(ctx, 1)

		ch := make(chan any)
		send(repo, storageMessage.SaveRecord{
			Key:     "abc",
			Channel: nil,
			Memo:    "myMemo",
			ReplyTo: ch,
		})

		reply, _ := receive(ch)
		assert.Equal(t, storageMessage.Ack{Memo: "myMemo"}, reply)

		send(repo, storageMessage.Terminate{Notify: make(chan struct{})})
	})
}

var defaultTimeout = 100 * time.Millisecond

func send(ch chan<- any, msg any) {
	goutil.SendWithTimeout[any](ch, msg, defaultTimeout)
}

func receive[T any](ch <-chan T) (T, bool) {
	return goutil.ReceiveWithTimeout[T](ch, defaultTimeout)
}
