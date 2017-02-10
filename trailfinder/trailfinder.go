package trailfinder

import (
	"encoding/json"
	"io"
)

type Trailfinder struct {
	Checkpoints []Checkpoint
}

type Checkpoint struct {
	Direction string `json:"direction"`
	Mode      string `json:"mode"`
}

func (tf *Trailfinder) Interpret(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&tf.Checkpoints)
	if err != nil {
		return err
	}
	return nil
}

// # Team 1
// 1. Download a json file from internets
// 2. Create a data structure which matches the file
// 3. Parse the json
// 4. Sort the data structure
