package main

import (
	"Airport/web/app/controller"
	docs "Airport/web/app/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Disable Console Color
	// gin.DisableConsoleColor()
	// Allow cross origin
	r.Use(cors.Default())

	// Get measures
	r.GET("/airport/:airportCode/measure", controller.GetMeasures)

	// Get average
	r.GET("/airport/:airportCode/averages", controller.GetAverages)

	return r
}

func main() {
	r := setupRouter()
	docs.SwaggerInfo.BasePath = "/"
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
