package controllers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"tesjwt.go/database"
	"tesjwt.go/helpers"
	"tesjwt.go/models"
)

// CreateSocialMedia godoc
// @Summary Create social media
// @Description Create social media of the user
// @Tags social media
// @Accept json
// @Produce json
// @Param name query string true "name"
// @Param social_media_url query string true "social_media_url"
// @Security BearerAuth
// @Success 201 {object} models.SocialMedia "Create social media success"
// @Failure 401 "Unauthorized"
// @Router /socialmedia [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

// UpdateSocialMedia godoc
// @Summary Update social media
// @Description Update social media identified by given id
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the social media"
// @Security BearerAuth
// @Success 200 {object} models.SocialMedia "Update social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/{socialmediaID} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialmediaID, _ := strconv.Atoi(c.Param("socialmediaID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialmediaID).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// DeleteSocialMedia godoc
// @Summary Delete social media
// @Description Delete social media identified by given ID
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the social media"
// @Security BearerAuth
// @Success 200 {string} string "Delete social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/{socialmediaID} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialmediaID, _ := strconv.Atoi(c.Param("socialmediaID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialmediaID).Delete(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SocialMedia deleted",
	})
}

// GetSocialMedia godoc
// @Summary Get social media
// @Description Get social media identified by given id
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the social media"
// @Security BearerAuth
// @Success 200 {object} models.SocialMedia "Get social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/{socialmediaID} [get]
func FindSocialMediaById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialmediaID, _ := strconv.Atoi(c.Param("socialmediaID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialmediaID).Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// GetAllSocialMedia godoc
// @Summary Get all social media
// @Description Get all social media in mygram
// @Tags social media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.SocialMedia "Get all social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia [get]
func FindAllSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	SocialMedia := []models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Debug().Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}
