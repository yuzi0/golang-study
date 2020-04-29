package block

import (	
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain_%s.db" 
const blocksBucket = "blocks" 

type Blockchain struct {
	//Blocks []*Block
	tip []byte
    Db  *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	/*
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	*/
	var lastHash []byte

	bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})

	newBlock := NewBlock(data, lastHash)

	bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		b.Put(newBlock.Hash, newBlock.Serialize())
		b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash
		return nil
	})
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	//return &Blockchain{[]*Block{NewGenesisBlock()}}

	var tip []byte
	db, _ := bolt.Open(dbFile, 0600, nil)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			// 제네시스 블록을 생성한다
			genesis := NewGenesisBlock()
			// DB에 저장한다
			b, _ := tx.CreateBucket([]byte(blocksBucket))
			b.Put(genesis.Hash, genesis.Serialize())
			// 제네시스 블록의 해시를 마지막 블록의 해시로 저장한다.
			b.Put([]byte("l"), genesis.Hash)
			// 제네시스 블록을 끝부분으로 하는 새로운 Blockchain 인스턴스를 생성한다.
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	bc := Blockchain{tip, db}
	return &bc
}
