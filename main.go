package main

import (
	"blockchain-golang/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("block a")
	chain.AddBlock("block b")

	for _, block := range chain.Blocks {
		fmt.Printf("Data in block: %x\n", block.Data)
		fmt.Printf("Previous hash in block: %s\n", block.PrevHash)
		fmt.Printf("Hash in block: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("IsValidPoW: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println()
	}
}
