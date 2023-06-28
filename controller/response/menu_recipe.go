package response

type DetailMenuResponse struct {
	Menu           string `json:"menu" gorm:"index"`
	IngredientName string `json:"name_ingredient"`
	Quantity       uint   `json:"quantity"`
}
