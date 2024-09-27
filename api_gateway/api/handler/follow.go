package handler

import (
	"context"
	"net/http"

	ap "github.com/axadjonovsardorbek/MiniTwitter/api_gateway/genproto/twit"

	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"

	"github.com/gin-gonic/gin"
)

// Follow godoc
// @Summary Following
// @Description Following
// @Tags follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param credentials body FollowCreateReq true "Follow credentials"
// @Success 200 {object} string "Successfully followed"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /follow [post]
func (h *Handler) Follow(c *gin.Context) {
	var req ap.FollowCreateReq
	var body FollowCreateReq

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

	req.FollowerId = id
	req.FollowedId = body.FollowedId

	_, err := h.srvs.Follow.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully followed"})
}

// UnFollow godoc
// @Summary UnFollowing
// @Description UnFollowing
// @Tags follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Follow Id"
// @Success 200 {object} string "Successfully unfollowed"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /unfollow [delete]
func (h *Handler) UnFollow(c *gin.Context) {
	var req ap.FollowCreateReq

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	followed_id := c.Query("id")

	if followed_id == "" || followed_id == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	req.FollowerId = id
	req.FollowedId = followed_id

	_, err := h.srvs.Follow.Delete(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully unfollowed"})
}

// Followers godoc
// @Summary Followers
// @Description Followers
// @Tags follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} ap.FollowGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /followers/list/my [get]
func (h *Handler) MyFollowersList(c *gin.Context) {
	var req ap.FollowGetAllReq

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	req.FollowedId = id

	res, err := h.srvs.Follow.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Follows godoc
// @Summary Follows
// @Description Follows
// @Tags follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} ap.FollowGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /follows/list/my [get]
func (h *Handler) MyFollowsList(c *gin.Context) {
	var req ap.FollowGetAllReq

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	req.FollowerId = id

	res, err := h.srvs.Follow.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Follows godoc
// @Summary Follows
// @Description Follows
// @Tags follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Follower ID"
// @Success 200 {object} string "Successfully removed"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /followers/delete [delete]
func (h *Handler) MyFollowersDelete(c *gin.Context) {
	var req ap.FollowCreateReq

	follower_id := c.Query("id")

	if follower_id == "" || follower_id == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	req.FollowedId = id
	req.FollowerId = follower_id

	_, err := h.srvs.Follow.Delete(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully removed"})
}

// Followers godoc
// @Summary Followers
// @Description Followers
// @Tags follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Followed ID"
// @Success 200 {object} ap.FollowGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /followers/list [get]
func (h *Handler) FollowersList(c *gin.Context) {
	var req ap.FollowGetAllReq

	followed_id := c.Query("id")

	if followed_id == "" || followed_id == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	req.FollowedId = followed_id

	res, err := h.srvs.Follow.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Follows godoc
// @Summary Follows
// @Description Follows
// @Tags follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Follower ID"
// @Success 200 {object} ap.FollowGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /follows/list [get]
func (h *Handler) FollowsList(c *gin.Context) {
	var req ap.FollowGetAllReq

	follower_id := c.Query("id")

	if follower_id == "" || follower_id == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	req.FollowerId = follower_id

	res, err := h.srvs.Follow.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
