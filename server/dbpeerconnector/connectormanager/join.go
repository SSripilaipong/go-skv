package connectormanager

import "go-skv/util/goutil"

func (m manager) Join() error {
	goutil.PanicUnhandledError(m.client.WaitForAllToBeDisconnected())
	return nil
}
