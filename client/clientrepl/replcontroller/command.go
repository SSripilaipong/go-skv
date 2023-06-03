package replcontroller

func (c *controller) generateCommandMapper() {
	c.commandMapper = map[string]func([]string) (string, error){
		"getvalue": c.handleGetValueCommand,
		"setvalue": c.handleSetValueCommand,
		"exit":     c.handleExitCommand,
	}
}
