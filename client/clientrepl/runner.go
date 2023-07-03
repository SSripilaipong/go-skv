package clientrepl

import (
	"go-skv/client/clientconnection"
	"go-skv/client/clientrepl/replcontroller"
	"go-skv/client/clientrepl/replloop"
	"go-skv/common/util/goutil"
)

func NewReplRunner(connectionFactory clientconnection.ConnectionFactory) func(string) error {
	ctrl := replcontroller.New(connectionFactory)

	return func(serverIp string) error {
		goutil.PanicUnhandledError(ctrl.Connect(serverIp))
		return replloop.StartLoop(ctrl)
	}
}
