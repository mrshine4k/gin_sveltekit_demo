package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	Id     primitive.ObjectID `bson:"id,omitempty"`
	Title  string             `bson:"Title,omitempty" validate:"required"`
	Artist string             `bson:"Artist,omitempty" validate:"required"`
	Price  float64            `bson:"Price,omitempty" validate:"required"`
}
