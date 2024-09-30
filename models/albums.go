package models

import "time"

type Album struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ArtistID    uint      `json:"artist_id" gorm:"not null"`
	Artist      Artist    `json:"artist" gorm:"foreignKey:ArtistID"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	ReleaseDate time.Time `json:"release_date" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	IsDeleted   bool      `json:"is_deleted" gorm:"default:false"`
}
