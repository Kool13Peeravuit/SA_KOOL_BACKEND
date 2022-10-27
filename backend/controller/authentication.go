package controller

import (
	"net/http"

	"github.com/Kool13Peeravuit/sa-project/entity"
	"github.com/Kool13Peeravuit/sa-project/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpPayload signup body
type SignUpPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Role  string `json:"role"`
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

// POST /login
func Login(c *gin.Context) {
	var payload LoginPayload
	var us entity.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM users WHERE email = ?", payload.Email).Scan(&us).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})
		return
	}

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(us.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginResponse{
		Role:  GetRoleName(us.ID),
		Token: signedToken,
		ID:    us.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

func GetRoleName(id uint) string {
	rn := entity.User{}
	tx := entity.DB().Preload("Role").First(&rn, id) //
	if tx.Error != nil {
		return "Role not found"
	} else if rn.Role.Name == "Employee" {
		return "admin"
	} else if rn.Role.Name == "Student" || rn.Role.Name == "Teacher" {
		return "user"
	}
	return "err"
}

// POST /create
func CreateLoginUser(c *gin.Context) {
	var payload SignUpPayload
	var us entity.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	us.FirstName = payload.Name
	us.Email = payload.Email
	us.Password = string(hashPassword)

	if err := entity.DB().Create(&us).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": us})
}
