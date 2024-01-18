package main

import (
	"GoMon/api"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	// Log error if .env file does not exist
	if err := godotenv.Load("envs/.env"); err != nil {
		fmt.Println("No .env file found")
		panic("No .env file found")
	}

	//Set mode in init function
	gin.SetMode(os.Getenv("GIN_MODE"))
}
func main() {

	var app *gin.Engine = api.Setup()
	log.Fatal(app.Run(":" + os.Getenv("SERVERPORT")))
	fmt.Println("Serving.... ")
}
