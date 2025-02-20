package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
}