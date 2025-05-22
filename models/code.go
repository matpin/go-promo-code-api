package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Code struct {
    ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Code   string             `json:"code,omitempty" bson:"code,omitempty"`
    Gift  string              `json:"gift,omitempty" bson:"gift,omitempty"`
    ExpireDate    string      `json:"expire_date,omitempty" bson:"expire_date,omitempty"`
}