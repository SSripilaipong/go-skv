package input_exit

import (
	"github.com/stretchr/testify/assert"
	"go-skv/client/clientrepl"
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

func Test_should_return_repl_closed_error(t *testing.T) {
	ctrl := clientrepltest.NewController()
	goutil.PanicUnhandledError(clientrepltest.DoConnect(ctrl))

	_, err := clientrepltest.DoExit(ctrl)

	assert.Equal(t, clientrepl.ReplClosedError{}, err)
}
