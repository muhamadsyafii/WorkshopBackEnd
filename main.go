package main

/*
 * Created by Muhamad Syafii
 * Friday, 28/01/2023
 * Learning Week Back End Developer
 * Copyright (c) 2023 by Moladin.
 * All Rights Reserved
 */

import (
	"finalProjectGOMoladin/model/request"
	"finalProjectGOMoladin/model/response"
	"finalProjectGOMoladin/router"
	"finalProjectGOMoladin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * Read Documentation from https://gorm.io/docs/
 * for model using package response (albums and songs)
 */

func main() {
	//utils.ConnectDatabase()
	//
	gin.SetMode(gin.ReleaseMode)
	//r := gin.Default()
	//
	//// regions albums
	//CreateAlbums(r)
	//FindAllAlbums(r)
	//FindAlbumsById(r)
	////endregion
	//
	//err := r.Run()
	//if err != nil {
	//	return
	//}
	router.Run()
}

func CreateAlbums(r *gin.Engine) {
	// Validate input
	var input request.Albums
	createAlbums := request.Albums{Name: input.Name, Year: input.Year}
	utils.DB.Create(&createAlbums)
	r.POST("/add-albums", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Berhasil Menambah Data",
			"data":    createAlbums,
		})
	})
}

func FindAlbumsById(r *gin.Engine) { // Get model if exist
	var albums []response.Albums
	r.GET("/albums/:id", func(c *gin.Context) {
		if err := utils.DB.Where("id = ?", c.Param("id")).First(&albums).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Gagal Menampilkan Data",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data":    albums,
			"success": true,
			"message": "Berhasil Menampilkan Data",
		})
	})
}

func FindAllAlbums(r *gin.Engine) {
	var albums []response.Albums
	utils.DB.Find(&albums)
	r.GET("/all-albums", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":    albums,
			"success": true,
			"message": "Berhasil menampilkan data albums",
		})
	})
}
