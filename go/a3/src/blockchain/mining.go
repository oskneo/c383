package blockchain

import (
	"work_queue"
)

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	start uint64
	end   uint64
	block *Block
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	// TODO
	end++
	ck0 := (end - start) / chunks
	ck1 := ck0 + 1

	i := start

	workqueue := work_queue.Create(uint(workers), uint(chunks))
	//Divide works into chunks parts and enqueue each part to concurrent worker
	for i < end {
		mw := miningWorker{
			start: i,
			block: &blk}
		if ck0 != 0 && (end-i)%ck0 == 0 {
			mw.end = i + ck0
			i += ck0
		} else {
			mw.end = i + ck1
			i += ck1
		}
		workqueue.Enqueue(mw)

	}
	//receive result and change its type to mingingresult
	//shutdown if the proof is found
	mrs := MiningResult{Found: false}
	dones := uint64(0)
	for {
		select {
		case rs := <-workqueue.Results:
			mrs = rs.(MiningResult)
			if mrs.Proof >= end-1 {
				return mrs
			} else if mrs.Found {
				workqueue.Shutdown()
				return mrs
			}
		case <-workqueue.Done:
			dones++
			if dones >= workers {
				return mrs
			}

		}
	}

	return mrs
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}

//the run function to calcultate proof in the range given then return the result
func (m miningWorker) Run() interface{} {
	mrs := MiningResult{Found: false}
	for i := m.start; i < m.end; i++ {
		m.block.SetProof(i)
		mrs.Proof = i
		if m.block.ValidHash() {
			mrs.Found = true
			return mrs
		}
	}
	return mrs
}
