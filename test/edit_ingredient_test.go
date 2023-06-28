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

func TestEditIngredient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/edit-ingredient/:id", controller.EditIngredient)

	// Membuat payload JSON untuk pengujian
	payload := []byte(`{
		"name": "cabe",
		"measurement": "gram"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/edit-ingredient/1", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dtos.MediaDto
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "success", response.Status)
	assert.Equal(t, "bahan berhasil di update", response.Message)

	ingredient := response.Data.(model.Ingredient)
	assert.Equal(t, "cabe", ingredient.Name)
	assert.Equal(t, "gram", ingredient.Measurement)
}
