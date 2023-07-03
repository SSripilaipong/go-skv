package storagerecord

func (r recordInteractor) Destroy() error {
	r.ctxCancel()
	<-r.stopped
	return nil
}
