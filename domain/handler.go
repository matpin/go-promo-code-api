package domain

import (
	"net/http"
)

type Handler interface {
	GetAllCodes(w http.ResponseWriter, r *http.Request)
	InsertCode(w http.ResponseWriter, r *http.Request)
	UpdateCode(w http.ResponseWriter, r *http.Request)
	DeleteCode(w http.ResponseWriter, r *http.Request)
}
