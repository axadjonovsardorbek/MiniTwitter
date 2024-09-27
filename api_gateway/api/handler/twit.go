package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	ap "github.com/axadjonovsardorbek/MiniTwitter/api_gateway/genproto/twit"

	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"

	"github.com/gin-gonic/gin"
)

// Twit godoc
// @Summary Twit
// @Description Twit
// @Tags twit
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param credentials body TwitCreateReq true "Twit credentials"
// @Success 200 {object} string "Successfully twitted"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /twit/create [post]
func (h *Handler) CreateTwit(c *gin.Context) {
	var req ap.TwitCreateReq
	var body TwitCreateReq

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
	req.Content = body.Content
	req.ImageUrl = body.ImageUrl

	_, err := h.srvs.Twit.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully twitted"})
}

// Twit godoc
// @Summary Twit
// @Description Twit
// @Tags twit
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param credentials body ReTwitCreateReq true "ReTwit credentials"
// @Success 200 {object} string "Successfully retwitted"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /retwit/create [post]
func (h *Handler) CreateReTwit(c *gin.Context) {
	var req ap.TwitCreateReq
	var body ReTwitCreateReq

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
	req.Content = body.Content
	req.ImageUrl = body.ImageUrl
	req.TwitId = body.TwitId

	_, err := h.srvs.Twit.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully retwitted"})
}

// Twit godoc
// @Summary Twit
// @Description Twit
// @Tags twit
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Twit ID"
// @Param credentials body TwitCreateReq true "Twit credentials"
// @Success 200 {object} string "Successfully updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /twit/update [put]
func (h *Handler) UpdateTwit(c *gin.Context) {
	var req ap.TwitUpdateReq
	var body TwitCreateReq

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

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserId = id
	req.Id = twit_id
	req.Content = body.Content
	req.ImageUrl = body.ImageUrl

	_, err := h.srvs.Twit.Update(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully updated"})
}

// Twit godoc
// @Summary Twit
// @Description Twit
// @Tags twit
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Twit ID"
// @Success 200 {object} string "Successfully deleted"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /twit/delete [delete]
func (h *Handler) DeleteTwit(c *gin.Context) {
	var req ap.TwitDeleteReq

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
	req.Id = twit_id

	_, err := h.srvs.Twit.Delete(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully deleted"})
}

// Twit godoc
// @Summary Twit
// @Description Twit
// @Tags twit
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} ap.TwitGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /twit/list/my [get]
func (h *Handler) MyTwitsList(c *gin.Context) {
	req := ap.TwitGetAllReq{
		Filter: &ap.Filter{},
	}

	limit := c.Query("limit")
	offset := c.Query("offset")

	limitValue, offsetValue, err := parsePaginationParams(c, limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		slog.Error("Error parsing pagination parameters: ", err)
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	req.UserId = id
	req.Filter.Limit = int32(limitValue)
	req.Filter.Offset = int32(offsetValue)

	res, err := h.srvs.Twit.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Twit godoc
// @Summary Twit
// @Description Twit
// @Tags twit
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Twit ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} ap.TwitGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /retwit/list [get]
func (h *Handler) GetReTwits(c *gin.Context) {
	req := ap.TwitGetAllReq{
		Filter: &ap.Filter{},
	}

	limit := c.Query("limit")
	offset := c.Query("offset")
	twit_id := c.Query("id")

	if twit_id == "" || twit_id == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	limitValue, offsetValue, err := parsePaginationParams(c, limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		slog.Error("Error parsing pagination parameters: ", err)
		return
	}

	req.Filter.Limit = int32(limitValue)
	req.Filter.Offset = int32(offsetValue)
	req.TwitId = twit_id

	res, err := h.srvs.Twit.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Twit godoc
// @Summary Twit
// @Description Twit
// @Tags twit
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param content query string false "Content"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} ap.TwitGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /twit/search [get]
func (h *Handler) SearchTwit(c *gin.Context) {
	req := ap.TwitGetAllReq{
		Filter: &ap.Filter{},
	}

	limit := c.Query("limit")
	offset := c.Query("offset")
	content := c.Query("content")

	if content == "" || content == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid content"})
		return
	}

	limitValue, offsetValue, err := parsePaginationParams(c, limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		slog.Error("Error parsing pagination parameters: ", err)
		return
	}

	req.Filter.Limit = int32(limitValue)
	req.Filter.Offset = int32(offsetValue)
	req.Content = content

	res, err := h.srvs.Twit.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Twit godoc
// @Summary Twit
// @Description Twit
// @Tags twit
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} ap.TwitGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unathorized"
// @Router /twit/list [get]
func (h *Handler) GetAll(c *gin.Context) {
	req := ap.TwitGetAllReq{
		Filter: &ap.Filter{},
	}

	limit := c.Query("limit")
	offset := c.Query("offset")

	limitValue, offsetValue, err := parsePaginationParams(c, limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		slog.Error("Error parsing pagination parameters: ", err)
		return
	}

	req.Filter.Limit = int32(limitValue)
	req.Filter.Offset = int32(offsetValue)

	res, err := h.srvs.Twit.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Twit godoc
// @Summary Twit
// @Description Twit
// @Tags twit
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string false "Twit ID"
// @Success 200 {object} ap.TwitRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /twit/get [get]
func (h *Handler) GetById(c *gin.Context) {
	twit_id := c.Query("id")

	res, err := h.srvs.Twit.GetById(context.Background(), &ap.ById{Id: twit_id})
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error getting twit by ID: ", err)

		return
	}

	c.JSON(200, res)
}

func parsePaginationParams(c *gin.Context, limit, offset string) (int, int, error) {
	limitValue := 10
	offsetValue := 0

	if limit != "" {
		parsedLimit, err := strconv.Atoi(limit)
		if err != nil {
			slog.Error("Invalid limit value", err)
			c.JSON(400, gin.H{"error": "Invalid limit value"})
			return 0, 0, err
		}
		limitValue = parsedLimit
	}

	if offset != "" {
		parsedOffset, err := strconv.Atoi(offset)
		if err != nil {
			slog.Error("Invalid offset value", err)
			c.JSON(400, gin.H{"error": "Invalid offset value"})
			return 0, 0, err
		}
		offsetValue = parsedOffset
	}

	return limitValue, offsetValue, nil
}
