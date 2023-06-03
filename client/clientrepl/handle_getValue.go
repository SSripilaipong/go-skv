package clientrepl

import (
	"context"
	"fmt"
	"go-skv/util/goutil"
)

func (c *Controller) handleGetValueCommand(params []string) (string, error) {
	key, err := stringTokenAt(params, 0)
	goutil.PanicUnhandledError(err)

	value, err := c.connection.GetValue(context.Background(), key)
	goutil.PanicUnhandledError(err)

	return fmt.Sprintf("%#v => %#v\n", key, value), nil
}
