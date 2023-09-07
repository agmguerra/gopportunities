package router

import (
	"github.com/agmguerra/gopportunities/docs"
	"github.com/agmguerra/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	//Initilize handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
	}

	// Initialize swagger
	docs.SwaggerInfo.BasePath = basePath
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
