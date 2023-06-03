package replcontroller

func (c *Controller) generateCommandMapper() {
	c.commandMapper = map[string]func([]string) (string, error){
		"getvalue": c.handleGetValueCommand,
		"setvalue": c.handleSetValueCommand,
		"exit":     c.handleExitCommand,
	}
}
