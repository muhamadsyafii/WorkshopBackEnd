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

func AlbumsRouters(rg *gin.RouterGroup) {
	rest := controller.New(service.NewAlbums(repository.NewAlbumsRepository()))
	rg.GET("/albums", rest.FindAllAlbums)
	rg.POST("/albums", rest.SaveAlbums)
	rg.PUT("/albums", rest.UpdateAlbums)
	rg.DELETE("/albums/:id", rest.DeleteAlbums)
	rg.GET("/albums/:id", rest.DetailAlbums)
}
