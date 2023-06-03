package input_setvalue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/client/clientrepl/clientrepltest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_close_connection(t *testing.T) {
	connection := &clientrepltest.ConnectionMock{}
	ctrl := clientrepltest.NewControllerWithConnectionFactory((&clientrepltest.ConnectionFactoryMock{Return: connection}).New())
	goutil.PanicUnhandledError(clientrepltest.DoConnect(ctrl))

	_, _ = clientrepltest.DoExit(ctrl)

	assert.True(t, connection.Close_IsCalled)
}
