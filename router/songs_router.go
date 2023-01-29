package router

import (
	"finalProjectGOMoladin/controller"
	"finalProjectGOMoladin/repository"
	"finalProjectGOMoladin/service"
	"github.com/gin-gonic/gin"
)

/*
 * Created by Muhamad Syafii
 * Friday, 28/01/2023
 * Learning Week Back End Developer
 * Copyright (c) 2023 by Moladin.
 * All Rights Reserved
 */

func SongsRouters(rg *gin.RouterGroup) {
	rest := controller.NewSongs(service.NewSongs(repository.NewSongRepository()))
	rg.GET("/songs/albums/:idAlbums", rest.FindSongsByAlbums)
	rg.POST("/songs", rest.SaveSongs)
	rg.PUT("/songs", rest.UpdateSongs)
	rg.DELETE("/songs/:id", rest.DeleteSongs)
	rg.GET("/songs/:id", rest.DetailSongs)
}
