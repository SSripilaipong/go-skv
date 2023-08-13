package cli

type Config struct {
	DbPort                int
	PeeringPort           int
	AdvertisedIp          string
	ExistingPeerAddresses []string
}
