package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelixFatahillah/resto/controller"
	"github.com/FelixFatahillah/resto/dtos"
	"github.com/FelixFatahillah/resto/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddMenu(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/add-menu", controller.AddMenu)

	payload := []byte(`{
		"name": "rendang",
		"category": "main",
		"recipe": [
			{
				"ingredient": "cabe",
				"quantity": 2
			},
			{
				"ingredient": "bawang",
				"quantity": 2
			}
		]
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/add-menu", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dtos.MediaDto
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "success", response.Status)
	assert.Equal(t, "Menu berhasil ditambahkan", response.Message)

	menu := response.Data.(model.MenuRecipe)
	assert.Equal(t, "rendang", menu.Name)
	assert.Equal(t, "main", menu.Category)
}
