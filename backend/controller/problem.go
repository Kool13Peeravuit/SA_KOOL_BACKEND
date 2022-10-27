package controller

import (
	"net/http"

	"github.com/Kool13Peeravuit/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// POST /Problems
func CreateProblem(c *gin.Context) {
	var Problem entity.Problem
	if err := c.ShouldBindJSON(&Problem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Problem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Problem})
}

// GET /Problem/:id
func GetProblem(c *gin.Context) {
	var Problem entity.Problem
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Problems WHERE id = ?", id).Scan(&Problem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Problem})
}

// GET /Problems
func ListProblems(c *gin.Context) {
	var Problems []entity.Problem
	if err := entity.DB().Raw("SELECT * FROM Problems").Scan(&Problems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Problems})
}

// DELETE /Problems/:id
func DeleteProblem(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Problems WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Problems
func UpdateProblem(c *gin.Context) {
	var Problem entity.Problem
	if err := c.ShouldBindJSON(&Problem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Problem.ID).First(&Problem); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem not found"})
		return
	}

	if err := entity.DB().Save(&Problem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Problem})
}
