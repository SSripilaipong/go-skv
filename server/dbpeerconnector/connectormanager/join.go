package connectormanager

import (
	"go-skv/common/util/goutil"
)

func (m manager) Join() error {
	goutil.PanicUnhandledError(m.client.WaitForAllToBeDisconnected())
	return nil
}
