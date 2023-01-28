package response

import "time"

type Songs struct {
	ID        uint `gorm:"primaryKey" json:"-"`
	AlbumId   uint
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	Albums    []Albums  `gorm:"foreignKey:ID;references:AlbumsId"`
}

//`gorm:"foreignKey:ID;references:AlbumsId"`
