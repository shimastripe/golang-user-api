package db

import (
	"log"
	"os"
	"path/filepath"

	"github.com/shimastripe/gouserapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect() *gorm.DB {
	dir := filepath.Dir("db/database.db")
	db, err := gorm.Open("sqlite3", dir+"/database.db")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.LogMode(false)
	if os.Getenv("DB") == "DEBUG" {
		db.LogMode(true)
	}
	if os.Getenv("AUTOMIGRATE") == "true" {
		db.AutoMigrate(&models.User{}, &models.Profile{}, &models.AccountName{}, &models.Email{}, &models.Nation{})
	}
	return db
}

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}
