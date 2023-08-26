package record

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	"go-skv/common/util/goutil"
	. "go-skv/server/storage/record/message"
	"testing"
	"time"
)

func Test_should_reply_value_when_request_with_get_value(t *testing.T) {
	factory := NewFactory(1)

	test.ContextScope(func(ctx context.Context) {
		record := factory.New(ctx, "Hello")

		replyChan := make(chan any)
		send(record, GetValue{Memo: "MMM", ReplyTo: replyChan})
		reply, _ := receive(replyChan)

		assert.Equal(t, Value{Value: "Hello", Memo: "MMM"}, reply)
	})
}

func Test_should_ack_when_set_value(t *testing.T) {
	factory := NewFactory(1)

	test.ContextScope(func(ctx context.Context) {
		record := factory.New(ctx, "Hello")

		replyChan := make(chan any)
		send(record, SetValue{Value: "World", Memo: "YeeHee", ReplyTo: replyChan})
		reply, _ := receive(replyChan)

		assert.Equal(t, Ack{Memo: "YeeHee"}, reply)
	})
}

var defaultTimeout = 100 * time.Millisecond

func send(ch chan<- any, msg any) {
	goutil.SendWithTimeout[any](ch, msg, defaultTimeout)
}

func receive[T any](ch <-chan T) (T, bool) {
	return goutil.ReceiveWithTimeout[T](ch, defaultTimeout)
}
