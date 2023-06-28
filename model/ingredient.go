package model

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name        string `json:"name"`
	Measurement string `json:"measurement"`
}
