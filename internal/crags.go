package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"hikitapp.org/crags/internal/configs"
	"hikitapp.org/crags/internal/models"
)

type CragsOperator interface {
	GetAllCrags(MongoDatabase configs.MongoDatabase) []models.CragDto
	GetCrag(MongoDatabase configs.MongoDatabase, id string) models.CragDto
}

func (c CragsManager) GetAllCrags() []models.CragDto {
	collection := c.Store.Client.Database("hikit").Collection("Crags")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	count, err := collection.EstimatedDocumentCount(context.TODO())
	fmt.Println(count)
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
		results = append(results, models.CragDto{ID: result.ID.Hex(),
			Name: result.Name, Picture: result.Picture})
	}

	return results

	//return []models.CragDto{
	//	{ID: "1234", Name: "my-name", Picture: "my-pic"},
	//}
	//return fetched
}

func (c CragsManager) GetCrag(mongodb configs.MongoDatabase, id string) models.CragDto {
	return models.CragDto{ID: "1234", Name: "my-name", Picture: "my-pic"}
}

type CragsManager struct {
	Store configs.MongoDatabase
}
