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
)

type AlbumsRepository interface {
	FindAlbums() []entity.Albums
	SaveAlbums(albums entity.Albums)
	UpdateAlbums(albums *entity.Albums)
	DeleteAlbums(albums *entity.Albums)
	FindAlbumsById(albums entity.Albums) entity.Albums
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
	//db.AutoMigrate(&entity.Albums{})
	return &AlbumsDb{
		conn: db,
	}
}

func (d *AlbumsDb) FindAlbums() []entity.Albums {
	var users []entity.Albums
	d.conn.Find(&users)
	return users
}

func (d *AlbumsDb) SaveAlbums(albums entity.Albums) {
	d.conn.Create(&albums)
}

func (d *AlbumsDb) DeleteAlbums(albums *entity.Albums) {
	d.conn.Delete(&albums)
}

func (d *AlbumsDb) UpdateAlbums(albums *entity.Albums) {
	d.conn.Save(&albums)
}

func (d *AlbumsDb) FindAlbumsById(albums entity.Albums) entity.Albums {
	d.conn.Find(&albums, albums.ID)
	return albums
}
