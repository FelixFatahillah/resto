package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelixFatahillah/resto/controller"
	"github.com/FelixFatahillah/resto/dtos"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddIngredient(t *testing.T) {
	r := gin.Default()

	r.POST("/add-ingredient", controller.AddIngredient)

	payload := `{"name": "cabe merah", "measurement": "pcs"}`
	req, _ := http.NewRequest("POST", "/add-ingredient", bytes.NewBuffer([]byte(payload)))

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var response dtos.MediaDto
	json.Unmarshal(res.Body.Bytes(), &response)

	assert.Equal(t, "success", response.Status)
	assert.Equal(t, "bahan berhasil ditambahkan", response.Message)
	assert.Equal(t, "salt", response.Data.(map[string]interface{})["name"])
	assert.Equal(t, "teaspoon", response.Data.(map[string]interface{})["measurement"])
}

func TestAddIngredient_InvalidPayload(t *testing.T) {
	r := gin.Default()

	r.POST("/ingredient", controller.AddIngredient)

	payload := `{"measurement": "teaspoon"}`
	req, _ := http.NewRequest("POST", "/ingredient", bytes.NewBuffer([]byte(payload)))

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)

	var response dtos.DtoException
	json.Unmarshal(res.Body.Bytes(), &response)

	assert.Equal(t, "failure", response.Status)
	assert.Equal(t, "input anda salah", response.Message)
}

func TestAddIngredient_ExistingIngredient(t *testing.T) {
	r := gin.Default()

	r.POST("/ingredient", controller.AddIngredient)

	payload := `{"name": "salt", "measurement": "teaspoon"}`
	req, _ := http.NewRequest("POST", "/ingredient", bytes.NewBuffer([]byte(payload)))

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotAcceptable, res.Code)

	var response dtos.DtoException
	json.Unmarshal(res.Body.Bytes(), &response)

	assert.Equal(t, "failure", response.Status)
	assert.Equal(t, "bahan sudah terdaftar", response.Message)
}
