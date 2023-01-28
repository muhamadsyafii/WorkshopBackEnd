package controller

import (
	"finalProjectGOMoladin/model/response"
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

type AlbumsController struct {
	serviceAlbums service.AlbumsService
}

func New(albumsService service.AlbumsService) *AlbumsController {
	return &AlbumsController{
		serviceAlbums: albumsService,
	}
}
func (c *AlbumsController) FindAllAlbums(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    c.serviceAlbums.FindAllAlbums(),
	})
}

func (c *AlbumsController) SaveAlbums(ctx *gin.Context) {
	var albums response.Albums
	if err := ctx.ShouldBindJSON(&albums); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.BindJSON(&albums)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    c.serviceAlbums.SaveAlbums(albums),
	})
}

func (c *AlbumsController) UpdateAlbums(ctx *gin.Context) {
	var albums response.Albums
	if err := ctx.ShouldBindJSON(&albums); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.BindJSON(&albums)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    c.serviceAlbums.UpdateAlbums(albums),
	})
}

func (c *AlbumsController) DeleteAlbums(ctx *gin.Context) {
	var albums response.Albums
	id, _ := strconv.Atoi(ctx.Param("id"))
	albums.ID = uint(id)
	c.serviceAlbums.DeleteAlbums(albums)
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "success",
	})
}

func (c *AlbumsController) DetailAlbums(ctx *gin.Context) {
	var albums response.Albums
	id, _ := strconv.Atoi(ctx.Param("id"))
	albums.ID = uint(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": id,
		"data":    c.serviceAlbums.DetailAlbums(albums),
	})
}
