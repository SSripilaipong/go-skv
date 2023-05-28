package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/client"
	"go-skv/tests/client/clienttest"
	"go-skv/util/goutil"
	"net"
	"testing"
)

func Test_should_pass_context(t *testing.T) {
	service := &clienttest.DbServiceServerMock{}
	cancelledContext, cancel := context.WithCancel(context.Background())
	cancel()

	var err error
	clienttest.RunServerWithService(service, func(addr net.Addr) {
		conn := client.NewConnection(addr.String())
		defer goutil.WillPanicUnhandledError(conn.Close)()

		_, err = conn.GetValue(cancelledContext, "")
	})

	assert.Equal(t, client.CancelledError{}, err)
}
