package entity

import "time"

type Songs struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AlbumId   uint      `json:"album_Id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Albums    []Albums  `gorm:"foreignKey:ID;references:AlbumsId"`
}
