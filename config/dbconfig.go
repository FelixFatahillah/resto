package config

import (
	"fmt"
	"log"
	"os"

	"github.com/FelixFatahillah/resto/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error config ENV")
	}

	db_username := os.Getenv("MYSQL_USER")
	db_password := os.Getenv("MYSQL_PASSWORD")
	db_host := os.Getenv("MYSQL_HOST")
	db_port := os.Getenv("MYSQL_PORT")
	db_name := os.Getenv("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_name)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed connect to database")
	}

	database.AutoMigrate(&model.Ingredient{})
	database.AutoMigrate(&model.MenuRecipe{})
	database.AutoMigrate(&model.DetailRecipe{})

	DB = database
	return database
}
