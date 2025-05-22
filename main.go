package main

import (
	"fmt"
	"go-promo-code-api/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	r := router.CodeRouter()

	port := os.Getenv("PORT")

	fmt.Println("Server running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
