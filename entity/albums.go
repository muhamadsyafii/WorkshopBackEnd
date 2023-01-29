package entity

import "time"

type Albums struct {
	ID        uint      `gorm:"primaryKey;type:INT UNSIGNED" json:"id"`
	Name      string    `json:"name" binding:"required"`
	Year      uint      `json:"year" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Songs     []Songs   `gorm:"foreignKey:AlbumId;references:ID" json:"albums"`
}
