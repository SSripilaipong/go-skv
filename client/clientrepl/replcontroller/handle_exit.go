package replcontroller

import "go-skv/util/goutil"

func (c *Controller) handleExitCommand([]string) (string, error) {
	goutil.PanicUnhandledError(c.connection.Close())
	return "", ReplClosedError{}
}
