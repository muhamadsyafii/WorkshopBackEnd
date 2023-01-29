package repository

/*
 * Created by Muhamad Syafii
 * Friday, 28/01/2023
 * Learning Week Back End Developer
 * Copyright (c) 2023 by Moladin.
 * All Rights Reserved
 */

import (
	"finalProjectGOMoladin/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type SongRepository interface {
	FindSongsByAlbums(idAlbums uint) ([]entity.Songs, error)
	SaveSongs(songs entity.Songs)
	UpdateSongs(songs *entity.Songs)
	DeleteSongs(songs *entity.Songs)
	FindSongsById(songs entity.Songs) entity.Songs
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
	//db.AutoMigrate(&entity.Songs{})
	return &SongsDb{
		conn: db,
	}
}

func (d *SongsDb) FindSongsByAlbums(idAlbums uint) ([]entity.Songs, error) {
	var songs []entity.Songs
	err := d.conn.Model(&entity.Songs{}).Scan(&songs).Find(&songs, "album_id=?", idAlbums).Error
	return songs, err
}

func (d *SongsDb) SaveSongs(songs entity.Songs) {
	d.conn.Create(&songs)
}

func (d *SongsDb) DeleteSongs(songs *entity.Songs) {
	d.conn.Delete(&songs)
}

func (d *SongsDb) UpdateSongs(songs *entity.Songs) {
	var requestSong entity.Songs
	requestSong.AlbumId = songs.AlbumId
	requestSong.Title = songs.Title
	requestSong.Author = songs.Author
	requestSong.CreatedAt = time.Now()
	d.conn.Where("id = ?", songs.ID).Updates(&requestSong)
}

func (d *SongsDb) FindSongsById(songs entity.Songs) entity.Songs {
	d.conn.Find(&songs, songs.ID)
	return songs
}
