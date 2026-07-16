package blockchain

const MaxDifficulty = 6

func (bc *Blockchain) AdjustDifficulty() {

	// Not enough blocks to adjust
	if len(bc.Blocks) <= AdjustmentInterval {
		return
	}

	latestBlock := bc.Blocks[len(bc.Blocks)-1]

	previousAdjustmentBlock :=
		bc.Blocks[len(bc.Blocks)-AdjustmentInterval]

	// Actual time taken for recent blocks
	actualTime :=
		latestBlock.Timestamp -
			previousAdjustmentBlock.Timestamp

	// Expected time
	expectedTime :=
		int64(AdjustmentInterval) *
			TargetBlockTime

		
		// Blocks are mined too quickly
	if actualTime < expectedTime/2 {

		if bc.Difficulty < MaxDifficulty {
			bc.Difficulty++
		}

		return
	}

	// Blocks are mined too slowly
	if actualTime > expectedTime*2 {

		// avoid difficulty becoming zero
		if bc.Difficulty > 1 {
			bc.Difficulty--
		}

		return
	}

}
