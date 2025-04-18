package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
}
