package storagerecord

func NewFactory(channelBufferSize int) Factory {
	return recordFactory{
		chBufferSize: channelBufferSize,
	}
}
