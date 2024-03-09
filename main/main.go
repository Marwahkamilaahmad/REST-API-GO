package main

import (
	"github.com/Marwahkamilaahmad/go-rest-api.git/controllers"
	"github.com/Marwahkamilaahmad/go-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default();

	models.ConnectDatabase()

	db := models.DB

	pasienController := controllers.NewPasienController(db)

	r.GET("/api/pasien",pasienController.GetPatients)
	r.GET("/api/pasien/:id", pasienController.GetPatient)
	r.POST("/api/pasien", pasienController.CreatePatient)
	r.PUT("/api/pasien/:id", pasienController.UpdatePatient)
	r.DELETE("/api/pasien", pasienController.DeletePatient)

	r.GET("/api/rumahsakit", controllers.IndexRS)
	r.GET("/api/rumahsakit/:id", controllers.ShowRS)
	r.POST("/api/rumahsakit", controllers.CreateRS)
	r.PUT("/api/rumahsakit/:id", controllers.UpdateRS)
	// r.DELETE("/api/rumahsakit", controllers.DeleteRS)

	r.Run()
}