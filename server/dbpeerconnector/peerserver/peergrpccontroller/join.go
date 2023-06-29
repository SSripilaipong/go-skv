package peergrpccontroller

func (c *controller) Join() {
	c.wg.Wait()
}
