package model

import "gorm.io/gorm"

type MenuRecipe struct {
	gorm.Model
	Name     string `json:"name"`
	Category string `json:"category"`
}
