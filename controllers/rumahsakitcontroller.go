package controllers

import (
	"net/http"

	"github.com/Marwahkamilaahmad/go-rest-api.git/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexRS(c *gin.Context){
	var rumahSakit []models.RumahSakit

	models.DB.Find(&rumahSakit)
	c.JSON(http.StatusOK, gin.H{"rumahSakit" : rumahSakit})

}
func ShowRS(c *gin.Context){
	var rumahSakit models.RumahSakit
	id := c.Param("id")

	if err := models.DB.First(&rumahSakit, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message" : "Data tidak ditemukan"})
			return
			default :
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"rumahSakit" : rumahSakit})

}

func CreateRS(c *gin.Context){
	var rumahSakit models.RumahSakit

	if err := c.ShouldBindJSON(&rumahSakit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	models.DB.Create(&rumahSakit)
	c.JSON(http.StatusOK, gin.H{"rumahSakit" : rumahSakit})

}
func UpdateRS(c *gin.Context){
	var rumahSakit models.RumahSakit
	id := c.Param("id")

	if err := c.ShouldBindJSON(&rumahSakit); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	if models.DB.Model(&rumahSakit).Where("id = ?", id).Updates(&rumahSakit).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : "tidak dapat diperbarui"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message" : "berhasil memperbarui data"})

}
// func DeleteRS(c *gin.Context){
// 	var rumahSakit models.RumahSakit

// 	var input struct {
// 		Id int64 `json:"id"`
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil{
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
// 		return
// 	}

// 	if models.DB.Delete(&rumahSakit, input.Id).RowsAffected==0{
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"tidak dapat menghapus data"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})

// }

