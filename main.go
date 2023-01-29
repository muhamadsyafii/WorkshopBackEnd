package main

/*
 * Created by Muhamad Syafii
 * Friday, 28/01/2023
 * Learning Week Back End Developer
 * Copyright (c) 2023 by Moladin.
 * All Rights Reserved
 */

import (
	"finalProjectGOMoladin/router"
	"github.com/gin-gonic/gin"
)

/*
 * Read Documentation from https://gorm.io/docs/
 * for model using package entity (albums and songs)
 */

func main() {
	/*
	* Switch if error mode from debug mode to ReleaseMode
	 */
	gin.SetMode(gin.DebugMode)
	router.Run()
}
