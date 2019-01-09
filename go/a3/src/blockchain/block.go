package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	// TODO
	b := Block{
		PrevHash:   []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Generation: 0,
		Difficulty: difficulty,
		Data:       "",
		Proof:      0}

	return b
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	return Block{
		PrevHash:   prev_block.Hash,
		Generation: prev_block.Generation + 1,
		Difficulty: prev_block.Difficulty,
		Data:       data,
		Proof:      0}
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	// TODO

	hstring := fmt.Sprintf("%s:%v:%v:%s:%v", hex.EncodeToString(blk.PrevHash), blk.Generation, blk.Difficulty, blk.Data, blk.Proof)

	output := sha256.Sum256([]byte(hstring))
	return output[:]

}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO

	bits := blk.Difficulty / 8
	remainder := blk.Difficulty % 8

	if remainder != 0 {
		bits++
	}

	//Check check the last bits byte
	//Check whether the byte have enough 0s by checking whether the remainder is 0
	//when use 2^remainder to divide for the (31-i)th byte

	for i := uint8(0); i < bits; i++ {
		if i == bits-1 && remainder != 0 {

			if uint(blk.Hash[31-i])%(1<<remainder) != 0 {
				return false
			}
		} else {
			//check whether the bytes after the (31-i)th byte is zero
			if blk.Hash == nil || uint(blk.Hash[31-i]) != 0 {
				return false
			}
		}
	}
	return true

}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
