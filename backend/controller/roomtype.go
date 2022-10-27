package controller

import (
	"net/http"

	"github.com/Kool13Peeravuit/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// POST /roomtypes
func CreateRoomType(c *gin.Context) {
	var roomtype entity.RoomType
	if err := c.ShouldBindJSON(&roomtype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// GET /roomtype/:id
func GetRoomType(c *gin.Context) {
	var roomtype entity.RoomType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM room_types WHERE id = ?", id).Scan(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// GET /roomtypes
func ListRoomTypes(c *gin.Context) {
	var roomtypes []entity.RoomType
	if err := entity.DB().Raw("SELECT * FROM room_types").Scan(&roomtypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtypes})
}

// DELETE /roomtypes/:id
func DeleteRoomType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_types not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /roomtypes
func UpdateRoomType(c *gin.Context) {
	var roomtype entity.RoomType
	if err := c.ShouldBindJSON(&roomtype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", roomtype.ID).First(&roomtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_type not found"})
		return
	}
	if err := entity.DB().Save(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtype})

}
