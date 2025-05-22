package handlers

import (
	"context"
	"encoding/json"
	"go-promo-code-api/config"
	"go-promo-code-api/models"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCodeResponse struct {
	Message string `json:"message,omitempty"`
}

func GetAllCodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := config.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var codes []models.Code
	for cursor.Next(ctx) {
        var code models.Code
        cursor.Decode(&code)
        codes = append(codes, code)
    }

    json.NewEncoder(w).Encode(codes)
}

func InsertCode (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := config.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var code models.Code

	if err := json.NewDecoder(r.Body).Decode(&code); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	_, err = collection.InsertOne(ctx, code)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    res := UpdateCodeResponse{Message: "Code inserted successfully"}
	json.NewEncoder(w).Encode(res)
}

func UpdateCode (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := config.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	params := mux.Vars(r)
	idParam := params["id"]

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var updateCode models.Code

	if err := json.NewDecoder(r.Body).Decode(&updateCode); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	filter := bson.M{"_id": id}
    update := bson.M{"$set": updateCode}
    _, err = collection.UpdateOne(ctx, filter, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	res := UpdateCodeResponse{Message: "Code updated successfully"}
    json.NewEncoder(w).Encode(res)
}

func DeleteCode (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := config.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	params := mux.Vars(r)
	idParam := params["id"]

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	 _, err = collection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	res := UpdateCodeResponse{Message: "Code deleted successfully"}
	json.NewEncoder(w).Encode(res)
}