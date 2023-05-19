package dbmanager

type PeerServerMock struct {
	Start_IsCalled bool
}

func (p *PeerServerMock) Start() error {
	p.Start_IsCalled = true
	return nil
}
