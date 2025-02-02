package controllers

import (
	"DvdRental/Config"
	"DvdRental/Models"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	// "log"
)

func GetAllActors(c *gin.Context) {
	var actor []models.Actor
	log.Println("Getting all actors")
	config.DB.Find(&actor)

	c.JSON(200, gin.H{
		"message": actor,
	})
}

func GetActors(c *gin.Context) {
	var actor models.Actor
	name := strings.Split(c.Param("actor"), " ")
	
	log.Println("Getting actor: ", name)

	config.DB.Where("actor.first_name=? AND actor.last_name =?", name[0], name[1]).
		Find(&actor)

	c.JSON(200, gin.H{
		"message": actor,
	})
}

func GetActorsByFilm(c *gin.Context) {
	var actors []models.Actor
	title := c.Query("film")
	log.Println("Getting actors by film: ", title)

	config.DB.Table("actor").
		Select("actor.first_name, actor.last_name").
		Joins("JOIN film_actor ON actor.actor_id = film_actor.actor_id").
		Joins("JOIN film ON film_actor.film_id = film.film_id").
		Where("film.title = ?", title).
		Find(&actors)

	var actorResponses []models.ActorResponse
	for _, actor := range actors {
		actorResponses = append(actorResponses, models.ActorResponse{
			FirstName: actor.FirstName,
			LastName: actor.LastName,
		})
	}

	c.JSON(200, gin.H{
		"film": title,
		"actors": actorResponses,
	})
}