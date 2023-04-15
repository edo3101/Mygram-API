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

// CreatePhoto godoc
// @Summary Create photo
// @Description Create photo to post in mygram
// @Tags photo
// @Accept json
// @Produce json
// @Param title query string true "title"
// @Param caption query string false "caption"
// @Param photo_url query string true "photo_url"
// @Security BearerAuth
// @Success 201 {object} models.Photo "Create photo success"
// @Failure 401 "Unauthorized"
// @Router /photo [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// UpdatePhoto godoc
// @Summary Update photo
// @Description Update photo identified by given ID
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {object} models.Photo{} "Update photo success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /photo/{photoID} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	PhotoID, _ := strconv.Atoi(c.Param("photoID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(PhotoID)

	err := db.Model(&Photo).Where("id = ?", PhotoID).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto godoc
// @Summary Delete photo
// @Description Delete photo identified by given ID
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {string} string "Delete photo success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /photo/{photoID} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	PhotoID, _ := strconv.Atoi(c.Param("photoID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(PhotoID)

	err := db.Model(&Photo).Where("id = ?", PhotoID).Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo deleted",
	})
}

// GetPhoto godoc
// @Summary Get photo
// @Description Get photo by ID
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {object} models.Photo{} "Get photo success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /photo/{photoID} [get]
func FindPhotoById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	PhotoID, _ := strconv.Atoi(c.Param("photoID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(PhotoID)

	err := db.Model(&Photo).Where("id = ?", PhotoID).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// GetAllPhotos godoc
// @Summary Get all photos
// @Description Get all existing photos
// @Tags photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Photo{} "Get all photos success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photos Not Found"
// @Router /photo [get]
func FindAllPhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Photo := []models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Debug().Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}
