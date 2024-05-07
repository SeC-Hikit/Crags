package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CragDto struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type Crag struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string
	Picture string
}

func New(id string, name string, picture string) *CragDto {
	return &CragDto{
		id, name, picture,
	}
}
