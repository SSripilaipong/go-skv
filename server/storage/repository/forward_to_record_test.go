package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	. "go-skv/server/storage/repository/message"
	"testing"
)

func TestForwardToRecord_should_reply_with_record_not_found_message_containing_memo(t *testing.T) {
	test.ContextScope(func(ctx context.Context) {
		handle := forwardToRecord(ctx, make(map[string]chan<- any))

		replyChan := make(chan any)
		go handle(ForwardToRecord{
			Key:     "xxx",
			Message: nil,
			Memo:    "Yeet",
			ReplyTo: replyChan,
		})
		reply, _ := receive(replyChan)

		assert.Equal(t, RecordNotFound{
			Key:  "xxx",
			Memo: "Yeet",
		}, reply)
	})
}

func TestForwardToRecord_should_reply_with_record_not_found_if_existing_keys_dont_match_the_requested_key(t *testing.T) {
	existingRecords := map[string]chan<- any{"": make(chan<- any)}

	test.ContextScope(func(ctx context.Context) {
		handle := forwardToRecord(ctx, existingRecords)

		replyChan := make(chan any)
		go handle(ForwardToRecord{
			Key:     "yyy",
			Message: nil,
			Memo:    "HeHe",
			ReplyTo: replyChan,
		})
		reply, _ := receive(replyChan)

		assert.Equal(t, RecordNotFound{
			Key:  "yyy",
			Memo: "HeHe",
		}, reply)
	})
}
