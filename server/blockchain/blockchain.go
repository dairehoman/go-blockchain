package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block :
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

// Blockchain :
type Blockchain struct {
	Blocks []*Block
}

// GenHash :
func (b *Block) GenHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// NewBlock :
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}

	block.GenHash()

	return block
}

// AddBlock :
func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

// NewBlockchain :
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// NewGenesisBlock :
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}
