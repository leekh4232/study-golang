package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// 1. Block은 블록체인의 각 요소를 나타냄
type Block struct {
	Timestamp     int64
	Data          string
	PrevBlockHash []byte
	Hash          []byte
}

// 2. SetHash는 블록의 데이터를 기반으로 해시를 계산하고 설정함
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := append(b.PrevBlockHash, append([]byte(b.Data), timestamp...)...)
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 3. NewBlock은 주어진 데이터와 이전 블록 해시로 새 블록을 생성함
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), data, prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// 4. Blockchain은 블록들의 체인을 나타냄
type Blockchain struct {
	blocks []*Block
}

// 5. NewBlockchain은 제네시스 블록과 함께 새 블록체인을 생성함
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", []byte{})
	return &Blockchain{[]*Block{genesisBlock}}
}

// 6. AddBlock은 블록체인에 새 블록을 추가함
func (bc *Blockchain) AddBlock(data string) {
	// 2. 체인의 마지막 블록을 가져옴
	prevBlock := bc.blocks[len(bc.blocks)-1]
	// 3. 이전 블록의 해시를 사용하여 새 블록 생성
	newBlock := NewBlock(data, prevBlock.Hash)
	// 4. 새 블록을 체인에 추가
	bc.blocks = append(bc.blocks, newBlock)
}

// 7. IsChainValid는 블록체인의 무결성을 검증함
func (bc *Blockchain) IsChainValid() bool {
	// 1. 제네시스 블록을 제외한 모든 블록을 순회
	for i := 1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		prevBlock := bc.blocks[i-1]

		// 2. 현재 블록의 데이터로 해시를 다시 계산하여 저장된 해시와 일치하는지 확인
		tempHash := currentBlock.Hash
		currentBlock.SetHash() // 해시 다시 계산
		if string(currentBlock.Hash) != string(tempHash) {
			return false // 데이터가 변조됨
		}
		currentBlock.Hash = tempHash // 원래 해시로 복원

		// 3. 이전 블록의 해시와 현재 블록의 PrevBlockHash가 일치하는지 확인
		if string(currentBlock.PrevBlockHash) != string(prevBlock.Hash) {
			return false // 체인 연결이 끊어짐
		}
	}
	return true
}

func main() {
	fmt.Println("============= step1 =============")
	bc := NewBlockchain()
	genesis := bc.blocks[0]
	fmt.Printf("Genesis Block Data: %s\n", genesis.Data)
	fmt.Printf("Genesis Block Hash: %s\n", hex.EncodeToString(genesis.Hash))

	fmt.Println("\n\n============= step2 =============")
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %s\n", hex.EncodeToString(block.PrevBlockHash))
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", hex.EncodeToString(block.Hash))
		fmt.Println("----------------------------------")
	}

	fmt.Println("\n\n============= step3 =============")
	if bc.IsChainValid() {
		fmt.Println("✅ 블록체인 무결성 검증 성공!")
	} else {
		fmt.Println("❌ 블록체인 무결성 검증 실패!")
	}
}
