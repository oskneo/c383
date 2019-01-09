package blockchain

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TODO: some useful tests of Blocks

func TestBlock(t *testing.T) {

	b0 := Initial(16)

	b0.Mine(1)
	fmt.Println(hex.EncodeToString(b0.Hash))

	b1 := b0.Next("message")
	b1.Mine(1)
	fmt.Println(hex.EncodeToString(b1.Hash))

	assert.Equal(t, hex.EncodeToString(b0.Hash), "6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000")
	assert.Equal(t, hex.EncodeToString(b1.Hash), "9b4417b36afa6d31c728eed7abc14dd84468fdb055d8f3cbe308b0179df40000")
}

func TestBlock2(t *testing.T) {

	b0 := Initial(7)
	b0.Mine(1)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))

	assert.Equal(t, hex.EncodeToString(b0.Hash), "379bf2fb1a558872f09442a45e300e72f00f03f2c6f4dd29971f67ea4f3d5300")
	assert.Equal(t, hex.EncodeToString(b1.Hash), "4a1c722d8021346fa2f440d7f0bbaa585e632f68fd20fed812fc944613b92500")
	assert.Equal(t, hex.EncodeToString(b2.Hash), "ba2f9bf0f9ec629db726f1a5fe7312eb76270459e3f5bfdc4e213df9e47cd380")

	assert.Equal(t, b0.Proof, uint64(385))

}

func TestBlock3(t *testing.T) {

	b0 := Initial(20)
	b0.Mine(1)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))

	assert.Equal(t, hex.EncodeToString(b0.Hash), "19e2d3b3f0e2ebda3891979d76f957a5d51e1ba0b43f4296d8fb37c470600000")
	assert.Equal(t, hex.EncodeToString(b1.Hash), "a42b7e319ee2dee845f1eb842c31dac60a94c04432319638ec1b9f989d000000")
	assert.Equal(t, hex.EncodeToString(b2.Hash), "6c589f7a3d2df217fdb39cd969006bc8651a0a3251ffb50470cbc9a0e4d00000")

}

func TestBlockChain(t *testing.T) {

	bc := new(Blockchain)

	block0 := Initial(20)
	for !block0.ValidHash() {
		block0.Mine(8)
	}
	blockpointer := &block0
	bc.Add(block0)

	for i := 0; i < 10; i++ {
		nextblock := blockpointer.Next("Hi!")
		for !nextblock.ValidHash() {
			nextblock.Mine(8)
		}
		blockpointer = &nextblock
		bc.Add(nextblock)

	}

	assert.True(t, bc.IsValid())
}

func TestConcurrency(t *testing.T) {

	start1 := time.Now()
	block1 := Initial(20)
	block1.Mine(1)
	end1 := time.Now()

	start0 := time.Now()
	block0 := Initial(20)
	block0.Mine(8)
	end0 := time.Now()

	assert.True(t, (float64(end0.Sub(start0))) < (float64(end1.Sub(start1)))*0.6)
}
