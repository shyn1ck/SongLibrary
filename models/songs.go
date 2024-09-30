package models

import (
	"time"
)

type Song struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ArtistID    uint      `json:"artist_id" gorm:"not null"`
	Artist      Artist    `json:"artist" gorm:"foreignKey:ArtistID"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	ReleaseDate time.Time `json:"release_date" gorm:"default:CURRENT_TIMESTAMP"`
	Text        string    `json:"text" gorm:"type:text"`
	Link        string    `json:"link" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	IsDeleted   bool      `json:"is_deleted" gorm:"default:false"`
}
