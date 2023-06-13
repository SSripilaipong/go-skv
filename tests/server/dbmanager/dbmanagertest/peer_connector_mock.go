package dbmanagertest

type PeerConnectorMock struct {
	Start_IsCalled bool
}

func (p *PeerConnectorMock) Start() error {
	p.Start_IsCalled = true
	return nil
}
