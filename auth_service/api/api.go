package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/axadjonovsardorbek/MiniTwitter/auth_service/api/docs"
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/api/handler"
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/api/middleware"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewApi(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/auth/user/register", h.Register)
	router.POST("/auth/user/login", h.Login)
	router.GET("/auth/user/profile/get", h.ProfileGetById)
	router.POST("/file-upload", h.UploadHandler)

	
	auth := router.Group("/auth", middleware.JWTMiddleware())
	user := auth.Group("/user", middleware.JWTMiddleware())
	user.GET("/profile", h.Profile)
	user.PUT("/update", h.UpdateUser)
	user.DELETE("/profile/delete", h.DeleteProfile)
	user.PUT("/password/change", h.ChangePassword)

	return router
}
