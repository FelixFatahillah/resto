package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelixFatahillah/resto/controller"
	"github.com/FelixFatahillah/resto/dtos"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSearchIngredient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/search-ingredient/:name", controller.SearchIngredient)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search-ingredient/cabe", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response dtos.DtoException
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "failure", response.Status)
	assert.Equal(t, "bahan tidak ditemukan", response.Message)
}
