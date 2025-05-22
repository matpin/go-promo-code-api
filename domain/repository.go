package domain

import (
	"go-promo-code-api/infrastructure/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	FindAll() ([]models.Code, error)
	Insert(code models.Code) error
	Update(id primitive.ObjectID, code models.Code) error
	Delete(id primitive.ObjectID) error
}
