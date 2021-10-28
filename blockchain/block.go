package blockchain

import (
	"fmt"
	"strings"
	"time"

	"github.com/lcmps/gopher-coin/hash"
)

const (
	Difficulty   = 2
	MiningReward = 100
)

// TransactionNFO will be implemented later on
type TransactionNFO struct {
	SentFrom string
	SentTo   string
	Amount   int
}

type Block struct {
	Timestamp    time.Time
	Hash         string
	LastHash     string
	Transactions []TransactionNFO
	Nonce        int
}

func (block *Block) CalculateHash() string {
	bytes := hash.EncodeToBytes(fmt.Sprintf("%s%s%d", block.Timestamp, block.LastHash, block.Nonce))
	return hash.EncodeToSha512(bytes)
}

func (block *Block) MineBlock(difficulty int) {
	for !strings.HasPrefix(block.Hash, strings.Repeat("0", difficulty)) {
		block.Nonce++
		block.Hash = block.CalculateHash()
	}
}
