package replloop

import (
	"bufio"
	"fmt"
	"go-skv/client/clientrepl/replcontroller"
	"go-skv/util/goutil"
	"os"
)

func StartLoop(controller replcontroller.Interface) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		input, err := reader.ReadString('\n')
		goutil.PanicUnhandledError(err)

		output, err := controller.Input(input)
		fmt.Print(output)

		switch err.(type) {
		case replcontroller.ReplClosedError:
			goto EndLoop
		}
	}
EndLoop:
	return nil
}
