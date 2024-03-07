package main

import (
	"github.com/Marwahkamilaahmad/go-rest-api.git/controllers"
	"github.com/Marwahkamilaahmad/go-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/api/pasien", controllers.Index)
	r.GET("/api/pasien/:id", controllers.Show)
	r.POST("/api/pasien", controllers.Create)
	r.PUT("/api/pasien/:id", controllers.Update)
	r.DELETE("/api/pasien", controllers.Delete)

	r.GET("/api/rumahsakit", controllers.IndexRS)
	r.GET("/api/rumahsakit/:id", controllers.ShowRS)
	r.POST("/api/rumahsakit", controllers.CreateRS)
	r.PUT("/api/rumahsakit/:id", controllers.UpdateRS)
	r.DELETE("/api/rumahsakit", controllers.DeleteRS)

	r.Run()
}