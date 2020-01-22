package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Type      string
	Firstname string
	Lastname  string
	Email     string
}
