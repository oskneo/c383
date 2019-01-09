package blockchain

import (
	"bytes"
	// "fmt"
	// "encoding/hex"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	// TODO
	chain.Chain = append(chain.Chain, blk)
}

func (chain Blockchain) IsValid() bool {
	// TODO

	for i, blk := range chain.Chain {
		if i == 0 && !bytes.Equal(blk.PrevHash, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {

			return false
		}
		if blk.Generation != uint64(i) {
			panic("0")
			return false
		}
		if blk.Difficulty != chain.Chain[0].Difficulty {
			panic("0")
			return false
		}
		if i != 0 && !bytes.Equal(blk.PrevHash, chain.Chain[i-1].Hash) {
			panic("0")
			return false
		}
		if !bytes.Equal(blk.CalcHash(), blk.Hash) {
			panic("0")
			return false
		}
		if !blk.ValidHash() {
			panic("0")
			return false
		}

	}
	return true
}
