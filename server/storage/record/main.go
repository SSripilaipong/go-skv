package record

func NewFactory(channelBufferSize int) Factory {
	return factory{channelBufferSize: channelBufferSize}
}
