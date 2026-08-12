package node

import "fmt"

func (n *Node) SyncFromPeer(peer string) error {
	status, err := fetchPeerStatus(peer)

	if err != nil {
		return fmt.Errorf(
			"failed to fetch peer status: %w",
			err,
		)
	}
	n.mu.RLock()
	localHeight := len(n.Blockchain.Blocks) - 1
	n.mu.RUnlock()
	if status.Height <= localHeight {
		return nil
	}
	missingBlocks, err := fetchMissingBlocks(
		peer,
		localHeight+1,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to fetch missing blocks: %w",
			err,
		)
	}
	for _, missingBlock := range missingBlocks {
		n.mu.Lock()

		if err := n.Blockchain.AddReceivedBlock(
			missingBlock,
		); err != nil {
			n.mu.Unlock()

			return fmt.Errorf(
				"invalid received block %d: %w",
				missingBlock.Index,
				err,
			)
		}

		n.seenBlocks[missingBlock.Hash] = true

		for _, tx := range missingBlock.Transactions {
			if tx.ID != "" {
				n.seenTransactions[tx.ID] = true
			}
		}

		n.mu.Unlock()
	}
	n.mu.RLock()
	finalHeight := len(n.Blockchain.Blocks) - 1
	n.mu.RUnlock()

	if finalHeight != status.Height {
		return fmt.Errorf(
			"sync incomplete: local height %d, peer height %d",
			finalHeight,
			status.Height,
		)
	}

	return nil
}
