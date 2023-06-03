package connect

import (
	"github.com/stretchr/testify/assert"
	"go-skv/client/clientrepl"
	"go-skv/tests/client/clientrepl/clientrepltest"
	"testing"
)

func Test_should_create_connection_with_address(t *testing.T) {
	connectionFactory := &clientrepltest.ConnectionFactoryMock{}
	ctrl := clientrepl.NewController(connectionFactory.New())

	address := "127.0.0.1:1234"
	_ = ctrl.Connect(address)

	assert.Equal(t, address, connectionFactory.Address)
}
