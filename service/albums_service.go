package service

/*
 * Created by Muhamad Syafii
 * Friday, 28/01/2023
 * Learning Week Back End Developer
 * Copyright (c) 2023 by Moladin.
 * All Rights Reserved
 */

import (
	"finalProjectGOMoladin/model/response"
	"finalProjectGOMoladin/repository"
)

type AlbumsService interface {
	FindAllAlbums() []response.Albums
	SaveAlbums(users response.Albums) response.Albums
	DeleteAlbums(users response.Albums)
	UpdateAlbums(users response.Albums) response.Albums
	DetailAlbums(users response.Albums) response.Albums
}

type albumsService struct {
	albumsRepository repository.AlbumsRepository
}

func NewAlbums(albumsRepository repository.AlbumsRepository) AlbumsService {
	return &albumsService{
		albumsRepository: albumsRepository,
	}
}

func (u *albumsService) FindAllAlbums() []response.Albums {
	return u.albumsRepository.FindAlbums()
}

func (u *albumsService) SaveAlbums(albums response.Albums) response.Albums {
	u.albumsRepository.SaveAlbums(albums)
	return albums
}

func (u *albumsService) DeleteAlbums(albums response.Albums) {
	u.albumsRepository.DeleteAlbums(&albums)
}

func (u *albumsService) UpdateAlbums(albums response.Albums) response.Albums {
	u.albumsRepository.UpdateAlbums(&albums)
	return albums
}

func (u *albumsService) DetailAlbums(albums response.Albums) response.Albums {
	res := u.albumsRepository.FindAlbumsById(albums)
	return res
}
