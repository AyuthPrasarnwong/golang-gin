package routes

import (
	"app/controllers"
	"app/newrelic"
	"app/routes/api/v1"
	//"app/routes/api/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute(handleNoRoute)

	newrelic.Setup(r)
	v1.Setup(r)
	//v2.Setup(r)

	return r
}

func handleNoRoute(c *gin.Context)  {

	errors := controllers.Errors{
		404,
		"Page not Found",
		nil,
	}

	c.JSON(404, controllers.ErrorResponse{
		errors,
	})
}
