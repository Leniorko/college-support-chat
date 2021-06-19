package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	log.Default().Println("Connectiont to database established!")

	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Println("Database pinged without errors!")

	return db
}

func checkDatabaseTables(db *gorm.DB) {

}
