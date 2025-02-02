package controllers

import (
	models "DvdRental/Models"
	"log"

	"github.com/gin-gonic/gin"
	"DvdRental/Config"
)

func GetFilmsByActor(c *gin.Context) {
	var filmResponses []models.FilmResponse
	firstName := c.Query("firstName")
	lastName := c.Query("lastName")
	log.Println("Getting films by actor: ", firstName, lastName)

	config.DB.Table("film").
		Select("film.title, film.description, film.release_year, TRIM(language.name) as language, category.name as category, film.length, film.rating, film.special_features").
		Joins("JOIN film_category on film.film_id = film_category.film_id").
		Joins("JOIN category on film_category.category_id = category.category_id").
		Joins("JOIN language on language.language_id = film.language_id").
		Joins("JOIN film_actor on film.film_id = film_actor.film_id ").
		Joins("JOIN actor on film_actor.actor_id = actor.actor_id").
		Where("actor.first_name = ? AND actor.last_name = ?", firstName, lastName).
		Find(&filmResponses)

	c.JSON(200, gin.H{
		"message": filmResponses,
	})
}

func GetAllFilms(c *gin.Context) {
	var films []models.FilmResponse
	log.Println("Getting all films")
	
	config.DB.Table("film").
		Select("film.title, film.description, film.release_year, TRIM(language.name) as language, category.name as category, film.length, film.rating, film.special_features").
		Joins("JOIN film_category on film.film_id = film_category.film_id").
		Joins("JOIN category on film_category.category_id = category.category_id").
		Joins("JOIN language on language.language_id = film.language_id").
		Joins("JOIN film_actor on film.film_id = film_actor.film_id ").
		Joins("JOIN actor on film_actor.actor_id = actor.actor_id").
		Find(&films)

	c.JSON(200, gin.H{
		"message": films,
	})
}

func GetFilmByTitle(c *gin.Context) {
	var film models.FilmResponse
	title := c.Param("film")
	log.Println("Getting film by title: ", title)

	config.DB.Table("film").
		Select("film.title, film.description, film.release_year, TRIM(language.name) as language, category.name as category, film.length, film.rating, film.special_features").
		Joins("JOIN film_category on film.film_id = film_category.film_id").
		Joins("JOIN category on film_category.category_id = category.category_id").
		Joins("JOIN language on language.language_id = film.language_id").
		Where("film.title = ?", title).
		Find(&film)
	
	c.JSON(200, gin.H{
		"message": film,
	})
}

func GetFilmsByLanguage(c *gin.Context) {
	var film []models.FilmResponse
	language := c.Query("language")
	log.Println("Getting films by language: ", language)

	config.DB.Table("film").
		Select("film.title, film.description, film.release_year, TRIM(language.name) as language, category.name as category, film.length, film.rating, film.special_features").
		Joins("JOIN film_category on film.film_id = film_category.film_id").
		Joins("JOIN category on film_category.category_id = category.category_id").
		Joins("JOIN language on language.language_id = film.language_id").
		Where("language.name = ?", language).
		Find(&film)

	c.JSON(200, gin.H{
		"message": film,
	})
}

func GetFilmsByCategory(c *gin.Context) {
	var film []models.FilmResponse
	category := c.Query("category")
	log.Println("Getting films by category: ", category)

	config.DB.Table("film").
		Select("film.title, film.description, film.release_year, TRIM(language.name) as language, category.name as category,film.length, film.rating, film.special_features").
		Joins("JOIN language on language.language_id = film.language_id").
		Joins("JOIN film_category on film.film_id = film_category.film_id").
		Joins("JOIN category on film_category.category_id = category.category_id").
		Where("category.name = ?", category).
		Find(&film)
	
	c.JSON(200, gin.H{
		"message": film,
	})
}