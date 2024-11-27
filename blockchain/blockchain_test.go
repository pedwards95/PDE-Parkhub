package blockchain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateBlockchain(t *testing.T) {
	ti := time.Now()

	block1 := &Block{
		Str:          "abcdefg",
		Time:         ti,
		Company:      "Test Inc.",
		PreviousHash: "0",
	}

	bc := CreateBlockchain(block1)

	assert.Equal(t, block1.Str, bc.chain[0].Str, "Should be the same block")
	block1.Str = "new string"
	assert.Equal(t, block1.Str, bc.chain[0].Str, "Should be the same block changed")
	assert.Equal(t, block1.Hash, bc.firstBlock.Hash)
}

func TestAddBlock(t *testing.T) {
	ti := time.Now()

	block1 := &Block{
		Str:     "abcdefg",
		Time:    ti,
		Company: "Test Inc.",
	}
	block2 := &Block{
		Str:     "123456",
		Time:    ti.Add(time.Hour),
		Company: "Test Inc.",
	}
	block3 := &Block{
		Str:     "oogabooga central",
		Time:    ti.Add(time.Hour * 2),
		Company: "Test Inc.",
	}

	bc := CreateBlockchain(block1)
	bc.AddBlock(block2)
	bc.AddBlock(block3)

	assert.Equal(t, block2.Hash, bc.chain[1].Hash)
	assert.Equal(t, block2.PreviousHash, bc.chain[0].Hash)
	assert.Equal(t, block3.Hash, bc.chain[2].Hash)
}

func TestIsValid(t *testing.T) {
	ti := time.Now()

	block1 := &Block{
		Str:     "abcdefg",
		Time:    ti,
		Company: "Test Inc.",
	}
	block2 := &Block{
		Str:     "123456",
		Time:    ti.Add(time.Hour),
		Company: "Test Inc.",
	}
	block3 := &Block{
		Str:     "oogabooga central",
		Time:    ti.Add(time.Hour * 2),
		Company: "Test Inc.",
	}

	bc := CreateBlockchain(block1)
	bc.AddBlock(block2)
	bc.AddBlock(block3)

	assert.True(t, bc.IsValid())
	block2.Str = "bad string"
	assert.False(t, bc.IsValid())
	block2.Str = "123456"
	assert.True(t, bc.IsValid())
	block3.PreviousHash = "wwwtttredfddsaaddvfrrr7cdbjhvbhvrbhj"
	assert.False(t, bc.IsValid())
}
