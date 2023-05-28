package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/client"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/tests/client/clienttest"
	"go-skv/util/goutil"
	"net"
	"testing"
)

func Test_should_call_get_value(t *testing.T) {
	service := &clienttest.DbServiceServerMock{}

	clienttest.RunServerWithService(service, func(addr net.Addr) {
		conn := client.NewConnection(addr.String())
		defer goutil.WillPanicUnhandledError(conn.Close)()

		_, err := conn.GetValue(context.Background(), "kkk")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "kkk", service.GetValue_Request.Key)
}

func Test_should_return_value(t *testing.T) {
	service := &clienttest.DbServiceServerMock{GetValue_Return: &dbgrpc.GetValueResponse{
		Value: goutil.Pointer("vvv"),
	}}

	var value string
	clienttest.RunServerWithService(service, func(addr net.Addr) {
		conn := client.NewConnection(addr.String())
		defer goutil.WillPanicUnhandledError(conn.Close)()

		var err error
		value, err = conn.GetValue(context.Background(), "")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "vvv", value)
}
