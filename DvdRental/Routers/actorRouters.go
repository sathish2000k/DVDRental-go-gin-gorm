package routers

import (
	"github.com/gin-gonic/gin"
	"DvdRental/Controllers"
)

func SetupActorRouters (router *gin.Engine) {
	actorRouter := router.Group("/actor")
	{
		actorRouter.GET("/all", controllers.GetAllActors)
		actorRouter.GET("/film", controllers.GetActorsByFilm)
		actorRouter.GET("/:actor", controllers.GetActors)
	}
}