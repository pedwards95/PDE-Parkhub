package main

import (
	"testing"

	"github.com/pedwards95/PDE-Parkhub/blockchain"
	"github.com/pedwards95/PDE-Parkhub/random"
	"github.com/stretchr/testify/assert"
)

func TestBlockChain(t *testing.T) {
	chain := blockchain.CreateBlockchain(random.GenerateRandomBlock())
	chain.AddBlock(random.GenerateRandomBlock())
	chain.AddBlock(random.GenerateRandomBlock())
	chain.AddBlock(random.GenerateRandomBlock())

	assert.True(t, chain.IsValid(), "Should be a valid chain")
}
