package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"tesjwt.go/database"
	_ "tesjwt.go/docs"
	"tesjwt.go/router"
)

func main() {
	if os.Getenv("ENV") != "prd" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("errpr loading .env ")
		}
	}
	database.StartDB()
	r := router.StartApp()
	log.Println("starting app...")
	r.Run(":5000")
}
