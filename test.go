package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func main() {
	func() {
		t := time.Now()
		var genesisBlock = Block{}

		genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), " "}
		//	spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
		fmt.Println(Blockchain)
		fmt.Println(" ")

		var msg int
		for i := 0; i < 3; i++ {
			fmt.Scanln(&msg)
			prevBlock := Blockchain[len(Blockchain)-1]
			newBlock := generateBlock(prevBlock, msg)

			Blockchain = append(Blockchain, newBlock)

			fmt.Println(newBlock)

		}
		fmt.Println(Blockchain)
	}()

}

// SHA256 hasing
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// create a new block using previous block's hash
func generateBlock(oldBlock Block, BPM int) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}
