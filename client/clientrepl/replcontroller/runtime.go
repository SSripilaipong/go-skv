package replcontroller

import "fmt"

func RunRuntimeRepl(serverIp string) error {
	var s string
	for {
		fmt.Printf(">>> ")
		_, err := fmt.Scanf("%s\n", &s)
		if err != nil {
			return err
		}
		fmt.Printf("%#v\n", s)

		if s == "exit" {
			break
		}
	}
	return nil
}
