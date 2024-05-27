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
	Routes            []RouteDto  `json:"routes"`
	SuggestedPeriod   []string    `json:"suggestedPeriod"`
	Municipality      []string    `json:"municipality"`
	AverageGrade      []string    `json:"averageGrade"`
	LatitudeLongitude Point       `json:"point"`
	Geometry          GeometryDto `json:"geometry"`
	Story             string      `json:"story"`
	ExternalResources []string    `json:"externalResources"`
	Creators          []string    `json:"creators"`
	Record            RecordDto   `json:"record"`
	Tags              []string    `json:"tags"`
}

type RecordDto struct {
	CreatedBy     string    `json:"createdBy"`
	LastUpdatedBy string    `json:"lastUpdatedBy"`
	LastUpdate    time.Time `json:"lastUpdate"`
	CreatedDate   time.Time `json:"createdDate"`
}

type GeometryDto struct {
	Polygon []Point `json:"polygon"`
}

type RouteDto struct {
	Name              string   `json:"name"`
	Grade             string   `json:"grade"`
	Length            int      `json:"length"`
	Description       string   `json:"description"`
	Media             []string `json:"media"`
	ClimbingStyle     string   `json:"climbingStyle"`
	LatitudeLongitude Point    `json:"point"`
	Story             string   `json:"polygon"`
}
