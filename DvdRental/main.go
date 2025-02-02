package main

import (
	"DvdRental/Routers"
	"log"
	"github.com/gin-gonic/gin"
	"DvdRental/Config"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	config.ConnectDB()
	r := gin.Default()

	log.Println("Server is running on port 8080")
	
	routers.SetupFilmRouters(r)
	routers.SetupActorRouters(r)
	
	r.Run(":8080")
}