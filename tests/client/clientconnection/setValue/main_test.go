package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/client/clientconnection"
	clientconnectiontest2 "go-skv/tests/client/clientconnection/clientconnectiontest"
	"go-skv/util/goutil"
	"net"
	"testing"
)

func Test_should_call_set_value(t *testing.T) {
	service := &clientconnectiontest2.DbServiceServerMock{}

	clientconnectiontest2.RunServerWithService(service, func(addr net.Addr) {
		conn, _ := clientconnection.New(addr.String())
		defer goutil.WillPanicUnhandledError(conn.Close)()

		err := conn.SetValue(context.Background(), "aaa", "bbb")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "aaa", service.SetValue_Request.Key)
	assert.Equal(t, "bbb", service.SetValue_Request.Value)
}
