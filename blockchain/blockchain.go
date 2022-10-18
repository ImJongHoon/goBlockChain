package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

type blockchain struct { //struct. 인스턴스 없음.
	blockList []*block
}

func (b *blockchain) AddBlock(data string) {
	b.blockList = append(b.blockList, createBlock(data))
}

func (b *blockchain) ReturnAllBlocks() []*block {
	return b.blockList
}

var b *blockchain
var once sync.Once

func getLastHash() string {
	totalBlocksNum := len(GetBlockchain().blockList)
	if totalBlocksNum == 0 { //block이 하나도 없으면
		return ""
	}
	return GetBlockchain().blockList[totalBlocksNum-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

//singleton 패턴. 약간 생성자 느낌인거 같기도 하고.
func GetBlockchain() *blockchain {
	if b == nil {
		//병렬적으로 처리하다가 동시 접속하는 걸 막기 위해.
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("First Block")
		})
	}

	//한번만 선언한다. 즉 처음 말고 한번 더 실행되는 경우 그냥 b 반환
	return b
}
