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
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
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
	albums := make([]entity.Albums, 0)
	err := d.conn.Find(&albums).Error
	if err != nil {
		fmt.Println("Erorororor")
		return nil
	}
	return albums
}

func (d *AlbumsDb) SaveAlbums(albums entity.Albums) {
	d.conn.Create(&albums)
}

func (d *AlbumsDb) DeleteAlbums(albums *entity.Albums) {
	d.conn.Delete(&albums)
}

func (d *AlbumsDb) UpdateAlbums(albums *entity.Albums) {
	var requestAlbums entity.Albums
	requestAlbums.Name = albums.Name
	requestAlbums.Year = albums.Year
	requestAlbums.CreatedAt = time.Now()
	d.conn.Where("id = ?", albums.ID).Updates(&requestAlbums)
}

func (d *AlbumsDb) FindAlbumsById(albums entity.Albums) entity.Albums {
	d.conn.Find(&albums, albums.ID)
	return albums
}
