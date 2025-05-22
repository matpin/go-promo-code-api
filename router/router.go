package router

import (
	"go-promo-code-api/app"
	handlers "go-promo-code-api/http"
	"go-promo-code-api/infrastructure/mongo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CodeRouter() http.Handler {
	client, err := mongo.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	repo := mongo.NewMongoRepository(client)

	service := app.NewCodeService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/codes", handler.GetAllCodes).Methods("GET")
	router.HandleFunc("/code", handler.InsertCode).Methods("POST")
	router.HandleFunc("/code/{id}", handler.UpdateCode).Methods("PUT")
	router.HandleFunc("/code/{id}", handler.DeleteCode).Methods("DELETE")

	return router
}
