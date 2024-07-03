package blockchain

import (
	"errors"

	"github.com/LGROW101/Block-Blockchain/block"
	"github.com/LGROW101/Block-Blockchain/transaction"
	"github.com/LGROW101/Block-Blockchain/wallet"
)

var (
	ErrWalletNotFound = errors.New("wallet not found")
)

type BlockChain struct {
	difficulty int
	blocks     []*block.Block
	wallets    map[string]*wallet.Wallet
}

func New(difficulty int) *BlockChain {
	return &BlockChain{
		difficulty: difficulty,
		blocks:     []*block.Block{block.New([]transaction.Transaction{}, "")},
		wallets:    make(map[string]*wallet.Wallet),
	}
}

func (bc *BlockChain) CreateWallet() string {
	wallet := wallet.New()
	address := wallet.GetAddress()
	bc.wallets[address] = wallet
	return address
}

func (bc *BlockChain) GetWallet(address string) (*wallet.Wallet, error) {
	if wallet, ok := bc.wallets[address]; ok {
		return wallet, nil
	}
	return nil, ErrWalletNotFound
}

func (bc *BlockChain) AddBlock(transactions []transaction.Transaction) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.New(transactions, prevBlock.Hash)
	newBlock.Mine(bc.difficulty)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *BlockChain) CreateTransaction(from, to string, amount float64) error {
	fromWallet, ok := bc.wallets[from]
	if !ok {
		return errors.New("sender wallet not found")
	}
	toWallet, ok := bc.wallets[to]
	if !ok {
		return errors.New("recipient wallet not found")
	}

	_ = fromWallet
	_ = toWallet

	tx := transaction.Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}

	bc.AddBlock([]transaction.Transaction{tx})
	return nil
}

func (bc *BlockChain) IsValid() bool {
	for i := 1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		prevBlock := bc.blocks[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}
	return true
}

func (bc *BlockChain) Blocks() []*block.Block {
	return bc.blocks
}
