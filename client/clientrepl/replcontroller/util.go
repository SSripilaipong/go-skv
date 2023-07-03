package replcontroller

import (
	"go-skv/common/util/goutil"
	"strings"
)

func stringTokenAt(tokens []string, index int) (string, error) {
	valueWithQuotes, err := goutil.ElementAt(tokens, index)
	if err != nil {
		return "", err
	}

	return strings.Trim(valueWithQuotes, "\""), nil
}
