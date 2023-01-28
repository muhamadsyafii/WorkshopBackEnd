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

type SongService interface {
	FindSongs() []entity.Songs
	SaveSongs(songs entity.Songs) entity.Songs
	UpdateSongs(songs entity.Songs) entity.Songs
	DeleteSongs(songs entity.Songs)
	DetailSongs(songs entity.Songs) entity.Songs
}

type songService struct {
	repository repository.SongRepository
}

func NewSongs(songRepository repository.SongRepository) SongService {
	return &songService{
		repository: songRepository,
	}
}

func (u *songService) FindSongs() []entity.Songs {
	return u.repository.FindSongs()
}

func (u *songService) SaveSongs(songs entity.Songs) entity.Songs {
	u.repository.SaveSongs(songs)
	return songs
}

func (u *songService) UpdateSongs(songs entity.Songs) entity.Songs {
	u.repository.UpdateSongs(&songs)
	return songs
}

func (u *songService) DeleteSongs(songs entity.Songs) {
	u.repository.DeleteSongs(&songs)
}
func (u *songService) DetailSongs(songs entity.Songs) entity.Songs {
	res := u.repository.FindSongsById(songs)
	return res
}
