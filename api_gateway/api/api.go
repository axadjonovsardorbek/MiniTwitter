package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/axadjonovsardorbek/MiniTwitter/api_gateway/api/docs"
	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/api/handler"
	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/api/middleware"
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
	
	follow := router.Group("/", middleware.JWTMiddleware())
	follow.POST("/follow", h.Follow)
	follow.DELETE("/unfollow", h.UnFollow)
	follow.GET("/followers/list/my", h.MyFollowersList)
	follow.GET("/follows/list/my", h.MyFollowsList)
	follow.DELETE("/followers/delete", h.MyFollowersDelete)
	follow.GET("/followers/list", h.FollowersList)
	follow.GET("/follows/list", h.FollowsList)

	like := router.Group("/", middleware.JWTMiddleware())
	like.POST("/like", h.Like)
	like.DELETE("/unlike", h.UnLike)
	like.GET("/likes/list/my", h.MyLikesList)
	like.GET("/likes/list", h.LikesList)

	twit := router.Group("/", middleware.JWTMiddleware())
	twit.POST("/twit/create", h.CreateTwit)
	twit.POST("/retwit/create", h.CreateReTwit)
	twit.PUT("/twit/update", h.UpdateTwit)
	twit.DELETE("/twit/delete", h.DeleteTwit)
	twit.GET("/twit/list/my", h.MyTwitsList)
	twit.GET("/retwit/list", h.GetReTwits)
	twit.GET("/twit/list", h.GetAll)
	twit.GET("/twit/search", h.SearchTwit)
	twit.GET("/twit/get", h.GetById)

	return router
}
