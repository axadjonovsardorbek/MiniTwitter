package handler

import (
	"context"
	"net/http"
	"regexp"

	ap "github.com/axadjonovsardorbek/MiniTwitter/auth_service/genproto/auth"

	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

const emailRegex = `^[a-zA-Z0-9._]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func isValidEmail(email string) bool {
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// Register godoc
// @Summary Register a user
// @Description Register user
// @Tags user
// @Accept json
// @Produce json
// @Param credentials body UserCreateReq true "User register credentials"
// @Success 200 {object} ap.TokenRes "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid username or password"
// @Router /auth/user/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req ap.UserCreateReq
	var body UserCreateReq

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidEmail(body.Email) {
		c.JSON(409, gin.H{"message": "Incorrect email"})
		return
	}

	req.Bio = body.Bio
	req.Email = body.Email
	req.ImageUrl = body.Image_url
	req.Username = body.Username
	req.Password = body.Password
	req.Name = body.Name

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(409, gin.H{"error": "Server error"})
		return
	}
	req.Password = string(hashedPassword)

	res, err := h.User.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user with username and password
// @Tags user
// @Accept json
// @Produce json
// @Param credentials body LoginReq true "User login credentials"
// @Success 200 {object} ap.TokenRes "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid username or password"
// @Router /auth/user/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req ap.LoginReq
	var body LoginReq

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Username = body.Username
	req.Password = body.Password

	res, err := h.User.Login(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary Update a user
// @Description Update a user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param credentials body UserUpdateReq true "User update credentials"
// @Success 200 {object} string "User updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unauthorized"
// @Router /auth/user/update [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var req ap.UserUpdateReq
	var body UserUpdateReq

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

	req.Id = id
	req.Email = body.Email
	req.Username = body.Username
	req.Name = body.Name
	req.ImageUrl = body.Image_url
	req.Bio = body.Bio

	_, err := h.User.Update(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// Get godoc
// @Summary Get a profile
// @Description Get a profile
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} ap.UserRes "Profile"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid username or password"
// @Router /auth/user/profile [get]
func (h *Handler) Profile(c *gin.Context) {
	var req ap.ById

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	req.Id = id

	res, err := h.User.GetById(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Get godoc
// @Summary Get a profile
// @Description Get a profile
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Id"
// @Success 200 {object} ap.UserRes "Profile"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid username or password"
// @Router /auth/user/profile/get [get]
func (h *Handler) ProfileGetById(c *gin.Context) {
	var req ap.ById

	id := c.Query("id")

	if id == "" || id == "string" {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	req.Id = id

	res, err := h.User.GetById(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} string "User deleted"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid username or password"
// @Router /auth/user/profile/delete [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	var req ap.ById

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	req.Id = id

	_, err := h.User.Delete(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// ChangePassword godoc
// @Summary ChangePassword
// @Description ChangePassword
// @Tags user
// @Accept json
// @Produce json
// @Param credentials body ChangePasswordReq true "Passwrod change credentials"
// @Success 200 {object} string "Changed password"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Invalid password"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /auth/user/password/change [put]
func (h *Handler) ChangePassword(c *gin.Context) {
	var req ap.ChangePasswordReq
	var body ChangePasswordReq

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	if body.OldPassword == "" || body.OldPassword == "string" {
		c.JSON(409, gin.H{"error": "Invalid password"})
		return
	}
	if body.NewPassword == "" || body.NewPassword == "string" {
		c.JSON(409, gin.H{"error": "Invalid password"})
		return
	}

	req.Id = id
	req.OldPassword = body.OldPassword
	req.NewPassword = body.NewPassword

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(409, gin.H{"error": "Server error"})
		return
	}

	req.NewPassword = string(hashedPassword)

	_, err = h.User.ChangePassword(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Changed password"})
}
