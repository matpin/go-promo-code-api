package main

import (
	"fmt"
	"go-promo-code-api/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	router := mux.NewRouter()

	router.HandleFunc("/codes", handlers.GetAllCodes).Methods("GET")
	router.HandleFunc("/code", handlers.InsertCode).Methods("POST")
	router.HandleFunc("/code/{id}", handlers.UpdateCode).Methods("PUT")
	router.HandleFunc("/code/{id}", handlers.DeleteCode).Methods("DELETE")

	port := os.Getenv("PORT")

	fmt.Println("Server running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}