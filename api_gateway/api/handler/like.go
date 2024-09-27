package handler

import (
	"context"
	"net/http"

	ap "github.com/axadjonovsardorbek/MiniTwitter/api_gateway/genproto/twit"

	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"

	"github.com/gin-gonic/gin"
)

// Like godoc
// @Summary Liking
// @Description Liking
// @Tags like
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param credentials body LikeCreateReq true "Like credentials"
// @Success 200 {object} string "Successfully liked"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /like [post]
func (h *Handler) Like(c *gin.Context) {
	var req ap.LikeCreateReq
	var body LikeCreateReq

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserId = id
	req.TwitId = body.TwitId

	_, err := h.srvs.Like.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully liked"})
}

// UnLike godoc
// @Summary UnLiking
// @Description UnLiking
// @Tags like
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Twit ID"
// @Success 200 {object} string "Successfully unliked"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /unlike [delete]
func (h *Handler) UnLike(c *gin.Context) {
	var req ap.LikeCreateReq

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	id := claims.(jwt.MapClaims)["user_id"].(string)

	twit_id := c.Query("id")

	if twit_id == "" || twit_id == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	req.UserId = id
	req.TwitId = twit_id

	_, err := h.srvs.Like.Delete(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully unliked"})
}

// Likes godoc
// @Summary Likes
// @Description Likes
// @Tags like
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} ap.LikeGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /likes/list/my [get]
func (h *Handler) MyLikesList(c *gin.Context) {
	var req ap.LikeGetAllReq

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	req.UserId = id

	res, err := h.srvs.Like.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Likes godoc
// @Summary Likes
// @Description Likes
// @Tags like
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Twit ID"
// @Success 200 {object} ap.LikeGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /likes/list [get]
func (h *Handler) LikesList(c *gin.Context) {
	var req ap.LikeGetAllReq

	twit_id := c.Query("id")

	if twit_id == "" || twit_id == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	req.TwitId = twit_id

	res, err := h.srvs.Like.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}