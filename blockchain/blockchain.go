package blockchain

import (
	"time"

	"github.com/lcmps/gopher-coin/hash"
)

// TransactionNFO will be implemented later on
// type TransactionNFO struct {
// 	SentFrom string
// 	SentTo   string
// }

type Block struct {
	Index     int
	Timestamp time.Time
	Hash      string
	LastHash  string
}

type BlockChain struct {
	Chain []Block
}

func New() BlockChain {
	return CreateGenesisBlock()
}

func CreateGenesisBlock() BlockChain {
	gen := Block{
		Index:     0,
		Timestamp: time.Now(),
		LastHash:  "",
	}
	gen.Hash = hash.EncodeToSha512(hash.EncodeToBytes(gen))

	return BlockChain{
		Chain: []Block{gen},
	}
}

func (chain *BlockChain) AddBlock() {
	b := Block{
		Index:     chain.GetLastBlock().Index + 1,
		Timestamp: time.Now(),
	}
	b.LastHash = chain.GetLastBlock().Hash
	b.Hash = b.CalculateHash()

	chain.Chain = append(chain.Chain, b)
}

func (chain *BlockChain) GetLastBlock() Block {
	return chain.Chain[len(chain.Chain)-1]
}

func (block *Block) CalculateHash() string {
	bytes := hash.EncodeToBytes(&block)
	return hash.EncodeToSha512(bytes)
}
