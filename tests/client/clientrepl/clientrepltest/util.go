package clientrepltest

import (
	"fmt"
	"go-skv/client/clientconnection"
	"go-skv/client/clientrepl"
)

func NewControllerWithConnectionFactory(connectionFactory clientconnection.ConnectionFactory) *clientrepl.Controller {
	return clientrepl.NewController(connectionFactory)
}

func DoConnect(ctrl *clientrepl.Controller) error {
	return DoConnectWithAddress(ctrl, "")
}

func DoConnectWithAddress(ctrl *clientrepl.Controller, address string) error {
	return ctrl.Connect(address)
}

func DoInputWithText(ctrl *clientrepl.Controller, text string) (string, error) {
	return ctrl.Input(text)
}

func DoGetValueInputWithKey(ctrl *clientrepl.Controller, key string) (string, error) {
	return DoInputWithText(ctrl, fmt.Sprintf(`getvalue "%s"`+"\n", key))
}
