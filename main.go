package main

import (
	"fmt"

	"github.com/pedwards95/PDE-Parkhub/blockchain"
	"github.com/pedwards95/PDE-Parkhub/random"
)

func main() {
	chain := blockchain.CreateBlockchain(random.GenerateRandomBlock())
	chain.PrintLastHashString()
	chain.AddBlock(random.GenerateRandomBlock())
	chain.PrintLastHashString()
	chain.AddBlock(random.GenerateRandomBlock())
	chain.PrintLastHashString()
	chain.AddBlock(random.GenerateRandomBlock())
	chain.PrintLastHashString()
	if chain.IsValid() {
		fmt.Println("BlockChain is currently valid!")
	}
}
