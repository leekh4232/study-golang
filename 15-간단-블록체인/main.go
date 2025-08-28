package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

/** 1. 블록과 체인을 위한 구조체 정의하기 */
type Block struct {
	Timestamp     int64  // 만든 시간
	Data          string // 담을 데이터
	PrevBlockHash []byte // 이전 블록 해시
	Hash          []byte // 현재 블록 해시
}

// 블록체인 구조체 정의
type Blockchain struct {
	blocks []*Block
}

// 블록의 해시 계산 함수
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := append(b.PrevBlockHash, append([]byte(b.Data), timestamp...)...)
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 새 블록 생성 함수
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), data, prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// 블록체인 생성 (제네시스 블록 포함)
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", []byte{})
	return &Blockchain{[]*Block{genesisBlock}}
}

/** 2. 블록 추가하기 */
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

/** 3. 블록체인 검증하기 */
func (bc *Blockchain) IsChainValid() bool {
	for i := 1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		prevBlock := bc.blocks[i-1]

		tempHash := currentBlock.Hash
		currentBlock.SetHash()
		if string(currentBlock.Hash) != string(tempHash) {
			return false
		}
		currentBlock.Hash = tempHash

		if string(currentBlock.PrevBlockHash) != string(prevBlock.Hash) {
			return false
		}
	}
	return true
}

/******** main() *********/
func main() {
	// 새 블록체인 생성
	bc := NewBlockchain()

	// 블록 추가
	bc.AddBlock("Send 1 BTC to Alice")
	bc.AddBlock("Send 2 BTC to Bob")

	// 블록 출력
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %s\n", hex.EncodeToString(block.PrevBlockHash))
		fmt.Printf("Time: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", hex.EncodeToString(block.Hash))
		fmt.Println("----------------------------------")
	}

	// 체인 검증
	if bc.IsChainValid() {
		fmt.Println("✅ 블록체인 무결성 검증 성공!")
	} else {
		fmt.Println("❌ 블록체인 무결성 검증 실패!")
	}
}
