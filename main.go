package main

import (
	"crypto/sha256"
	"fmt"
)

//linked list 느낌
type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blockList []block
}

//메소드 안에서 사용할 메소드를 선언하는 것과 유사하다.
func (b *blockchain) getLastHash() string {
	if len(b.blockList) > 0 {
		return b.blockList[len(b.blockList)-1].hash
	}

	return ""
}

//포인터 써주기. main에서 만든 객체에 직접 접근하려면.
//b가 인자로 struct를 받는 것 같은데 포인터로 안 하면 인자를 그냥 값을 복사해서 받는 듯.
func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", ""} // data 넣기

	//맨 처음 블럭이 아닌 경우 이전 블럭의 해시값 넣기
	newBlock.prevHash = b.getLastHash()

	//위 두 정보를 합쳐서 이번 블록의 해시 만들기.
	//아래 hash는 byte단위의 해시
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))

	newBlock.hash = fmt.Sprintf("%x", hash) //hash를 16진수로만든 후 Spring 자료형으로 반환하라.
	//위에까지가 새로운 블럭을 생성. 아래서부터가 원래 블록체인에 추가.

	b.blockList = append(b.blockList, newBlock)
}

func (b *blockchain) listBlocks() {
	//index는 안 쓰고 blockList의 단위인 block을 처음부터 끝까지 하나씩 사용.
	for _, block := range b.blockList {
		fmt.Printf("%s %s %s \n", block.data, block.hash, block.prevHash)
	}

}

func main() {
	chain := blockchain{}

	chain.addBlock("First Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")

	chain.listBlocks()
}
