package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	"go-skv/server/storage/message"
	"testing"
)

func TestForwardToRecord_should_reply_with_record_not_found_message_containing_memo(t *testing.T) {
	test.ContextScope(func(ctx context.Context) {
		handle := forwardToRecord(ctx, make(map[string]chan<- any))

		replyChan := make(chan any)
		go handle(message.ForwardToRecord{
			Key:     "xxx",
			Message: nil,
			Memo:    "Yeet",
			ReplyTo: replyChan,
		})
		reply, _ := receive(replyChan)

		assert.Equal(t, message.RecordNotFound{
			Key:  "xxx",
			Memo: "Yeet",
		}, reply)
	})
}
