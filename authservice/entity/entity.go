package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Type      string             `json:"type"  binding:"required"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Email     string             `json:"email"  binding:"required"`
	Password  string             `json:"password"`
	Token     string             `json:"token"  binding:"required"`
}

type Payload struct {
	ID        primitive.ObjectID `bson:"_id"`
	Type      string             `json:"type"  binding:"required"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Email     string             `json:"email"  binding:"required"`
	Token     string             `json:"token"  binding:"required"`
}

type Login struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
