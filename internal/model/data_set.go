package model

import (
	"gorm.io/gorm"
)

type Dataset struct {
	gorm.Model
	Instruction string
	InputVal    string
	OutputVal   string
	LlmSource   string
}

func (d *Dataset) TableName() string {
	return "data_sets"
}
