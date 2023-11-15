package validation

import (
	"blockchainproject/blockchain"
)

func isBlockValid(newBlock, prevBlock blockchain.Block) bool {
	if prevBlock.Index+1 != newBlock.Index {
		return false
	}

	if prevBlock.Hash != newBlock.PrevHash {
		return false
	}

	if blockchain.CalculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
