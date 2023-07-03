package connect

import (
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	replcontrollertest2 "go-skv/tests/client/clientrepl/replcontroller/replcontrollertest"
	"testing"
)

func Test_should_create_connection_with_address(t *testing.T) {
	connectionFactory := &replcontrollertest2.ConnectionFactoryMock{}
	ctrl := replcontrollertest2.NewControllerWithConnectionFactory(connectionFactory.New())

	address := "127.0.0.1:1234"
	goutil.PanicUnhandledError(replcontrollertest2.DoConnectWithAddress(ctrl, address))

	assert.Equal(t, address, connectionFactory.Address)
}
