package form

import "gorm.io/gorm"

type MenuRecipe struct {
	gorm.Model
	Name     string         `json:"name" binding:"required"`
	Category string         `json:"category" binding:"required"`
	Recipe   []DetailRecipe `json:"recipe" binding:"required" gorm:"foreignKey:Menu"`
}

type DetailRecipe struct {
	Menu       string `json:"menu" gorm:"index"`
	Ingredient string `json:"ingredient" binding:"required"`
	Quantity   uint   `json:"quantity" binding:"required"`
}
