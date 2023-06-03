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

	_, err := clientrepltest.DoGetValueInputWithKey(ctrl, "abc")
	goutil.PanicUnhandledError(err)

	assert.Equal(t, "abc", connection.GetValue_key)
}

func Test_should_return_display_text(t *testing.T) {
	connection := &clientrepltest.ConnectionMock{GetValue_Value: "World"}
	ctrl := clientrepltest.NewControllerWithConnectionFactory((&clientrepltest.ConnectionFactoryMock{Return: connection}).New())
	goutil.PanicUnhandledError(clientrepltest.DoConnect(ctrl))

	display, err := clientrepltest.DoGetValueInputWithKey(ctrl, "Hello")
	goutil.PanicUnhandledError(err)

	assert.Equal(t, `"Hello" => "World"`+"\n", display)
}
