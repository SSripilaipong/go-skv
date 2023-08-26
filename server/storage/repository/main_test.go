package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	"go-skv/common/util/goutil"
	. "go-skv/server/storage/repository/message"
	"testing"
	"time"
)

func Test_should_notify_termination(t *testing.T) {
	var isNotified bool
	done := make(chan struct{})

	test.ContextScope(func(ctx context.Context) {
		repo := newRepository(ctx, 1)

		send(repo, Terminate{Notify: done})

		_, isNotified = receive(done)
	})

	assert.True(t, isNotified)
}

func Test_should_acknowledge_save_with_memo(t *testing.T) {
	test.ContextScope(func(ctx context.Context) {
		repo := newRepository(ctx, 1)

		ch := make(chan any)
		send(repo, SaveRecord{
			Key:     "",
			Channel: nil,
			Memo:    "myMemo",
			ReplyTo: ch,
		})

		reply, _ := receive(ch)
		assert.Equal(t, Ack{Memo: "myMemo"}, reply)

		send(repo, Terminate{Notify: make(chan struct{})})
	})
}

func Test_should_forward_message_to_saved_record(t *testing.T) {
	test.ContextScope(func(ctx context.Context) {
		repo := newRepository(ctx, 1)

		recordChan := make(chan any)
		send(repo, SaveRecord{
			Key:     "abc",
			Channel: recordChan,
			Memo:    "",
			ReplyTo: make(chan<- any, 1),
		})

		send(repo, ForwardToRecord{
			Key:     "abc",
			Message: "Hello Record",
			Memo:    "",
			ReplyTo: make(chan<- any, 1),
		})

		forwardedMessage, _ := receive(recordChan)
		assert.Equal(t, "Hello Record", forwardedMessage)

		send(repo, Terminate{Notify: make(chan struct{})})
	})
}

var defaultTimeout = 100 * time.Millisecond

func send(ch chan<- any, msg any) {
	goutil.SendWithTimeout[any](ch, msg, defaultTimeout)
}

func receive[T any](ch <-chan T) (T, bool) {
	return goutil.ReceiveWithTimeout[T](ch, defaultTimeout)
}
