package validation

import (
	"blockchainproject/blockchain"
	"fmt"
)

func ReplaceChain(newBlocks []blockchain.Block) {
	if len(newBlocks) > len(blockchain.Blockchain) && isChainValid(newBlocks) {
		blockchain.Blockchain = newBlocks
		fmt.Println("Blockchain replaced with the longer and valid chain. ")
	} else {
		fmt.Println("Received blockchain is not longer or valid. Keeping the current blockchain. ")
	}
}

func isChainValid(chain []blockchain.Block) bool {
	for i := 1; i < len(chain); i++ {
		if !isBlockValid(chain[i], chain[i-1]) {
			return false
		}
	}
	return true
}
