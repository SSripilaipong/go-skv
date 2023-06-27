package connectormanager

func (m manager) Stop() error {
	m.cancelSubCtx()
	return nil
}
