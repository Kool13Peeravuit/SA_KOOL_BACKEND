package controller

import (
	"net/http"

	"github.com/Kool13Peeravuit/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// POST /researchrooms
func CreateResearchRoom(c *gin.Context) {
	var researchroom entity.ResearchRoom

	if err := c.ShouldBindJSON(&researchroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&researchroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchroom})

}

// GET /researchroom/:id
func GetResearchRoom(c *gin.Context) {
	var researchroom entity.ResearchRoom
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&researchroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "researchroom not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchroom})
}

// GET /researchrooms
func ListResearchRooms(c *gin.Context) {
	var researchrooms []entity.ResearchRoom
	if err := entity.DB().Preload("RoomType").Raw("SELECT * FROM research_rooms").Find(&researchrooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchrooms})
}

// DELETE /researchrooms/:id
func DeleteResearchRoom(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM research_rooms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ResearchRoom not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /researchrooms
func UpdateResearchRoom(c *gin.Context) {
	var researchroom entity.ResearchRoom
	if err := c.ShouldBindJSON(&researchroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", researchroom.ID).First(&researchroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ResearchRoom not found"})
		return
	}
	if err := entity.DB().Save(&researchroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchroom})
}
