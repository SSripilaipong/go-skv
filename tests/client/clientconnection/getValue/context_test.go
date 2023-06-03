package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/client/clientconnection"
	clientconnectiontest2 "go-skv/tests/client/clientconnection/clientconnectiontest"
	"go-skv/util/goutil"
	"net"
	"testing"
)

func Test_should_return_error_when_context_is_cancelled(t *testing.T) {
	service := &clientconnectiontest2.DbServiceServerMock{}
	cancelledContext, cancel := context.WithCancel(context.Background())
	cancel()

	var err error
	clientconnectiontest2.RunServerWithService(service, func(addr net.Addr) {
		conn, _ := clientconnection.New(addr.String())
		defer goutil.WillPanicUnhandledError(conn.Close)()

		_, err = conn.GetValue(cancelledContext, "")
	})

	assert.Equal(t, clientconnection.CancelledError{}, err)
}
