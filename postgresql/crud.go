package postgresql

import (
	"gen-SFT-Dataset/internal/model"
	"gorm.io/gorm"
)

func Transaction(instruction, inputVal, outputVal, llmSource string) error {
	data := model.Dataset{
		Instruction: instruction,
		InputVal:    inputVal,
		OutputVal:   outputVal,
		LlmSource:   llmSource,
	}

	if err := GetInstance().
		Transaction(func(tx *gorm.DB) error {
			return tx.Create(&data).Error
		}); err != nil {
		return err
	}
	return nil
}
