package main

import (
	"blockchainproject/blockchain"
	"blockchainproject/mining"
	"blockchainproject/validation"
	"fmt"
	"time"
)

func main() {
	genesisBlock := blockchain.Block{Index: 0, Timestamp: time.Now().Unix(), Transaction: []blockchain.Transaction{}, PrevHash: "", Hash: ""}
	blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)

	for i := 0; i < 5; i++ {
		transactions := []blockchain.Transaction{
			{Sender: fmt.Sprintf("Sender%d", i), Recipient: fmt.Sprintf("Recipient%d", i), Amount: float64(i) + 1.0},
		}

		newBlock := mining.MineBlock(blockchain.Blockchain[len(blockchain.Blockchain)-1], transactions)
		blockchain.Blockchain = append(blockchain.Blockchain, newBlock)

		fmt.Printf("Block #%d has been added to the blockchain!\n", newBlock.Index)
		fmt.Printf("Timestamp: %v\n", time.Unix(newBlock.Timestamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("Hash: %s\n", newBlock.Hash)
	}

	newChain := append(blockchain.Blockchain[:2], blockchain.GenerateBlock(blockchain.Blockchain[1], []blockchain.Transaction{{Sender: "Sender", Recipient: "Recipient", Amount: 1.0}}))
	validation.ReplaceChain(newChain)

	fmt.Println("\nBlockchain:")
	for _, block := range blockchain.Blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Println("Transactions:")
		for _, tx := range block.Transaction {
			fmt.Printf("\tSender: %s, Recipient: %s, Amount: %f\n", tx.Sender, tx.Recipient, tx.Amount)
		}
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("---------------------------------")
	}
}
