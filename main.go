package main

import (
	"fmt"

	"github.com/lcmps/gopher-coin/blockchain"
)

func main() {
	bchan := blockchain.New()
	bchan.AddBlock()
	bchan.AddBlock()

	for _, block := range bchan.Chain {
		fmt.Println(
			"Block position:", block.Index,
			"\nTime: ", block.Timestamp,
			"\nThis block hash: ", block.Hash,
			"\nLast block hash: ", block.LastHash,
		)
	}
}
