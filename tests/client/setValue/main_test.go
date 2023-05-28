package setValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/client"
	"go-skv/tests/client/clienttest"
	"go-skv/util/goutil"
	"net"
	"testing"
)

func Test_should_call_set_value(t *testing.T) {
	service := &clienttest.DbServiceServerMock{}

	clienttest.RunServerWithService(service, func(addr net.Addr) {
		conn := client.NewConnection(addr.String())
		defer goutil.WillPanicUnhandledError(conn.Close)()

		err := conn.SetValue(context.Background(), "aaa", "bbb")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "aaa", service.SetValue_Request.Key)
	assert.Equal(t, "bbb", service.SetValue_Request.Value)
}
