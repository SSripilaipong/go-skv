package input_setvalue

import (
	"github.com/stretchr/testify/assert"
	replcontrollertest2 "go-skv/tests/client/clientrepl/replcontroller/replcontrollertest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_call_set_value_with_key_and_value(t *testing.T) {
	connection := &replcontrollertest2.ConnectionMock{}
	ctrl := replcontrollertest2.NewControllerWithConnectionFactory((&replcontrollertest2.ConnectionFactoryMock{Return: connection}).New())
	goutil.PanicUnhandledError(replcontrollertest2.DoConnect(ctrl))

	_, err := replcontrollertest2.DoSetValueInputWithKeyAndValue(ctrl, "Go", "Lang")
	goutil.PanicUnhandledError(err)

	assert.Equal(t, "Go", connection.SetValue_key)
	assert.Equal(t, "Lang", connection.SetValue_value)
}
