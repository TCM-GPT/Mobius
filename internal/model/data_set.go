package model

import "gorm.io/gorm"

type Dataset struct {
	*gorm.Model
	Instruction string `json:"instruction"`
	Input       string `json:"input"`
	Output      string `json:"output"`
}

func (d *Dataset) TableName() string {
	return "data_sets"
}
