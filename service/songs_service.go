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

type SongService interface {
	FindSongs() []response.Songs
	SaveSongs(songs response.Songs) response.Songs
	UpdateSongs(songs response.Songs) response.Songs
	DeleteSongs(songs response.Songs)
	DetailSongs(songs response.Songs) response.Songs
}

type songService struct {
	repository repository.SongRepository
}

func NewSong(songRepository repository.SongRepository) SongService {
	return &songService{
		repository: songRepository,
	}
}

func (u *songService) FindSongs() []response.Songs {
	return u.repository.FindSongs()
}

func (u *songService) SaveSongs(songs response.Songs) response.Songs {
	u.repository.SaveSongs(songs)
	return songs
}

func (u *songService) UpdateSongs(songs response.Songs) response.Songs {
	u.repository.UpdateSongs(&songs)
	return songs
}

func (u *songService) DeleteSongs(songs response.Songs) {
	u.repository.DeleteSongs(&songs)
}
func (u *songService) DetailSongs(songs response.Songs) response.Songs {
	res := u.repository.FindSongsById(songs)
	return res
}
