package repository

/*
 * Created by Muhamad Syafii
 * Friday, 28/01/2023
 * Learning Week Back End Developer
 * Copyright (c) 2023 by Moladin.
 * All Rights Reserved
 */

import (
	"finalProjectGOMoladin/model/response"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SongRepository interface {
	FindSongs() []response.Songs
	SaveSongs(songs response.Songs)
	UpdateSongs(songs *response.Songs)
	DeleteSongs(songs *response.Songs)
	FindSongsById(songs response.Songs) response.Songs
}

type SongsDb struct {
	conn *gorm.DB
}

func NewSongRepository() SongRepository {
	dsn := "root:root@tcp(127.0.0.1:3306)/workshop_be_db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error")
	}
	//db.AutoMigrate(&response.Songs{})
	return &SongsDb{
		conn: db,
	}
}

func (d *SongsDb) FindSongs() []response.Songs {
	var songs []response.Songs
	d.conn.Find(&songs)
	return songs
}

func (d *SongsDb) SaveSongs(songs response.Songs) {
	d.conn.Create(&songs)
}

func (d *SongsDb) DeleteSongs(songs *response.Songs) {
	d.conn.Delete(&songs)
}

func (d *SongsDb) UpdateSongs(songs *response.Songs) {
	d.conn.Save(&songs)
}

func (d *SongsDb) FindSongsById(songs response.Songs) response.Songs {
	d.conn.Find(&songs, songs.ID)
	return songs
}
