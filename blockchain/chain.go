package blockchain

import (
	"fmt"
	"time"

	"github.com/lcmps/gopher-coin/hash"
)

type BlockChain struct {
	Chain               []Block
	PendingTransactions []TransactionNFO
}

func New() BlockChain {
	return CreateGenesisBlock()
}

func CreateGenesisBlock() BlockChain {
	gen := Block{
		Timestamp: time.Now(),
		LastHash:  "",
		Nonce:     0,
	}
	gen.Hash = hash.EncodeToSha512(hash.EncodeToBytes(gen))

	return BlockChain{
		Chain: []Block{gen},
	}
}

func (chain *BlockChain) MinePendingTransactions(address string) {
	b := Block{
		Timestamp:    time.Now(),
		Nonce:        0,
		Transactions: chain.PendingTransactions,
	}
	b.LastHash = chain.GetLastBlock().Hash
	b.MineBlock(Difficulty)
	fmt.Println("Block mined:", b.Hash)

	chain.Chain = append(chain.Chain, b)

	chain.PendingTransactions = []TransactionNFO{
		{
			SentFrom: "",
			SentTo:   address,
			Amount:   MiningReward,
		},
	}
}

func (chain *BlockChain) GetBalance(address string) int {
	balance := 0
	for _, block := range chain.Chain {
		for _, transaction := range block.Transactions {
			if transaction.SentFrom == address {
				balance -= transaction.Amount
			}
			if transaction.SentTo == address {
				balance += transaction.Amount
			}
		}
	}
	return balance
}

func (BlockChain *BlockChain) CreateTransaction(transaction TransactionNFO) {
	BlockChain.PendingTransactions = append(BlockChain.PendingTransactions, transaction)
}

func (chain *BlockChain) GetLastBlock() Block {
	return chain.Chain[len(chain.Chain)-1]
}

func (chain *BlockChain) IsValid() bool {
	for i := 1; i < len(chain.Chain); i++ {

		if chain.Chain[i].Hash != chain.Chain[i].CalculateHash() {
			return false
		}
		if chain.Chain[i].LastHash != chain.Chain[i-1].Hash {
			return false
		}
	}
	return true
}
