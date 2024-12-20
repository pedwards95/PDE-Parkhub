package blockchain

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

// Blockchain ...
type Blockchain struct {
	firstBlock *Block
	chain      []*Block
}

// CreateBlockchain makes a new Blockchain object
func CreateBlockchain(block *Block) Blockchain {
	// always calculate the hash when adding a new block
	block.Hash = block.CalculateHash()
	// previous will not exist, because this is the first node
	block.PreviousHash = "0"
	return Blockchain{
		firstBlock: block,
		chain:      []*Block{block},
	}
}

// AddBlock adds a new Block object to the blockchain
func (b *Blockchain) AddBlock(block *Block) {
	// get the last in the chain, to set previous hash
	lastBlock := b.chain[len(b.chain)-1]
	block.PreviousHash = lastBlock.Hash
	// always calculate the hash when adding a new block
	block.Hash = block.CalculateHash()
	// add to end of chain
	b.chain = append(b.chain, block)
}

// IsValid checks if the whole chain is still valid
func (b *Blockchain) IsValid() bool {
	// start at index 1 because the first node will not have a previous hash to compare
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		// make sure all hashes match along the chain, both to themselves and to the PreviousHash field
		if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

// PrintLastHashString prints the hash of the last Block in the chain, and prints QR code to file
func (b *Blockchain) PrintLastHashString() {
	lastBlock := b.chain[len(b.chain)-1]
	fileName := fmt.Sprintf("block%d.png", len(b.chain))
	// create QR code
	err := qrcode.WriteFile(lastBlock.Hash, qrcode.Medium, 256, fileName)
	if err != nil {
		fmt.Println("Error writing QR code!")
		return
	}
	msg := fmt.Sprintf("Last block's hash string is: %+v, Printed QR code to file %s", lastBlock.Hash, fileName)
	fmt.Println(msg)
}
