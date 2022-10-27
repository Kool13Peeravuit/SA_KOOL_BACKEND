package controller

import (
	"net/http"

	"github.com/Kool13Peeravuit/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// POST /Place_Classs
func CreatePlace_Class(c *gin.Context) {
	var Place_Class entity.Place_Class
	if err := c.ShouldBindJSON(&Place_Class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Place_Class).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Place_Class})
}

// GET /Place_Class/:id
func GetPlace_Class(c *gin.Context) {
	var Place_Class entity.Place_Class
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&Place_Class); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Place_Class not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Place_Class})
}

func ListPlace_Class(c *gin.Context) {
	var Place_Class []entity.Place_Class
	if err := entity.DB().Raw("SELECT * FROM Place_Classes").Scan(&Place_Class).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Place_Class})
}

// DELETE /Place_Class/:id
func DeletePlace_Class(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Place_Classes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Place_Class not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Place_Classs
func UpdatePlace_Class(c *gin.Context) {
	var Place_Class entity.Place_Class
	if err := c.ShouldBindJSON(&Place_Class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Place_Class.ID).First(&Place_Class); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Place_Classes not found"})
		return
	}

	if err := entity.DB().Save(&Place_Class).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Place_Class})
}
