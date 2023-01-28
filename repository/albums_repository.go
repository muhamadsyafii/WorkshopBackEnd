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

type AlbumsRepository interface {
	FindAlbums() []response.Albums
	SaveAlbums(albums response.Albums)
	UpdateAlbums(albums *response.Albums)
	DeleteAlbums(albums *response.Albums)
	FindAlbumsById(albums response.Albums) response.Albums
}

type AlbumsDb struct {
	conn *gorm.DB
}

func NewAlbumsRepository() AlbumsRepository {
	dsn := "root:root@tcp(127.0.0.1:3306)/workshop_be_db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error")
	}
	//db.AutoMigrate(&response.Albums{})
	return &AlbumsDb{
		conn: db,
	}
}

func (d *AlbumsDb) FindAlbums() []response.Albums {
	var users []response.Albums
	d.conn.Find(&users)
	return users
}

func (d *AlbumsDb) SaveAlbums(albums response.Albums) {
	d.conn.Create(&albums)
}

func (d *AlbumsDb) DeleteAlbums(albums *response.Albums) {
	d.conn.Delete(&albums)
}

func (d *AlbumsDb) UpdateAlbums(albums *response.Albums) {
	d.conn.Save(&albums)
}

func (d *AlbumsDb) FindAlbumsById(albums response.Albums) response.Albums {
	d.conn.Find(&albums, albums.ID)
	return albums
}
