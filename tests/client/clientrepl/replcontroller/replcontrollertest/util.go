package replcontrollertest

import (
	"fmt"
	"go-skv/client/clientconnection"
	"go-skv/client/clientrepl/replcontroller"
)

func NewController() *replcontroller.Controller {
	factory := &ConnectionFactoryMock{Return: &ConnectionMock{}}
	return NewControllerWithConnectionFactory(factory.New())
}

func NewControllerWithConnectionFactory(connectionFactory clientconnection.ConnectionFactory) *replcontroller.Controller {
	return replcontroller.NewController(connectionFactory)
}

func DoConnect(ctrl *replcontroller.Controller) error {
	return DoConnectWithAddress(ctrl, "")
}

func DoConnectWithAddress(ctrl *replcontroller.Controller, address string) error {
	return ctrl.Connect(address)
}

func DoInputWithText(ctrl *replcontroller.Controller, text string) (string, error) {
	return ctrl.Input(text)
}

func DoGetValueInputWithKey(ctrl *replcontroller.Controller, key string) (string, error) {
	return DoInputWithText(ctrl, fmt.Sprintf(`getvalue "%s"`+"\n", key))
}

func DoSetValueInputWithKeyAndValue(ctrl *replcontroller.Controller, key string, value string) (string, error) {
	return DoInputWithText(ctrl, fmt.Sprintf(`setvalue "%s" "%s"`+"\n", key, value))
}

func DoExit(ctrl *replcontroller.Controller) (string, error) {
	return DoInputWithText(ctrl, "exit\n")
}
