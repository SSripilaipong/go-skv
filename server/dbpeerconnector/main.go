package dbpeerconnector

func New() Interface {
	return &connector{}
}

type connector struct{}

func (p *connector) Start() error {
	return nil
}

func (p *connector) Stop() error {
	return nil
}

var _ Interface = &connector{}
