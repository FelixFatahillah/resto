package response

type IngredientResponse struct {
	ID          uint   `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Measurement string `gorm:"column:measurement" json:"measurement"`
}
