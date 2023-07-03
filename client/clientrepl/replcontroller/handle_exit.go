package replcontroller

import (
	"go-skv/common/util/goutil"
)

func (c *controller) handleExitCommand([]string) (string, error) {
	goutil.PanicUnhandledError(c.connection.Close())
	return "", ReplClosedError{}
}
