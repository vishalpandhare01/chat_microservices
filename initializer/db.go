package initializer

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	var dbUri = os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(mysql.Open(dbUri))
	if err != nil {
		log.Fatal("Failer to connect Databae ")
	}
	fmt.Println("database connected successfully 🚀")
}
