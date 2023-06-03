package input_setvalue

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/client/clientrepl/clientrepltest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_call_set_value_with_key_and_value(t *testing.T) {
	connection := &clientrepltest.ConnectionMock{}
	ctrl := clientrepltest.NewControllerWithConnectionFactory((&clientrepltest.ConnectionFactoryMock{Return: connection}).New())
	goutil.PanicUnhandledError(clientrepltest.DoConnect(ctrl))

	_, err := clientrepltest.DoSetValueInputWithKeyAndValue(ctrl, "Go", "Lang")
	goutil.PanicUnhandledError(err)

	assert.Equal(t, "Go", connection.SetValue_key)
	assert.Equal(t, "Lang", connection.SetValue_value)
}
