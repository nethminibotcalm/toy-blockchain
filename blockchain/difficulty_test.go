package blockchain_test

import (
	"testing"

	"toy-blockchain/block"
	"toy-blockchain/blockchain"
)


func TestDifficultyIncrease(t *testing.T) {

	bc := blockchain.NewBlockchain()


	bc.Difficulty = 4


	// Simulate blocks mined too quickly
	bc.Blocks = []block.Block{

		{
			Index:      0,
			Timestamp: 1000,
			Difficulty: 4,
		},

		{
			Index:      1,
			Timestamp: 1001,
			Difficulty: 4,
		},

		{
			Index:      2,
			Timestamp: 1002,
			Difficulty: 4,
		},

		{
			Index:      3,
			Timestamp: 1003,
			Difficulty: 4,
		},

		{
			Index:      4,
			Timestamp: 1004,
			Difficulty: 4,
		},

		{
			Index:      5,
			Timestamp: 1005,
			Difficulty: 4,
		},
	}


	oldDifficulty := bc.Difficulty


	bc.AdjustDifficulty()


	if bc.Difficulty <= oldDifficulty {

		t.Fatal(
			"difficulty should increase when blocks are too fast",
		)
	}

}
func TestDifficultyDecrease(t *testing.T) {

	bc := blockchain.NewBlockchain()

	bc.Difficulty = 5


	bc.Blocks = []block.Block{

		{
			Index:      0,
			Timestamp: 1000,
			Difficulty: 5,
		},

		{
			Index:      1,
			Timestamp: 1050,
			Difficulty: 5,
		},

		{
			Index:      2,
			Timestamp: 1100,
			Difficulty: 5,
		},

		{
			Index:      3,
			Timestamp: 1150,
			Difficulty: 5,
		},

		{
			Index:      4,
			Timestamp: 1200,
			Difficulty: 5,
		},

		{
			Index:      5,
			Timestamp: 1250,
			Difficulty: 5,
		},
	}


	oldDifficulty := bc.Difficulty


	bc.AdjustDifficulty()


	if bc.Difficulty >= oldDifficulty {

		t.Fatal(
			"difficulty should decrease when blocks are too slow",
		)
	}

}