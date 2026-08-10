package node

import (
	"flag"
	"strings"
)

type Config struct {
	Address string
	Peers   []string
}

func ParseConfig(args []string) (Config, error) {
	flags := flag.NewFlagSet("node", flag.ContinueOnError)
	address := flags.String(
		"address",
		"localhost:8001",
		"HTTP address used by this node",
	)

	peersText := flags.String(
		"peers",
		"",
		"comma-separated peer addresses",
	)
	if err := flags.Parse(args); err != nil {
		return Config{}, err
	}
	peers := []string{}

	for _, peer := range strings.Split(*peersText, ",") {
		peer = strings.TrimSpace(peer)

		if peer != "" {
			peers = append(peers, peer)
		}
	}
	return Config{
		Address: *address,
		Peers:   peers,
	}, nil
}
