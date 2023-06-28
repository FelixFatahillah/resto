package form

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Measurement string `json:"measurement" binding:"required"`
}
