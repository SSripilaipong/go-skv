package repositoryroutine

func (m *manager) mainLoop() {
	for {
		select {
		case raw := <-m.ch:
			m.handleMessage(raw)
		case <-m.ctx.Done():
			goto stop
		}
	}
stop:
	m.stopped <- struct{}{}
}
