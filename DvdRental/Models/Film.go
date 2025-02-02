package models

import "time"

type Film struct {
	FilmId int `gorm:"primaryKey;column:film_id"`
	Title string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	ReleaseYear string `gorm:"column:release_year"`
	LanguageId int `gorm:"column:language_id"`
	RentalDuration int `gorm:"column:rental_duration"`
	RentalRate float64 `gorm:"column:rental_rate"`
	Length int `gorm:"column:length"`
	ReplacementCost float64 `gorm:"column:replacement_cost"`
	Rating string `gorm:"column:rating"`
	LastUpdate time.Time `gorm:"column:last_update"`
	SpecialFeatures string `gorm:"column:special_features"`
	FullText string `gorm:"column:fulltext"`
}

func (Film) TableName() string {
	return "film"
}

type FilmResponse struct {
	Title string `json:"title"`
	Description string `json:"description"`
	ReleaseYear string `json:"release_year"`
	Language    string `json:"language"`
	Category	string `json:"category"`
	Length 		int    `json:"length"`
	Rating 		string `json:"rating"`
	SpecialFeatures string `json:"special_features"`
}
