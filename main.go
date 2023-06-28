package main

import (
	"log"
	"os"

	"github.com/FelixFatahillah/resto/config"
	"github.com/FelixFatahillah/resto/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error load ENV")
	}

	port := os.Getenv("PORT")
	r := gin.Default()

	// Setup db
	config.SetupDatabaseConnection()

	r.GET("/show-ingredient", controller.ShowIngredient)
	r.POST("/add-ingredient", controller.AddIngredient)
	r.GET("/search-ingredient/:name", controller.SearchIngredient)
	r.PUT("/edit-ingredient/:id", controller.EditIngredient)
	r.DELETE("/delete-ingredient/:id", controller.DeleteIngredient)

	r.POST("/add-menu", controller.AddMenu)
	r.GET("/show-menu", controller.ShowMenu)
	r.GET("/detail-menu/:menu", controller.ShowDetailMenu)
	r.GET("/search-menu/:name", controller.SearchMenu)
	r.PUT("/edit-menu/:id", controller.EditMenu)
	r.DELETE("/delete-menu/:id", controller.DeleteMenu)

	log.Fatal(r.Run(":" + port))
}
