package model

import "gorm.io/gorm"

type DetailRecipe struct {
	gorm.Model
	Menu           string `json:"menu" gorm:"index"`
	IngredientName string `json:"name_ingredient"`
	Quantity       uint   `json:"quantity"`
}
