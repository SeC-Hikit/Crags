package mapper

import "hikitapp.org/crags/internal/models"

type CragMapper interface {
	toDto(crag models.Crag) models.CragDto
	toEntity(crag models.CragDto) models.Crag
}

type CragMapperImpl struct {
}

func (c CragMapperImpl) toDto(crag models.Crag) models.CragDto {
	var routesDto []models.RouteDto
	for _, route := range crag.Routes {
		routesDto = append(routesDto, models.RouteDto{
			Name:              route.Name,
			Grade:             route.Grade,
			Length:            route.Length,
			Description:       route.Description,
			Media:             route.Media,
			ClimbingStyle:     route.ClimbingStyle,
			LatitudeLongitude: route.LatitudeLongitude,
			Story:             route.Story,
		})
	}

	return models.CragDto{
		ID:                crag.ID.Hex(),
		Name:              crag.Name,
		Description:       crag.Description,
		Access:            crag.Access,
		Photos:            crag.Photos,
		RockTypes:         crag.RockTypes,
		Orientation:       crag.Orientation,
		Routes:            routesDto,
		SuggestedPeriod:   crag.SuggestedPeriod,
		Municipality:      crag.Municipality,
		AverageGrade:      crag.AverageGrade,
		LatitudeLongitude: crag.LatitudeLongitude,
		Geometry:          models.GeometryDto{Polygon: crag.Geometry.Polygon},
		Story:             crag.Story,
		ExternalResources: crag.ExternalResources,
		Creators:          crag.Creators,
		Record: models.RecordDto{CreatedBy: crag.Record.CreatedBy,
			LastUpdatedBy: crag.Record.LastUpdatedBy, LastUpdate: crag.Record.LastUpdate,
			CreatedDate: crag.Record.CreatedDate},
		Tags: crag.Tags,
	}
}

func (c CragMapperImpl) toEntity(crag models.CragDto) models.Crag {
	//TODO implement me
	panic("implement me")
}
