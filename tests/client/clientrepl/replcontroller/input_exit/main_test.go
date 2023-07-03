package input_exit

import (
	"github.com/stretchr/testify/assert"
	"go-skv/client/clientrepl/replcontroller"
	"go-skv/common/util/goutil"
	replcontrollertest2 "go-skv/tests/client/clientrepl/replcontroller/replcontrollertest"
	"testing"
)

func Test_should_close_connection(t *testing.T) {
	connection := &replcontrollertest2.ConnectionMock{}
	ctrl := replcontrollertest2.NewControllerWithConnectionFactory((&replcontrollertest2.ConnectionFactoryMock{Return: connection}).New())
	goutil.PanicUnhandledError(replcontrollertest2.DoConnect(ctrl))

	_, _ = replcontrollertest2.DoExit(ctrl)

	assert.True(t, connection.Close_IsCalled)
}

func Test_should_return_repl_closed_error(t *testing.T) {
	ctrl := replcontrollertest2.NewController()
	goutil.PanicUnhandledError(replcontrollertest2.DoConnect(ctrl))

	_, err := replcontrollertest2.DoExit(ctrl)

	assert.Equal(t, replcontroller.ReplClosedError{}, err)
}
