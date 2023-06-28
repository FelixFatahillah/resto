package controller

import (
	"net/http"
	"strings"

	"github.com/FelixFatahillah/resto/config"
	"github.com/FelixFatahillah/resto/controller/response"
	"github.com/FelixFatahillah/resto/dtos"
	"github.com/FelixFatahillah/resto/form"
	"github.com/FelixFatahillah/resto/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddIngredient(c *gin.Context) {
	var payload form.Ingredient

	// get body json
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, dtos.DtoException{
			Status:  "failure",
			Message: "input anda salah",
		})
		return
	}
	name := strings.ToLower(payload.Name)
	measurement := strings.ToLower(payload.Measurement)

	data := config.DB.Table("ingredients").Where("name = ?", name).First(&payload)
	if data.Error != nil {
		if data.Error == gorm.ErrRecordNotFound {
			add_data := model.Ingredient{
				Name:        name,
				Measurement: measurement,
			}
			config.DB.Create(&add_data)

			c.JSON(http.StatusOK, dtos.MediaDto{
				Status:  "success",
				Message: "bahan berhasil ditambahkan",
				Data:    add_data,
			})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.DtoException{
				Status:  "failure",
				Message: "internal server error",
			})
		}
	} else {
		c.JSON(http.StatusNotAcceptable, dtos.DtoException{
			Status:  "failure",
			Message: "bahan sudah terdaftar",
		})
	}
}

func ShowIngredient(c *gin.Context) {
	var result []response.IngredientResponse

	config.DB.Table("ingredients").Where("deleted_at IS NULL").Find(&result)

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, dtos.DtoException{
			Status:  "failure",
			Message: "bahan tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, dtos.MediaDto{
		Status:  "success",
		Message: "bahan ditemukan",
		Data:    result,
	})
}

func SearchIngredient(c *gin.Context) {
	var result []response.IngredientResponse

	config.DB.Table("ingredients").Where("deleted_at IS NULL AND name LIKE ?", "%"+c.Param("name")+"%").Find(&result)

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, dtos.DtoException{
			Status:  "failure",
			Message: "bahan tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, dtos.MediaDto{
		Status:  "success",
		Message: "bahan ditemukan",
		Data:    result,
	})
}

func EditIngredient(c *gin.Context) {
	var result model.Ingredient
	if err := config.DB.Table("ingredients").Where("id = ?", c.Param("id")).First(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, dtos.DtoException{
			Status:  "failure",
			Message: "bahan tidak ditemukan",
		})
		return
	}

	var input form.Ingredient
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dtos.DtoException{
			Status:  "failure",
			Message: "input anda salah",
		})
		return
	}

	config.DB.Model(&result).Updates(input)

	c.JSON(http.StatusOK, dtos.MediaDto{
		Status:  "success",
		Message: "bahan berhasil di update",
		Data:    result,
	})

}

func DeleteIngredient(c *gin.Context) {
	var result model.Ingredient
	if err := config.DB.Table("ingredients").Unscoped().Where("id = ?", c.Param("id")).First(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, dtos.DtoException{
			Status:  "failure",
			Message: "bahan tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&result)

	c.JSON(http.StatusOK, dtos.MediaDto{
		Status:  "success",
		Message: "bahan berhasil dihapus",
		Data:    result,
	})
}
