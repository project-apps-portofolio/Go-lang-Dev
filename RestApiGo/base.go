package main

import (
	config "restapigo/config"
	controllers "restapigo/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Panggil Config DBinit

	db := config.DBinit()
	apps := &controllers.InDB{DB: db}

	// Router Gin Default

	router := gin.Default()

	// With Group
	v1 := router.Group("api/v1")
	{
		v1.GET("/persons", apps.GetPerson)
		v1.POST("/persons/create", apps.PostPerson)
	}

	// Without Group

	router.GET("/persons", apps.GetPerson )
	router.POST("/persons/create", apps.GetPerson )

	router.Run(":3000")
	
}