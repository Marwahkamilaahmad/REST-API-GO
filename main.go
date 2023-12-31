package main

import (

	"github.com/Marwahkamilaahmad/go-rest-api.git/controllers"
	"github.com/Marwahkamilaahmad/go-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

func restapi(){

	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/api/products", controllers.Index)
	r.GET("/api/products/:id", controllers.Show)
	r.POST("/api/products", controllers.Create)
	r.PUT("/api/products/:id", controllers.Update)
	r.DELETE("/api/products", controllers.Delete)

	r.Run()
}