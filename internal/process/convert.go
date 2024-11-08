package process

import (
	"gen-SFT-Dataset/config"
	"gen-SFT-Dataset/internal/model"
	"gen-SFT-Dataset/mysql"
	"gen-SFT-Dataset/postgresql"
)

type K struct {
	Instruction string `json:"instruction"`
	InputVal    string `json:"inputVal"`
	OutputVal   string `json:"outputVal"`
	LlmSource   string `json:"llmSource"`
}

// Convert a content to JSON type and insert in to database.
func Convert(generateContext, models string) error {

	jsonArray, err := ExtractJSONArray(generateContext)
	if err != nil {
		return err
	}

	for _, object := range jsonArray {
		data := model.Dataset{
			Instruction: object.Instruction,
			InputVal:    object.InputVal,
			OutputVal:   object.OutputVal,
			LlmSource:   models,
		}
		switch config.GetInstance().AppConfig.UseDb {
		case "mysql":
			err = mysql.Transaction(data.Instruction, data.InputVal, data.OutputVal, data.LlmSource)
			if err != nil {
				return err
			}
			break
		case "postgresql":
			err = postgresql.Transaction(data.Instruction, data.InputVal, data.OutputVal, data.LlmSource)
			if err != nil {
				return err
			}
			break
		default:
			err = postgresql.Transaction(data.Instruction, data.InputVal, data.OutputVal, data.LlmSource)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
