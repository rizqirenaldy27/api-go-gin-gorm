package models

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func DatabaseConfig() {

	// decalare setup Database
	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	DBURL := DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName
	database, err := gorm.Open(Dbdriver, DBURL)

	if err != nil {
		log.Println("This is the error:", err)
	} else {
		log.Printf("Database %s Connected", Dbdriver)
	}

	database.AutoMigrate(&Request{})

	DB = database
}
