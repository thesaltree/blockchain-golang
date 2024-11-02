package blockchain

import (
	"math/rand"
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

func CreateBlock(data string, prevHash string) *Block {
	rand.Seed(time.Now().UnixNano())
	initialNonce := rand.Intn(10000)

	block := &Block{"", data, prevHash, initialNonce}

	newPow := NewProofOfWork(block)

	nonce, hash := newPow.MineBlock()

	block.Hash = string(hash[:])
	block.Nonce = nonce

	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", "")
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func (chain *Blockchain) AddBlock(data string) {
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, lastBlock.PrevHash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
