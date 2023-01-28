package service

/*
 * Created by Muhamad Syafii
 * Friday, 28/01/2023
 * Learning Week Back End Developer
 * Copyright (c) 2023 by Moladin.
 * All Rights Reserved
 */

import (
	"finalProjectGOMoladin/entity"
	"finalProjectGOMoladin/repository"
)

type AlbumsService interface {
	FindAllAlbums() []entity.Albums
	SaveAlbums(users entity.Albums) entity.Albums
	DeleteAlbums(users entity.Albums)
	UpdateAlbums(users entity.Albums) entity.Albums
	DetailAlbums(users entity.Albums) entity.Albums
}

type albumsService struct {
	albumsRepository repository.AlbumsRepository
}

func NewAlbums(albumsRepository repository.AlbumsRepository) AlbumsService {
	return &albumsService{
		albumsRepository: albumsRepository,
	}
}

func (u *albumsService) FindAllAlbums() []entity.Albums {
	return u.albumsRepository.FindAlbums()
}

func (u *albumsService) SaveAlbums(albums entity.Albums) entity.Albums {
	u.albumsRepository.SaveAlbums(albums)
	return albums
}

func (u *albumsService) DeleteAlbums(albums entity.Albums) {
	u.albumsRepository.DeleteAlbums(&albums)
}

func (u *albumsService) UpdateAlbums(albums entity.Albums) entity.Albums {
	u.albumsRepository.UpdateAlbums(&albums)
	return albums
}

func (u *albumsService) DetailAlbums(albums entity.Albums) entity.Albums {
	res := u.albumsRepository.FindAlbumsById(albums)
	return res
}
