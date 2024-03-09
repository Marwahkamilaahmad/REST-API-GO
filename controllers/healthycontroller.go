package controllers

import (
	"net/http"

	"github.com/Marwahkamilaahmad/go-rest-api.git/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type pasienController struct {
  db *gorm.DB // Replace with your database connection instance
}

func NewPasienController(db *gorm.DB) *pasienController {
  return &pasienController{db: db}
}

// GetPatients retrieves all patients from the database
func (c *pasienController) GetPatients(ctx *gin.Context) {
  var patients []models.Pasien
  if err := c.db.Find(&patients).Error; err != nil {
    ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error fetching patients"})
    return
  }
  ctx.JSON(http.StatusOK, gin.H{"patients": patients})
}

// GetPatient retrieves a specific patient by ID
func (c *pasienController) GetPatient(ctx *gin.Context) {
  id := ctx.Param("id")
  var patient models.Pasien
  if err := c.db.First(&patient, id).Error; err != nil {
    if err == gorm.ErrRecordNotFound {
      ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Patient not found"})
    } else {
      ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error fetching patient"})
    }
    return
  }
  ctx.JSON(http.StatusOK, gin.H{"patient": patient})
}

// CreatePatient creates a new patient in the database
func (c *pasienController) CreatePatient(ctx *gin.Context) {
  var patient models.Pasien
  if err := ctx.ShouldBindJSON(&patient); err != nil {
    ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
    return
  }

  // Implement validation logic here to ensure data integrity

  if err := c.db.Create(&patient).Error; err != nil {
    ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error creating patient"})
    return
  }

  ctx.JSON(http.StatusOK, gin.H{"message": "Patient created successfully", "id": patient.ID})
}

// UpdatePatient updates an existing patient in the database
func (c *pasienController) UpdatePatient(ctx *gin.Context) {
  id := ctx.Param("id")
  var patient models.Pasien
  if err := ctx.ShouldBindJSON(&patient); err != nil {
    ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
    return
  }

  // Implement validation logic here to ensure data integrity

  if c.db.Model(&patient).Where("id = ?", id).Updates(&patient).RowsAffected == 0 {
    ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Patient not found or update failed"})
    return
  }

  ctx.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully"})
}

// DeletePatient deletes a patient from the database
func (c *pasienController) DeletePatient(ctx *gin.Context) {
  id := ctx.Param("id")
  var patient models.Pasien
  if err := c.db.Delete(&patient, id).Error; err != nil {
    ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting patient"})
    return
  }

  ctx.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
