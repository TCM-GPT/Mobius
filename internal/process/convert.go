package process

import (
	"gen-SFT-Dataset/internal/model"
	"gen-SFT-Dataset/mysql"
)

type K struct {
	Instruction string `json:"instruction"`
	Input       string `json:"input"`
	Output      string `json:"output"`
}

// Convert a content to JSON type and insert in to database.
func Convert(generateContext string) error {

	jsonArray, err := ExtractJSONArray(generateContext)
	if err != nil {
		return err
	}

	for _, object := range jsonArray {
		data := model.Dataset{
			Instruction: object.Instruction,
			Input:       object.Input,
			Output:      object.Output,
		}
		err = mysql.Transaction(data.Instruction, data.Input, data.Output)
		if err != nil {
			return err
		}
	}
	return nil
}
