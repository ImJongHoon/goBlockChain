package main

import (
	"fmt"

	"github.com/ImJongHoon/goBlockChain/blockchain"
)

//linked list 느낌

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.ReturnAllBlocks() {
		fmt.Printf("%s %s %s \n\n", block.Data, block.Hash, block.PrevHash)
	}
}
