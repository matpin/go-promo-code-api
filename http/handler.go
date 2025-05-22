package http

import (
	"encoding/json"
	"go-promo-code-api/app"
	"go-promo-code-api/infrastructure/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	Service *app.CodeService
}

type UpdateCodeResponse struct {
	Message string `json:"message,omitempty"`
}

func NewHandler(service *app.CodeService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetAllCodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	codes, err := h.Service.GetAllCodes()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(codes)
}

func (h *Handler) InsertCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var code models.Code
	if err := json.NewDecoder(r.Body).Decode(&code); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	err := h.Service.InsertCode(code)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	res := UpdateCodeResponse{Message: "Code inserted successfully"}
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) UpdateCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)["id"]

	id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateCode models.Code
	if err := json.NewDecoder(r.Body).Decode(&updateCode); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	err = h.Service.UpdateCode(id, updateCode)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	res := UpdateCodeResponse{Message: "Code updated successfully"}
    json.NewEncoder(w).Encode(res)
}

func (h *Handler) DeleteCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)["id"]

	id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteCode(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	res := UpdateCodeResponse{Message: "Code deleted successfully"}
	json.NewEncoder(w).Encode(res)
}
