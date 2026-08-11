package node

import (
	"encoding/json"
	"net/http"
)

type PeersResponse struct {
	Peers []string `json:"peers"`
}

func (n *Node) handlePeers(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(PeersResponse{
		Peers: n.Config.Peers,
	})
}
