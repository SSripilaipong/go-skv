package getValue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-skv/goutil"
	"go-skv/server/dbusecase"
	"go-skv/server/storage"
	"testing"
)

func Test_should_send_get_value_message_to_storage(t *testing.T) {
	storageChan := make(chan storage.Packet, 2)
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(storageChan))

	_, _ = execute(&dbusecase.GetValueRequest{Key: "Go"})

	packet, ok := goutil.ReceiveNoBlock(storageChan)
	if !ok {
		panic(fmt.Errorf("unexpected error"))
	}

	assert.Equal(t, storage.GetValueMessage{Key: "Go"}, packet.Message())
}
