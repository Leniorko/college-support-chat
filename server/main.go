package main

import (
	"college-support-chat-backend/middleware/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	appDatabase := db.Init()
	log.Default().Println(appDatabase)
}
