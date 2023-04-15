package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tesjwt.go/controllers"
	_ "tesjwt.go/docs"
	"tesjwt.go/middlewares"
)

// @title Mygram API
// @version 1.0
// @description This is a api to add photos, comments, and store the social media of users
// @termsOfService http://swagger.io/terms
// @contact.name API Support
// @contact.email redhomayan@gmail.com
// @license.name Apache 2.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @license.url http://www.apache.org/licenses/license-2.0.html
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		// Create
		userRouter.POST("/register", controllers.UserRegister)
		// Read
		userRouter.POST("/login", controllers.UserLogin)
	}

	socialmediaRouter := r.Group("/socialmedia")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		// Create
		socialmediaRouter.POST("/", controllers.CreateSocialMedia)
		// Read
		socialmediaRouter.GET("/", controllers.FindAllSocialMedia)
		// Update
		socialmediaRouter.PUT("/:socialmediaID", middlewares.Authorization(), controllers.UpdateSocialMedia)
		// Delete
		socialmediaRouter.DELETE("/:socialmediaID", middlewares.Authorization(), controllers.DeleteSocialMedia)
		// Read
		socialmediaRouter.GET("/:socialmediaID", middlewares.Authorization(), controllers.FindSocialMediaById)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		// Create
		photoRouter.POST("/", controllers.CreatePhoto)
		// Read
		photoRouter.GET("/", controllers.FindAllPhoto)
		// Update
		photoRouter.PUT("/:photoID", middlewares.Authorization(), controllers.UpdatePhoto)
		// Delete
		photoRouter.DELETE("/:photoID", middlewares.Authorization(), controllers.DeletePhoto)
		// Read
		photoRouter.GET("/:photoID", middlewares.Authorization(), controllers.FindPhotoById)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		// Create
		commentRouter.POST("/", controllers.CreateComment)
		// Read
		commentRouter.GET("/", controllers.FindAllComment)
		// Update
		commentRouter.PUT("/:commentID", middlewares.Authorization(), controllers.UpdateComment)
		// Delete
		commentRouter.DELETE("/:commentID", middlewares.Authorization(), controllers.DeleteComment)
		// Read
		commentRouter.GET("/:commentID", middlewares.Authorization(), controllers.FindCommentById)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
