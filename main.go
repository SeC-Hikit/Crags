package main

import (
	"fmt"
	"hikitapp.org/crags/internal"
	"hikitapp.org/crags/internal/configs"
	"net/http"
)
import "github.com/gin-gonic/gin"

const (
	APIBasePath = "/api/v1"
)

func main() {
	r := gin.Default()
	db := configs.NewMongoDatabase()
	c := internal.CragsManager{Store: db}
	r.GET(fmt.Sprintf("%s/crags", APIBasePath),
		func(context *gin.Context) { getAllCragsResponse(context, c) })
	r.GET(fmt.Sprintf("%s/crags/:id", APIBasePath), func(context *gin.Context) {

	})
	r.GET(fmt.Sprintf("%s/routes/:id", APIBasePath), func(context *gin.Context) {

	})

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getAllCragsResponse(c *gin.Context, cragsManager internal.CragsManager) {
	c.IndentedJSON(http.StatusOK, cragsManager.GetAllCrags())
}
