package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/LGROW101/Block-Blockchain/transaction"
)

type Block struct {
	Transactions []transaction.Transaction `json:"transactions"`
	Hash         string                    `json:"hash"`
	PrevHash     string                    `json:"prevHash"`
	Timestamp    string                    `json:"timestamp"`
	Nonce        int                       `json:"nonce"`
}

func (b *Block) CalculateHash() string {
	data := b.PrevHash + b.getTransactionsString() + b.Timestamp + strconv.Itoa(b.Nonce)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (b *Block) getTransactionsString() string {
	var txStr string
	for _, tx := range b.Transactions {
		txStr += tx.ToString()
	}
	return txStr
}

func New(transactions []transaction.Transaction, prevHash string) *Block {
	block := &Block{
		Transactions: transactions,
		PrevHash:     prevHash,
		Timestamp:    time.Now().String(),
	}
	block.Hash = block.CalculateHash()
	return block
}

func (b *Block) Mine(difficulty int) {
	target := strings.Repeat("0", difficulty)

	for !strings.HasPrefix(b.Hash, target) {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
}
