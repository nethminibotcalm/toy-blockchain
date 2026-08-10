package node

import "testing"

func TestParseConfig(t *testing.T) {
	args := []string{
		"-address", "localhost:8002",
		"-peers", "localhost:8001, localhost:8003",
	}

	config, err := ParseConfig(args)

	if err != nil {
		t.Fatalf("expected configuration to parse: %v", err)
	}
	if config.Address != "localhost:8002" {
		t.Fatalf(
			"expected address localhost:8002, got %s",
			config.Address,
		)
	}
	if len(config.Peers) != 2 {
		t.Fatalf(
			"expected 2 peers, got %d",
			len(config.Peers),
		)
	}

	if config.Peers[0] != "localhost:8001" {
		t.Errorf("unexpected first peer: %s", config.Peers[0])
	}

	if config.Peers[1] != "localhost:8003" {
		t.Errorf("unexpected second peer: %s", config.Peers[1])
	}
}
