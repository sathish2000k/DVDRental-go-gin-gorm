package models

type Actor struct {
	ActorId int `gorm:"primaryKey;column:actor_id"`
	FirstName string `gorm:"column:first_name"`
	LastName string `gorm:"column:last_name"`
	LastUpdate string `gorm:"column:last_update"`
}

func (Actor) TableName() string {
	return "actor"
}

type ActorResponse struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}