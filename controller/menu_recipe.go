package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/FelixFatahillah/resto/config"
	"github.com/FelixFatahillah/resto/dtos"
	"github.com/FelixFatahillah/resto/form"
	"github.com/FelixFatahillah/resto/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddMenu(c *gin.Context) {
	var payload form.MenuRecipe

	// get body json
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, dtos.DtoException{
			Status:  "failure",
			Message: "input anda salah",
		})
		return
	}

	name := strings.ToLower(payload.Name)
	category := strings.ToLower(payload.Category)

	data := config.DB.Table("menu_recipes").Where("name = ?", name).First(&payload)

	if data.Error != nil {
		if data.Error == gorm.ErrRecordNotFound {
			add_data := model.MenuRecipe{
				Name:     name,
				Category: category,
			}
			config.DB.Create(&add_data)

			for _, recipe_data := range payload.Recipe {
				add_detail_recipe := model.DetailRecipe{
					Menu:           name,
					IngredientName: recipe_data.Ingredient,
					Quantity:       recipe_data.Quantity,
				}

				err := config.DB.Table("detail_recipes").Create(&add_detail_recipe).Error
				if err != nil {
					log.Fatal(err)
				}
			}

			c.JSON(http.StatusOK, dtos.MediaDto{
				Status:  "success",
				Message: "Menu berhasil ditambahkan",
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
			Message: "menu sudah terdaftar",
		})
	}
}

func ShowMenu(c *gin.Context) {
	var menu []model.MenuRecipe

	config.DB.Table("menu_recipes").Where("deleted_at IS NULL").Find(&menu)

	if len(menu) == 0 {
		c.JSON(http.StatusNotFound, dtos.DtoException{
			Status:  "failure",
			Message: "bahan tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, dtos.MediaDto{
		Status:  "success",
		Message: "Menu ditemukan",
		Data:    menu,
	})
}

func ShowDetailMenu(c *gin.Context) {
	var result []model.DetailRecipe
	config.DB.Table("detail_recipes").Where("deleted_at IS NULL AND menu = ?", c.Param("menu")).Find(&result)

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, dtos.DtoException{
			Status:  "failure",
			Message: "bahan tidak ditemukan",
		})
		return
	}

	recipe := struct {
		Menu   string `json:"menu"`
		Recipe []struct {
			Ingredient string `json:"ingredient"`
			Quantity   int    `json:"quantity"`
		} `json:"recipe"`
	}{
		Menu: c.Param("menu"),
	}

	for _, detail := range result {
		recipe.Recipe = append(recipe.Recipe, struct {
			Ingredient string `json:"ingredient"`
			Quantity   int    `json:"quantity"`
		}{
			Ingredient: detail.IngredientName,
			Quantity:   int(detail.Quantity),
		})
	}

	c.JSON(http.StatusOK, dtos.MediaDto{
		Status:  "success",
		Message: "detail resep ditemukan",
		Data:    recipe,
	})
}

func SearchMenu(c *gin.Context) {
	var result []model.MenuRecipe

	config.DB.Table("menu_recipes").Where("deleted_at IS NULL AND name LIKE ?", "%"+c.Param("name")+"%").Find(&result)

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, dtos.DtoException{
			Status:  "failure",
			Message: "menu tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, dtos.MediaDto{
		Status:  "success",
		Message: "menu ditemukan",
		Data:    result,
	})
}

func EditMenu(c *gin.Context) {
	var result model.MenuRecipe
	if err := config.DB.Table("menu_recipes").Where("id = ?", c.Param("id")).First(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, dtos.DtoException{
			Status:  "failure",
			Message: "bahan tidak ditemukan",
		})
		return
	}

	var payload form.MenuRecipe

	// get body json
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, dtos.DtoException{
			Status:  "failure",
			Message: "input anda salah",
		})
		return
	}

	if err := config.DB.Table("detail_recipes").Where("menu = ?", result.Name).Delete(&model.DetailRecipe{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, dtos.DtoException{
			Status:  "failure",
			Message: "gagal update data",
		})
		return
	}

	config.DB.Model(&result).Updates(payload)

	for _, recipe_data := range payload.Recipe {
		add_detail_recipe := model.DetailRecipe{
			Menu:           payload.Name,
			IngredientName: recipe_data.Ingredient,
			Quantity:       recipe_data.Quantity,
		}

		err := config.DB.Table("detail_recipes").Create(&add_detail_recipe).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(http.StatusOK, dtos.MediaDto{
		Status:  "success",
		Message: "menu ditemukan",
		Data:    result,
	})
}

func DeleteMenu(c *gin.Context) {
	var result model.MenuRecipe
	if err := config.DB.Table("menu_recipes").Unscoped().Where("id = ?", c.Param("id")).First(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, dtos.DtoException{
			Status:  "failure",
			Message: "menu tidak ditemukan",
		})
		return
	}

	if err := config.DB.Table("detail_recipes").Where("menu = ?", result.Name).Delete(&model.DetailRecipe{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, dtos.DtoException{
			Status:  "failure",
			Message: "gagal delete data",
		})
		return
	}

	config.DB.Delete(&result)

	c.JSON(http.StatusOK, dtos.MediaDto{
		Status:  "success",
		Message: "menu berhasil dihapus",
		Data:    result,
	})
}
