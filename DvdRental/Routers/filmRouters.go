package routers

import (
	"github.com/gin-gonic/gin"
	"DvdRental/Controllers"
)

func SetupFilmRouters (router *gin.Engine) {
	filmRouter := router.Group("/film")
	{
		filmRouter.GET("/actor", controllers.GetFilmsByActor)
		filmRouter.GET("/all", controllers.GetAllFilms)
		filmRouter.GET(("/:film"), controllers.GetFilmByTitle)
		filmRouter.GET("/language", controllers.GetFilmsByLanguage)
		filmRouter.GET("/category", controllers.GetFilmsByCategory)
	}
}