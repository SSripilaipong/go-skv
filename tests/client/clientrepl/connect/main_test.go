package connect

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/client/clientrepl/clientrepltest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_create_connection_with_address(t *testing.T) {
	connectionFactory := &clientrepltest.ConnectionFactoryMock{}
	ctrl := clientrepltest.NewControllerWithConnectionFactory(connectionFactory.New())

	address := "127.0.0.1:1234"
	goutil.PanicUnhandledError(clientrepltest.DoConnectWithAddress(ctrl, address))

	assert.Equal(t, address, connectionFactory.Address)
}
