package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block ...
type Block struct {
	Str          string
	Time         time.Time
	Company      string
	PreviousHash string
	Hash         string
}

// CalculateHash gets the hash of the current Block object
func (b Block) CalculateHash() string {
	// gather all the data of the block
	blockData := b.Str + b.Time.String() + b.Company + b.PreviousHash
	// hash using sha256
	blockHash := sha256.Sum256([]byte(blockData))

	return fmt.Sprintf("%x", blockHash)
}
