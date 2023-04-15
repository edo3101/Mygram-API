package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tesjwt.go/database"
	"tesjwt.go/helpers"
	"tesjwt.go/models"
)

var appJSON = "application/json"

// UserRegister godoc
// @Summary Register user
// @Description Register new user
// @Tags user
// @Accept json
// @Produce json
// @Param username query string true "username"
// @Param email query string true "email"
// @Param password query string true "password"
// @Param age query int true "age"
// @Success 201 {object} models.User "Register success response"
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.Username,
		"age":       User.Age,
	})
}

// UserLogin godoc
// @Summary Login user
// @Description Login user by email
// @Tags user
// @Accept json
// @Produce json
// @Param email query string true "email"
// @Param password query string true "password"
// @Success 200 {object} interface{} "Login response"
// @Failure 401 "Unauthorized"
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
