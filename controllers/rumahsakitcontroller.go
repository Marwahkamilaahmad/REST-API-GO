package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Marwahkamilaahmad/go-rest-api.git/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexRS(c *gin.Context){
	var pasien []models.Pasien

	models.DB.Find(&pasien)
	c.JSON(http.StatusOK, gin.H{"pasien" : pasien})

}
func ShowRS(c *gin.Context){
	var pasien models.Pasien
	id := c.Param("id")

	if err := models.DB.First(&pasien, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message" : "Data tidak ditemukan"})
			return
			default :
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"pasien" : pasien})

}

func CreateRS(c *gin.Context){
	var pasien models.Pasien

	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	models.DB.Create(&pasien)
	c.JSON(http.StatusOK, gin.H{"pasien" : pasien})

}
func UpdateRS(c *gin.Context){
	var pasien models.Pasien
	id := c.Param("id")

	if err := c.ShouldBindJSON(&pasien); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	if models.DB.Model(&pasien).Where("id = ?", id).Updates(&pasien).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : "tidak dapat diperbarui"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message" : "berhasil memperbarui data"})

}
func DeleteRS(c *gin.Context){
	var pasien models.Pasien

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}
	id, _ := input.Id.Int64()

	if models.DB.Delete(&pasien, id).RowsAffected==0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"tidak dapat menghapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})

}

