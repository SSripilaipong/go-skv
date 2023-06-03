package replcontrollertest

import (
	"fmt"
	"go-skv/client/clientconnection"
	"go-skv/client/clientrepl/replcontroller"
)

func NewController() replcontroller.Interface {
	factory := &ConnectionFactoryMock{Return: &ConnectionMock{}}
	return NewControllerWithConnectionFactory(factory.New())
}

func NewControllerWithConnectionFactory(connectionFactory clientconnection.ConnectionFactory) replcontroller.Interface {
	return replcontroller.New(connectionFactory)
}

func DoConnect(ctrl replcontroller.Interface) error {
	return DoConnectWithAddress(ctrl, "")
}

func DoConnectWithAddress(ctrl replcontroller.Interface, address string) error {
	return ctrl.Connect(address)
}

func DoInputWithText(ctrl replcontroller.Interface, text string) (string, error) {
	return ctrl.Input(text)
}

func DoGetValueInputWithKey(ctrl replcontroller.Interface, key string) (string, error) {
	return DoInputWithText(ctrl, fmt.Sprintf(`getvalue "%s"`+"\n", key))
}

func DoSetValueInputWithKeyAndValue(ctrl replcontroller.Interface, key string, value string) (string, error) {
	return DoInputWithText(ctrl, fmt.Sprintf(`setvalue "%s" "%s"`+"\n", key, value))
}

func DoExit(ctrl replcontroller.Interface) (string, error) {
	return DoInputWithText(ctrl, "exit\n")
}
