package models

import "time"

type CragDto struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	Access            string      `json:"access"`
	Photos            []string    `json:"photos"`
	RockTypes         []string    `json:"rockTypes"`
	Orientation       []string    `json:"orientation"`
	Routes            []string    `json:"routes"`
	SuggestedPeriod   []string    `json:"suggestedPeriod"`
	Municipality      []string    `json:"municipality"`
	AverageGrade      []string    `json:"averageGrade"`
	LatitudeLongitude PointDto    `json:"point"`
	Geometry          GeometryDto `json:"geometry"`
	Story             string      `json:"story"`
	ExternalResources []string    `json:"externalResources"`
	Creators          []string    `json:"creators"`
	Record            RecordDto   `json:"record"`
	Tags              []string    `json:"tags"`
}

type PointDto struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RecordDto struct {
	CreatedBy     string    `json:"createdBy"`
	LastUpdatedBy string    `json:"lastUpdatedBy"`
	LastUpdate    time.Time `json:"lastUpdate"`
	CreatedDate   time.Time `json:"createdDate"`
}

type GeometryDto struct {
	Polygon []PointDto `json:"polygon"`
}

type RoutesDto struct {
	Name              string   `json:"name"`
	Grade             string   `json:"grade"`
	Length            int      `json:"length"`
	Description       string   `json:"description"`
	Media             []string `json:"media"`
	ClimbingStyle     string   `json:"climbingStyle"`
	LatitudeLongitude PointDto `json:"point"`
	Story             string   `json:"polygon"`
}
