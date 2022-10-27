package controller

import (
	"fmt"
	"net/http"

	"github.com/Kool13Peeravuit/sa-project/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST //watch_Tlts
func CreateProblemReport(c *gin.Context) {
	var User entity.User
	var ProblemReport entity.ProblemReport
	var Problem entity.Problem
	var Place_Class entity.Place_Class
	var Tlt entity.Toilet
	var RdZone entity.ReadingZone
	var ReschRoom entity.ResearchRoom
	var Com entity.Computer

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร ProblemReport
	if err := c.ShouldBindJSON(&ProblemReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา User
	if tx := entity.DB().Where("id = ?", ProblemReport.USER_ID).First(&User); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// 9: ค้นหา Problem ด้วย id
	if tx := entity.DB().Where("id = ?", ProblemReport.Problem_ID).First(&Problem); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem not found"})
		return

	}

	// 10: ค้นด้วยไอดีชื่อสถานที่(PlaceName_id)
	fmt.Println("what", ProblemReport.Tlt_id)
	fmt.Println("what", ProblemReport.RdZone_id)
	fmt.Println("what", ProblemReport.ReschRoom_id)
	fmt.Println("what", ProblemReport.Com_id)

	if ProblemReport.Tlt_id != nil {
		if tx := entity.DB().Where("id = ?", ProblemReport.Tlt_id).First(&Tlt); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Toilet not found"})
			return
		}
	}

	if ProblemReport.RdZone_id != nil {
		if tx := entity.DB().Where("id = ?", ProblemReport.RdZone_id).First(&RdZone); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Reding Zone not found"})
			return
		}
	}

	if ProblemReport.ReschRoom_id != nil {
		if tx := entity.DB().Where("id = ?", ProblemReport.ReschRoom_id).First(&ReschRoom); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Research Room not found"})
			return
		}
	}

	if ProblemReport.Com_id != nil {
		if tx := entity.DB().Where("id = ?", ProblemReport.Com_id).First(&Com); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Reading Zone not found"})
			return
		}
	}

	//********************************************************************************

	// 11: ค้นหา Place_Class ด้วย id
	if tx := entity.DB().Where("id = ?", ProblemReport.Place_Class_ID).First(&Place_Class); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Place_Class not found"})
		return
	}

	// 12: สร้าง ProblemReport
	wv := entity.ProblemReport{
		User:        User,        // โยงความสัมพันธ์กับ Entity User
		Problem:     Problem,     // โยงความสัมพันธ์กับ Entity Problem
		Tlt:         Tlt,         // โยงความสัมพันธ์กับ Entity Toilet
		RdZone:      RdZone,      // โยงความสัมพันธ์กับ Entity ReadingZone
		ReschRoom:   ReschRoom,   // โยงความสัมพันธ์กับ Entity ResearchRoom
		Com:         Com,         // โยงความสัมพันธ์กับ Entity Computer
		Place_Class: Place_Class, // โยงความสัมพันธ์กับ Entity Place_Class
		Comment:     ProblemReport.Comment,
		Date:        ProblemReport.Date,
		/*
			WatchedTime: ProblemReport.WatchedTime, // ตั้งค่าฟิลด์ watchedTime
		*/
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(wv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

/**************************************************************************************/

// GET /ProblemReport/:id *********************************************** !!!!

func GetProblemReport(c *gin.Context) {
	var ProblemReport entity.ProblemReport
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&ProblemReport); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem_reports not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ProblemReport})
}

// GET /ProblemReport *********************************************** !!!!
func ListProblemReport(c *gin.Context) {
	var ProblemReports []entity.ProblemReport
	if err := entity.DB().Preload("User").Preload("Problem").Preload("RdZone").Preload("Tlt").Preload("ReschRoom").Preload("Com").Preload("Place_Class").Raw("SELECT * FROM problem_reports").Find(&ProblemReports).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ProblemReports})
}

// DELETE /watch_videos/:id  *********************************************** !!!!
func DeleteProblemReport(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM problem_reports WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem_reports not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos  *********************************************** !!!!
func UpdateProblemReport(c *gin.Context) {
	var ProblemReport entity.ProblemReport
	if err := c.ShouldBindJSON(&ProblemReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", ProblemReport.ID).First(&ProblemReport); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem_reports not found"})
		return
	}

	if err := entity.DB().Save(&ProblemReport).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ProblemReport})
}
