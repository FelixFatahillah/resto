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

func TestSearchMenu(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/search-menu/:name", controller.SearchMenu)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search-menu/rendang", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dtos.MediaDto
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "success", response.Status)
	assert.Equal(t, "menu ditemukan", response.Message)
}
