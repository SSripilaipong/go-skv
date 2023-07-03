package replcontroller

import (
	"context"
	"go-skv/common/util/goutil"
)

func (c *controller) handleSetValueCommand(params []string) (string, error) {
	key, err := stringTokenAt(params, 0)
	goutil.PanicUnhandledError(err)

	value, err := stringTokenAt(params, 1)
	goutil.PanicUnhandledError(err)

	err = c.connection.SetValue(context.Background(), key, value)
	goutil.PanicUnhandledError(err)

	return "", nil
}
