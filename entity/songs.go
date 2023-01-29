package entity

import "time"

type Songs struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AlbumId   uint      `json:"albumId"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//Albums    []Albums  `gorm:"foreignKey:ID;references:AlbumId" json:"albums"`
}
