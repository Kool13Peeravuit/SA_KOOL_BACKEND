package controller

import (
	"net/http"

	"github.com/Kool13Peeravuit/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// POST /computers
func CreateComputer(c *gin.Context) {
	var computer entity.Computer
	var computer_os entity.Computer_os

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร computer_reservation
	if err := c.ShouldBindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา computer_os ด้วย id
	if tx := entity.DB().Where("id = ?", computer.Computer_os_id).First(&computer_os); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := c.ShouldBindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&computer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": computer})

	// 13: สร้าง COMPUTER
	com := entity.Computer{
		Computer_os: computer_os, // โยงความสัมพันธ์กับ Entity User

	}

	// 14: บันทึก
	if err := entity.DB().Create(&com).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": com})
}

// GET /computers/:id
func GetComputer(c *gin.Context) {
	var computer entity.Computer
	id := c.Param("id")
	if err := entity.DB().Preload("Computer_os").Raw("SELECT * FROM computers WHERE id = ?", id).Find(&computer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer})
}

// GET /computers
func ListComputers(c *gin.Context) {
	var computers []entity.Computer
	if err := entity.DB().Preload("Computer_os").Raw("SELECT * FROM computers").Find(&computers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computers})
}

// DELETE /computers/:id
func DeleteComputer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM computers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /computers
func UpdateComputer(c *gin.Context) {
	var computer entity.Computer
	if err := c.ShouldBindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", computer.ID).First(&computer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer not found"})
		return
	}

	if err := entity.DB().Save(&computer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer})
}
