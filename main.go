package main

import (
	"goProject3/config"
	"goProject3/models"
	"goProject3/routes"

	"github.com/go-playground/validator/v10"

	docs "goProject3/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /

func main() {
	models.Connection = models.GetSingleInstance()
	defer config.CloseDatabase(models.Connection)

	models.Validate = validator.New()

	docs.SwaggerInfo.BasePath = "/"

	router := routes.SetUpRouter(models.Connection)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":3000")
}
