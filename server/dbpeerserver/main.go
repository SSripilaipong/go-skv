package dbpeerserver

func New() Interface {
	return &peerServer{}
}

type peerServer struct{}

func (p *peerServer) Start() error {
	return nil
}
