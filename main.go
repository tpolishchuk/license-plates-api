package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"license-plates-api/plate"
)

func initDB() *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_PORT_CONTAINER"),
		os.Getenv("MYSQL_DATABASE"))

	db, err := gorm.Open("mysql", dbURL)
	
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&plate.Plate{})

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	plateAPI := InitPlateAPI(db)

	r := gin.Default()

	r.GET("/plates", plateAPI.FindAll)
	r.GET("/plates/:id", plateAPI.FindByID)
	r.POST("/plates", plateAPI.Create)
	r.PUT("/plates/:id", plateAPI.Update)
	r.DELETE("/plates/:id", plateAPI.Delete)

	err := r.Run()
	
	if err != nil {
		panic(err)
	}
}
