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

type CreateCommentReq struct {
	Message string `json:"message" form:"message" valid:"required~Message is required"`
	PhotoID uint   `json:"photo_id" form:"photo_id" valid:"required~Photo is required"`
}

type UpdateCommentReq struct {
	Message string `json:"message" form:"message" valid:"required~Message is required"`
}

// CreateComment godoc
// @Summary Create comment
// @Description Create comment for photo identified by given id
// @Tags comment
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Param message query string true "message"
// @Security BearerAuth
// @Success 201 {object} models.Comment "Create comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /comment/{photoId} [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	req := CreateCommentReq{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&req)
	} else {
		c.ShouldBind(&req)
	}
	Comment := models.Comment{
		UserID:  userID,
		PhotoID: req.PhotoID,
		Message: req.Message,
	}

	err := db.First(&models.Photo{}, req.PhotoID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "photo not found",
			"error":   err.Error(),
		})
		return
	}

	err = db.Debug().Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// UpdateComment godoc
// @Summary Update comment
// @Description Update comment identified by given id
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Security BearerAuth
// @Success 200 {object} models.Comment "Update comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comment Not Found"
// @Router /comment/{commentID} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	req := UpdateCommentReq{}

	CommentID, _ := strconv.Atoi(c.Param("commentID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&req)
	} else {
		c.ShouldBind(&req)
	}

	Comment := models.Comment{
		Message: req.Message,
	}

	err := db.Debug().Model(&Comment).Where("id = ?", CommentID).Updates(models.Comment{Message: req.Message}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Comment Updated",
	})
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment identified by given ID
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Security BearerAuth
// @Success 200 {string} string "Delete comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comment Not Found"
// @Router /comment/{commentID} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	CommentID, _ := strconv.Atoi(c.Param("commentID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(CommentID)

	err := db.Model(&Comment).Where("id = ?", CommentID).Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted",
	})
}

// GetAllCommentsForPhoto godoc
// @Summary Get all comments for specific photo
// @Description Get all comments for photo with given id
// @Tags comment
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {object} []models.Comment "Get all comments success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comments Not Found"
// @Router /comment/{photoID} [get]
func FindCommentById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	CommentID, _ := strconv.Atoi(c.Param("commentID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(CommentID)

	err := db.Model(&Comment).Where("id = ?", CommentID).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// GetAllComments godoc
// @Summary Get all comments
// @Description Get all comments in mygram
// @Tags comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Comment "Get all comments success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comments Not Found"
// @Router /comment [get]
func FindAllComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Comment := []models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Debug().Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}
