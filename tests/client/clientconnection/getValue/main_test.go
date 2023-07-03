package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/client/clientconnection"
	goutil2 "go-skv/common/util/goutil"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/tests/client/clientconnection/clientconnectiontest"
	"net"
	"testing"
)

func Test_should_call_get_value(t *testing.T) {
	service := &clientconnectiontest.DbServiceServerMock{}

	clientconnectiontest.RunServerWithService(service, func(addr net.Addr) {
		conn, _ := clientconnection.New(addr.String())
		defer goutil2.WillPanicUnhandledError(conn.Close)()

		_, err := conn.GetValue(context.Background(), "kkk")
		goutil2.PanicUnhandledError(err)
	})

	assert.Equal(t, "kkk", service.GetValue_Request.Key)
}

func Test_should_return_value(t *testing.T) {
	service := &clientconnectiontest.DbServiceServerMock{GetValue_Return: &dbgrpc.GetValueResponse{
		Value: goutil2.Pointer("vvv"),
	}}

	var value string
	clientconnectiontest.RunServerWithService(service, func(addr net.Addr) {
		conn, _ := clientconnection.New(addr.String())
		defer goutil2.WillPanicUnhandledError(conn.Close)()

		var err error
		value, err = conn.GetValue(context.Background(), "")
		goutil2.PanicUnhandledError(err)
	})

	assert.Equal(t, "vvv", value)
}
