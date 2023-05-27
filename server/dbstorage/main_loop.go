package dbstorage

func (s *storage) mainLoop() {
	for {
		select {
		case raw := <-s.ch:
			if message, isSetMessage := raw.(SetValueMessage); isSetMessage {
				s.handleSetValueMessage(message)
			} else if message, isGetMessage := raw.(GetValueMessage); isGetMessage {
				s.handleGetValueMessage(message)
			}
		case <-s.ctx.Done():
			goto stop
		}
	}
stop:
	s.stopped <- struct{}{}
}