package servercli

import (
	"github.com/urfave/cli/v2"
)

func newCliApp(dep dependency) *cli.App {
	return &cli.App{
		EnableBashCompletion: true,
		Commands:             commands(dep),
	}
}

var dbPortFlagName = "db-port"
var peeringPortFlagName = "peer-port"
var advertisedIpFlagName = "advertised-ip"
var existingPeerAddresses = "peers"

func commands(dep dependency) []*cli.Command {
	return []*cli.Command{
		{
			Name: "start",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  dbPortFlagName + ",p",
					Value: 5555,
					Usage: "port for db client to connect",
				},
				&cli.IntFlag{
					Name:  peeringPortFlagName + ",q",
					Value: 5556,
					Usage: "port for other peers to join network",
				},
				&cli.StringFlag{
					Name:  advertisedIpFlagName + ",h",
					Value: "localhost",
					Usage: "ip that other peers can reach to",
				},
				&cli.StringSliceFlag{
					Name:  existingPeerAddresses,
					Usage: "addresses of other peers that can be connected to",
				},
			},
			Action: func(cCtx *cli.Context) error {
				return dep.Start(Config{
					DbPort:                cCtx.Int(dbPortFlagName),
					PeeringPort:           cCtx.Int(peeringPortFlagName),
					AdvertisedIp:          cCtx.String(advertisedIpFlagName),
					ExistingPeerAddresses: cCtx.StringSlice(existingPeerAddresses),
				})
			},
		},
	}
}
