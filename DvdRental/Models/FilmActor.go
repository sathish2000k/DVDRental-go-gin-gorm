package models

type FilmActor struct {
	ActorId int `gorm:"column:actor_id"`
	FilmId int `gorm:"column:film_id"`
	LastUpdate string `gorm:"column:last_update"`
}