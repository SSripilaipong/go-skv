package input_getvalue

import (
	"github.com/stretchr/testify/assert"
	replcontrollertest2 "go-skv/tests/client/clientrepl/replcontroller/replcontrollertest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_call_get_value_with_key(t *testing.T) {
	connection := &replcontrollertest2.ConnectionMock{}
	ctrl := replcontrollertest2.NewControllerWithConnectionFactory((&replcontrollertest2.ConnectionFactoryMock{Return: connection}).New())
	goutil.PanicUnhandledError(replcontrollertest2.DoConnect(ctrl))

	_, err := replcontrollertest2.DoGetValueInputWithKey(ctrl, "abc")
	goutil.PanicUnhandledError(err)

	assert.Equal(t, "abc", connection.GetValue_key)
}

func Test_should_return_display_text(t *testing.T) {
	connection := &replcontrollertest2.ConnectionMock{GetValue_Value: "World"}
	ctrl := replcontrollertest2.NewControllerWithConnectionFactory((&replcontrollertest2.ConnectionFactoryMock{Return: connection}).New())
	goutil.PanicUnhandledError(replcontrollertest2.DoConnect(ctrl))

	display, err := replcontrollertest2.DoGetValueInputWithKey(ctrl, "Hello")
	goutil.PanicUnhandledError(err)

	assert.Equal(t, `"Hello" => "World"`+"\n", display)
}
