package node

import "net/http"

func (n *Node) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", n.handleStatus)
	return mux
}
func (n *Node) Start() error {
	return http.ListenAndServe(
		n.Config.Address,
		n.Handler(),
	)
}
