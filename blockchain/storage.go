package blockchain

import (
	"encoding/json"
	"os"
)

func (bc *Blockchain) SaveToFile(filename string) error {

	data, err := json.MarshalIndent(bc, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
func LoadFromFile(filename string) (*Blockchain, error) {

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var bc Blockchain

	err = json.Unmarshal(data, &bc)

	if err != nil {
		return nil, err
	}

	return &bc, nil
}