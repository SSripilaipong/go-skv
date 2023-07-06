package clientsidepeer

func (t interactor) Join() error {
	t.wg.Wait()
	return nil
}
