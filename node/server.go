package node

import "net/http"

func (n *Node) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", n.handleStatus)
	mux.HandleFunc("/peers", n.handlePeers)
	mux.HandleFunc("/mempool", n.handleMempool)
	mux.HandleFunc("/balances", n.handleBalances)
	mux.HandleFunc("/chain", n.handleChain)
	mux.HandleFunc("/transactions", n.handleTransaction)
	return mux
}
func (n *Node) Start() error {
	return http.ListenAndServe(
		n.Config.Address,
		n.Handler(),
	)
}
