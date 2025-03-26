package initializer

import (
	"fmt"
	"log"
)

func MigrateTables() {
	err := DB.AutoMigrate()
	if err != nil {
		log.Fatal("Migration failed")
	}
	fmt.Println("Table migrate successfully")
}
