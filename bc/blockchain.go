package bc

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateNewBlock(oldBlock Block) Block {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = oldBlock.Hash
	return newBlock
}

func GenerateBlock(oldBlock Block, BPM int) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}

// CalculateHash Calculate SHA256 hash for a new block
func CalculateHash(block Block) string {
	data := strconv.Itoa(block.Index) + block.Timestamp + block.Hash + block.PrevHash
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// IsBlockValid make sure block is valid by checking index, and comparing the hash of the previous block
func IsBlockValid(oldBlock Block, newBlock Block) bool {
	if (oldBlock.Index != newBlock.Index+1) || (newBlock.PrevHash != oldBlock.Hash) {
		return false
	} else if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
