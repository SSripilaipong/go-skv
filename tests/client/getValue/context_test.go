package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/client/clientconnection"
	"go-skv/tests/client/clienttest"
	"go-skv/util/goutil"
	"net"
	"testing"
)

func Test_should_return_error_when_context_is_cancelled(t *testing.T) {
	service := &clienttest.DbServiceServerMock{}
	cancelledContext, cancel := context.WithCancel(context.Background())
	cancel()

	var err error
	clienttest.RunServerWithService(service, func(addr net.Addr) {
		conn, _ := clientconnection.New(addr.String())
		defer goutil.WillPanicUnhandledError(conn.Close)()

		_, err = conn.GetValue(cancelledContext, "")
	})

	assert.Equal(t, clientconnection.CancelledError{}, err)
}
