package db

import (
	"context"
	"fmt"
	"time"

	model "github.com/dev-sandip/go-auth/cmd/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitMongoCollections(client *mongo.Client) {
	userCollection = client.Database("goauth").Collection("users")
}
func CreateUser(user *model.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return result, nil
}
