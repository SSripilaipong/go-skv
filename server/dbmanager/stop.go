package dbmanager

import (
	"go-skv/util/goutil"
)

func (m manager) Stop() error {
	goutil.PanicUnhandledError(m.dbServer.Stop())
	goutil.PanicUnhandledError(m.dbStorage.Stop())
	goutil.PanicUnhandledError(m.peerConnector.Join())
	return nil
}
