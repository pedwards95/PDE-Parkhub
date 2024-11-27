package blockchain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHash(t *testing.T) {
	ti := time.Now()

	block1 := &Block{
		Str:          "abcdefg",
		Time:         ti,
		Company:      "Test Inc.",
		PreviousHash: "cbehjgf5e5656jvhber76547dsfg845gvhd6fg3cgf",
	}

	block2 := &Block{
		Str:          "abcdefg",
		Time:         ti,
		Company:      "Test Inc.",
		PreviousHash: "cbehjgf5e5656jvhber76547dsfg845gvhd6fg3cgf",
	}

	hash1 := block1.CalculateHash()

	time.Sleep(1000)

	hash2 := block2.CalculateHash()

	assert.Equal(t, hash1, hash2, "Both blocks should hash out the same")
	assert.Equal(t, 64, len(hash1), "hash should be a set length of 64")
}
