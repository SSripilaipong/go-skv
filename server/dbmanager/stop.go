package dbmanager

import (
	"go-skv/common/util/goutil"
)

func (m manager) Stop() error {
	m.cancelCtx()

	goutil.PanicUnhandledError(m.dbServer.Stop())
	goutil.PanicUnhandledError(m.dbStorage.Join())
	goutil.PanicUnhandledError(m.peerConnector.Join())
	return nil
}
