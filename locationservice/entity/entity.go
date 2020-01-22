package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	User    primitive.ObjectID `bson:"user, omitempty"`
	Latitude            string `json:"latitude"  binding:"required"`
	Longitude       string `json:"longitude" binding:"required"`
}

type LocationDB struct {
	ID        string `bson:"_id, omitempty"`
	User    primitive.ObjectID `bson:"user, omitempty"`
	Latitude            string `json:"latitude"  binding:"required"`
	Longitude       string `json:"longitude" binding:"required"`
}

type Response struct {
	Data    LocationDB `json: "data"`
	Message string `json:"message"`
}
