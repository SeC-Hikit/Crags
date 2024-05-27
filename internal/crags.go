package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hikitapp.org/crags/internal/configs"
	"hikitapp.org/crags/internal/models"
)

type CragsService interface {
	GetAllCrags() []models.CragDto
	GetCrag(id string) models.CragDto
	InsertCrag(dto models.CragDto)
}

func (c CragsManager) GetAllCrags() []models.CragDto {
	collection := c.Store.Database.Collection("Crags")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return []models.CragDto{}
	}

	var fetched []models.Crag
	if err = cursor.All(context.TODO(), &fetched); err != nil {
		panic(err)
	}

	// Prints the fetched of the find operation as structs
	var results []models.CragDto

	for _, result := range fetched {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
		results = append(results,
			models.CragDto{ID: result.ID.Hex(), Name: result.Name, Picture: result.Picture})
	}

	return results
}

func (c CragsManager) InsertCrag(dto models.CragDto) {
	crag := models.Crag{Name: dto.Name, Picture: dto.Picture}
	collection := c.Store.Database.Collection("Crags")
	_, err := collection.InsertOne(context.TODO(), bson.D{{"name", crag.Name}, {"picture", crag.Picture}})
	if err != nil {
		// TODO manage errors
	}
}

func (c CragsManager) GetCrag(id string) models.CragDto {
	collection := c.Store.Database.Collection("Crags")
	hex, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		//	TODO wtf with error handling
	}

	one := collection.FindOne(context.TODO(), bson.D{{"_id", hex}})

	var foundCrag models.Crag
	err := one.Decode(&foundCrag)
	if err != nil {
		panic("Issue with decoding from DB")
	}

	//TODO implement this
	panic("Not implemented!")
	//return models.CragDto{
	//	ID:      foundCrag.ID.Hex(),
	//	Name:    foundCrag.Name,
	//	Picture: foundCrag.Picture}
}

type CragsManager struct {
	Store configs.MongoDatabase
}
