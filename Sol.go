package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Block
type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func GetSHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

//
func CalculateHash(inputBlock *Block) string {
	reg := inputBlock.transactions
	var concat_transactions = strings.Join(reg, "")
	var final = concat_transactions + inputBlock.prevHash
	return GetSHA256Hash(final)
}

//
func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {
	var block1 = Block{transactions: transactionsToInsert, prevPointer: chainHead, prevHash: chainHead.currentHash}
	block1.currentHash = CalculateHash(&block1)
	return &block1
	//insert new block and return head pointer
}
func PrintBlock(chainHead *Block) {
	fmt.Println("Transactions:", chainHead.transactions)
}

//
func SearchTransaction(transact []string, tran string) int {
	var flag = false
	var index = -1
	for i := 0; i < len(transact) && flag == false; i++ {
		if transact[i] == tran {
			index = i
			flag = true
		}
	}
	return index
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	var chain = chainHead
	var flag = false
	for chain != nil && flag == false {
		if SearchTransaction(chain.transactions, oldTrans) != -1 {
			chain.transactions[SearchTransaction(chain.transactions, oldTrans)] = newTrans
			chain.currentHash = CalculateHash(chain)
			flag = true
		}
		chain = chain.prevPointer
	}
	if flag == false {
		fmt.Println("Transaction not found")
	} else {
		fmt.Println("Transaction changed successfully")
	}

}

//
func ListBlocks(chainHead *Block) {

	//dispaly the data(transaction) inside all blocks
	var chain = chainHead
	for chain != nil {
		PrintBlock(chain)
		chain = chain.prevPointer
	}

}

//

func VerifyChain(chainHead *Block) {
	var chain = chainHead
	var flag = false
	for (chain != nil) && (flag == false) {
		if chain.prevPointer != nil {
			if chain.prevHash != chain.prevPointer.currentHash {
				flag = true
			}
		}
		chain = chain.prevPointer
	}
	//check whether "Block chain is compromised" or "Block chain is unchanged"
	if flag {
		fmt.Println("Block chain is compromised")
	} else {
		fmt.Println("Block chain is unchanged")
	}
}

func main() {
	var genesis = Block{transactions: []string{"a", "b", "c"}, prevPointer: nil, prevHash: "00000"}
	genesis.currentHash = CalculateHash(&genesis)
	var chainHead = &genesis
	chainHead = InsertBlock([]string{"d", "e", "f"}, chainHead)
	chainHead = InsertBlock([]string{"g", "h", "i"}, chainHead)
	ListBlocks(chainHead)
	VerifyChain(chainHead)
	ChangeBlock("j", "s", chainHead)
	ChangeBlock("a", "z", chainHead)
	ListBlocks(chainHead)
	VerifyChain(chainHead)
}
