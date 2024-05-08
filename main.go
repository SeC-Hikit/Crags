package main

import (
	"fmt"
	"hikitapp.org/crags/internal"
	"hikitapp.org/crags/internal/configs"
	"hikitapp.org/crags/internal/models"
	"net/http"
)
import "github.com/gin-gonic/gin"

const (
	APIBasePath = "/api/v1"
)

func main() {
	r := gin.Default()
	db := configs.NewMongoDatabase()
	cragsManager := internal.CragsManager{Store: db}
	r.GET(fmt.Sprintf("%s/crags", APIBasePath),
		func(context *gin.Context) { getAllCragsResponse(context, cragsManager) })
	r.GET(fmt.Sprintf("%s/crags/:id", APIBasePath),
		func(context *gin.Context) { getCrag(context, cragsManager) })
	r.POST(fmt.Sprintf("%s/crags", APIBasePath),
		func(context *gin.Context) { insertCrag(context, cragsManager) })

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getAllCragsResponse(c *gin.Context, cragsManager internal.CragsManager) {
	c.IndentedJSON(http.StatusOK, cragsManager.GetAllCrags())
}

func getCrag(c *gin.Context, cragsManager internal.CragsManager) {
	c.IndentedJSON(http.StatusOK, cragsManager.GetCrag(c.Param("id")))
}

func insertCrag(context *gin.Context, manager internal.CragsManager) {
	var dto models.CragDto
	err := context.BindJSON(&dto)
	if err != nil {
		// TODO
	}
	manager.InsertCrag(dto)
	context.Done()
}
