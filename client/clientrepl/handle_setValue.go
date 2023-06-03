package clientrepl

import (
	"context"
	"go-skv/util/goutil"
)

func (c *Controller) handleSetValueCommand(params []string) (string, error) {
	key, err := stringTokenAt(params, 0)
	goutil.PanicUnhandledError(err)

	value, err := stringTokenAt(params, 1)
	goutil.PanicUnhandledError(err)

	err = c.connection.SetValue(context.Background(), key, value)
	goutil.PanicUnhandledError(err)

	return "", nil
}
