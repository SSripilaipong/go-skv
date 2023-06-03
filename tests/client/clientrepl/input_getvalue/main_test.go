package input_getvalue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/client/clientrepl/clientrepltest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_call_get_value_with_key(t *testing.T) {
	connection := &clientrepltest.ConnectionMock{}
	ctrl := clientrepltest.NewControllerWithConnectionFactory((&clientrepltest.ConnectionFactoryMock{Return: connection}).New())
	goutil.PanicUnhandledError(clientrepltest.DoConnect(ctrl))

	goutil.PanicUnhandledError(clientrepltest.DoInputWithText(ctrl, `getvalue "abc"`+"\n"))

	assert.Equal(t, "abc", connection.GetValue_key)
}
