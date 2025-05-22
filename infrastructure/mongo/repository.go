package mongo

import (
	"context"
	"go-promo-code-api/domain"
	"go-promo-code-api/infrastructure/models"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) domain.Repository {
	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	return &mongoRepo{collection: collection}
}

func (r *mongoRepo) FindAll() ([]models.Code, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var codes []models.Code
	for cursor.Next(ctx) {
		var code models.Code
		if err := cursor.Decode(&code); err != nil {
			return nil, err
		}
		codes = append(codes, code)
	}
	return codes, nil
}

func (r *mongoRepo) Insert(code models.Code) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, code)
	return err
}

func (r *mongoRepo) Update(id primitive.ObjectID, code models.Code) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": code}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepo) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
