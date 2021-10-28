package main

import (
	"fmt"

	"github.com/lcmps/gopher-coin/blockchain"
	"github.com/lcmps/gopher-coin/wallet"
)

func main() {
	testkey()
	transactionTest()
}

func testkey() {
	pk, sk := wallet.GenerateKeyHash()
	fmt.Println("Private key: ", sk)
	fmt.Println("Public key: ", pk)
}

func transactionTest() {
	bchan := blockchain.New()
	bchan.CreateTransaction((blockchain.TransactionNFO{
		SentFrom: "You",
		SentTo:   "Me",
		Amount:   100,
	}))

	bchan.CreateTransaction((blockchain.TransactionNFO{
		SentFrom: "Me",
		SentTo:   "You",
		Amount:   50,
	}))

	fmt.Println("Mining")

	bchan.MinePendingTransactions("You")

	for _, block := range bchan.Chain {
		fmt.Println(
			"\nTime: ", block.Timestamp,
			"\nThis block hash: ", block.Hash,
			"\nLast block hash: ", block.LastHash,
			"\n\nBlock Transactions\n\n", block.Transactions,
		)
	}

	fmt.Println("\n Is blockchain valid?", bchan.IsValid())
}
