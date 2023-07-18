package router

import (
	"log"

	"github.com/getground/tech-tasks/backend/handlers"
	"github.com/gin-gonic/gin"

)

var Router *gin.Engine

func errorHandlingMiddleWare(log *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		log.Printf("unexpected error: %s\n", err.Error())
	}
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", handlers.HealthHandler)
	router.GET("/guest_list", handlers.GetGuestLists)
	router.GET("/guests", handlers.GetArrivedGuestLists)
	router.GET("/seats_empty", handlers.GetEmptySeats)

	router.POST("/guest_list/:name", handlers.AddGuest)
	router.POST("/tables", handlers.AddTable)

	router.PUT("/guests/:name", handlers.CheckInGuest)

	router.DELETE("/guests/:name", handlers.DeleteGuest)

	return router
}

func Setup() {
	
	Router = SetupRouter()
}