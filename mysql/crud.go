package mysql

import (
	"gen-SFT-Dataset/internal/model"
	"gorm.io/gorm"
)

func Transaction(instruction, input, output string) error {
	data := model.Dataset{
		Instruction: instruction,
		Input:       input,
		Output:      output,
	}

	if err := GetInstance().
		Transaction(func(tx *gorm.DB) error {
			return tx.Create(&data).Error
		}); err != nil {
		return err
	}
	return nil
}
