package storagerepository

type manager struct {
	ch      chan command
	stopped chan struct{}
}

func (m manager) Join() error {
	<-m.stopped
	return nil
}
