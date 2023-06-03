package clientrepltest

import (
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

func DoInputWithText(ctrl *clientrepl.Controller, text string) error {
	return ctrl.Input(`GetValue "abc"` + "\n")
}
