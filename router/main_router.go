package router

import "github.com/gin-gonic/gin"

/*
 * Created by Muhamad Syafii
 * Friday, 28/01/2023
 * Learning Week Back End Developer
 * Copyright (c) 2023 by Moladin.
 * All Rights Reserved
 */

var (
	router = gin.Default()
)

func Run() {
	routes()
	router.Run()
}

func routes() {
	v1 := router.Group("/api/v1")
	SongsRouters(v1)
	AlbumsRouters(v1)
}
