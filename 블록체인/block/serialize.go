package block

import (
	"bytes"
	"encoding/gob"
)

func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	encoder.Encode(b)

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	decoder.Decode(&block)

	return &block
}