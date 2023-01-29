package controller

import (
	"finalProjectGOMoladin/entity"
	"finalProjectGOMoladin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
 * Created by Muhamad Syafii
 * Friday, 28/01/2023
 * Learning Week Back End Developer
 * Copyright (c) 2023 by Moladin.
 * All Rights Reserved
 */

type SongController struct {
	serviceSongs service.SongService
}

func NewSongs(songService service.SongService) *SongController {
	return &SongController{
		serviceSongs: songService,
	}
}

func (c *SongController) FindSongsByAlbums(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	idAlbums, err := strconv.Atoi(ctx.Param("idAlbums"))
	data, err := c.serviceSongs.FindSongsByAlbums(uint(idAlbums))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    data,
	})
}

func (c *SongController) SaveSongs(ctx *gin.Context) {
	var songs entity.Songs
	ctx.Header("Content-Type", "application/json")
	if err := ctx.Bind(&songs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    c.serviceSongs.SaveSongs(songs),
	})
}

func (c *SongController) UpdateSongs(ctx *gin.Context) {
	var songs entity.Songs
	if err := ctx.ShouldBindJSON(&songs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.BindJSON(&songs)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    c.serviceSongs.UpdateSongs(songs),
	})
}

func (c *SongController) DeleteSongs(ctx *gin.Context) {
	var songs entity.Songs
	id, _ := strconv.Atoi(ctx.Param("id"))
	songs.ID = uint(id)
	ctx.Header("Content-Type", "application/json")
	c.serviceSongs.DeleteSongs(songs)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menghapus Songs",
	})
}

func (c *SongController) DetailSongs(ctx *gin.Context) {
	var songs entity.Songs
	id, err := strconv.Atoi(ctx.Param("id"))
	songs.ID = uint(id)
	ctx.Header("Content-Type", "application/json")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal Menampilkan Data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    c.serviceSongs.DetailSongs(songs),
	})
}
