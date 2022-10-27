package controller

import (
	"net/http"

	"github.com/Kool13Peeravuit/sa-project/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// POST /users
func CreateUser(c *gin.Context) {
	var user entity.User
	var memberclass entity.MemberClass
	var province entity.Province
	var role entity.Role

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา member ด้วย id
	if tx := entity.DB().Where("id = ?", user.MemberClassID).First(&memberclass); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "memberclass not found"})
		return
	}

	// 10: ค้นหา province ด้วย id
	if tx := entity.DB().Where("id = ?", user.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	// 11: ค้นหา role ด้วย id
	if tx := entity.DB().Where("id = ?", user.RoleID).First(&role); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found"})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	user.Password = string(hashPassword)

	// 12: สร้าง user
	us := entity.User{
		MemberClass: memberclass,
		Province:    province,
		Role:        role,
		Pin:         user.Pin,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Civ:         user.Civ,
		Phone:       user.Phone,
		Email:       user.Email,
		Password:    user.Password,
		Address:     user.Address,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&us).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": us})
}

// GET /user/:id
func GetUser(c *gin.Context) {
	var user entity.User
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users
func ListUser(c *gin.Context) {
	var user []entity.User
	if err := entity.DB().Preload("MemberClass").Preload("Province").Preload("Role").Raw("SELECT * FROM users").Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /users
func UpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", user.ID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
