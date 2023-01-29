package controller

import (
	"finalProjectGOMoladin/entity"
	"finalProjectGOMoladin/service"
	"fmt"
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
	albums := c.serviceAlbums.FindAllAlbums()
	fmt.Println(albums)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    albums,
	})
}

func (c *AlbumsController) SaveAlbums(ctx *gin.Context) {
	var albums entity.Albums
	ctx.Header("Content-Type", "application/json")
	ctx.BindJSON(&albums)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    c.serviceAlbums.SaveAlbums(albums),
	})
}

func (c *AlbumsController) UpdateAlbums(ctx *gin.Context) {
	var albums entity.Albums
	if err := ctx.ShouldBindJSON(&albums); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.BindJSON(&albums)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    c.serviceAlbums.UpdateAlbums(albums),
	})
}

func (c *AlbumsController) DeleteAlbums(ctx *gin.Context) {
	var albums entity.Albums
	id, _ := strconv.Atoi(ctx.Param("id"))
	albums.ID = uint(id)
	ctx.Header("Content-Type", "application/json")
	c.serviceAlbums.DeleteAlbums(albums)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menghapus Albums",
	})
}

func (c *AlbumsController) DetailAlbums(ctx *gin.Context) {
	var albums entity.Albums
	id, err := strconv.Atoi(ctx.Param("id"))
	albums.ID = uint(id)
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
		"data":    c.serviceAlbums.DetailAlbums(albums),
	})
}
