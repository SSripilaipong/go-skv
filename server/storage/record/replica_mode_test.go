package record

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	. "go-skv/server/storage/record/message"
	"testing"
)

func TestReplicaMode_should_reply_value_when_request_with_replica_get_value(t *testing.T) {
	factory := NewFactory(1)

	test.ContextScope(func(ctx context.Context) {
		record := factory.NewReplica(ctx, "HelloReplica")

		replyChan := make(chan any)
		send(record, GetValue{Memo: "Memo Nemo", ReplyTo: replyChan})
		reply, _ := receive(replyChan)

		assert.Equal(t, Value{Value: "HelloReplica", Memo: "Memo Nemo"}, reply)
	})
}
