package getValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/client/clientconnection"
	"go-skv/server/dbserver/dbgrpc"
	clientconnectiontest2 "go-skv/tests/client/clientconnection/clientconnectiontest"
	"go-skv/util/goutil"
	"net"
	"testing"
)

func Test_should_call_get_value(t *testing.T) {
	service := &clientconnectiontest2.DbServiceServerMock{}

	clientconnectiontest2.RunServerWithService(service, func(addr net.Addr) {
		conn, _ := clientconnection.New(addr.String())
		defer goutil.WillPanicUnhandledError(conn.Close)()

		_, err := conn.GetValue(context.Background(), "kkk")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "kkk", service.GetValue_Request.Key)
}

func Test_should_return_value(t *testing.T) {
	service := &clientconnectiontest2.DbServiceServerMock{GetValue_Return: &dbgrpc.GetValueResponse{
		Value: goutil.Pointer("vvv"),
	}}

	var value string
	clientconnectiontest2.RunServerWithService(service, func(addr net.Addr) {
		conn, _ := clientconnection.New(addr.String())
		defer goutil.WillPanicUnhandledError(conn.Close)()

		var err error
		value, err = conn.GetValue(context.Background(), "")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "vvv", value)
}
